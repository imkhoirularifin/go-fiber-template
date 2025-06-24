package xkafka

import (
	"go-fiber-template/lib/config"

	"github.com/rs/zerolog/log"
)

func Setup(kafkaCfg config.KafkaConfig) *Client {
	cfg := DefaultConfig()
	cfg.Brokers = kafkaCfg.Brokers
	cfg.ConsumerGroup = kafkaCfg.GroupId

	client, err := NewClient(cfg)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create Kafka client")
	}

	return client
}
