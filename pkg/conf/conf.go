package conf

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Debug  bool   `envconfig:"DEBUG"`
	Addr   string `envconfig:"ADDR" default:":8080"`
	Stage  string `envconfig:"STAGE" default:"dev"`
	Branch string `envconfig:"BRANCH"`
}

func (cfg *Config) logging() error {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	if cfg.Stage != "prod" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	return nil
}

func (cfg *Config) parseDbSecrets() error {
	return fmt.Errorf("TODO")
}

func NewConfig() (*Config, error) {
	cfg := new(Config)
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Error().Msgf("%s", err)
		return cfg, err
	}

	err = cfg.parseDbSecrets()
	if err != nil {
		log.Error().Msgf("%s", err)
		return cfg, err
	}

	err = cfg.logging()
	if err != nil {
		log.Error().Msgf("%s", err)
		return cfg, err
	}

	return cfg, nil
}
