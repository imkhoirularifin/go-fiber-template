package infrastructure

import (
	"go-fiber-template/lib/config"
	"go-fiber-template/lib/xlogger"
	"go-fiber-template/lib/xvalidator"
)

var (
	cfg config.AppConfig
	val *xvalidator.Validator
)

func init() {
	cfg = config.Setup()
	xlogger.Setup(cfg)
	setupValidator()
}
