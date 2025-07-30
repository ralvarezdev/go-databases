package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PrepareFindOneOptions prepares the find one options
func PrepareFindOneOptions(
	projection interface{},
	sort interface{},
) *options.FindOneOptions {
	// Create the find options
	findOptions := options.FindOne()

	// Set the projection
	if projection != nil {
		findOptions.SetProjection(projection)
	}

	// Set the sort
	if sort != nil {
		findOptions.SetSort(sort)
	}

	return findOptions
}

// PrepareFindOptions prepares the find options
func PrepareFindOptions(
	projection interface{},
	sort interface{},
	limit int64,
	skip int64,
) *options.FindOptions {
	// Create the find options
	findOptions := options.Find()

	// Set the projection
	if projection != nil {
		findOptions.SetProjection(projection)
	}

	// Set the sort
	if sort != nil {
		findOptions.SetSort(sort)
	}

	// Set the limit
	if limit > 0 {
		findOptions.SetLimit(limit)
	}

	// Set the skip
	if skip > 0 {
		findOptions.SetSkip(skip)
	}

	return findOptions
}

// PrepareUpdateOptions prepares the update options
func PrepareUpdateOptions(upsert bool) *options.UpdateOptions {
	// Create the update options
	updateOptions := options.Update()

	// Set the upsert
	updateOptions.SetUpsert(upsert)

	return updateOptions
}

// PrepareFindOneAndUpdateOptions prepares the find one and update options
func PrepareFindOneAndUpdateOptions(
	projection interface{},
	sort interface{},
	upsert bool,
	returnDocument options.ReturnDocument,
) *options.FindOneAndUpdateOptions {
	// Create the find one and update options
	findOneAndUpdateOptions := options.FindOneAndUpdate()

	// Set the projection
	if projection != nil {
		findOneAndUpdateOptions.SetProjection(projection)
	}

	// Set the sort
	if sort != nil {
		findOneAndUpdateOptions.SetSort(sort)
	}

	// Set the upsert
	findOneAndUpdateOptions.SetUpsert(upsert)

	// Set the return document
	findOneAndUpdateOptions.SetReturnDocument(returnDocument)

	return findOneAndUpdateOptions
}
