package pgxpool

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	godatabases "github.com/ralvarezdev/go-databases"
)

type (
	// PoolHandler interface
	PoolHandler interface {
		Connect() (*pgxpool.Pool, error)
		Pool() (*pgxpool.Pool, error)
		Disconnect()
	}

	// DefaultPoolHandler struct
	DefaultPoolHandler struct {
		config Config
		pool   *pgxpool.Pool
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
func (d *DefaultPoolHandler) Pool() (*pgxpool.Pool, error) {
	// Check if the connection is established
	if d.pool == nil {
		return nil, godatabases.ErrNotConnected
	}

	return d.pool, nil
}

// Disconnect closes the connection pool
func (d *DefaultPoolHandler) Disconnect() {
	// Check if the connection is established
	if d.pool == nil {
		return
	}

	// Close the connection pool
	d.pool.Close()
}
