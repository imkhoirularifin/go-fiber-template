package infrastructure

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
)

// startEmailConsumer starts the email consumer in a background goroutine
func startEmailConsumer(ctx context.Context) {
	go func() {
		emailTopics := []string{"auth.login", "email.notifications"}
		if err := emailService.StartEmailConsumer(ctx, emailTopics); err != nil {
			log.Error().Err(err).Msg("Failed to start email consumer")
		}
	}()
}

// stopEmailConsumer gracefully stops the email consumer
func stopEmailConsumer(cancel context.CancelFunc) {
	log.Info().Msg("Stopping email consumer...")

	// Cancel context to stop the email consumer
	cancel()

	// Give the consumer some time to stop gracefully
	time.Sleep(1 * time.Second)

	log.Info().Msg("Email consumer stopped")
}
