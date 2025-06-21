package email

import (
	"context"
	"encoding/json"
	"go-fiber-template/internal/domain/interfaces"

	"github.com/IBM/sarama"
)

type emailConsumerHandler struct {
	emailService interfaces.EmailService
}

// NewEmailConsumerHandler creates a new email consumer handler
func NewEmailConsumerHandler(emailService interfaces.EmailService) *emailConsumerHandler {
	return &emailConsumerHandler{
		emailService: emailService,
	}
}

func (h *emailConsumerHandler) HandleMessage(msg *sarama.ConsumerMessage) error {
	var emailConfig interfaces.EmailConfig
	if err := json.Unmarshal(msg.Value, &emailConfig); err != nil {
		return err
	}

	if err := h.emailService.Send(&emailConfig); err != nil {
		return err
	}

	return nil
}

// StartEmailConsumer starts consuming email messages from Kafka topics
func (s *service) StartEmailConsumer(ctx context.Context, topics []string) error {
	handler := NewEmailConsumerHandler(s)
	return s.kafkaClient.Consume(ctx, topics, handler)
}
