package config

import (
	"go-fiber-template/lib/common"

	apitally "github.com/apitally/apitally-go/fiber"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog/log"
)

func FiberCfg(cfg AppConfig) fiber.Config {
	return fiber.Config{
		AppName:               cfg.AppName,
		ErrorHandler:          common.ErrorHandler,
		DisableStartupMessage: true,
	}
}

var CorsConfig = cors.Config{
	AllowOrigins:     "http://localhost:3000",
	AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	AllowHeaders:     "*",
	AllowCredentials: true,
}

// map apitally environment
var mapEnv = map[string]string{
	"development": "dev",
	"production":  "prod",
}

func ApitallyCfg(cfg AppConfig) *apitally.Config {
	return &apitally.Config{
		ClientId: cfg.Apitally.ClientId,
		Env:      mapEnv[cfg.GoEnv],
		RequestLoggingConfig: &apitally.RequestLoggingConfig{
			Enabled:            true,
			LogQueryParams:     true,
			LogRequestHeaders:  true,
			LogRequestBody:     true,
			LogResponseHeaders: true,
			LogResponseBody:    true,
			LogPanic:           true,
		},
	}
}

func FiberZerologCfg(cfg AppConfig) fiberzerolog.Config {
	return fiberzerolog.Config{
		Logger:          &log.Logger,
		Fields:          cfg.LogFields,
		WrapHeaders:     true,
		FieldsSnakeCase: true,
	}
}
