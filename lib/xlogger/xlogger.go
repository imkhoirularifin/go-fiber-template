package xlogger

import (
	"go-fiber-template/lib/config"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Setup(cfg config.AppConfig) {
	prodLogger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	devLogger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()

	if cfg.GoEnv == "production" {
		log.Logger = prodLogger
	} else {
		log.Logger = devLogger
	}
}
