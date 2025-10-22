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
		DB() *sql.DB
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
		QueryRow(query *string, params ...any) *sql.Row
		QueryRowWithCtx(
			ctx any,
			query *string,
			params ...any,
		) *sql.Row
		ScanRow(row *sql.Row, destinations ...any) error
	}
)
