package middleware

import (
	"errors"
	"go-fiber-template/lib/config"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Logger(cfg config.AppConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		stop := time.Now()

		level := zerolog.InfoLevel
		latency := stop.Sub(start)
		status := c.Response().StatusCode()
		ip := c.IP()
		method := c.Method()
		path := c.OriginalURL()

		errMsg := ""
		if err != nil {
			errMsg = err.Error()
			level = zerolog.ErrorLevel
			status = fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				status = e.Code
			}
		}

		if status >= 500 {
			level = zerolog.ErrorLevel
		} else if status >= 400 {
			level = zerolog.WarnLevel
		}

		log.WithLevel(level).Int("status", status).Str("latency", latency.String()).Str("ip", ip).Str("method", method).Str("path", path).Str("error", errMsg).Send()

		return err
	}
}
