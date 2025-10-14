package redis

import (
	"sync"

	"github.com/go-redis/redis/v8"
	godatabases "github.com/ralvarezdev/go-databases"
	"golang.org/x/net/context"
)

type (
	// DefaultConnHandler struct
	DefaultConnHandler struct {
		client        *redis.Client
		clientOptions *redis.Options
		mutex         sync.Mutex
	}
)

// NewDefaultConnHandler creates a new connection
//
// Parameters:
//
// - config Config: configuration for the connection
//
// Returns:
//
// - *DefaultConnHandler: connection handler
// - error: error if the config is nil
func NewDefaultConnHandler(config Config) (
	*DefaultConnHandler,
	error,
) {
	// Check if the config is nil
	if config == nil {
		return nil, godatabases.ErrNilConfig
	}

	// Define the Redis options
	clientOptions := &redis.Options{
		Addr:     config.URI(),
		Password: config.Password(),
		DB:       config.Database(),
	}

	return &DefaultConnHandler{
		clientOptions: clientOptions,
	}, nil
}

// Connect returns a new Redis client
//
// Returns:
//
// - *redis.Client: Redis client
// - error: error if the connection fails or is already established
func (d *DefaultConnHandler) Connect() (*redis.Client, error) {
	if d == nil {
		return nil, godatabases.ErrNilConnHandler
	}

	// Mutex lock to ensure thread safety
	d.mutex.Lock()
	defer d.mutex.Unlock()

	// Check if the connection is already established
	if d.client != nil {
		return d.client, godatabases.ErrAlreadyConnected
	}

	// Create a new Redis client
	client := redis.NewClient(d.clientOptions)

	// Ping the Redis server to check the connection
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, godatabases.ErrPingFailed
	}

	// Set client
	d.client = client

	return client, nil
}

// Client returns the Redis client
//
// Returns:
//
// - *redis.Client: Redis client
// - error: error if the connection is not established
func (d *DefaultConnHandler) Client() (*redis.Client, error) {
	if d == nil {
		return nil, godatabases.ErrNilConnHandler
	}

	// Mutex lock to ensure thread safety
	d.mutex.Lock()
	defer d.mutex.Unlock()

	// Check if the connection is established
	if d.client == nil {
		return nil, godatabases.ErrNotConnected
	}

	return d.client, nil
}

// Disconnect closes the Redis client connection
//
// Returns:
//
// - error: error if the disconnection fails
func (d *DefaultConnHandler) Disconnect() error {
	if d == nil {
		return godatabases.ErrNilConnHandler
	}

	// Mutex lock to ensure thread safety
	d.mutex.Lock()
	defer d.mutex.Unlock()

	// Check if the connection is established
	if d.client == nil {
		return nil
	}

	// Close the connection
	if err := d.client.Close(); err != nil {
		return godatabases.ErrFailedToDisconnect
	}

	// Set the client to nil
	d.client = nil
	return nil
}
