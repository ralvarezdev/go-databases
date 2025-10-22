package mongodb

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type (
	// Collection represents a MongoDB collection
	Collection struct {
		name    string
		Indexes []*mongo.IndexModel
	}
)

// NewCollection creates a new MongoDB collection
//
// Parameters:
//
//   - name: the name of the collection
//   - indexes: the indexes to create for the collection
//
// Returns:
//
//   - *Collection: the new collection
func NewCollection(
	name string,
	indexes []*mongo.IndexModel,
) *Collection {
	return &Collection{
		name,
		indexes,
	}
}

// CreateCollection creates the collection
//
// Parameters:
//
//   - database: the MongoDB database
func (c Collection) CreateCollection(database *mongo.Database) (
	collection *mongo.Collection, err error,
) {
	// Get the collection
	collection = database.Collection(c.name)

	// Create the indexes
	if createErr := c.createIndexes(collection); createErr != nil {
		return nil, createErr
	}

	return collection, nil
}

// createIndexes creates the indexes for the collection
//
// Parameters:
//
//   - collection: the MongoDB collection
//
// Returns:
//
//   - error: if there was an error creating the indexes
func (c Collection) createIndexes(collection *mongo.Collection) (err error) {
	if c.Indexes == nil {
		return nil
	}

	for _, index := range c.Indexes {
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
	return nil
}
