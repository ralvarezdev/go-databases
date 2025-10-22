package sql

import (
	"context"
	"database/sql"

	godatabases "github.com/ralvarezdev/go-databases"
)

type (
	// TransactionFn is the function type for transactions
	TransactionFn func(tx *sql.Tx) error
)

// CreateTransaction creates a transaction for the database
//
// Parameters:
//
//   - ctx: The context for the transaction
//   - db: The database connection
//   - fn: The function to execute within the transaction
//   - opts: The transaction options
//
// Returns:
//
//   - error: An error if the transaction fails
func CreateTransaction(
	ctx context.Context,
	db *sql.DB,
	fn TransactionFn,
	opts *sql.TxOptions,
) error {
	// Check if the connection is nil
	if db == nil {
		return godatabases.ErrNilConnection
	}

	// Start a transaction
	tx, err := db.BeginTx(ctx, opts)
	if err != nil {
		return err
	}

	// Execute the transaction function
	if fnErr := fn(tx); fnErr != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}
		return fnErr
	}

	// Commit the transaction
	return tx.Commit()
}
