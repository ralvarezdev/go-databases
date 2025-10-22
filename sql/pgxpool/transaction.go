package pgxpool

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	godatabases "github.com/ralvarezdev/go-databases"
)

type (
	// TransactionFn is the function type for transactions with context
	TransactionFn func(ctx context.Context, tx pgx.Tx) error
)

// CreateTransaction creates a transaction for the database with context
//
// Parameters:
//
//   - ctx: The context for the transaction
//   - pool: The pgxpool.Pool instance
//   - fn: The function to execute within the transaction
//
// Returns:
//
//   - error: An error if the transaction fails, otherwise nil
func CreateTransaction(
	ctx context.Context,
	pool *pgxpool.Pool,
	fn TransactionFn,
) error {
	// Check if the pool is nil
	if pool == nil {
		return godatabases.ErrNilPool
	}

	// Start a transaction
	tx, err := pool.Begin(ctx)
	if err != nil {
		return err
	}

	// Execute the transaction function
	if fnErr := fn(ctx, tx); fnErr != nil {
		err = tx.Rollback(ctx)
		if err != nil {
			return err
		}
		return fnErr
	}

	// Commit the transaction
	return tx.Commit(ctx)
}
