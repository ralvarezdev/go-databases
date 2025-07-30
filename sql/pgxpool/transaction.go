package pgxpool

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	godatabases "github.com/ralvarezdev/go-databases"
)

// CreateTransactionWithCtx creates a transaction for the database with context
func CreateTransactionWithCtx(
	ctx context.Context,
	pool *pgxpool.Pool,
	fn func(ctx context.Context, tx pgx.Tx) error,
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
func CreateTransaction(
	pool *pgxpool.Pool,
	fn func(ctx context.Context, tx pgx.Tx) error,
) error {
	return CreateTransactionWithCtx(context.Background(), pool, fn)
}
