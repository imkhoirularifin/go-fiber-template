package database

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

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

// Default config values
const (
	DefaultDriver   = "sqlite3"
	DefaultDsn      = "file::memory:?cache=shared"
	DefaultLogLevel = "silent"
)

// Database holds the database connection and configuration.
type Database struct {
	config Config
	db     *gorm.DB
}

// setConfig sets the configuration for the database connection.
func (d *Database) setConfig(config Config) {
	d.config = config

	if d.config.Driver == "" {
		d.config.Driver = DefaultDriver
	}
	if d.config.Dsn == "" {
		d.config.Dsn = DefaultDsn
	}
	if d.config.LogLevel == "" {
		d.config.LogLevel = DefaultLogLevel
	}
}

// New creates a new database struct with the given configuration.
func New(config Config) *Database {
	database := &Database{
		config: Config{},
		db:     nil,
	}
	database.setConfig(config)

	return database
}

// Setup initializes the database connection and returns a gorm.DB instance.
func (d *Database) Setup() (*gorm.DB, error) {
	logger, err := getLogger(d.config.LogLevel)
	if err != nil {
		return nil, err
	}

	dialector, err := getDialector(d.config.Driver, d.config.Dsn)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger:  logger,
		NowFunc: nowFunc,
	})
	if err != nil {
		return nil, err
	}
	d.db = db

	err = d.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Ping checks if the database connection is alive.
func (d *Database) Ping() error {
	if d.db == nil {
		return nil
	}

	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Ping()
	if err != nil {
		return err
	}

	log.Info().Msg("Database connection is healthy")
	return nil
}

// Close closes the database connection.
func (d *Database) Close() error {
	if d.db == nil {
		return nil
	}

	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Close()
	if err != nil {
		return err
	}

	log.Info().Msg("Database connection closed")
	return nil
}
