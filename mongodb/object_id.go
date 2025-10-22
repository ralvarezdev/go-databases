package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetObjectIDFromString gets the object ID from the string
//
// Parameters:
//
//   - id: the string ID to convert
//
// Returns:
//
//   - *primitive.ObjectID: the object ID
//   - error: if any error occurred
func GetObjectIDFromString(id string) (*primitive.ObjectID, error) {
	// Check if the ID is empty
	if id == "" {
		return nil, mongo.ErrNoDocuments
	}

	// Create the Object ID from the ID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return &objectID, nil
}
