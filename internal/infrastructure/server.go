package infrastructure

import (
	"fmt"
	"go-fiber-template/lib/common"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	app.Use(logger.New())
	app.Use(recover.New())
	api := app.Group("/api/v1")
	registerRoutes(api)

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
	app.Shutdown()
	log.Info().Msg("Running cleanup tasks")

	// Your cleanup tasks here
	// db.Close()
	// redisConn.Close()

	log.Info().Msg("Server shutdown complete")
}
