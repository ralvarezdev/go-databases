package redis

import (
	"github.com/go-redis/redis/v8"
	godatabases "github.com/ralvarezdev/go-databases"
	"golang.org/x/net/context"
)

type (
	// ConnHandler interface
	ConnHandler interface {
		Connect() (*redis.Client, error)
		Client() (*redis.Client, error)
		Disconnect()
	}

	// DefaultConnHandler struct
	DefaultConnHandler struct {
		client        *redis.Client
		clientOptions *redis.Options
	}
)

// NewDefaultConnHandler creates a new connection
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
func (d *DefaultConnHandler) Connect() (*redis.Client, error) {
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
func (d *DefaultConnHandler) Client() (*redis.Client, error) {
	// Check if the connection is established
	if d.client == nil {
		return nil, godatabases.ErrNotConnected
	}

	return d.client, nil
}

// Disconnect closes the Redis client connection
func (d *DefaultConnHandler) Disconnect() {
	// Check if the connection is established
	if d.client == nil {
		return
	}

	// Close the connection
	if err := d.client.Close(); err != nil {
		panic(godatabases.ErrFailedToDisconnect)
	}
}
