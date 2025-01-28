package sql

import (
	"database/sql"
	godatabases "github.com/ralvarezdev/go-databases"
)

type (
	// ConnHandler interface
	ConnHandler interface {
		Connect() (*sql.DB, error)
		DB() (*sql.DB, error)
		Disconnect()
	}

	// DefaultConnHandler struct
	DefaultConnHandler struct {
		config Config
		db     *sql.DB
	}
)

// NewDefaultConnHandler creates a new connection
func NewDefaultConnHandler(
	driverName, dataSourceName string,
	config Config,
) (*DefaultConnHandler, error) {
	// Check if the configuration is nil
	if config == nil {
		return nil, godatabases.ErrNilConfig
	}

	return &DefaultConnHandler{
		config: config,
		db:     nil,
	}, nil
}

// Connect returns a new SQL connection
func (d *DefaultConnHandler) Connect() (*sql.DB, error) {
	// Open a new connection
	db, err := sql.Open(d.config.DriverName(), d.config.DataSourceName())
	if err != nil {
		return nil, err
	}

	// Set the maximum open connections
	db.SetMaxOpenConns(d.config.MaxOpenConnections())

	// Set the maximum idle connections
	db.SetMaxIdleConns(d.config.MaxIdleConnections())

	// Set the connection max lifetime
	db.SetConnMaxLifetime(d.config.ConnectionMaxLifetime())

	// Set the connection max idle time
	db.SetConnMaxIdleTime(d.config.ConnectionMaxIdleTime())

	// Set client
	d.db = db

	return db, nil
}

// DB returns the SQL connection
func (d *DefaultConnHandler) DB() (*sql.DB, error) {
	if d.db == nil {
		return nil, godatabases.ErrNotConnected
	}

	return d.db, nil
}

// Disconnect closes the SQL connection
func Disconnect(db *sql.DB) error {
	return db.Close()
}
