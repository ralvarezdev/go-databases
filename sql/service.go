package sql

import (
	"context"
	"database/sql"

	godatabases "github.com/ralvarezdev/go-databases"
)

type (
	// DefaultService is the default service struct
	DefaultService struct {
		db *sql.DB
	}
)

// NewDefaultService creates a new default service
//
// Parameters:
//
// *db *sql.DB: the database connection
//
// Returns:
//
// *DefaultService: the default service
// error: if there was an error creating the service
func NewDefaultService(db *sql.DB) (
	instance *DefaultService,
	err error,
) {
	// Check if the connection is nil
	if db == nil {
		return nil, godatabases.ErrNilConnection
	}

	return &DefaultService{
		db,
	}, nil
}

// DB returns the database
//
// Returns:
//
// *sql.DB: the database
func (d *DefaultService) DB() *sql.DB {
	if d == nil {
		return nil
	}
	return d.db
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
	return CreateTransaction(ctx, d.db, fn, opts)
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

	// Run the exec
	return d.db.ExecContext(ctx, *query, params...)
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
func (d *DefaultService) QueryRow(
	query *string,
	params ...any,
) *sql.Row {
	if d == nil {
		return nil
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
func (d *DefaultService) QueryRowWithCtx(
	ctx context.Context,
	query *string,
	params ...any,
) *sql.Row {
	if d == nil {
		return nil
	}

	// Check if the query is nil
	if query == nil {
		return nil
	}

	// Run the query row
	return d.db.QueryRowContext(ctx, *query, params...)
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
