# Infrastructure Architecture

This document explains the modular infrastructure architecture of the Go Fiber Template application.

## Overview

The infrastructure package has been refactored into smaller, focused modules to improve maintainability and separation of concerns.

## File Structure

```
internal/infrastructure/
├── server.go      # Main application entry point
├── app.go         # Fiber app setup and configuration
├── consumer.go    # Email consumer management
├── cleanup.go     # Resource cleanup and shutdown
├── signal.go      # Signal handling and graceful shutdown
├── container.go   # Dependency injection container
└── router.go      # Route registration
```

## Module Responsibilities

### 1. `server.go` - Main Entry Point

- **Purpose**: Orchestrates the application startup and shutdown
- **Responsibilities**:
  - Creates the main application context
  - Coordinates startup sequence
  - Handles graceful shutdown coordination

```go
func Run() {
    app := setupApp()
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    startEmailConsumer(ctx)
    startServer(app)
    waitForShutdownSignal(app, ctx, cancel)
}
```

### 2. `app.go` - Application Setup

- **Purpose**: Configures the Fiber web framework
- **Responsibilities**:
  - Creates and configures Fiber app
  - Sets up middleware stack
  - Registers routes

```go
func setupApp() *fiber.App {
    app := fiber.New(config.FiberCfg(cfg))
    setupMiddleware(app)
    setupRoutes(app)
    return app
}
```

### 3. `consumer.go` - Email Consumer Management

- **Purpose**: Manages Kafka email consumer lifecycle
- **Responsibilities**:
  - Starts email consumer in background
  - Handles graceful consumer shutdown
  - Manages consumer topics

```go
func startEmailConsumer(ctx context.Context) {
    emailTopics := []string{"auth.login", "email.notifications"}
    // Start consumer in background
}

func stopEmailConsumer(cancel context.CancelFunc) {
    // Graceful shutdown
}
```

### 4. `cleanup.go` - Resource Cleanup

- **Purpose**: Handles application resource cleanup
- **Responsibilities**:
  - Shuts down HTTP server gracefully
  - Closes database connections
  - Closes Kafka client

```go
func shutdownServer(app *fiber.App) {
    // Graceful server shutdown
}

func cleanupResources() {
    // Cleanup all resources
}
```

### 5. `signal.go` - Signal Handling

- **Purpose**: Manages OS signals and shutdown coordination
- **Responsibilities**:
  - Listens for shutdown signals (SIGINT, SIGTERM)
  - Coordinates graceful shutdown sequence
  - Starts HTTP server in background

```go
func waitForShutdownSignal(app *fiber.App, ctx context.Context, cancel context.CancelFunc) {
    // Signal handling and shutdown coordination
}
```

### 6. `container.go` - Dependency Injection

- **Purpose**: Manages application dependencies
- **Responsibilities**:
  - Initializes all services
  - Manages dependency lifecycle
  - Provides service instances

### 7. `router.go` - Route Registration

- **Purpose**: Registers HTTP routes
- **Responsibilities**:
  - Defines API routes
  - Sets up route handlers
  - Configures route middleware

## Benefits of This Architecture

### 1. **Separation of Concerns**

Each file has a single, well-defined responsibility, making the code easier to understand and maintain.

### 2. **Improved Testability**

Smaller, focused functions are easier to unit test in isolation.

### 3. **Better Maintainability**

Changes to one aspect (e.g., signal handling) don't affect other parts of the codebase.

### 4. **Enhanced Readability**

The main `Run()` function provides a clear overview of the application startup sequence.

### 5. **Easier Debugging**

Issues can be isolated to specific modules, making debugging more efficient.

## Startup Sequence

1. **Application Setup** (`app.go`)

   - Create Fiber app
   - Configure middleware
   - Register routes

2. **Context Creation** (`server.go`)

   - Create cancellable context for graceful shutdown

3. **Email Consumer** (`consumer.go`)

   - Start Kafka consumer in background

4. **HTTP Server** (`signal.go`)

   - Start HTTP server in background

5. **Signal Handling** (`signal.go`)
   - Wait for shutdown signals

## Shutdown Sequence

1. **Signal Received** (`signal.go`)

   - Intercept SIGINT/SIGTERM

2. **Consumer Shutdown** (`consumer.go`)

   - Cancel context
   - Wait for graceful shutdown

3. **Server Shutdown** (`cleanup.go`)

   - Shutdown HTTP server with timeout

4. **Resource Cleanup** (`cleanup.go`)
   - Close database connections
   - Close Kafka client

## Best Practices

### 1. **Context Management**

- Always use cancellable contexts for background operations
- Cancel contexts before shutdown to prevent resource leaks

### 2. **Graceful Shutdown**

- Give background operations time to complete
- Use timeouts to prevent hanging

### 3. **Error Handling**

- Log errors appropriately
- Don't let errors in one component crash the entire application

### 4. **Resource Management**

- Always close resources in the correct order
- Use defer statements for cleanup

## Adding New Components

To add a new component to the infrastructure:

1. **Create a new file** for the component's responsibilities
2. **Add startup function** to initialize the component
3. **Add shutdown function** to clean up the component
4. **Update main Run() function** to include the new component
5. **Update signal handling** to include shutdown coordination

Example:

```go
// new_component.go
func startNewComponent(ctx context.Context) {
    // Initialize component
}

func stopNewComponent() {
    // Cleanup component
}

// Update server.go
func Run() {
    // ... existing code ...
    startNewComponent(ctx)
    // ... existing code ...
}

// Update signal.go
func waitForShutdownSignal(app *fiber.App, ctx context.Context, cancel context.CancelFunc) {
    // ... existing code ...
    stopNewComponent()
    // ... existing code ...
}
```

This modular architecture makes the application more maintainable, testable, and easier to understand while preserving all the original functionality.
