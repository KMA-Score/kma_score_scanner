package exporter

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"os"
)

func ExportJson(subjectScoreMap map[string]SubjectStudentCore, output string) {
	// Export to JSON
	v, _ := json.MarshalIndent(subjectScoreMap, "", "    ")

	// Write to file
	err := os.WriteFile("outputAll.json", v, 0644)

	if err != nil {
		log.Error().Err(err).Msg("Failed to write to file")
		return
	}
}
