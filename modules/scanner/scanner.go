package scanner

import (
	"encoding/json"
	"github.com/KMA-Score/kma_score_scanner/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"strings"
)

type SubjectStudentCore struct {
	SubjectName   string
	SubjectCode   string
	SubjectCredit string
	StudentScores []StudentSubjectScore
}

type StudentSubjectScore struct {
	StudentCode       string
	StudentName       string
	StudentClass      string
	StudentScoreTP1   float32
	StudentScoreTP2   float32
	StudentScoreExam  float32
	StudentScoreFinal float32
	StudentScoreStr   string
}

func HandleScanCommand(input string, output string) {
	// Determine input is file or directory
	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	fileStat, err := f.Stat()
	if err != nil {
		panic(err)
	}

	var scannedScores = map[string]SubjectStudentCore{}

	if fileStat.IsDir() {
		// Scan all files in directory
		files, err := f.Readdir(0)

		if err != nil {
			log.Error().Err(err).Msg("Failed to read directory")
			return
		}

		log.Info().Msgf("Found %d files", len(files))

		for _, file := range files {
			log.Info().Msgf("Scanning file %s", file.Name())

			// Scan file
			scanOutput := scanFile(filepath.Join(input, file.Name()))

			// Merge scores
			for key, value := range scanOutput {
				oldVal := scannedScores[key]

				if oldVal.SubjectName == "" {
					oldVal.SubjectName = value.SubjectName
				}
				if oldVal.SubjectCredit == "" {
					oldVal.SubjectCredit = value.SubjectCredit
				}

				oldVal.StudentScores = append(oldVal.StudentScores, value.StudentScores...)
				scannedScores[key] = oldVal
			}
		}
	} else {
		// Scan only one file
		scannedScores = scanFile(input)
	}

	// Export to JSON
	v, _ := json.MarshalIndent(scannedScores, "", "    ")

	// Write to file
	err = os.WriteFile("outputAll.json", v, 0644)
	if err != nil {
		log.Error().Err(err).Msg("Failed to write to file")
		return
	}
}

func scanFile(input string) map[string]SubjectStudentCore {
	// Open file
	f, err := os.Open(input)

	if err != nil {
		log.Fatal().Err(err).Msgf("Cannot open file %s", input)
		os.Exit(1)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(f)

	if err != nil {
		log.Fatal().Err(err).Msg("Cannot load HTML document")
		os.Exit(1)
	}

	// Find the review items. Which is <p> tag with text "Số TC:" and <table> tag
	allSelection := doc.Find("p:contains(\"Số TC:\"), table")

	// Global map to store all scores and subject info
	var ssScores = make(map[string]SubjectStudentCore)

	// This variable is used to store the subject code for the current table. Prevent the subject code is not available in some tables
	globalSubjectCode := ""

	for i := range allSelection.Nodes {
		s := allSelection.Eq(i)

		// Check if the current selection is <p> tag
		if s.Is("p") {
			// Split the string by "Số TC:"
			parts := strings.Split(s.Text(), "Số TC:")

			// The first part contains the course name, split it by ":"
			courseParts := strings.Split(parts[0], ":")
			subjectName := strings.TrimSpace(courseParts[1])

			// The second part contains the course credit and code, split it by "Mã học phần:"
			creditParts := strings.Split(parts[1], "Mã học phần:")
			subjectCredit := strings.TrimSpace(creditParts[0])
			subjectCode := strings.TrimSpace(creditParts[1])

			// Check subjectCode exist in map
			_, ok := ssScores[subjectCode]

			if !ok {
				ssSubject := SubjectStudentCore{}
				ssSubject.SubjectName = subjectName
				ssSubject.SubjectCode = subjectCode
				ssSubject.SubjectCredit = subjectCredit

				ssScores[subjectCode] = ssSubject

				// Save global subject code for later use. Because the subject code is not available in some tables
				globalSubjectCode = subjectCode
			}

		} else {
			var tableStudentScores []StudentSubjectScore
			trNodes := s.Find("tr")

			// Skip table of content table (First appear on semester 2 2020-2021)
			if strings.HasPrefix(trNodes.Eq(0).Text(), "TT") {
				continue
			}

			// Loop each row of table
			for rowIndex := range trNodes.Nodes {
				trSelection := trNodes.Eq(rowIndex)

				// Skip header
				if strings.Contains(trSelection.Text(), "STT") {
					continue
				}

				tdNodes := trSelection.Find("td")

				col2 := tdNodes.Eq(2).Text() // student code

				// Skip empty row
				if col2 == "" {
					continue
				}

				col3 := tdNodes.Eq(3).Text()   // student name
				col4 := tdNodes.Eq(4).Text()   // student name or class
				col5 := tdNodes.Eq(5).Text()   // class
				col6 := tdNodes.Eq(6).Text()   // tp1
				col7 := tdNodes.Eq(7).Text()   // tp2
				col8 := tdNodes.Eq(8).Text()   // exam
				col9 := tdNodes.Eq(9).Text()   // final
				col10 := tdNodes.Eq(10).Text() // str

				studentSubjectScore := StudentSubjectScore{}

				/*
					In original file, the student name is split into 2 columns.
					But when export to PDF and convert back to HTML, sometimes, the student name is split into two columns or one.
					We need to check if the column 4 is a class or not by checking if it starts with "CT", "AT", "DT"
				*/
				if !utils.CheckStudentClassColIsReal(col4) {
					studentSubjectScore.StudentCode = col2
					studentSubjectScore.StudentName = utils.ProcessStudentName(col3 + " " + col4)
					studentSubjectScore.StudentClass = col5
					studentSubjectScore.StudentScoreTP1 = utils.ProcessFloatScore(col6)
					studentSubjectScore.StudentScoreTP2 = utils.ProcessFloatScore(col7)
					studentSubjectScore.StudentScoreExam = utils.ProcessFloatScore(col8)
					studentSubjectScore.StudentScoreFinal = utils.ProcessFloatScore(col9)
					studentSubjectScore.StudentScoreStr = utils.CleanStudentStringScore(col10)
				} else {
					studentSubjectScore.StudentCode = col2
					studentSubjectScore.StudentName = utils.ProcessStudentName(col3)
					studentSubjectScore.StudentClass = col4
					studentSubjectScore.StudentScoreTP1 = utils.ProcessFloatScore(col5)
					studentSubjectScore.StudentScoreTP2 = utils.ProcessFloatScore(col6)
					studentSubjectScore.StudentScoreExam = utils.ProcessFloatScore(col7)
					studentSubjectScore.StudentScoreFinal = utils.ProcessFloatScore(col8)
					studentSubjectScore.StudentScoreStr = utils.CleanStudentStringScore(col9)
				}

				tableStudentScores = append(tableStudentScores, studentSubjectScore)
			}

			// Get subject object then add scores to it
			subjectFromMap := ssScores[globalSubjectCode]
			subjectFromMap.StudentScores = append(subjectFromMap.StudentScores, tableStudentScores...)
			ssScores[globalSubjectCode] = subjectFromMap
		}
	}

	return ssScores
}
