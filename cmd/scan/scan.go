package scan

import (
	"github.com/KMA-Score/kma_score_scanner/modules/scanner"
	"github.com/spf13/cobra"
)

var ScanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan scores and students' data from HTML and export to TSV",
	Long:  `Scan scores and students' data from HTML and export to TSV`,
	Run: func(cmd *cobra.Command, args []string) {
		if !cmd.Flags().Changed("input") {
			_ = cmd.Help()
			return
		}

		input, _ := cmd.Flags().GetString("input")
		output, _ := cmd.Flags().GetString("output")

		scanner.HandleScanCommand(input, output)
	},
}

func init() {
	ScanCmd.Flags().StringP("input", "i", "", "Input path for HTML file or directory")
	ScanCmd.Flags().StringP("output", "o", "", "Output directory path for TSV file")
}
