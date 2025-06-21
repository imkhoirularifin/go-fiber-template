package infrastructure

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// waitForShutdownSignal waits for shutdown signals and coordinates graceful shutdown
func waitForShutdownSignal(app *fiber.App, cancel context.CancelFunc) {
	// Setup signal handling
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Wait for shutdown signal
	<-c
	log.Info().Msg("Received shutdown signal")

	// Stop email consumer
	stopEmailConsumer(cancel)

	// Shutdown server
	shutdownServer(app)

	// Cleanup resources
	cleanupResources()

	log.Info().Msg("Application shutdown complete")
}

// startServer starts the HTTP server in a background goroutine
func startServer(app *fiber.App) {
	go func() {
		log.Info().Msgf("Server is running on port %s", cfg.Port)
		if err := app.Listen(fmt.Sprintf(":%s", cfg.Port)); err != nil {
			log.Error().Err(err).Msg("Failed to start server")
		}
	}()
}
