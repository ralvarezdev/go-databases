package pgxpool

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	godatabases "github.com/ralvarezdev/go-databases"
)

type (
	// DefaultPoolHandler struct
	DefaultPoolHandler struct {
		config Config
		pool   *pgxpool.Pool
		mutex  sync.Mutex
	}
)

// NewDefaultPoolHandler creates a new connection
func NewDefaultPoolHandler(
	config Config,
) (*DefaultPoolHandler, error) {
	// Check if the configuration is nil
	if config == nil {
		return nil, godatabases.ErrNilConfig
	}

	return &DefaultPoolHandler{
		config: config,
	}, nil
}

// Connect returns a new connection pool
func (d *DefaultPoolHandler) Connect() (*pgxpool.Pool, error) {
	if d == nil {
		return nil, godatabases.ErrNilConnHandler
	}

	// Lock the mutex to ensure thread safety
	d.mutex.Lock()
	defer d.mutex.Unlock()

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
func (d *DefaultPoolHandler) Pool() (*pgxpool.Pool, error) {
	if d == nil {
		return nil, godatabases.ErrNilPoolHandler
	}

	// Lock the mutex to ensure thread safety
	d.mutex.Lock()
	defer d.mutex.Unlock()

	// Check if the connection is established
	if d.pool == nil {
		return nil, godatabases.ErrNotConnected
	}

	return d.pool, nil
}

// Disconnect closes the connection pool
func (d *DefaultPoolHandler) Disconnect() {
	if d == nil {
		return
	}

	// Lock the mutex to ensure thread safety
	d.mutex.Lock()
	defer d.mutex.Unlock()

	// Check if the connection is established
	if d.pool == nil {
		return
	}

	// Close the connection pool
	d.pool.Close()

	// Set the pool to nil
	d.pool = nil
}
