package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	// ConnHandler interface
	ConnHandler interface {
		Connect() (*mongo.Client, error)
		Client() (*mongo.Client, error)
		Disconnect() error
	}
)
