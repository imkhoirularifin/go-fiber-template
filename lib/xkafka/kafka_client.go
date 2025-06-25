package xkafka

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/IBM/sarama"
	"github.com/rs/zerolog/log"
)

// Config holds Kafka client configuration.
type Config struct {
	// Default: []string{"localhost:9092"}
	Brokers []string

	// Default: 10 * time.Second
	ProducerTimeout time.Duration

	// Default: "default-group"
	ConsumerGroupID string

	// Default: 10 * time.Second
	ConsumerTimeout time.Duration

	// SaramaConfig holds the Sarama configuration.
	// Default: sarama.NewConfig() with sensible defaults
	// Set this to nil to use the default Sarama configuration.
	SaramaConfig *sarama.Config
}

// defaultConfig returns a default Kafka configuration.
var DefaultConfig = Config{
	Brokers:         []string{"localhost:9092"},
	ProducerTimeout: 10 * time.Second,
	ConsumerGroupID: "default-group",
	ConsumerTimeout: 10 * time.Second,
	SaramaConfig:    defaultSaramaConfig(),
}

// defaultConfig returns a default sarama kafka configuration.
func defaultSaramaConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Consumer.Return.Errors = true
	config.Version = sarama.V2_8_0_0 // Adjust based on your Kafka version
	return config
}

// setConfig sets the Kafka client configuration.
func setConfig(config ...Config) Config {
	if len(config) == 0 {
		return DefaultConfig
	}

	// Override default config with provided configs
	cfg := config[0]

	// Set default values if not provided
	if len(cfg.Brokers) == 0 {
		cfg.Brokers = DefaultConfig.Brokers
	}
	if cfg.ProducerTimeout == 0 {
		cfg.ProducerTimeout = DefaultConfig.ProducerTimeout
	}
	if cfg.ConsumerGroupID == "" {
		cfg.ConsumerGroupID = DefaultConfig.ConsumerGroupID
	}
	if cfg.ConsumerTimeout == 0 {
		cfg.ConsumerTimeout = DefaultConfig.ConsumerTimeout
	}
	if cfg.SaramaConfig == nil {
		cfg.SaramaConfig = defaultSaramaConfig()
	}

	return cfg
}

// Client wraps Kafka producer and consumer group functionality.
type Client struct {
	config        Config
	producer      sarama.SyncProducer // Use SyncProducer
	consumerGroup sarama.ConsumerGroup
	closed        chan struct{}
	wg            sync.WaitGroup
}

// NewClient creates a new Kafka client.
func NewClient(config ...Config) (*Client, error) {
	cfg := setConfig(config...)

	producer, err := sarama.NewSyncProducer(cfg.Brokers, cfg.SaramaConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create sync producer: %w", err)
	}

	client := &Client{
		config:   cfg,
		producer: producer,
		closed:   make(chan struct{}),
	}

	return client, nil
}

// Produce sends a message to the specified topic.
func (c *Client) Produce(ctx context.Context, topic string, value []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(value),
	}
	select {
	case <-c.closed:
		return fmt.Errorf("client closed")
	default:
	}
	_, _, err := c.producer.SendMessage(msg)
	return err
}

// ConsumerHandler defines the interface for handling consumed messages.
type ConsumerHandler interface {
	HandleMessage(*sarama.ConsumerMessage) error
}

// Consume starts consuming messages from the specified topics using the provided handler.
func (c *Client) Consume(ctx context.Context, topics []string, handler ConsumerHandler) error {
	if c.config.ConsumerGroupID == "" {
		return fmt.Errorf("consumer group not specified")
	}

	consumerGroup, err := sarama.NewConsumerGroup(c.config.Brokers, c.config.ConsumerGroupID, c.config.SaramaConfig)
	if err != nil {
		return fmt.Errorf("failed to create consumer group: %w", err)
	}
	c.consumerGroup = consumerGroup

	consumer := &consumer{
		handler: handler,
		ready:   make(chan bool),
	}

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for {
			select {
			case <-ctx.Done():
				log.Info().Msg("Consumer context cancelled, stopping consumption")
				return
			default:
				if err := c.consumerGroup.Consume(ctx, topics, consumer); err != nil {
					// Don't log errors when context is cancelled (normal shutdown)
					if ctx.Err() == nil {
						log.Error().Err(err).Msg("Consumer error")
					}
				}
				if ctx.Err() != nil {
					return
				}
				consumer.ready = make(chan bool)
			}
		}
	}()

	<-consumer.ready
	return nil
}

// Close shuts down the Kafka client gracefully.
func (c *Client) Close() error {
	close(c.closed)

	if c.producer != nil {
		if err := c.producer.Close(); err != nil {
			log.Error().Err(err).Msg("Failed to close producer")
		}
	}

	if c.consumerGroup != nil {
		if err := c.consumerGroup.Close(); err != nil {
			// Don't log errors when consumer group is already closed (normal during shutdown)
			if err.Error() != "kafka: tried to use consumer group that was closed" {
				log.Error().Err(err).Msg("Failed to close consumer group")
			}
		}
	}

	c.wg.Wait()
	log.Info().Msg("Kafka client closed")
	return nil
}

// consumer implements sarama.ConsumerGroupHandler.
type consumer struct {
	handler ConsumerHandler
	ready   chan bool
}

// Setup is run at the beginning of a new session.
func (c *consumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

// Cleanup is run at the end of a session.
func (c *consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim processes messages from a single partition.
func (c *consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		if err := c.handler.HandleMessage(message); err != nil {
			log.Error().Err(err).Msg("Handler error")
		}
		session.MarkMessage(message, "")
	}
	return nil
}
