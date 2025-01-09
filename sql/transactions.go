package sql

import (
	"database/sql"
	godatabases "github.com/ralvarezdev/go-databases"
)

// CreateTransaction creates a transaction for the database
func CreateTransaction(db *sql.DB, fn func(tx *sql.Tx) error) error {
	// Check if the database is nil
	if db == nil {
		return godatabases.ErrNilDatabase
	}

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Execute the transaction function
	fnErr := fn(tx)
	if fnErr != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}
		return fnErr
	}

	// Commit the transaction
	return tx.Commit()
}
