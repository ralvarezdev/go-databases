package service

import (
	"database/sql"
	godatabases "github.com/ralvarezdev/go-databases"
	godatabasessql "github.com/ralvarezdev/go-databases/sql"
	"strings"
)

type (
	// Service is the interface for the service
	Service interface {
		DB() *sql.DB
		Migrate(queries ...string) error
		RunTransaction(fn func(tx *sql.Tx) error) error
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
func NewDefaultService(db *sql.DB) (
	instance *DefaultService,
	err error,
) {
	// Check if the database is nil
	if db == nil {
		return nil, godatabases.ErrNilDatabase
	}

	// Create the instance
	instance = &DefaultService{
		db: db,
	}

	// Migrate the database
	err = instance.Migrate()
	if err != nil {
		return nil, err
	}

	return instance, nil
}

// DB returns the database
func (d *DefaultService) DB() *sql.DB {
	return d.db
}

// Migrate migrates the database
func (d *DefaultService) Migrate(queries ...string) error {
	// Check if there are no queries
	if len(queries) == 0 {
		return nil
	}

	// Join the migrate queries into a single migration
	migration := strings.Join(
		queries,
		"\n",
	)

	// Execute the migration
	_, err := d.db.Exec(migration)
	return err
}

// RunTransaction runs a transaction
func (d *DefaultService) RunTransaction(fn func(tx *sql.Tx) error) error {
	return godatabasessql.CreateTransaction(d.db, fn)
}

// Exec executes a query with parameters and returns the result
func (d *DefaultService) Exec(query *string, params ...interface{}) (
	sql.Result,
	error,
) {
	// Check if the query is nil
	if query == nil {
		return nil, godatabasessql.ErrNilQuery
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
		return godatabasessql.ErrNilRow
	}

	// Scan the row
	return row.Scan(destinations...)
}
