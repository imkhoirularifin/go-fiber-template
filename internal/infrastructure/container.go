package infrastructure

import (
	"go-fiber-template/lib/config"
	"go-fiber-template/lib/xlogger"
)

var (
	cfg config.AppConfig
)

func init() {
	cfg = config.Setup()
	xlogger.Setup(cfg)
}
