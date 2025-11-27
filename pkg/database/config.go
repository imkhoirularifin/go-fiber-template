package database

// Config holds the configuration for the database connection.
type Config struct {
	// Driver is the database driver.
	// Possible values: "sqlite3", "mysql", "postgres".
	//
	// Default: sqlite3
	Driver string

	// Dsn is the Data Source Name for the database connection.
	//
	// Default: file::memory:?cache=shared
	Dsn string

	// LogLevel is the logging level for the database connection.
	// Possible values: "silent", "error", "warn", "info".
	//
	// Default: silent
	LogLevel string
}

// DefaultConfig provides default values for the database configuration.
var DefaultConfig = Config{
	Driver:   "sqlite3",
	Dsn:      "file::memory:?cache=shared",
	LogLevel: "silent",
}

// setConfig sets the configuration for the database connection.
func setConfig(config ...Config) Config {
	if len(config) == 0 {
		return DefaultConfig
	}

	// Override default config with provided configs
	cfg := config[0]

	// Set default values if not provided
	if cfg.Driver == "" {
		cfg.Driver = DefaultConfig.Driver
	}
	if cfg.Dsn == "" {
		cfg.Dsn = DefaultConfig.Dsn
	}
	if cfg.LogLevel == "" {
		cfg.LogLevel = DefaultConfig.LogLevel
	}
	return cfg
}
