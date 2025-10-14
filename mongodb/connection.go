package mongodb

import (
	"sync"

	godatabases "github.com/ralvarezdev/go-databases"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

type (
	// DefaultHandler struct
	DefaultHandler struct {
		ctx           context.Context
		cancel        context.CancelFunc
		clientOptions *options.ClientOptions
		client        *mongo.Client
		mutex         sync.Mutex
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
//   - *DefaultHandler: DefaultHandler struct
//   - error: error if any
func NewDefaultHandler(config *Config) (
	*DefaultHandler,
	error,
) {
	// Check if the config is nil
	if config == nil {
		return nil, godatabases.ErrNilConfig
	}

	// Set client options
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	clientOptions := options.Client().ApplyURI(config.URI)

	return &DefaultHandler{
		cancel:        cancel,
		ctx:           ctx,
		clientOptions: clientOptions,
	}, nil
}

// Connect returns a new MongoDB client
//
// Returns:
//
//   - *mongo.Client: MongoDB client
//   - error: error if any
func (d *DefaultHandler) Connect() (*mongo.Client, error) {
	if d == nil {
		return nil, godatabases.ErrNilConnHandler
	}

	// Lock the mutex to ensure thread safety
	d.mutex.Lock()
	defer d.mutex.Unlock()

	// Check if the connection is already established
	if d.client != nil {
		return d.client, godatabases.ErrAlreadyConnected
	}

	// Connect to MongoDB
	client, err := mongo.Connect(d.ctx, d.clientOptions)

	// Create MongoDB Connection struct
	if err != nil {
		return nil, godatabases.ErrConnectionFailed
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, godatabases.ErrPingFailed
	}

	// Set client
	d.client = client

	return client, nil
}

// Client returns the MongoDB client
//
// Returns:
//
//   - *mongo.Client: MongoDB client
//   - error: error if any
func (d *DefaultHandler) Client() (*mongo.Client, error) {
	if d == nil {
		return nil, godatabases.ErrNilConnHandler
	}

	// Lock the mutex to ensure thread safety
	d.mutex.Lock()
	defer d.mutex.Unlock()

	// Check if the connection is established
	if d.client == nil {
		return nil, godatabases.ErrNotConnected
	}

	return d.client, nil
}

// Disconnect closes the MongoDB client connection
//
// Returns:
//
//   - error: error if any
func (d *DefaultHandler) Disconnect() error {
	if d == nil {
		return nil
	}

	// Lock the mutex to ensure thread safety
	d.mutex.Lock()
	defer d.mutex.Unlock()

	// Check if the connection is established
	if d.client == nil {
		return nil
	}

	// Close the connection
	d.cancel()
	if err := d.client.Disconnect(d.ctx); err != nil {
		return err
	}

	// Set the client to nil
	d.client = nil
	return nil
}
