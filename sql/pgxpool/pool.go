package pgxpool

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	godatabases "github.com/ralvarezdev/go-databases"
)

type (
	// DefaultHandler struct
	DefaultHandler struct {
		config *Config
		pool   *pgxpool.Pool
		mutex  sync.Mutex
	}
)

// NewDefaultHandler creates a new connection
//
// Parameters:
//
//   - config *Config: configuration for the connection
//
// Returns:
//
//   - *DefaultHandler: connection handler
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

// Connect returns a new connection pool
func (d *DefaultHandler) Connect() (*pgxpool.Pool, error) {
	if d == nil {
		return nil, godatabases.ErrNilHandler
	}

	// Lock the mutex to ensure thread safety
	d.mutex.Lock()
	defer d.mutex.Unlock()

	// Check if the connection is already established
	if d.IsConnected() {
		return d.pool, nil
	}

	// Get the parsed configuration
	config, err := d.config.ParsedConfig()
	if err != nil {
		return nil, err
	}

	// Create a new connection pool
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	// Set the connection pool
	d.pool = pool

	return pool, nil
}

// Pool returns the connection pool
//
// Returns:
//
//   - *pgxpool.Pool: the connection pool
//   - error: if the connection is not established
func (d *DefaultHandler) Pool() (*pgxpool.Pool, error) {
	if d == nil {
		return nil, godatabases.ErrNilHandler
	}

	// Lock the mutex to ensure thread safety
	d.mutex.Lock()
	defer d.mutex.Unlock()

	// Check if the connection is established
	if !d.IsConnected() {
		return nil, godatabases.ErrNotConnected
	}

	return d.pool, nil
}

// IsConnected checks if the connection is established
//
// Returns:
//
//   - bool: true if the connection is established, false otherwise
func (d *DefaultHandler) IsConnected() bool {
	if d == nil {
		return false
	}
	return d.pool != nil
}

// Disconnect closes the connection pool
func (d *DefaultHandler) Disconnect() {
	if d == nil {
		return
	}

	// Lock the mutex to ensure thread safety
	d.mutex.Lock()
	defer d.mutex.Unlock()

	// Check if the connection is established
	if !d.IsConnected() {
		return
	}

	// Close the connection pool
	d.pool.Close()

	// Set the pool to nil
	d.pool = nil
}
