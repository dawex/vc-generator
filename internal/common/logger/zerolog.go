package logger

import (
	"os"

	"github.com/dawex/vc-generator/internal/common/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

// New : initialize logger
func NewZerolog(app_config config.Config) {
	// Initialize Zerolog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	if app_config.Server.Env == "DEV" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		log.Info().Msg("Env is DEV, pretty logging enabled")
	}

	// Set Log level from config file
	switch logsLevel := app_config.Logs.Level; logsLevel {
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Info().Msg("Log level is INFO")
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Info().Msg("Log level is DEBUG")
	case "TRACE":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
		log.Info().Msg("Log level is TRACE")
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Info().Msg("Config empty - Log level default is INFO")
	}
}
