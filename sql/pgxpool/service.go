package pgxpool

import (
	"context"
	"time"

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
		CreateTransaction(fn TransactionWithCtxFn) error
		CreateTransactionWithCtx(
			ctx context.Context,
			fn TransactionWithCtxFn,
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
		QueryRow(query *string, params ...interface{}) (pgx.Row, error)
		QueryRowWithCtx(
			ctx context.Context,
			query *string,
			params ...interface{},
		) (pgx.Row, error)
		ScanRow(row pgx.Row, destinations ...interface{}) error
		SetStatTicker(
			duration time.Duration,
			fn func(*pgxpool.Stat),
		)
		ClearStatTicker()
	}

	// DefaultService is the default service struct
	DefaultService struct {
		pool       *pgxpool.Pool
		statTicker *time.Ticker
	}
)

// NewDefaultService creates a new default service
//
// Parameters:
//
//   - pool: the pgxpool.Pool instance
//
// Returns:
//
//   - *DefaultService: the DefaultService instance
//   - error: if any error occurs
func NewDefaultService(pool *pgxpool.Pool) (
	instance *DefaultService,
	err error,
) {
	// Check if the pool is nil
	if pool == nil {
		return nil, godatabases.ErrNilPool
	}

	return &DefaultService{
		pool: pool,
	}, nil
}

// Pool returns the pool
//
// Returns:
//
//   - *pgxpool.Pool: the pgxpool.Pool instance
func (d *DefaultService) Pool() *pgxpool.Pool {
	if d == nil {
		return nil
	}
	return d.pool
}

