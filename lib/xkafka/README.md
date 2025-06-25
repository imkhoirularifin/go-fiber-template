# XKafka Library

A simple and robust Kafka client library for Go applications built on top of Sarama. This library provides easy-to-use producer and consumer functionality with sensible defaults and graceful shutdown handling.

## Features

- ✅ **Simple API**: Easy-to-use producer and consumer interfaces
- ✅ **Sensible Defaults**: Pre-configured with production-ready settings
- ✅ **Graceful Shutdown**: Proper cleanup of resources
- ✅ **Error Handling**: Comprehensive error handling and logging
- ✅ **Context Support**: Full context support for cancellation
- ✅ **Consumer Groups**: Built-in consumer group support
- ✅ **Sync Producer**: Reliable message delivery with sync producer

## Installation

```bash
go get github.com/IBM/sarama
```

## Quick Start

### Basic Setup

```go
package main

import (
    "context"
    "log"
    "time"

    "your-project/lib/xkafka"
)

func main() {
    // Create Kafka client with default configuration
    client, err := xkafka.NewClient()
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    // Your application logic here
}
```

### Custom Configuration

```go
// Custom configuration
config := xkafka.Config{
    Brokers:         []string{"kafka1:9092", "kafka2:9092"},
    ProducerTimeout: 5 * time.Second,
    ConsumerGroupID: "my-app-group",
    ConsumerTimeout: 5 * time.Second,
}

client, err := xkafka.NewClient(config)
if err != nil {
    log.Fatal(err)
}
defer client.Close()
```

## Producer Examples

### Simple Message Production

```go
package main

import (
    "context"
    "log"

    "your-project/lib/xkafka"
)

func main() {
    client, err := xkafka.NewClient()
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    ctx := context.Background()

    // Send a simple message
    message := []byte("Hello, Kafka!")
    err = client.Produce(ctx, "my-topic", message)
    if err != nil {
        log.Printf("Failed to produce message: %v", err)
    }

    log.Println("Message sent successfully!")
}
```

### Batch Message Production

```go
func sendBatchMessages(client *xkafka.Client) {
    ctx := context.Background()
    messages := []string{
        "First message",
        "Second message",
        "Third message",
    }

    for _, msg := range messages {
        err := client.Produce(ctx, "batch-topic", []byte(msg))
        if err != nil {
            log.Printf("Failed to send message '%s': %v", msg, err)
            continue
        }
        log.Printf("Sent: %s", msg)
    }
}
```

## Consumer Examples

### Simple Message Consumer

```go
package main

import (
    "context"
    "log"

    "github.com/IBM/sarama"
    "your-project/lib/xkafka"
)

// MessageHandler implements the ConsumerHandler interface
type MessageHandler struct{}

func (h *MessageHandler) HandleMessage(msg *sarama.ConsumerMessage) error {
    log.Printf("Received message from topic %s: %s", msg.Topic, string(msg.Value))
    return nil
}

func main() {
    client, err := xkafka.NewClient()
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    ctx := context.Background()
    handler := &MessageHandler{}
    topics := []string{"my-topic"}

    // Start consuming messages
    err = client.Consume(ctx, topics, handler)
    if err != nil {
        log.Fatal(err)
    }

    // Keep the application running
    select {}
}
```

### Advanced Consumer with Error Handling

```go
type AdvancedHandler struct {
    processedCount int
}

func (h *AdvancedHandler) HandleMessage(msg *sarama.ConsumerMessage) error {
    h.processedCount++

    log.Printf("Processing message #%d from topic %s", h.processedCount, msg.Topic)

    // Simulate some processing
    if err := h.processMessage(msg); err != nil {
        log.Printf("Failed to process message: %v", err)
        return err // Return error to indicate processing failure
    }

    log.Printf("Successfully processed message #%d", h.processedCount)
    return nil
}

func (h *AdvancedHandler) processMessage(msg *sarama.ConsumerMessage) error {
    // Your message processing logic here
    // For example: parse JSON, save to database, etc.
    return nil
}
```

## Configuration Options

### Default Configuration

```go
var DefaultConfig = Config{
    Brokers:         []string{"localhost:9092"},
    ProducerTimeout: 10 * time.Second,
    ConsumerGroupID: "default-group",
    ConsumerTimeout: 10 * time.Second,
    SaramaConfig:    defaultSaramaConfig(),
}
```

### Custom Sarama Configuration

```go
import "github.com/IBM/sarama"

// Custom Sarama configuration
saramaConfig := sarama.NewConfig()
saramaConfig.Producer.Return.Successes = true
saramaConfig.Producer.Return.Errors = true
saramaConfig.Consumer.Return.Errors = true
saramaConfig.Version = sarama.V2_8_0_0

config := xkafka.Config{
    Brokers:      []string{"kafka:9092"},
    SaramaConfig: saramaConfig,
}

client, err := xkafka.NewClient(config)
```

