package mongodb

import (
	"errors"
)

const (
	ErrFailedToCreateIndex = "failed to create index '%v': %v"
)

var (
	ErrFailedToCreateDocument = errors.New("failed to create document")
	ErrFailedToStartSession   = errors.New("failed to start session")
	ErrNilClient              = errors.New("mongodb client cannot be nil")
)
