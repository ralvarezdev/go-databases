package mongodb

import (
	godatabases "github.com/ralvarezdev/go-databases"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"time"
)

type (
	// ConnectionHandler interface
	ConnectionHandler interface {
		Connect() (*mongo.Client, error)
		GetClient() (*mongo.Client, error)
		Disconnect()
	}

	// Config struct
	Config struct {
		Uri     string
		Timeout time.Duration
	}

	// DefaultConnectionHandler struct
	DefaultConnectionHandler struct {
		Ctx           context.Context
		Cancel        context.CancelFunc
		ClientOptions *options.ClientOptions
		Client        *mongo.Client
	}
)

// NewDefaultConnectionHandler creates a new connection
func NewDefaultConnectionHandler(config *Config) (*DefaultConnectionHandler, error) {
	// Check if the config is nil
	if config == nil {
		return nil, NilClientError
	}

	// Set client options
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	clientOptions := options.Client().ApplyURI(config.Uri)

	return &DefaultConnectionHandler{
		Cancel:        cancel,
		Ctx:           ctx,
		ClientOptions: clientOptions,
		Client:        nil,
	}, nil
}

// Connect returns a new MongoDB client
func (d *DefaultConnectionHandler) Connect() (*mongo.Client, error) {
	// Check if the connection is already established
	if d.Client != nil {
		return d.Client, godatabases.AlreadyConnectedError
	}

	// Connect to MongoDB
	client, err := mongo.Connect(d.Ctx, d.ClientOptions)

	// Create MongoDB Connection struct
	if err != nil {
		return nil, godatabases.FailedToConnectError
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, godatabases.FailedToPingError
	}

	// Set client
	d.Client = client

	return client, nil
}

// GetClient returns the MongoDB client
func (d *DefaultConnectionHandler) GetClient() (*mongo.Client, error) {
	// Check if the connection is established
	if d.Client == nil {
		return nil, godatabases.NotConnectedError
	}

	return d.Client, nil
}

// Disconnect closes the MongoDB client connection
func (d *DefaultConnectionHandler) Disconnect() {
	defer func() {
		// Check if the connection is established
		if d.Client == nil {
			return
		}

		// Close the connection
		d.Cancel()
		if err := d.Client.Disconnect(d.Ctx); err != nil {
			panic(godatabases.FailedToDisconnectError)
		}
	}()
}
