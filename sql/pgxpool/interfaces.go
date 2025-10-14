package pgxpool

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	// Handler interface
	Handler interface {
		Connect() (*pgxpool.Pool, error)
		Pool() (*pgxpool.Pool, error)
		Disconnect()
	}

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
)
