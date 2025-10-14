package sql

import (
	"database/sql"
	"time"
)

type (
	// Config interface
	Config interface {
		DriverName() string
		DataSourceName() string
		MaxOpenConnections() int
		MaxIdleConnections() int
		ConnectionMaxLifetime() time.Duration
		ConnectionMaxIdleTime() time.Duration
	}

	// ConnHandler interface
	ConnHandler interface {
		Connect() (*sql.DB, error)
		DB() (*sql.DB, error)
		Disconnect()
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
