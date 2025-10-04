package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

// CreateTransactionOptions creates the transaction options
//
// Returns:
//
// *options.TransactionOptions: the transaction options
func CreateTransactionOptions() *options.TransactionOptions {
	wc := writeconcern.Majority()
	return options.Transaction().SetWriteConcern(wc)
}

// CreateSession creates a new session
//
// Parameters:
//
// *client *mongo.Client: the MongoDB client
//
// Returns:
//
// (mongo.Session, error): the MongoDB session and an error if any
func CreateSession(client *mongo.Client) (mongo.Session, error) {
	// Check if the client is nil
	if client == nil {
		return nil, ErrNilClient
	}

	return client.StartSession()
}

// CreateTransaction creates a new transaction
//
// Parameters:
//
// *client *mongo.Client: the MongoDB client
// *queries func(sc mongo.SessionContext) error: the queries to execute in the transaction
//
// Returns:
//
// error: an error if any
func CreateTransaction(
	client *mongo.Client,
	queries func(sc mongo.SessionContext) error,
) error {
	// Create the session
	clientSession, err := CreateSession(client)
	if err != nil {
		return err
	}
	defer clientSession.EndSession(context.Background())

	// Create the transaction options
	transactionOptions := CreateTransactionOptions()

	// Start the transaction
	err = mongo.WithSession(
		context.Background(),
		clientSession,
		func(sc mongo.SessionContext) error {
			if err = clientSession.StartTransaction(transactionOptions); err != nil {
				return err
			}

			// Call the queries
			err = queries(sc)
			if err != nil {
				_ = clientSession.AbortTransaction(sc)
				return err
			}

			if err = clientSession.CommitTransaction(sc); err != nil {
				_ = clientSession.CommitTransaction(sc)
				return err
			}

			return nil
		},
	)
	return err
}
