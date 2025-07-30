package mongodb

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

// Collection represents a MongoDB collection
type Collection struct {
	name    string
	Indexes *[]*mongo.IndexModel
}

// NewCollection creates a new MongoDB collection
func NewCollection(
	name string,
	indexes *[]*mongo.IndexModel,
) *Collection {
	return &Collection{
		name,
		indexes,
	}
}

// CreateCollection creates the collection
func (c *Collection) CreateCollection(database *mongo.Database) (
	collection *mongo.Collection, err error,
) {
	// Get the collection
	collection = database.Collection(c.name)

	// Create the indexes
	if err = c.createIndexes(collection); err != nil {
		return nil, err
	}

	return collection, nil
}

// createIndexes creates the indexes for the collection
func (c *Collection) createIndexes(collection *mongo.Collection) (err error) {
	if c.Indexes != nil {
		for _, index := range *c.Indexes {
			// Check if the index is nil
			if index == nil {
				continue
			}

			// Create the index
			_, err = collection.Indexes().CreateOne(
				context.Background(), *index,
			)
			if err != nil {
				return fmt.Errorf(ErrFailedToCreateIndex, *index, err)
			}
		}
	}
	return nil
}
