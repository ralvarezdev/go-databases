package sql

import (
	"context"
	"database/sql"
)

type (
	// Handler interface
	Handler interface {
		Connect() (*sql.DB, error)
		IsConnected() bool
		DB() (*sql.DB, error)
		Disconnect() error
	}

	// Service is the interface for the service
	Service interface {
		Handler
		CreateTransaction(
			ctx context.Context,
			fn TransactionFn,
			opts *sql.TxOptions,
		) error
		Exec(query *string, params ...any) (sql.Result, error)
		ExecWithCtx(
			ctx any,
			query *string,
			params ...any,
		) (sql.Result, error)
		QueryRow(query *string, params ...any) (*sql.Row, error)
		QueryRowWithCtx(
			ctx any,
			query *string,
			params ...any,
		) (*sql.Row, error)
		ScanRow(row *sql.Row, destinations ...any) error
	}
)
