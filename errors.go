package go_databases

import "errors"

var (
	ErrNilConfig           = errors.New("config cannot be nil")
	ErrNilConnection       = errors.New("connection cannot be nil")
	ErrNilPool             = errors.New("pool cannot be nil")
	ErrConnectionFailed    = errors.New("failed to connect to database")
	ErrPingFailed          = errors.New("failed to ping database")
	ErrNotConnected        = errors.New("connection to database not established")
	ErrFailedToDisconnect  = errors.New("failed to disconnect from database")
	ErrEmptyDriverName     = errors.New("driver name cannot be empty")
	ErrEmptyDataSourceName = errors.New("data source name cannot be empty")
	ErrNilQuery            = errors.New("sql query cannot be nil")
	ErrNilRow              = errors.New("sql row cannot be nil")
	ErrNilHandler          = errors.New("connection handler cannot be nil")
	ErrNilService          = errors.New("database service cannot be nil")
)
