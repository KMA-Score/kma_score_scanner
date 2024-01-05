package exporter

import (
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"github.com/rs/zerolog/log"
	"golang.org/x/exp/maps"
	"io"
	"os"
)

type StudentInfo struct {
	StudentCode  string `csv:"studentCode"`
	StudentName  string `csv:"studentName"`
	StudentClass string `csv:"studentClass"`
}

type SubjectInfo struct {
	SubjectName   string `csv:"subjectName"`
	SubjectCode   string `csv:"subjectCode"`
	SubjectCredit string `csv:"subjectCredit"`
}

type SubjectStudentCore struct {
	SubjectName   string
	SubjectCode   string
	SubjectCredit string
	StudentScores []StudentScore
}

type StudentScore struct {
	StudentCode       string  `csv:"studentCode"`
	StudentName       string  `csv:"-"`
	StudentClass      string  `csv:"-"`
	StudentScoreTP1   float32 `csv:"studentScoreTP1"`
	StudentScoreTP2   float32 `csv:"studentScoreTP2"`
	StudentScoreExam  float32 `csv:"studentScoreExam"`
	StudentScoreFinal float32 `csv:"studentScoreFinal"`
	StudentScoreStr   string  `csv:"studentScoreStr"`
	SubjectCode       string  `csv:"subjectCode"`
}

func TsvExport(subjectScoreMap map[string]SubjectStudentCore, output string) {

	// Setup csv writer to use tab as delimiter
	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(out)
		writer.Comma = '\t'
		return gocsv.NewSafeCSVWriter(writer)
	})

	var studentsDataMap = make(map[string]StudentInfo)
	var subjectsDataMap = make(map[string]SubjectInfo)
	var studentScores []StudentScore

	for _, value := range subjectScoreMap {
		_, subjectExistYet := subjectsDataMap[value.SubjectCode]

		if !subjectExistYet {
			subjectInfo := SubjectInfo{
				SubjectName:   value.SubjectName,
				SubjectCode:   value.SubjectCode,
				SubjectCredit: value.SubjectCredit,
			}

			subjectsDataMap[value.SubjectCode] = subjectInfo
		}

		for _, studentScore := range value.StudentScores {

			// Append to student data if not exist
			_, studentExistYet := studentsDataMap[studentScore.StudentCode]
			if !studentExistYet {
				studentInfo := StudentInfo{
					StudentCode:  studentScore.StudentCode,
					StudentName:  studentScore.StudentName,
					StudentClass: studentScore.StudentClass,
				}

				studentsDataMap[studentScore.StudentCode] = studentInfo
			}

			// Append to student score with subject code
			studentScore.SubjectCode = value.SubjectCode
			studentScores = append(studentScores, studentScore)
		}
	}

	studentsData := maps.Values(studentsDataMap)
	subjectsData := maps.Values(subjectsDataMap)

	studentsTsvContent, err := gocsv.MarshalString(&studentsData)
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal TSV students data")
	}
	saveToFile(output, "students.tsv", studentsTsvContent)
	log.Info().Msg("Exported students data to TSV")

	subjectsTsvContent, err := gocsv.MarshalString(&subjectsData)
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal TSV subjects data")
	}
	saveToFile(output, "subjects.tsv", subjectsTsvContent)
	log.Info().Msg("Exported students data to TSV")

	studentScoresTsvContent, err := gocsv.MarshalString(&studentScores)
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal TSV student scores data")
	}
	saveToFile(output, "scores.tsv", studentScoresTsvContent)
	log.Info().Msg("Exported students data to TSV")
}

func saveToFile(output string, fileName string, content string) {
	// Create file
	err := os.WriteFile(output+fileName, []byte(content), 0644)

	if err != nil {
		log.Error().Err(err).Msgf("Failed to create file %s", fileName)
	}
}
