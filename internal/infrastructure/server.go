package infrastructure

import (
	"context"
)

// Run starts the application with all its components
func Run() {
	// Setup the Fiber application
	app := setupApp()

	// Create a cancellable context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start email consumer
	startEmailConsumer(ctx)

	// Start HTTP server
	startServer(app)

	// Wait for shutdown signal and handle graceful shutdown
	waitForShutdownSignal(app, cancel)
}
