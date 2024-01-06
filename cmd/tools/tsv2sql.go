package tools

import (
	"github.com/KMA-Score/kma_score_scanner/modules/converter/tsv2sql"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var Tsv2SqlCmd = &cobra.Command{
	Use:     "tsv2sql [flags] [input path]",
	Aliases: []string{"t2s, t2sql, tsvtosql"},
	Short:   "Convert TSV to SQL commands",
	Long: `Convert TSV to SQL commands to modify database.
INPUT path must be a directory containing TSV files (students.tsv, scores.tsv, subjects.tsv).`,
	Example: "kma_score_scanner tools tsv2sql ./input -o ./output",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputPath := args[0]

		if inputPath == "" {
			_ = cmd.Help()
			return
		}

		outputPath, err := cmd.Flags().GetString("output")

		if err != nil {
			log.Error().Msgf("Error while getting output path: %s", err.Error())
			return
		}

		tsv2sql.HandleCommand(inputPath, outputPath)
	},
}

func init() {
	Tsv2SqlCmd.Flags().StringP("output", "o", "", "Output path for directory")
}
