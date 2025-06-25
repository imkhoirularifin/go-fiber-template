# Three-File Infrastructure Architecture

This document explains the clean, three-file infrastructure architecture of the Go Fiber Template application.

## Overview

The infrastructure package has been organized into three focused files that provide clear separation of concerns while maintaining simplicity.

## File Structure

```
internal/infrastructure/
├── container.go   # Dependency injection and service setup
├── app.go         # Fiber app configuration and middleware
├── server.go      # Server lifecycle and signal handling
└── README.md      # This documentation
```

## Architecture

### 1. `container.go` - Dependency Management

**Purpose**: Manages all application dependencies and service initialization

**Responsibilities**:

- Loads configuration
- Sets up logging and validation
- Initializes database connection
- Sets up Kafka client
- Creates repositories and services
- Provides dependency container

```go
type Container struct {
    Config         config.AppConfig
    DB             *gorm.DB
    KafkaClient    *xkafka.Client
    AuthService    interfaces.AuthService
    UserService    interfaces.UserService
    EmailService   interfaces.EmailService
    ProductService interfaces.ProductService
}

func NewContainer() *Container {
    // Initialize all dependencies
    // Return container with all services
}
```

### 2. `app.go` - Fiber Application Setup

**Purpose**: Configures the Fiber web framework and routes

**Responsibilities**:

- Creates Fiber app instance
- Sets up middleware stack
- Registers all routes
- Provides access to server and container

```go
type App struct {
    container *Container
    server    *fiber.App
}

func NewApp(container *Container) *App {
    // Setup Fiber app with middleware
    // Register all routes
    // Return configured app
}
```

### 3. `server.go` - Server Lifecycle

**Purpose**: Handles server startup, shutdown, and signal management

**Responsibilities**:

- Starts email consumer
- Starts HTTP server
- Handles graceful shutdown
- Manages signal handling
- Performs resource cleanup

```go
type Server struct {
    app *App
}

func (s *Server) Start() {
    // Start email consumer
    // Start HTTP server
    // Wait for shutdown signal
    // Graceful shutdown
}
```

## Key Features

### 1. **Clear Separation of Concerns**

Each file has a single, well-defined responsibility:

- **Container**: "What services do we need?"
- **App**: "How do we configure the web framework?"
- **Server**: "How do we start and stop everything?"

### 2. **Simple Flow**

The application startup follows a clear, linear flow:

```go
func Run() {
    container := NewContainer()    // 1. Setup dependencies
    app := NewApp(container)      // 2. Configure web app
    server := NewServer(app)      // 3. Create server
    server.Start()                // 4. Start everything
}
```

### 3. **Easy to Understand**

- **container.go**: All your services in one place
- **app.go**: All your routes and middleware in one place
- **server.go**: All your startup/shutdown logic in one place

## Benefits

### 1. **Maintainability**

- Easy to find and modify specific functionality
- Clear boundaries between different concerns
- No scattered initialization logic

### 2. **Testability**

- Each component can be tested independently
- Easy to mock dependencies
- Clear interfaces between components

### 3. **Readability**

- Each file is focused and concise
- Clear naming conventions
- Logical organization

### 4. **Extensibility**

- Easy to add new services to container
- Easy to add new routes to app
- Easy to modify server behavior

## Usage

### Starting the Application

```go
// In main.go
func main() {
    infrastructure.Run()
}
```

### Adding New Services

1. **Add to Container** (`container.go`):

```go
type Container struct {
    // ... existing fields
    NewService interfaces.NewService
}

func NewContainer() *Container {
    // ... existing setup
    newService := new.NewService(...)

    return &Container{
        // ... existing fields
        NewService: newService,
    }
}
```

2. **Add Routes** (`app.go`):

```go
func setupRoutes(app *fiber.App, container *Container) {
    api := app.Group("/api/v1")
    // ... existing routes
    new.NewHttpHandler(api.Group("/new"), container.NewService)
}
```

### Customizing Server Behavior

Modify `server.go` to customize:

- Startup sequence
- Shutdown timeout
- Signal handling
- Resource cleanup

## Best Practices

### 1. **Keep Files Focused**

- Each file should have one clear purpose
- Don't mix concerns between files
- Use clear, descriptive names

### 2. **Dependency Flow**

- Container → App → Server
- Dependencies flow in one direction
- No circular dependencies

### 3. **Error Handling**

- Handle errors at the appropriate level
- Log errors meaningfully
- Don't let errors crash the application

### 4. **Resource Management**

- Always clean up resources in server.go
- Use timeouts for graceful shutdown
- Close connections properly

## Example: Adding a Notification Service

### 1. Update Container

```go
// container.go
type Container struct {
    // ... existing fields
    NotificationService interfaces.NotificationService
}

func NewContainer() *Container {
    // ... existing setup
    notificationService := notification.NewService(kafkaClient)

    return &Container{
        // ... existing fields
        NotificationService: notificationService,
    }
}
```

### 2. Update App

```go
// app.go
func setupRoutes(app *fiber.App, container *Container) {
    api := app.Group("/api/v1")
    // ... existing routes
    notification.NewHttpHandler(api.Group("/notifications"), container.NotificationService)
}
```

### 3. Update Server (if needed)

```go
// server.go
func (s *Server) Start() {
    // ... existing startup
    // Add notification consumer if needed
}
```

This three-file architecture provides the perfect balance between simplicity and organization, making your application easy to understand, maintain, and extend.
