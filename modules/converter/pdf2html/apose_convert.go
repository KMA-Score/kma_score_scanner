package pdf2html

import (
	"context"
	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2401/api"
	"github.com/aspose-words-cloud/aspose-words-cloud-go/v2401/api/models"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"io"
	"os"
)

func CreateApi() (wordsApi *api.WordsApiService, ctx context.Context) {
	config := models.Configuration{
		BaseUrl:      viper.Get("aspose.baseUrl").(string),
		ClientId:     viper.Get("aspose.clientId").(string),
		ClientSecret: viper.Get("aspose.clientSecret").(string),
	}

	wordsApi, ctx, err := api.CreateWordsApi(&config)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create api")

	}

	return wordsApi, ctx
}

func ConvertPDFtoHTML(wordsApi *api.WordsApiService, ctx context.Context, inputFile string, outputFile string) {
	doc, err := os.Open(inputFile)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open input file")
	}

	options := map[string]interface{}{}

	format := "html"

	request := &models.ConvertDocumentRequest{
		Document:  doc,
		Format:    &format,
		Optionals: options,
	}

	convert, err := wordsApi.ConvertDocument(ctx, request)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to convert document")
	}

	defer convert.Body.Close()

	// Save output file
	b, _ := io.ReadAll(convert.Body)
	err = os.WriteFile(outputFile, b, 0644)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to write output file")
	}
}