## Complete Example: Producer and Consumer

```go
package main

import (
    "context"
    "log"
    "time"

    "github.com/IBM/sarama"
    "your-project/lib/xkafka"
)

type MyHandler struct{}

func (h *MyHandler) HandleMessage(msg *sarama.ConsumerMessage) error {
    log.Printf("Consumer received: %s", string(msg.Value))
    return nil
}

func main() {
    // Create client
    client, err := xkafka.NewClient()
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Start consumer in background
    go func() {
        handler := &MyHandler{}
        topics := []string{"test-topic"}

        if err := client.Consume(ctx, topics, handler); err != nil {
            log.Printf("Consumer error: %v", err)
        }
    }()

    // Give consumer time to start
    time.Sleep(2 * time.Second)

    // Send some messages
    messages := []string{"Hello", "World", "Kafka"}
    for _, msg := range messages {
        err := client.Produce(ctx, "test-topic", []byte(msg))
        if err != nil {
            log.Printf("Failed to produce message: %v", err)
        } else {
            log.Printf("Produced: %s", msg)
        }
        time.Sleep(500 * time.Millisecond)
    }

    // Keep running for a while to see messages
    time.Sleep(5 * time.Second)
    log.Println("Shutting down...")
}
```

## Error Handling

### Producer Errors

```go
err := client.Produce(ctx, "topic", []byte("message"))
if err != nil {
    switch {
    case err.Error() == "client closed":
        log.Println("Client was closed, cannot send message")
    default:
        log.Printf("Producer error: %v", err)
    }
}
```

### Consumer Errors

```go
func (h *MyHandler) HandleMessage(msg *sarama.ConsumerMessage) error {
    // Your processing logic
    if err := h.processMessage(msg); err != nil {
        // Log the error but don't panic
        log.Printf("Failed to process message: %v", err)

        // You can choose to return the error to indicate processing failure
        // or return nil to continue processing other messages
        return err
    }
    return nil
}
```

## Graceful Shutdown

The library handles graceful shutdown automatically:

```go
func main() {
    client, err := xkafka.NewClient()
    if err != nil {
        log.Fatal(err)
    }

    // The client will be properly closed when the application exits
    defer client.Close()

    // Your application logic
    // When the application receives a shutdown signal,
    // the client.Close() will be called automatically
}
```

## Best Practices

### 1. **Always Close the Client**

```go
client, err := xkafka.NewClient()
if err != nil {
    log.Fatal(err)
}
defer client.Close() // Always close the client
```

### 2. **Use Context for Cancellation**

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// Use ctx for both producer and consumer operations
err := client.Produce(ctx, "topic", []byte("message"))
```

### 3. **Handle Consumer Errors Gracefully**

```go
func (h *MyHandler) HandleMessage(msg *sarama.ConsumerMessage) error {
    // Don't panic on errors, log them and continue
    if err := h.processMessage(msg); err != nil {
        log.Printf("Processing error: %v", err)
        // Return error only if you want to stop processing
        return nil // Continue processing other messages
    }
    return nil
}
```

### 4. **Use Appropriate Timeouts**

```go
config := xkafka.Config{
    ProducerTimeout: 5 * time.Second,  // Shorter for interactive apps
    ConsumerTimeout: 10 * time.Second, // Longer for batch processing
}
```

## Troubleshooting

### Common Issues

1. **Connection Refused**

   - Check if Kafka brokers are running
   - Verify broker addresses in configuration

2. **Consumer Group Errors**

   - Ensure unique consumer group IDs for different applications
   - Check if consumer group exists

3. **Message Delivery Failures**
   - Check topic exists and has proper permissions
   - Verify producer configuration

### Debug Mode

Enable Sarama debug logging:

```go
import "github.com/IBM/sarama"

sarama.Logger = log.New(os.Stdout, "[Sarama] ", log.LstdFlags)
```

## API Reference

### Types

- `Config`: Kafka client configuration
- `Client`: Main Kafka client
- `ConsumerHandler`: Interface for message handlers

### Methods

- `NewClient(config ...Config) (*Client, error)`: Create new client
- `Produce(ctx context.Context, topic string, value []byte) error`: Send message
- `Consume(ctx context.Context, topics []string, handler ConsumerHandler) error`: Start consuming
- `Close() error`: Gracefully close client

This library provides a simple yet powerful interface for working with Kafka in Go applications. It handles the complexity of Sarama while providing a clean, easy-to-use API.
