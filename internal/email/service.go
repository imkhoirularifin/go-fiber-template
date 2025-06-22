package email

import (
	"context"
	"fmt"
	"go-fiber-template/internal/domain/interfaces"
	"go-fiber-template/lib/xkafka"
)

type service struct {
	kafkaClient *xkafka.Client
}

func (s *service) Send(config *interfaces.EmailConfig) error {
	// For now just print in the console
	fmt.Printf("Sending email to: %s\n", config.To)
	fmt.Printf("Subject: %s\n", config.Subject)
	fmt.Printf("Body: %s\n", config.Body)

	return nil
}

// StartEmailConsumer starts consuming email messages from Kafka topics
func (s *service) StartEmailConsumer(ctx context.Context, topics []string) error {
	handler := NewEmailConsumerHandler(s)
	return s.kafkaClient.Consume(ctx, topics, handler)
}

func NewService(kafkaClient *xkafka.Client) interfaces.EmailService {
	return &service{
		kafkaClient: kafkaClient,
	}
}