// Migrate migrates the database
//
// Parameters:
//
//   - queries: the SQL queries to execute
//
// Returns:
//
//   - error: if any error occurs
func (d *DefaultService) Migrate(queries ...string) error {
	if d == nil {
		return godatabases.ErrNilService
	}

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
//
// Parameters:
//
//   - fn: the function to execute within the transaction
//
// Returns:
//
//   - error: if any error occurs
func (d *DefaultService) CreateTransaction(
	fn TransactionWithCtxFn,
) error {
	if d == nil {
		return godatabases.ErrNilService
	}
	return CreateTransaction(d.pool, fn)
}

// CreateTransactionWithCtx creates a transaction for the database with a context
//
// Parameters:
//
//   - ctx: the context to use
//   - fn: the function to execute within the transaction
//
// Returns:
//
//   - error: if any error occurs
func (d *DefaultService) CreateTransactionWithCtx(
	ctx context.Context,
	fn TransactionWithCtxFn,
) error {
	if d == nil {
		return godatabases.ErrNilService
	}
	return CreateTransactionWithCtx(ctx, d.pool, fn)
}

// ExecWithCtx executes a query with parameters and returns the result with a context
//
// Parameters:
//
//   - ctx: the context to use
//   - query: the SQL query to execute
//   - params: the parameters for the SQL query
//
// Returns:
//
//   - *pgconn.CommandTag: the command tag result
//   - error: if any error occurs
func (d *DefaultService) ExecWithCtx(
	ctx context.Context,
	query *string,
	params ...interface{},
) (
	*pgconn.CommandTag,
	error,
) {
	if d == nil {
		return nil, godatabases.ErrNilService
	}

	// Check if the query is nil
	if query == nil {
		return nil, godatabases.ErrNilQuery
	}

	// Run the exec
	commandTag, err := d.pool.Exec(ctx, *query, params...)
	if err != nil {
		return nil, err
	}
	return &commandTag, nil
}

// Exec executes a query with parameters and returns the result
//
// Parameters:
//
//   - query: the SQL query to execute
//   - params: the parameters for the SQL query
//
// Returns:
//
//   - *pgconn.CommandTag: the command tag result
//   - error: if any error occurs
func (d *DefaultService) Exec(query *string, params ...interface{}) (
	*pgconn.CommandTag,
	error,
) {
	if d == nil {
		return nil, godatabases.ErrNilService
	}
	return d.ExecWithCtx(context.Background(), query, params...)
}

// QueryWithCtx runs a query with parameters and returns the result with a context
//
// Parameters:
//
//   - ctx: the context to use
//   - query: the SQL query to execute
//   - params: the parameters for the SQL query
//
// Returns:
//
//   - pgx.Rows: the result rows
func (d *DefaultService) QueryWithCtx(
	ctx context.Context,
	query *string,
	params ...interface{},
) (pgx.Rows, error) {
	if d == nil {
		return nil, godatabases.ErrNilService
	}

	// Check if the query is nil
	if query == nil {
		return nil, godatabases.ErrNilQuery
	}

	// Run the query
	return d.pool.Query(ctx, *query, params...)
}

// Query runs a query with parameters and returns the result
//
// Parameters:
//
//   - query: the SQL query to execute
//   - params: the parameters for the SQL query
//
// Returns:
//
//   - pgx.Rows: the result rows
//   - error: if any error occurs
func (d *DefaultService) Query(
	query *string,
	params ...interface{},
) (pgx.Rows, error) {
	if d == nil {
		return nil, godatabases.ErrNilService
	}
	return d.QueryWithCtx(context.Background(), query, params...)
}

// QueryRowWithCtx runs a query row with parameters and returns the result row with a context
//
// Parameters:
//
//   - ctx: the context to use
//   - query: the SQL query to execute
//   - params: the parameters for the SQL query
//
// Returns:
//
//   - pgx.Row: the result row
//   - error: if any error occurs
func (d *DefaultService) QueryRowWithCtx(
	ctx context.Context,
	query *string,
	params ...interface{},
) (pgx.Row, error) {
	if d == nil {
		return nil, godatabases.ErrNilService
	}

	// Check if the query is nil
	if query == nil {
		return nil, godatabases.ErrNilQuery
	}

	// Run the query row
	return d.pool.QueryRow(ctx, *query, params...), nil
}

// QueryRow runs a query row with parameters and returns the result row
//
// Parameters:
//
//   - query: the SQL query to execute
//   - params: the parameters for the SQL query
//
// Returns:
//
//   - pgx.Row: the result row
//   - error: if any error occurs
func (d *DefaultService) QueryRow(
	query *string,
	params ...interface{},
) (pgx.Row, error) {
	return d.QueryRowWithCtx(context.Background(), query, params...)
}

// ScanRow scans a row
//
// Parameters:
//
// - row: the pgx.Row instance
// - destinations: the destinations to scan the row into
//
// Returns:
//
// - error: if any error occurs
func (d *DefaultService) ScanRow(
	row pgx.Row,
	destinations ...interface{},
) error {
	if d == nil {
		return godatabases.ErrNilService
	}

	// Check if the row is nil
	if row == nil {
		return godatabases.ErrNilRow
	}

	// Scan the row
	return row.Scan(destinations...)
}

// SetStatTicker sets a stat ticker
//
// Parameters:
//
//   - duration: the duration of the ticker
//   - fn: the function to execute on each tick, receiving the pgxpool.Stat
func (d *DefaultService) SetStatTicker(
	duration time.Duration,
	fn func(*pgxpool.Stat),
) {
	if d == nil {
		return
	}

	// Check if the stat ticker is nil
	if d.statTicker != nil {
		d.statTicker.Stop()
	}

	// Set the stat ticker
	d.statTicker = time.NewTicker(duration)

	// Set the stat ticker
	go func() {
		for {
			select {
			case <-d.statTicker.C:
				fn(d.pool.Stat())
			}
		}
	}()
}

// ClearStatTicker clears the stat ticker
func (d *DefaultService) ClearStatTicker() {
	if d == nil {
		return
	}
	if d.statTicker != nil {
		d.statTicker.Stop()
	}
}
