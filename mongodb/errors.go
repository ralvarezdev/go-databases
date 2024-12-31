package mongodb

import "errors"

var (
	ErrFailedToCreateDocument = errors.New("failed to create document")
	ErrFailedToStartSession   = errors.New("failed to start session")
	ErrFailedToCreateIndex    = "failed to create index '%v': %v"
	ErrNilConfig              = errors.New("mongodb connection config cannot be nil")
	ErrNilClient              = errors.New("mongodb client cannot be nil")
)