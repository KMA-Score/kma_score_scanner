package main

import (
	"github.com/KMA-Score/kma_score_scanner/cmd"
	"github.com/KMA-Score/kma_score_scanner/utils"
)

func main() {
	// Create logger and config
	utils.CreateLogger()
	utils.CreateConfig()

	// Execute root command
	cmd.Execute()
}
