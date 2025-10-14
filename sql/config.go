package sql

import (
	"time"

	godatabases "github.com/ralvarezdev/go-databases"
)

type (
	// Config struct
	Config struct {
		DriverName            string
		DataSourceName        string
		MaxOpenConnections    int
		MaxIdleConnections    int
		ConnectionMaxLifetime time.Duration
		ConnectionMaxIdleTime time.Duration
	}
)

// NewConfig creates a new configuration for the connection
//
// Parameters:
//
//   - driverName: the name of the driver
//   - dataSourceName: the data source name
//   - maxOpenConnections: the maximum number of open connections
//   - maxIdleConnections: the maximum number of idle connections
//   - connectionMaxIdleTime: the maximum idle time for a connection
//   - connectionMaxLifetime: the maximum lifetime for a connection
//
// Returns:
//
//   - *Config: the connection configuration
//   - error: if any error occurs
func NewConfig(
	driverName,
	dataSourceName string,
	maxOpenConnections,
	maxIdleConnections int,
	connectionMaxIdleTime,
	connectionMaxLifetime time.Duration,
) (*Config, error) {
	// Check if the driver name ir data source name is empty
	if driverName == "" {
		return nil, godatabases.ErrEmptyDriverName
	}
	if dataSourceName == "" {
		return nil, godatabases.ErrEmptyDataSourceName
	}

	return &Config{
		driverName,
		dataSourceName,
		maxOpenConnections,
		maxIdleConnections,
		connectionMaxIdleTime,
		connectionMaxLifetime,
	}, nil
}
