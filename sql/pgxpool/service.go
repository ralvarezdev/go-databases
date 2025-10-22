package pgxpool

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	godatabases "github.com/ralvarezdev/go-databases"
)

type (
	// DefaultService is the default service struct
	DefaultService struct {
		Handler
		statTicker *time.Ticker
		logger     *slog.Logger
	}
)

// NewDefaultService creates a new default service
//
// Parameters:
//
//   - config *Config: configuration for the connection
//   - logger *slog.Logger: the logger instance
//
// Returns:
//
//   - *DefaultService: the DefaultService instance
//   - error: if any error occurs
func NewDefaultService(config *Config, logger *slog.Logger) (
	instance *DefaultService,
	err error,
) {
	// Create the handler
	handler, err := NewDefaultHandler(config)
	if err != nil {
		return nil, err
	}

	if logger != nil {
		logger = logger.With(slog.String("database", "pgxpool"), slog.String("service", "DefaultService"))
	}

	return &DefaultService{
		Handler: handler,
		logger:  logger,
	}, nil
}

// CreateTransaction creates a transaction for the database with a context
//
// Parameters:
//
//   - ctx: the context to use
//   - fn: the function to execute within the transaction
//
// Returns:
//
//   - error: if any error occurs
func (d *DefaultService) CreateTransaction(
	ctx context.Context,
	fn TransactionFn,
) error {
	if d == nil {
		return godatabases.ErrNilService
	}

	// Get the pool
	pool, err := d.Pool()
	if err != nil {
		return err
	}

	// Create the transaction
	return CreateTransaction(ctx, pool, fn)
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
	params ...any,
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

	// Get the pool
	pool, err := d.Pool()
	if err != nil {
		return nil, err
	}

	// Run the exec
	commandTag, err := pool.Exec(ctx, *query, params...)
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
func (d *DefaultService) Exec(query *string, params ...any) (
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
	params ...any,
) (pgx.Rows, error) {
	if d == nil {
		return nil, godatabases.ErrNilService
	}

	// Check if the query is nil
	if query == nil {
		return nil, godatabases.ErrNilQuery
	}

	// Get the pool
	pool, err := d.Pool()
	if err != nil {
		return nil, err
	}

	// Run the query
	return pool.Query(ctx, *query, params...)
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
	params ...any,
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
	params ...any,
) (pgx.Row, error) {
	if d == nil {
		return nil, godatabases.ErrNilService
	}

	// Check if the query is nil
	if query == nil {
		return nil, godatabases.ErrNilQuery
	}

	// Get the pool
	pool, err := d.Pool()
	if err != nil {
		return nil, err
	}

	// Run the query row
	return pool.QueryRow(ctx, *query, params...), nil
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
	params ...any,
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
	destinations ...any,
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
//   - ctx: the context to use
//   - duration: the duration of the ticker
//   - fn: the function to execute on each tick, receiving the pgxpool.Stat
func (d *DefaultService) SetStatTicker(
	ctx context.Context,
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
			case <-ctx.Done():
				return // Exit the goroutine when the context is done
			case <-d.statTicker.C:
				// Get the pool
				pool, err := d.Pool()
				if err != nil {
					d.logger.Error("Failed to get pool for stat ticker", slog.String("error", err.Error()))
					return
				}

				fn(pool.Stat())
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
