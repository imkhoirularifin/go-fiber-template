package infrastructure

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

type Server struct {
	app *App
}

func NewServer(app *App) *Server {
	return &Server{
		app: app,
	}
}

func (s *Server) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		emailTopics := []string{"auth.login"}
		if err := s.app.GetContainer().EmailService.StartEmailConsumer(ctx, emailTopics); err != nil {
			log.Error().Err(err).Msg("Failed to start email consumer")
		}
	}()

	go func() {
		log.Info().Msgf("Server is running on port %s", s.app.GetContainer().Config.Port)
		if err := s.app.GetServer().Listen(fmt.Sprintf(":%s", s.app.GetContainer().Config.Port)); err != nil {
			log.Error().Err(err).Msg("Failed to start server")
		}
	}()

	s.waitForShutdownSignal(cancel)
}

func (s *Server) waitForShutdownSignal(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Info().Msg("Received shutdown signal")

	s.Shutdown(cancel)
}

func (s *Server) Shutdown(cancel context.CancelFunc) {
	cancel()
	time.Sleep(1 * time.Second)

	log.Info().Msg("Shutting down server...")
	if err := s.app.GetServer().ShutdownWithTimeout(3 * time.Second); err != nil {
		log.Error().Err(err).Msg("Failed to gracefully shutdown server")
	}

	s.cleanupResources()

	log.Info().Msg("Application shutdown complete")
}

func (s *Server) cleanupResources() {
	log.Info().Msg("Running cleanup tasks...")

	sqlDB, err := s.app.GetContainer().DB.DB()
	if err == nil {
		if err := sqlDB.Close(); err != nil {
			log.Error().Err(err).Msg("Failed to close database connection")
		}
	}

	s.app.GetContainer().KafkaClient.Close()

	log.Info().Msg("Cleanup tasks completed")
}
