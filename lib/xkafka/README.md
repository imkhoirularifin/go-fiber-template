# Kafka Go Library

A lightweight Go library for interacting with Apache Kafka using the [Sarama](https://github.com/IBM/sarama) package. This library simplifies producing and consuming messages across microservices, providing a reusable abstraction with support for async producers and consumer groups.

## Features

- **Simplified Interface**: Easy-to-use methods for producing and consuming Kafka messages.
- **Async Producer**: High-throughput message production with error and success handling.
- **Consumer Groups**: Scalable and fault-tolerant message consumption.
- **Graceful Shutdown**: Supports context cancellation and resource cleanup.
- **Configurable**: Flexible configuration for brokers, timeouts, and Sarama settings.

## Installation

1. Initialize a Go module in your project:

   ```bash
   go mod init your-module-name
   ```

2. Add the Kafka library dependency (assuming the library is hosted at `github.com/yourusername/kafka-lib`):

   ```bash
   go get github.com/yourusername/kafka-lib
   ```

3. Install the Sarama dependency:

   ```bash
   go get github.com/IBM/sarama
   ```

4. Ensure your Kafka cluster is running and accessible (e.g., at `localhost:9092`).

## Example Usage

Below is an example of how to use the Kafka library in a Go microservice to produce and consume messages.

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/yourusername/kafka-lib"
	"github.com/IBM/sarama"
)

// myHandler implements the ConsumerHandler interface
type myHandler struct{}

func (h *myHandler) HandleMessage(msg *sarama.ConsumerMessage) error {
	fmt.Printf("Received message: Topic=%s, Key=%s, Value=%s\n", msg.Topic, string(msg.Key), string(msg.Value))
	return nil
}

func main() {
	// Configure the Kafka client
	cfg := kafka.DefaultConfig()
	cfg.Brokers = []string{"kafka-broker:9092"} // Replace with your Kafka broker(s)
	cfg.ConsumerGroup = "my-consumer-group"

	// Create a new Kafka client
	client, err := kafka.NewClient(cfg)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Start consuming messages in a goroutine
	go func() {
		if err := client.Consume(context.Background(), []string{"my-topic"}, &myHandler{}); err != nil {
			log.Fatalf("Consumer error: %v", err)
		}
	}()

	// Produce a message
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Produce(ctx, "my-topic", []byte("key"), []byte("Hello, Kafka!")); err != nil {
		log.Printf("Failed to produce message: %v", err)
	}

	// Keep the service running
	select {}
}
```

## Configuration

The library uses a `Config` struct to customize Kafka settings. Use `kafka.DefaultConfig()` to get a default configuration with sensible defaults, or create a custom configuration:

```go
cfg := &kafka.Config{
	Brokers:         []string{"broker1:9092", "broker2:9092"},
	ProducerTimeout: 10 * time.Second,
	ConsumerTimeout: 10 * time.Second,
	ConsumerGroup:   "my-consumer-group",
	SaramaConfig:    sarama.NewConfig(), // Customize Sarama settings here
}
```

### Key Configuration Options

- **Brokers**: List of Kafka broker addresses.
- **ProducerTimeout**: Timeout for producing messages.
- **ConsumerTimeout**: Timeout for consumer operations.
- **ConsumerGroup**: Name of the consumer group for message consumption.
- **SaramaConfig**: Sarama configuration for advanced settings (e.g., Kafka version, TLS, retries).

## Error Handling

- Producer errors and successes are logged in a background goroutine.
- Consumer errors are logged during message processing.
- Implement the `ConsumerHandler` interface to handle consumed messages and errors as needed.

## Dependencies

- [Sarama](https://github.com/IBM/sarama): Apache Kafka client library for Go.
- Go 1.18 or higher.

## Notes

- Ensure your Kafka cluster is running and accessible at the specified broker addresses.
- Adjust the Kafka version in `SaramaConfig` to match your cluster (default is `V2_8_0_0`).
- For production, configure TLS and authentication in `SaramaConfig` as needed.

## License

MIT License. See [LICENSE](LICENSE) for details.
