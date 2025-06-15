package infrastructure

import (
	"nothing-go/lib/config"
	"nothing-go/lib/xlogger"

	"github.com/rs/zerolog/log"
)

var (
	cfg config.AppConfig
)

func init() {
	var err error

	cfg, err = config.Setup()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	xlogger.Setup()
}
