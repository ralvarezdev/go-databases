package pgxpool

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	godatabases "github.com/ralvarezdev/go-databases"
)

type (
	// Service is the interface for the service
	Service interface {
		Pool() *pgxpool.Pool
		Migrate(queries ...string) error
		CreateTransaction(fn func(ctx context.Context, tx pgx.Tx) error) error
		CreateTransactionWithCtx(
			ctx context.Context,
			fn func(ctx context.Context, tx pgx.Tx) error,
		) error
		Exec(query *string, params ...interface{}) (*pgconn.CommandTag, error)
		ExecWithCtx(
			ctx context.Context,
			query *string,
			params ...interface{},
		) (*pgconn.CommandTag, error)
		Query(query *string, params ...interface{}) (pgx.Rows, error)
		QueryWithCtx(
			ctx context.Context,
			query *string,
			params ...interface{},
		) (pgx.Rows, error)
		QueryRow(query *string, params ...interface{}) pgx.Row
		QueryRowWithCtx(
			ctx context.Context,
			query *string,
			params ...interface{},
		) pgx.Row
		ScanRow(row pgx.Row, destinations ...interface{}) error
	}

	// DefaultService is the default service struct
	DefaultService struct {
		pool *pgxpool.Pool
	}
)

// NewDefaultService creates a new default service
func NewDefaultService(pool *pgxpool.Pool) (
	instance *DefaultService,
	err error,
) {
	// Check if the pool is nil
	if pool == nil {
		return nil, godatabases.ErrNilPool
	}

	return &DefaultService{
		pool,
	}, nil
}

// Pool returns the pool
func (d *DefaultService) Pool() *pgxpool.Pool {
	return d.pool
}

// Migrate migrates the database
func (d *DefaultService) Migrate(queries ...string) error {
	// Check if there are no queries
	if len(queries) == 0 {
		return nil
	}

	// Create a new transaction
	return d.CreateTransaction(
		func(ctx context.Context, tx pgx.Tx) error {
			// Execute the migration
			for _, query := range queries {
				if _, err := tx.Exec(context.Background(), query); err != nil {
					return err
				}
			}
			return nil
		},
	)
}

// CreateTransaction creates a transaction for the database
func (d *DefaultService) CreateTransaction(
	fn func(
		ctx context.Context,
		tx pgx.Tx,
	) error,
) error {
	return CreateTransaction(d.pool, fn)
}

// CreateTransactionWithCtx creates a transaction for the database with a context
func (d *DefaultService) CreateTransactionWithCtx(
	ctx context.Context,
	fn func(ctx context.Context, tx pgx.Tx) error,
) error {
	return CreateTransactionWithCtx(ctx, d.pool, fn)
}

// ExecWithCtx executes a query with parameters and returns the result with a context
func (d *DefaultService) ExecWithCtx(
	ctx context.Context,
	query *string,
	params ...interface{},
) (
	*pgconn.CommandTag,
	error,
) {
	// Check if the query is nil
	if query == nil {
		panic(godatabases.ErrNilQuery)
	}

	// Run the exec
	commandTag, err := d.pool.Exec(ctx, *query, params...)
	if err != nil {
		return nil, err
	}
	return &commandTag, nil
}

// Exec executes a query with parameters and returns the result
func (d *DefaultService) Exec(query *string, params ...interface{}) (
	*pgconn.CommandTag,
	error,
) {
	return d.ExecWithCtx(context.Background(), query, params...)
}

// QueryWithCtx runs a query with parameters and returns the result with a context
func (d *DefaultService) QueryWithCtx(
	ctx context.Context,
	query *string,
	params ...interface{},
) (pgx.Rows, error) {
	// Check if the query is nil
	if query == nil {
		panic(godatabases.ErrNilQuery)
	}

	// Run the query
	return d.pool.Query(ctx, *query, params...)
}

// Query runs a query with parameters and returns the result
func (d *DefaultService) Query(
	query *string,
	params ...interface{},
) (pgx.Rows, error) {
	return d.QueryWithCtx(context.Background(), query, params...)
}

// QueryRowWithCtx runs a query row with parameters and returns the result row with a context
func (d *DefaultService) QueryRowWithCtx(
	ctx context.Context,
	query *string,
	params ...interface{},
) pgx.Row {
	// Check if the query is nil
	if query == nil {
		panic(godatabases.ErrNilQuery)
	}

	// Run the query row
	return d.pool.QueryRow(ctx, *query, params...)
}

// QueryRow runs a query row with parameters and returns the result row
func (d *DefaultService) QueryRow(
	query *string,
	params ...interface{},
) pgx.Row {
	return d.QueryRowWithCtx(context.Background(), query, params...)
}

// ScanRow scans a row
func (d *DefaultService) ScanRow(
	row pgx.Row,
	destinations ...interface{},
) error {
	// Check if the row is nil
	if row == nil {
		return godatabases.ErrNilRow
	}

	// Scan the row
	return row.Scan(destinations...)
}
