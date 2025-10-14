package pgxpool

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	godatabases "github.com/ralvarezdev/go-databases"
)

type (
	// Config struct
	Config struct {
		DataSourceName              string
		MaxOpenConnections          int
		MaxIdleConnections          int
		ConnectionMaxLifetime       time.Duration
		ConnectionMaxIdleTime       time.Duration
		HealthCheckPeriod           time.Duration
		ConnectionMaxLifetimeJitter time.Duration
	}
)

// NewConfig creates a new pool configuration
//
// Parameters:
//
//   - dataSourceName: the data source name
//   - maxOpenConnections: the maximum number of open connections
//   - maxIdleConnections: the maximum number of idle connections
//   - connectionMaxIdleTime: the maximum idle time for a connection
//   - connectionMaxLifetime: the maximum lifetime for a connection
//   - healthCheckPeriod: the period for health checks
//   - connectionMaxLifetimeJitter: the jitter for the maximum lifetime of a connection
//
// Returns:
//
//   - *Config: the pool configuration
//   - error: if any error occurs
func NewConfig(
	dataSourceName string,
	maxOpenConnections,
	maxIdleConnections int,
	connectionMaxIdleTime,
	connectionMaxLifetime,
	healthCheckPeriod,
	connectionMaxLifetimeJitter time.Duration,
) (*Config, error) {
	// Check if the data source name is empty
	if dataSourceName == "" {
		return nil, godatabases.ErrEmptyDataSourceName
	}

	return &Config{
		DataSourceName:              dataSourceName,
		MaxOpenConnections:          maxOpenConnections,
		MaxIdleConnections:          maxIdleConnections,
		ConnectionMaxIdleTime:       connectionMaxIdleTime,
		ConnectionMaxLifetime:       connectionMaxLifetime,
		HealthCheckPeriod:           healthCheckPeriod,
		ConnectionMaxLifetimeJitter: connectionMaxLifetimeJitter,
	}, nil
}

// ParsedConfig returns the parsed configuration
//
// Returns:
//
//   - *pgxpool.Config: the parsed configuration
func (c *Config) ParsedConfig() (*pgxpool.Config, error) {
	if c == nil {
		return nil, godatabases.ErrNilPoolConfig
	}

	// Create a new parsed configuration
	parsedConfig, err := pgxpool.ParseConfig(c.DataSourceName)
	if err != nil {
		return nil, err
	}

	// Set the configuration values
	parsedConfig.MaxConnIdleTime = c.ConnectionMaxIdleTime
	parsedConfig.MaxConnLifetime = c.ConnectionMaxLifetime
	parsedConfig.MaxConns = int32(c.MaxOpenConnections)
	parsedConfig.MinConns = int32(c.MaxIdleConnections)
	parsedConfig.HealthCheckPeriod = c.HealthCheckPeriod
	parsedConfig.MaxConnLifetimeJitter = c.ConnectionMaxLifetimeJitter
	return parsedConfig, nil
}
