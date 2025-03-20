package config

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Config structure to hold the configuration values
type Config struct {
	Server   Server   `mapstructure:"server" validate:"required"`
	Db       Db       `mapstructure:"db" validate:"required"`
	Security Security `mapstructure:"security" validate:"required"`
	Issuer   Issuer   `mapstructure:"issuer" validate:"required"`
	Logs     Logs     `mapstructure:"logs" validate:"required"`
}

type Server struct {
	Env  string `mapstructure:"env" validate:"required,oneof=DEV PROD"`
	Addr string `mapstructure:"addr" validate:"required"`
}

type Security struct {
	Seed string `mapstructure:"seed" validate:"required"`
}

type Issuer struct {
	ID   string `mapstructure:"id" validate:"required"`
	Name string `mapstructure:"name" validate:"required"`
}

type Db struct {
	Addr string `mapstructure:"addr" validate:"required"`
}

type Logs struct {
	Level string `mapstructure:"level" validate:"required,oneof=INFO DEBUG TRACE"`
}

func LoadConfig(configName string) Config {
	// Set the name of the config file (without extension)
	viper.SetConfigName(configName)
	// Set the path to look for the config file
	viper.AddConfigPath("./config/")       // config folder from current directory
	viper.AddConfigPath("./config/test/")  // config folder from current directory
	viper.AddConfigPath("../config/")      // config folder from current directory
	viper.AddConfigPath("../config/test/") // config folder from current directory
	// Set the file type to YAML
	viper.SetConfigType("yaml")

	// Automatic environment variables
	viper.AutomaticEnv()

	// Replace `.` in env variables with `_` (ex: env var DATABASE_HOST is replaced by DATABASE.HOST to match struct config )
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Attempt to read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		// If no configuration file is found, fallback to env variables
		if viper_err, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warn().Msgf("%s, falling back to environment variables", viper_err.Error())
		} else {
			// Handle other errors
			log.Fatal().Err(err).Msg("Error reading config file")
		}
	} else {
		log.Info().Msg("Config file loaded successfully")
	}

	// Unmarshal the configuration into the struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal().Err(err).Msg("Unable to decode config")
	}

	// Validate the config
	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		log.Fatal().Err(err).Msg("Unable to validate config")
	}

	return config
}
