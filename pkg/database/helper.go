package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Helper function to return the current time in UTC
func nowFunc() time.Time {
	return time.Now().UTC()
}

// getLogger returns a logger interface based on the provided log level.
// It supports silent, error, warn, and info log levels.
func getLogger(logLevel string) (logger.Interface, error) {
	switch logLevel {
	case "silent":
		return logger.Default.LogMode(logger.Silent), nil
	case "error":
		return logger.Default.LogMode(logger.Error), nil
	case "warn":
		return logger.Default.LogMode(logger.Warn), nil
	case "info":
		return logger.Default.LogMode(logger.Info), nil
	default:
		return logger.Default.LogMode(logger.Silent), fmt.Errorf("invalid log level: %s, set to default level: silent", logLevel)
	}
}

// getDialector returns a Gorm Dialector based on the provided driver and DSN.
// It supports SQLite, MySQL, and PostgreSQL.
func getDialector(driver, dsn string) (gorm.Dialector, error) {
	switch driver {
	case "sqlite3":
		return sqlite.Open(dsn), nil
	case "mysql":
		return mysql.Open(dsn), nil
	case "postgres":
		return postgres.Open(dsn), nil
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", driver)
	}
}
