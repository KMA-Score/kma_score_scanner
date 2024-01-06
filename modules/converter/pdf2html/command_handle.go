package pdf2html

import (
	"github.com/briandowns/spinner"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func HandlePDFtoHTMLCommand(input string, output string) {
	// Sanity check
	if input == "" || output == "" {
		log.Fatal().Msg("No input or output specified")
		return
	}

	// If input is a directory, make sure output is also a directory
	inputFile, err := os.Open(input)
	if err != nil {
		log.Error().Err(err).Msg("Failed to open input file")
		return
	}

	defer inputFile.Close()

	inputFileInfo, err := inputFile.Stat()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get stat of input file")
		return
	}

	// Create api module
	wordsApi, ctx := CreateApi()

	// Convert pdf to html
	spin := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	spin.Start()

	if !inputFileInfo.IsDir() {
		ConvertPDFtoHTML(wordsApi, ctx, input, output)
		return
	} else {
		files, err := inputFile.Readdir(0)

		if err != nil {
			log.Error().Err(err).Msg("Failed to read directory")
			return
		}

		// Filter html files
		var pdfFiles []os.FileInfo

		for _, file := range files {
			if strings.HasSuffix(file.Name(), ".pdf") {
				pdfFiles = append(pdfFiles, file)
			}
		}

		log.Info().Msgf("Found %d files", len(pdfFiles))

		for _, file := range pdfFiles {
			ConvertPDFtoHTML(wordsApi, ctx, filepath.Join(input, file.Name()), filepath.Join(output, file.Name()+".html"))
		}
	}

	spin.Stop()

	log.Info().Msg("Done")
}
