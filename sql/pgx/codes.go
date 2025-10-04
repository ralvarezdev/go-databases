package pgx

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

// IsUniqueViolationError checks if the error is a unique violation error
//
// Parameters:
//
//   - err: the error to check
//
// Returns:
//
//   - bool: true if the error is a unique violation error, false otherwise
//   - string: the name of the constraint that was violated, empty string if not a unique violation error
func IsUniqueViolationError(err error) (bool, string) {
	var pqErr *pgconn.PgError
	if errors.As(err, &pqErr) && pqErr.Code == UniqueViolationCode {
		return true, pqErr.ConstraintName
	}
	return false, ""
}
