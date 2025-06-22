package infrastructure

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// shutdownServer gracefully shuts down the Fiber server
func shutdownServer(app *fiber.App) {
	log.Info().Msg("Shutting down server...")

	err := app.ShutdownWithTimeout(3 * time.Second)
	if err != nil {
		log.Error().Err(err).Msg("Failed to gracefully shutdown server")
	}

	log.Info().Msg("Server shutdown complete")
}

// cleanupResources performs cleanup of all application resources
func cleanupResources() {
	log.Info().Msg("Running cleanup tasks...")

	// Close database connection
	err := dbInstance.Close()
	if err != nil {
		log.Error().Err(err).Msg("Failed to close database connection")
	}

	// Close Kafka client
	kafkaClient.Close()

	log.Info().Msg("Cleanup tasks completed")
}
