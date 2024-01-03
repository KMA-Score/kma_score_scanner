package cmd

import (
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.KMA_Score_Scanner.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
