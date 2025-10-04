package mongodb

import (
	godatabases "github.com/ralvarezdev/go-databases"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

type (
	// ConnHandler interface
	ConnHandler interface {
		Connect() (*mongo.Client, error)
		Client() (*mongo.Client, error)
		Disconnect()
	}

	// DefaultConnHandler struct
	DefaultConnHandler struct {
		ctx           context.Context
		cancel        context.CancelFunc
		clientOptions *options.ClientOptions
		client        *mongo.Client
	}
)

// NewDefaultConnHandler creates a new connection
//
// Parameters:
//
//   - config: Config interface
//
// Returns:
//
//   - *DefaultConnHandler: DefaultConnHandler struct
//   - error: error if any
func NewDefaultConnHandler(config Config) (
	*DefaultConnHandler,
	error,
) {
	// Check if the config is nil
	if config == nil {
		return nil, godatabases.ErrNilConfig
	}

	// Set client options
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout())
	clientOptions := options.Client().ApplyURI(config.URI())

	return &DefaultConnHandler{
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
func (d *DefaultConnHandler) Connect() (*mongo.Client, error) {
	if d == nil {
		return nil, godatabases.ErrNilConnHandler
	}

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
func (d *DefaultConnHandler) Client() (*mongo.Client, error) {
	if d == nil {
		return nil, godatabases.ErrNilConnHandler
	}

	// Check if the connection is established
	if d.client == nil {
		return nil, godatabases.ErrNotConnected
	}

	return d.client, nil
}

// Disconnect closes the MongoDB client connection
func (d *DefaultConnHandler) Disconnect() {
	if d == nil {
		return
	}

	// Check if the connection is established
	if d.client == nil {
		return
	}

	// Close the connection
	d.cancel()
	if err := d.client.Disconnect(d.ctx); err != nil {
		panic(godatabases.ErrFailedToDisconnect)
	}
}
