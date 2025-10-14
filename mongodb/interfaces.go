package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	// Handler interface
	Handler interface {
		Connect() (*mongo.Client, error)
		Client() (*mongo.Client, error)
		Disconnect() error
	}
)
