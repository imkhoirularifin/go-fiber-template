package xlogger

import (
	"os"

	"github.com/rs/zerolog"
)

func Setup() {
	// pretty logger
	zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
}
