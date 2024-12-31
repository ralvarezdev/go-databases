package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FieldIndex struct {
	Name  string
	Order Order
}

// NewFieldIndex creates a new field index
func NewFieldIndex(name string, order Order) *FieldIndex {
	return &FieldIndex{Name: name, Order: order}
}

// NewUniqueIndex creates a new unique field index model
func NewUniqueIndex(fieldIndex FieldIndex, unique bool) *mongo.IndexModel {
	return &mongo.IndexModel{
		Keys:    bson.D{{fieldIndex.Name, fieldIndex.Order.OrderInt()}},
		Options: options.Index().SetUnique(unique),
	}
}

// NewTTLIndex creates a new TTL index model
func NewTTLIndex(fieldName string, expireAfterSeconds int32) *mongo.IndexModel {
	return &mongo.IndexModel{
		Keys:    bson.D{{fieldName, 1}},
		Options: options.Index().SetExpireAfterSeconds(expireAfterSeconds),
	}
}

// NewCompoundFieldIndex creates a new compound field index model
func NewCompoundFieldIndex(
	fieldIndexes []*FieldIndex, unique bool,
) *mongo.IndexModel {
	// Create the keys
	keys := bson.D{}
	for _, fieldIndex := range fieldIndexes {
		keys = append(
			keys,
			bson.E{Key: fieldIndex.Name, Value: fieldIndex.Order.OrderInt()},
		)
	}

	// Create the index model
	return &mongo.IndexModel{
		Keys:    keys,
		Options: options.Index().SetUnique(unique),
	}
}