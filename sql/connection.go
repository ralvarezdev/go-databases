package sql

import (
	"database/sql"
	"time"
)

type (
	// Config struct
	Config struct {
		DB                    *sql.DB
		MaxOpenConnections    *int
		MaxIdleConnections    *int
		ConnectionMaxLifetime *time.Duration
	}
)

// NewConfig creates a new configuration
func NewConfig(
	maxOpenConnections *int,
	maxIdleConnections *int,
	connectionMaxLifetime *time.Duration,
) *Config {
	return &Config{
		MaxOpenConnections:    maxOpenConnections,
		MaxIdleConnections:    maxIdleConnections,
		ConnectionMaxLifetime: connectionMaxLifetime,
	}
}

// Connect returns a new SQL connection
func Connect(
	driverName, dataSourceName string,
	config *Config,
) (*sql.DB, error) {
	// Open a new connection
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	// Check if the configuration is nil
	if config == nil {
		return db, nil
	}

	// Set the maximum open connections
	if config.MaxOpenConnections != nil {
		db.SetMaxOpenConns(*config.MaxOpenConnections)
	}

	// Set the maximum idle connections
	if config.MaxIdleConnections != nil {
		db.SetMaxIdleConns(*config.MaxIdleConnections)
	}

	// Set the connection max lifetime
	if config.ConnectionMaxLifetime != nil {
		db.SetConnMaxLifetime(*config.ConnectionMaxLifetime)
	}
	return db, nil
}

// Close closes the SQL connection
func Close(db *sql.DB) error {
	return db.Close()
}

// Connect returns a new SQL connection
func (c *Config) Connect(
	driverName, dataSourceName string,
) (*sql.DB, error) {
	return Connect(driverName, dataSourceName, c)
}

// Close closes the SQL connection
func (c *Config) Close(db *sql.DB) error {
	return Close(db)
}
