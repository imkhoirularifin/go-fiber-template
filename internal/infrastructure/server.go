package infrastructure

import (
	"fmt"
	"go-fiber-template/lib/common"
	"go-fiber-template/lib/config"
	"os"
	"os/signal"
	"syscall"
	"time"

	apitally "github.com/apitally/apitally-go/fiber"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/log"
)

func Run() {
	app := fiber.New(config.FiberCfg(cfg))

	app.Use(fiberi18n.New(config.I18nConfig))
	app.Use(apitally.Middleware(app, config.ApitallyCfg(cfg)))
	app.Use(fiberzerolog.New(config.FiberZerologCfg(cfg)))
	app.Use(recover.New())
	app.Use(cors.New(config.CorsConfig))

	api := app.Group("/api/v1")
	registerRoutes(api)
	app.Use(common.NotFoundHandler)

	go func() {
		log.Info().Msgf("Server is running on port %s", cfg.Port)
		if err := app.Listen(fmt.Sprintf(":%s", cfg.Port)); err != nil {
			log.Error().Err(err).Msg("Failed to start server")
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Info().Msg("Shutting down server")
	err := app.ShutdownWithTimeout(3 * time.Second)
	if err != nil {
		log.Error().Err(err).Msg("Failed to gracefully shutdown server")
	}
	log.Info().Msg("Running cleanup tasks")

	// Your cleanup tasks here
	err = dbInstance.Close()
	if err != nil {
		log.Error().Err(err).Msg("Failed to close database connection")
	}

	log.Info().Msg("Server shutdown complete")
}
