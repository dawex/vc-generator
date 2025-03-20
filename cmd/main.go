package main

import (
	"os"

	"github.com/dawex/vc-generator/internal/app"
	"github.com/dawex/vc-generator/internal/common/config"
	"github.com/dawex/vc-generator/internal/common/server"
	"github.com/rs/zerolog/log"
)

const app_name = "vc-generator"
const app_config_filename = "config"

func main() {
	configFileName := app_config_filename
	if len(os.Args) < 2 {
		log.Info().Msgf("Config file name not specified, using default value %s (usage: go run main.go <configFileName>)", app_config_filename)
	} else {
		configFileName = os.Args[1]
	}

	// Init config
	app_config := config.LoadConfig(configFileName)

	// Run server
	_, handler := app.NewApp(app_name, app_config)
	if err := server.Run(app_name, app_config, handler); err != nil {
		log.Fatal().Err(err).Msg("Error starting server")
	}
}
