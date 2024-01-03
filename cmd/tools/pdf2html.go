package tools

import (
	"github.com/KMA-Score/kma_score_scanner/modules/converter/pdf2html"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var Pdf2htmlCmd = &cobra.Command{
	Use:   "pdf2html",
	Short: "Convert PDF to HTML",
	Long:  `Convert PDF to HTML using Aspose Cloud API (Self-hosted or cloud hosted).`,
	Run: func(cmd *cobra.Command, args []string) {
		inputPath, err := cmd.Flags().GetString("input")

		if err != nil {
			log.Error().Msgf("Error while getting input path: %s", err.Error())
			return
		}

		outputPath, err := cmd.Flags().GetString("output")

		if err != nil {
			log.Error().Msgf("Error while getting output path: %s", err.Error())
			return
		}

		// If input or output path is not set, show help
		if !cmd.Flags().Changed("input") || !cmd.Flags().Changed("output") {
			_ = cmd.Help()
			return
		}

		pdf2html.HandlePDFtoHTMLCommand(inputPath, outputPath)
	},
}

func init() {
	Pdf2htmlCmd.Flags().StringP("input", "i", "", "Input path for file or directory")
	Pdf2htmlCmd.Flags().StringP("output", "o", "", "Output path for file or directory")
}
