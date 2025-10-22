package mongodb

import (
	"context"
	"fmt"

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
// ctx context.Context: the context
// *client *mongo.Client: the MongoDB client
// *queries func(sc mongo.SessionContext) error: the queries to execute in the transaction
//
// Returns:
//
// error: an error if any
func CreateTransaction(
	ctx context.Context,
	client *mongo.Client,
	queries func(sc mongo.SessionContext) error,
) error {
	// Create the session
	clientSession, err := CreateSession(client)
	if err != nil {
		return err
	}
	defer clientSession.EndSession(ctx)

	// Create the transaction options
	transactionOptions := CreateTransactionOptions()

	// Start the transaction
	return mongo.WithSession(
		ctx,
		clientSession,
		//nolint:contextcheck // The session context is correctly propagated as per mongo.WithSession signature
		func(sc mongo.SessionContext) error {
			if txErr := clientSession.StartTransaction(transactionOptions); txErr != nil {
				return txErr
			}

			// Call the queries
			if queriesErr := queries(sc); queriesErr != nil {
				abortErr := clientSession.AbortTransaction(sc)
				return fmt.Errorf("transaction aborted due to queries error: %w, abort error: %w", queriesErr, abortErr)
			}

			if commitErr := clientSession.CommitTransaction(sc); commitErr != nil {
				return fmt.Errorf("transaction commit error: %w", commitErr)
			}

			return nil
		},
	)
}
