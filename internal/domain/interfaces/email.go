package interfaces

import "context"

type EmailConfig struct {
	To      string
	Subject string
	Body    string
}

type EmailService interface {
	Send(config *EmailConfig) error
	StartEmailConsumer(ctx context.Context, topics []string) error
}
