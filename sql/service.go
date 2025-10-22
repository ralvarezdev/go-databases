package sql

import (
	"context"
	"database/sql"

	godatabases "github.com/ralvarezdev/go-databases"
)

type (
	// DefaultService is the default service struct
	DefaultService struct {
		Handler
	}
)

// NewDefaultService creates a new default service
//
// Parameters:
//
//   - config: the configuration for the connection
//
// Returns:
//
// *DefaultService: the default service
// error: if there was an error creating the service
func NewDefaultService(config *Config) (
	instance *DefaultService,
	err error,
) {
	// Create the handler
	handler, err := NewDefaultHandler(config)
	if err != nil {
		return nil, err
	}

	return &DefaultService{
		Handler: handler,
	}, nil
}

// CreateTransaction creates a transaction for the database
//
// Parameters:
//
// - ctx: The context for the transaction
// - fn: The function to execute within the transaction
// - opts: The transaction options
//
// Returns:
//
// - error: An error if the transaction fails
func (d *DefaultService) CreateTransaction(
	ctx context.Context,
	fn TransactionFn,
	opts *sql.TxOptions,
) error {
	if d == nil {
		return godatabases.ErrNilService
	}

	// Get the database connection
	db, err := d.DB()
	if err != nil {
		return err
	}

	// Create the transaction
	return CreateTransaction(ctx, db, fn, opts)
}

// Exec executes a query with parameters and returns the result
//
// Parameters:
//
//   - query: the query to execute
//
// - params: the parameters for the query
//
// Returns:
//
//   - sql.Result: the result of the execution
func (d *DefaultService) Exec(query *string, params ...any) (
	sql.Result,
	error,
) {
	if d == nil {
		return nil, godatabases.ErrNilService
	}
	return d.ExecWithCtx(context.Background(), query, params...)
}

// ExecWithCtx executes a query with parameters and returns the result with a context
//
// Parameters:
//
// - ctx: the context to use
// - query: the query to execute
// - params: the parameters for the query
//
// Returns:
//
// - sql.Result: the result of the execution
// - error: if any error occurs
func (d *DefaultService) ExecWithCtx(
	ctx context.Context,
	query *string,
	params ...any,
) (
	sql.Result,
	error,
) {
	if d == nil {
		return nil, godatabases.ErrNilService
	}

	// Check if the query is nil
	if query == nil {
		return nil, godatabases.ErrNilQuery
	}

	// Get the database connection
	db, err := d.DB()
	if err != nil {
		return nil, err
	}

	// Run the exec
	return db.ExecContext(ctx, *query, params...)
}

// QueryRow runs a query row with parameters and returns the result row
//
// Parameters:
//
// - query: the query to execute
// - params: the parameters for the query
//
// Returns:
//
// - *sql.Row: the result row
// - error: if any error occurs
func (d *DefaultService) QueryRow(
	query *string,
	params ...any,
) (*sql.Row, error) {
	if d == nil {
		return nil, godatabases.ErrNilService
	}
	return d.QueryRowWithCtx(context.Background(), query, params...)
}

// QueryRowWithCtx runs a query row with parameters and returns the result row with a context
//
// Parameters:
//
// - ctx: the context to use
// - query: the query to execute
// - params: the parameters for the query
//
// Returns:
//
// - *sql.Row: the result row
// - error: if any error occurs
func (d *DefaultService) QueryRowWithCtx(
	ctx context.Context,
	query *string,
	params ...any,
) (*sql.Row, error) {
	if d == nil {
		return nil, godatabases.ErrNilService
	}

	// Check if the query is nil
	if query == nil {
		return nil, godatabases.ErrNilQuery
	}

	// Get the database connection
	db, err := d.DB()
	if err != nil {
		return nil, err
	}

	// Run the query row
	return db.QueryRowContext(ctx, *query, params...), nil
}

// ScanRow scans a row
//
// Parameters:
//
// - row: the row to scan
// - destinations: the destinations to scan into
//
// Returns:
//
// - error: if any error occurs
func (d *DefaultService) ScanRow(
	row *sql.Row,
	destinations ...any,
) error {
	// Check if the row is nil
	if row == nil {
		return godatabases.ErrNilRow
	}

	// Scan the row
	return row.Scan(destinations...)
}
