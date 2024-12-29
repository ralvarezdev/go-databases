package go_databases

import "errors"

var (
	ErrAlreadyConnected   = errors.New("connection to database already established")
	ErrConnectionFailed   = errors.New("failed to connect to database")
	ErrPingFailed         = errors.New("failed to ping database")
	ErrNotConnected       = errors.New("connection to database not established")
	ErrFailedToDisconnect = errors.New("failed to disconnect from database")
)
