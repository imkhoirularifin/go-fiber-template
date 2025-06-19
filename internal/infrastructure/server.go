package infrastructure

import (
	"fmt"
	"go-fiber-template/lib/common"
	"go-fiber-template/lib/config"
	"go-fiber-template/lib/middleware"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/log"
)

func Run() {
	app := fiber.New(
		fiber.Config{
			ErrorHandler:          common.ErrorHandler,
			DisableStartupMessage: true,
		},
	)

	app.Use(recover.New())
	app.Use(middleware.Logger(cfg))
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
	err := app.ShutdownWithTimeout(2 * time.Second)
	if err != nil {
		log.Error().Err(err).Msg("Failed to gracefully shutdown server")
	}
	log.Info().Msg("Running cleanup tasks")

	// Your cleanup tasks here

	log.Info().Msg("Server shutdown complete")
}
