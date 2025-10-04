package sql

import (
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
//   - db: The database connection
//   - fn: The function to execute within the transaction
//
// Returns:
//
//   - error: An error if the transaction fails
func CreateTransaction(db *sql.DB, fn TransactionFn) error {
	// Check if the connection is nil
	if db == nil {
		return godatabases.ErrNilConnection
	}

	// Start a transaction
	tx, err := db.Begin()
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
