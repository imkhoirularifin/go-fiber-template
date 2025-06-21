package xkafka

import "github.com/rs/zerolog/log"

func Setup() *Client {
	cfg := DefaultConfig()
	cfg.Brokers = []string{"localhost:9092"}
	cfg.ConsumerGroup = "go-fiber-template"

	client, err := NewClient(cfg)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create Kafka client")
	}

	return client
}
