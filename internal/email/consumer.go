package email

import (
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

// HandleMessage processes the email message from Kafka
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
