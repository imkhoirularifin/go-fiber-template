# Database Package

A flexible and configurable database connection manager built on top of GORM, supporting multiple database drivers with configurable logging and connection settings.

## Features

- Multiple database driver support (SQLite, MySQL, PostgreSQL)
- Configurable logging levels
- Connection health monitoring
- Graceful connection handling
- UTC timestamp support
- Default configurations for quick setup

## Usage

### Basic Usage

```go
import "go-fiber-template/lib/database"

func main() {
    // Create database configuration
    config := database.Config{
        Driver:   "mysql",
        Dsn:      "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local",
        LogLevel: "info",
    }

    // Initialize database
    db := database.New(config)

    // Setup connection
    gormDB, err := db.Setup()
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to setup database")
    }
    defer db.Close()

    // Use gormDB for database operations
    // ...
}
```

### Default Configuration

The package provides sensible defaults:

```go
const (
    DefaultDriver   = "sqlite3"
    DefaultDsn      = "file::memory:?cache=shared"
    DefaultLogLevel = "silent"
)
```

## Configuration Options

| Option   | Description      | Default                    | Possible Values            |
| -------- | ---------------- | -------------------------- | -------------------------- |
| Driver   | Database driver  | sqlite3                    | sqlite3, mysql, postgres   |
| Dsn      | Data Source Name | file::memory:?cache=shared | Driver-specific DSN string |
| LogLevel | Logging level    | silent                     | silent, error, warn, info  |

## Database Drivers

### SQLite

```go
config := database.Config{
    Driver: "sqlite3",
    Dsn:    "file::memory:?cache=shared", // In-memory database
}
```

### MySQL

```go
config := database.Config{
    Driver: "mysql",
    Dsn:    "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local",
}
```

### PostgreSQL

```go
config := database.Config{
    Driver: "postgres",
    Dsn:    "host=localhost user=user password=password dbname=dbname port=5432 sslmode=disable",
}
```

## Logging Levels

- `silent`: No logging
- `error`: Only log errors
- `warn`: Log warnings and errors
- `info`: Log all database operations

## Methods

### New(config Config) \*Database

Creates a new database instance with the provided configuration.

### Setup() (\*gorm.DB, error)

Initializes the database connection and returns a GORM DB instance.

### Ping() error

Checks if the database connection is healthy.

### Close() error

Closes the database connection gracefully.

## Best Practices

1. Always call `Close()` when shutting down the application
2. Use appropriate log levels in production (silent/error) and development (info)
3. Configure connection timeouts and pool settings in the DSN when needed
4. Handle database errors appropriately
5. Use UTC timestamps for consistency (automatically handled by the package)

## Error Handling

The package provides clear error messages for:

- Invalid database drivers
- Invalid log levels
- Connection failures
- Configuration issues

## Dependencies

- GORM
- Database-specific drivers (sqlite3, mysql, postgres)
- zerolog for logging
