package sql

import (
	"database/sql"
	"sync"

	godatabases "github.com/ralvarezdev/go-databases"
)

type (
	// DefaultHandler struct
	DefaultHandler struct {
		config *Config
		db     *sql.DB
		mutex  sync.Mutex
	}
)

// NewDefaultHandler creates a new connection
//
// Parameters:
//
//   - config: the configuration for the connection
//
// Returns:
//
//   - *DefaultHandler: the connection handler
func NewDefaultHandler(
	config *Config,
) (*DefaultHandler, error) {
	// Check if the configuration is nil
	if config == nil {
		return nil, godatabases.ErrNilConfig
	}

	return &DefaultHandler{
		config: config,
	}, nil
}

// Connect returns a new SQL connection
//
// Returns:
//
//   - *sql.DB: the SQL connection
//   - error: if any error occurred
func (d *DefaultHandler) Connect() (*sql.DB, error) {
	if d == nil {
		return nil, godatabases.ErrNilConnHandler
	}

	// Lock the mutex to ensure thread safety
	d.mutex.Lock()
	defer d.mutex.Unlock()

	// Open a new connection
	db, err := sql.Open(d.config.DriverName, d.config.DataSourceName)
	if err != nil {
		return nil, err
	}

	// Set the maximum open connections
	db.SetMaxOpenConns(d.config.MaxOpenConnections)

	// Set the maximum idle connections
	db.SetMaxIdleConns(d.config.MaxIdleConnections)

	// Set the connection max lifetime
	db.SetConnMaxLifetime(d.config.ConnectionMaxLifetime)

	// Set the connection max idle time
	db.SetConnMaxIdleTime(d.config.ConnectionMaxIdleTime)

	// Set client
	d.db = db

	return db, nil
}

// DB returns the SQL connection
//
// Returns:
//
//   - *sql.DB: the SQL connection
//   - error: if any error occurred
func (d *DefaultHandler) DB() (*sql.DB, error) {
	if d == nil {
		return nil, godatabases.ErrNilConnHandler
	}

	// Lock the mutex to ensure thread safety
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.db == nil {
		return nil, godatabases.ErrNotConnected
	}

	return d.db, nil
}

// Disconnect closes the SQL connection
//
// Returns:
//
//   - error: if any error occurred
func (d *DefaultHandler) Disconnect() error {
	if d == nil {
		return godatabases.ErrNilConnHandler
	}

	// Lock the mutex to ensure thread safety
	d.mutex.Lock()
	defer d.mutex.Unlock()

	// Check if the connection is established
	if d.db == nil {
		return nil
	}

	// Close the connection
	if err := d.db.Close(); err != nil {
		return err
	}

	// Set the connection to nil
	d.db = nil
	return nil
}
