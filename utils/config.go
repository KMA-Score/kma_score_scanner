package utils

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
)

func CreateConfig() {
	// Check if config file exists
	_, err := os.Stat("config.json")

	if os.IsNotExist(err) {
		log.Warn().Msg("Config file does not exist. Create default config file...")

		// Create default config file by copying from config.default.json
		_, err := Copy("config.example.json", "config.json")

		if err != nil {
			log.Err(err).Msg("Failed to create default config file. You can create it manually by copying config.example.json to config.json")
			return
		}

		os.Exit(1)
	}

	viper.AddConfigPath(".")      // Register config file path
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("json")   // Look for specific type

	err = viper.ReadInConfig()

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read config file")
		return
	}
}
