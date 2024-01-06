package tools

import (
	"github.com/KMA-Score/kma_score_scanner/modules/converter/pdf2html"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var Pdf2htmlCmd = &cobra.Command{
	Use:     "pdf2html [flags] [input path]",
	Aliases: []string{"p2h, pdftohtml, p2html"},
	Short:   "Convert PDF to HTML",
	Long:    `Convert PDF to HTML using Aspose Cloud API (Self-hosted or cloud hosted).`,
	Example: "kma_score_scanner tools pdf2html ./input -o ./output",
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

		// If input or output path is not set, show help
		if !cmd.Flags().Changed("output") {
			_ = cmd.Help()
			return
		}

		pdf2html.HandlePDFtoHTMLCommand(inputPath, outputPath)
	},
}

func init() {
	Pdf2htmlCmd.Flags().StringP("output", "o", "", "Output path for file or directory")
}
