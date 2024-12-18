package mongodb

import "errors"

var (
	FailedToCreateDocumentError = errors.New("failed to create document")
	FailedToStartSessionError   = errors.New("failed to start session")
	FailedToCreateIndexError    = "failed to create index: %v"
	NilConfigError              = errors.New("mongodb connection config cannot be nil")
	NilClientError              = errors.New("mongodb client cannot be nil")
)
