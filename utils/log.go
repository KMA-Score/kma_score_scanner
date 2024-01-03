package utils

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
	"time"
)

func getLogLevelColor(logLevel string) color.Attribute {
	switch logLevel {
	case "debug":
		return color.FgMagenta
	case "info":
		return color.FgBlue
	case "warn":
		return color.FgYellow
	case "error":
		return color.FgRed
	case "fatal":
		return color.FgRed
	default:
		return color.FgWhite
	}
}

func CreateLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}

	logLevelRn := ""

	output.FormatLevel = func(i interface{}) string {
		logLevel := i.(string)
		c := getLogLevelColor(logLevel)
		d := color.New(c)

		// Save the log level for later use
		logLevelRn = logLevel

		return d.Sprintf("[%s]", strings.ToUpper(logLevel))
	}

	output.FormatMessage = func(i interface{}) string {
		c := getLogLevelColor(logLevelRn)
		d := color.New(c)

		return d.Sprintf("%s", i)
	}

	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}

	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	log.Logger = log.Output(output)
}
