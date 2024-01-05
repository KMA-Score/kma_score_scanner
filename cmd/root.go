package cmd

import (
	"github.com/KMA-Score/kma_score_scanner/cmd/scan"
	"github.com/KMA-Score/kma_score_scanner/cmd/tools"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "kma_score_scanner",
	Short: "A CLI tool use to scan score from PDF",
	Long: `A Fast and Flexible CLI tool use to scan score from PDF.
Source code: https://github.com/KMA-Score
Authors: Lucas & Arahiko`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(tools.ToolsCmd)
	rootCmd.AddCommand(scan.ScanCmd)
}
