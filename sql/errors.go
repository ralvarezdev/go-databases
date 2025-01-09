package sql

import (
	"errors"
)

var (
	ErrNilQuery = errors.New("sql query cannot be nil")
	ErrNilRow   = errors.New("sql row cannot be nil")
)
