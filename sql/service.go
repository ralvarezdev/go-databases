package sql

import (
	"database/sql"

	godatabases "github.com/ralvarezdev/go-databases"
)

type (
	// Service is the interface for the service
	Service interface {
		DB() *sql.DB
		Migrate(queries ...string) error
		CreateTransaction(fn TransactionFn) error
		Exec(query *string, params ...interface{}) (sql.Result, error)
		QueryRow(query *string, params ...interface{}) *sql.Row
		ScanRow(row *sql.Row, destinations ...interface{}) error
	}

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

// Migrate migrates the database
//
// Parameters:
//
// *queries ...string: the queries to migrate
//
// Returns:
//
// error: if there was an error migrating the database
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
		func(tx *sql.Tx) error {
			// Execute the migration
			for _, query := range queries {
				if _, err := tx.Exec(query); err != nil {
					return err
				}
			}
			return nil
		},
	)
}

// CreateTransaction creates a transaction for the database
func (d *DefaultService) CreateTransaction(fn TransactionFn) error {
	return CreateTransaction(d.db, fn)
}

// Exec executes a query with parameters and returns the result
func (d *DefaultService) Exec(query *string, params ...interface{}) (
	sql.Result,
	error,
) {
	// Check if the query is nil
	if query == nil {
		return nil, godatabases.ErrNilQuery
	}

	// Run the exec
	return d.db.Exec(*query, params...)
}

// QueryRow runs a query row with parameters and returns the result row
func (d *DefaultService) QueryRow(
	query *string,
	params ...interface{},
) *sql.Row {
	// Check if the query is nil
	if query == nil {
		return nil
	}

	// Run the query row
	return d.db.QueryRow(*query, params...)
}

// ScanRow scans a row
func (d *DefaultService) ScanRow(
	row *sql.Row,
	destinations ...interface{},
) error {
	// Check if the row is nil
	if row == nil {
		return godatabases.ErrNilRow
	}

	// Scan the row
	return row.Scan(destinations...)
}
