package config

import (
	"go-fiber-template/lib/common"
	"go-fiber-template/lib/constant"
	"go-fiber-template/lib/utils"
	"time"

	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/language"
)

func FiberCfg(cfg AppConfig) fiber.Config {
	return fiber.Config{
		AppName:               cfg.AppName,
		ErrorHandler:          common.ErrorHandler,
		DisableStartupMessage: true,
	}
}

func CorsCfg(cfg AppConfig) cors.Config {
	return cors.Config{
		AllowOrigins:     cfg.Cors.AllowOrigins,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "*",
		AllowCredentials: cfg.Cors.AllowCredentials,
	}
}

func FiberZerologCfg(cfg AppConfig) fiberzerolog.Config {
	return fiberzerolog.Config{
		Logger:          &log.Logger,
		Fields:          cfg.LogFields,
		WrapHeaders:     true,
		FieldsSnakeCase: true,
		SkipURIs:        []string{"/api/v1/ping"},
		Next:            common.SkipLog,
	}
}

var mapLanguageTags = map[string]language.Tag{
	"id":    language.Indonesian,
	"en-US": language.AmericanEnglish,
}

var I18nConfig = &fiberi18n.Config{
	RootPath:        "./lib/i18n",
	AcceptLanguages: []language.Tag{language.Indonesian, language.AmericanEnglish},
	DefaultLanguage: mapLanguageTags[constant.DefaultLanguage],
}

var CacheCfg = cache.Config{
	Expiration:           1 * time.Minute,
	CacheHeader:          "X-Cache",
	CacheControl:         false,
	KeyGenerator:         utils.CacheKeyWithQueryAndHeaders,
	ExpirationGenerator:  nil,
	StoreResponseHeaders: false,
	Storage:              nil,
	MaxBytes:             0,
	Methods:              []string{fiber.MethodGet, fiber.MethodHead},
}
