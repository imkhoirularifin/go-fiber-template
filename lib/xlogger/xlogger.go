package xlogger

import (
	"go-fiber-template/lib/config"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Setup(cfg config.AppConfig) {
	var l zerolog.Logger
	if cfg.GoEnv == "development" {
		l = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Timestamp().Logger()
		l.Level(zerolog.DebugLevel)
	} else {
		l = zerolog.New(os.Stderr).With().Timestamp().Logger()
	}
	log.Logger = l
}
