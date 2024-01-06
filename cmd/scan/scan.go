package scan

import (
	"github.com/KMA-Score/kma_score_scanner/modules/scanner"
	"github.com/spf13/cobra"
)

var ScanCmd = &cobra.Command{
	Use:     "scan [flags] [input path]",
	Aliases: []string{"s"},
	Short:   "Scan scores and students' data from HTML and export to TSV",
	Long:    `Scan scores and students' data from HTML and export to TSV`,
	Example: "kma_score_scanner scan ./input -o ./output",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		input := args[0]

		if input == "" {
			_ = cmd.Help()
			return
		}

		output, _ := cmd.Flags().GetString("output")

		scanner.HandleScanCommand(input, output)
	},
}

func init() {
	ScanCmd.Flags().StringP("output", "o", "", "Output directory path for TSV file")
}
