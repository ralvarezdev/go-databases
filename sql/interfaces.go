package sql

import (
	"database/sql"
)

type (
	// Handler interface
	Handler interface {
		Connect() (*sql.DB, error)
		DB() (*sql.DB, error)
		Disconnect() error
	}

	// Service is the interface for the service
	Service interface {
		DB() *sql.DB
		Migrate(queries ...string) error
		CreateTransaction(fn TransactionFn) error
		Exec(query *string, params ...interface{}) (sql.Result, error)
		QueryRow(query *string, params ...interface{}) *sql.Row
		ScanRow(row *sql.Row, destinations ...interface{}) error
	}
)
