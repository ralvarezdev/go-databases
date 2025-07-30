package sql

import (
	"time"

	godatabases "github.com/ralvarezdev/go-databases"
)

type (
	// Config interface
	Config interface {
		DriverName() string
		DataSourceName() string
		MaxOpenConnections() int
		MaxIdleConnections() int
		ConnectionMaxLifetime() time.Duration
		ConnectionMaxIdleTime() time.Duration
	}

	// ConnConfig struct
	ConnConfig struct {
		driverName            string
		dataSourceName        string
		maxOpenConnections    int
		maxIdleConnections    int
		connectionMaxLifetime time.Duration
		connectionMaxIdleTime time.Duration
	}
)

// NewConnConfig creates a new configuration for the connection
func NewConnConfig(
	driverName,
	dataSourceName string,
	maxOpenConnections,
	maxIdleConnections int,
	connectionMaxIdleTime,
	connectionMaxLifetime time.Duration,
) (*ConnConfig, error) {
	// Check if the driver name ir data source name is empty
	if driverName == "" {
		return nil, godatabases.ErrEmptyDriverName
	}
	if dataSourceName == "" {
		return nil, godatabases.ErrEmptyDataSourceName
	}

	return &ConnConfig{
		driverName,
		dataSourceName,
		maxOpenConnections,
		maxIdleConnections,
		connectionMaxIdleTime,
		connectionMaxLifetime,
	}, nil
}

// DriverName returns the driver name
func (c *ConnConfig) DriverName() string {
	return c.driverName
}

// DataSourceName returns the data source name
func (c *ConnConfig) DataSourceName() string {
	return c.dataSourceName
}

// MaxOpenConnections returns the maximum open connections
func (c *ConnConfig) MaxOpenConnections() int {
	return c.maxOpenConnections
}

// MaxIdleConnections returns the maximum idle connections
func (c *ConnConfig) MaxIdleConnections() int {
	return c.maxIdleConnections
}

// ConnectionMaxLifetime returns the connection max lifetime
func (c *ConnConfig) ConnectionMaxLifetime() time.Duration {
	return c.connectionMaxLifetime
}

// ConnectionMaxIdleTime returns the connection max idle time
func (c *ConnConfig) ConnectionMaxIdleTime() time.Duration {
	return c.connectionMaxIdleTime
}
