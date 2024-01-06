package tsv2sql

import (
	"encoding/csv"
	"fmt"
	"github.com/KMA-Score/kma_score_scanner/modules/exporter"
	"github.com/gocarina/gocsv"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"path"
	"strings"
)

func HandleCommand(inputPath string, outputPath string) {
	if outputPath == "" {
		outputPath = inputPath
	}

	dir, err := os.Open(inputPath)

	if err != nil {
		log.Error().Err(err).Msg("Failed to open input directory")
		return
	}

	defer dir.Close()

	dirInfo, err := dir.Stat()

	if err != nil {
		log.Error().Err(err).Msg("Failed to get stat of input directory")
		return
	}

	if !dirInfo.IsDir() {
		log.Error().Msg("Input path is not a directory")
		return
	}

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = '\t'
		return r
	})

	convertStudents(inputPath, outputPath)
	convertSubjects(inputPath, outputPath)
	convertScores(inputPath, outputPath)
}

func convertStudents(inputPath string, outputPath string) {
	tsvFile, err := os.Open(path.Join(inputPath, "students.tsv"))

	if err != nil {
		log.Warn().Msg("students.tsv not found. Skipping...")
		return
	}

	defer tsvFile.Close()

	// Read TSV file

	var tsvContent []exporter.StudentInfo
	err = gocsv.UnmarshalFile(tsvFile, &tsvContent)

	if err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal students.tsv file")
		return
	}

	sqlCommands := make([]string, 0)

	for _, student := range tsvContent {
		sqlCommand := fmt.Sprintf("INSERT INTO Students (Id,Name, Class) VALUES (\"%s\", \"%s\", \"%s\") ON DUPLICATE KEY UPDATE Name=\"%s\",Class=\"%s\";",
			student.StudentCode, student.StudentName, student.StudentClass, student.StudentName, student.StudentClass)
		sqlCommands = append(sqlCommands, sqlCommand)
	}

	saveToFile(outputPath, "output-students.sql", sqlCommands)
}

func convertSubjects(inputPath string, outputPath string) {
	tsvFile, err := os.Open(path.Join(inputPath, "subjects.tsv"))

	if err != nil {
		log.Warn().Msg("subjects.tsv not found. Skipping...")
		return
	}

	defer tsvFile.Close()

	// Read TSV file

	var tsvContent []exporter.SubjectInfo
	err = gocsv.UnmarshalFile(tsvFile, &tsvContent)

	if err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal subjects.tsv file")
		return
	}

	sqlCommands := make([]string, 0)

	for _, items := range tsvContent {
		sqlCommand := fmt.Sprintf("INSERT INTO Subjects (Id,Name,NumberOfCredits) VALUES (\"%s\", \"%s\", %s) "+
			"ON DUPLICATE KEY UPDATE Name=\"%s\",NumberOfCredits=%s;",
			items.SubjectCode, items.SubjectName, items.SubjectCredit, items.SubjectName, items.SubjectCredit)
		sqlCommands = append(sqlCommands, sqlCommand)
	}

	saveToFile(outputPath, "output-subjects.sql", sqlCommands)
}

func convertScores(inputPath string, outputPath string) {
	tsvFile, err := os.Open(path.Join(inputPath, "scores.tsv"))

	if err != nil {
		log.Warn().Msg("scores.tsv not found. Skipping...")
		return
	}

	defer tsvFile.Close()

	// Read TSV file

	var tsvContent []exporter.StudentScore
	err = gocsv.UnmarshalFile(tsvFile, &tsvContent)

	if err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal subjects.tsv file")
		return
	}

	sqlCommands := make([]string, 0)

	for _, items := range tsvContent {
		sqlCommand := fmt.Sprintf("INSERT INTO Scores "+
			"(StudentId,SubjectId,FirstComponentScore,SecondComponentScore,ExamScore,AvgScore,AlphabetScore) "+
			"VALUES (\"%s\", \"%s\", %.1f, %.1f, %.1f ,%.1f, \"%s\") "+
			"ON DUPLICATE KEY UPDATE "+
			"FirstComponentScore=%.f,SecondComponentScore=%.f,ExamScore=%.1f,AvgScore=%.1f,AlphabetScore=\"%s\";",
			items.StudentCode, items.SubjectCode, items.StudentScoreTP1, items.StudentScoreTP2, items.StudentScoreExam,
			items.StudentScoreFinal, items.StudentScoreStr, items.StudentScoreTP1, items.StudentScoreTP2,
			items.StudentScoreExam, items.StudentScoreFinal, items.StudentScoreStr)
		sqlCommands = append(sqlCommands, sqlCommand)
	}

	saveToFile(outputPath, "output-scores.sql", sqlCommands)
}

func saveToFile(output string, fileName string, content []string) {
	err := os.MkdirAll(output, os.ModePerm)
	if err != nil {
		log.Warn().Err(err).Msgf("Error when create directory %s", output)
	}

	err = os.WriteFile(path.Join(output, fileName), []byte(strings.Join(content, "\n")), 0644)

	if err != nil {
		log.Error().Err(err).Msgf("Failed to create file %s", fileName)
	}

	log.Info().Msgf("Exported %s to %s", fileName, output)
}
