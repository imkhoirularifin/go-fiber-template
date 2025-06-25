package xkafka

import (
	"go-fiber-template/lib/config"

	"github.com/rs/zerolog/log"
)

func Setup(kafkaCfg config.KafkaConfig) *Client {
	client, err := NewClient(Config{
		Brokers:         kafkaCfg.Brokers,
		ConsumerGroupID: kafkaCfg.GroupId,
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create Kafka client")
	}

	return client
}
