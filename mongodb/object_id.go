package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetObjectIdFromString gets the object ID from the string
//
// Parameters:
//
//   - id: the string ID to convert
//
// Returns:
//
//   - *primitive.ObjectID: the object ID
//   - error: if any error occurred
func GetObjectIdFromString(id string) (*primitive.ObjectID, error) {
	// Check if the ID is empty
	if id == "" {
		return nil, mongo.ErrNoDocuments
	}

	// Create the Object ID from the ID
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return &objectId, nil
}
