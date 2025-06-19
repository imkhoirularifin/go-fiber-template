package infrastructure

import (
	"go-fiber-template/lib/config"
	"go-fiber-template/lib/xlogger"
	"go-fiber-template/lib/xvalidator"
)

var (
	cfg config.AppConfig
)

func init() {
	cfg = config.Setup()
	xlogger.Setup(cfg)
	xvalidator.Setup()
}
