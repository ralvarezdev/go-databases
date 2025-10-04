package pgxpool

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	godatabases "github.com/ralvarezdev/go-databases"
)

type (
	// TransactionWithCtxFn is the function type for transactions with context
	TransactionWithCtxFn func(ctx context.Context, tx pgx.Tx) error
)

// CreateTransactionWithCtx creates a transaction for the database with context
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
func CreateTransactionWithCtx(
	ctx context.Context,
	pool *pgxpool.Pool,
	fn TransactionWithCtxFn,
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

// CreateTransaction creates a transaction for the database
//
// Parameters:
//
//   - pool: The pgxpool.Pool instance
//   - fn: The function to execute within the transaction
//
// Returns:
//
//   - error: An error if the transaction fails, otherwise nil
func CreateTransaction(
	pool *pgxpool.Pool,
	fn TransactionWithCtxFn,
) error {
	return CreateTransactionWithCtx(context.Background(), pool, fn)
}
