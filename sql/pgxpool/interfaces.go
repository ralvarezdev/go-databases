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
		IsConnected() bool
		Pool() (*pgxpool.Pool, error)
		Disconnect()
	}

	// Service is the interface for the service
	Service interface {
		Handler
		CreateTransaction(fn TransactionFn) error
		Exec(query *string, params ...any) (*pgconn.CommandTag, error)
		ExecWithCtx(
			ctx context.Context,
			query *string,
			params ...any,
		) (*pgconn.CommandTag, error)
		Query(query *string, params ...any) (pgx.Rows, error)
		QueryWithCtx(
			ctx context.Context,
			query *string,
			params ...any,
		) (pgx.Rows, error)
		QueryRow(query *string, params ...any) (pgx.Row, error)
		QueryRowWithCtx(
			ctx context.Context,
			query *string,
			params ...any,
		) (pgx.Row, error)
		ScanRow(row pgx.Row, destinations ...any) error
		SetStatTicker(
			ctx context.Context,
			duration time.Duration,
			fn func(*pgxpool.Stat),
		)
		ClearStatTicker()
	}
)
