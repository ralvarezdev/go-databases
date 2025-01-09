package sql

import (
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
)

// IsUniqueViolationError checks if the error is a unique violation error
func IsUniqueViolationError(err error) (bool, string) {
	var pqErr *pgconn.PgError
	if errors.As(err, &pqErr) && pqErr.Code == UniqueViolationCode {
		return true, pqErr.ColumnName
	}
	return false, ""
}
