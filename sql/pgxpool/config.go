package pgxpool

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	godatabases "github.com/ralvarezdev/go-databases"
)

type (
	// Config interface
	Config interface {
		DataSourceName() string
		MaxOpenConnections() int
		MaxIdleConnections() int
		ConnectionMaxLifetime() time.Duration
		ConnectionMaxIdleTime() time.Duration
		HealthCheckPeriod() time.Duration
		ConnectionMaxLifetimeJitter() time.Duration
		ParsedConfig() (*pgxpool.Config, error)
	}

	// PoolConfig struct
	PoolConfig struct {
		dataSourceName              string
		maxOpenConnections          int
		maxIdleConnections          int
		connectionMaxLifetime       time.Duration
		connectionMaxIdleTime       time.Duration
		healthCheckPeriod           time.Duration
		connectionMaxLifetimeJitter time.Duration
		parsedConfig                *pgxpool.Config
	}
)

// NewPoolConfig creates a new pool configuration
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
//   - *PoolConfig: the pool configuration
//   - error: if any error occurs
func NewPoolConfig(
	dataSourceName string,
	maxOpenConnections,
	maxIdleConnections int,
	connectionMaxIdleTime,
	connectionMaxLifetime,
	healthCheckPeriod,
	connectionMaxLifetimeJitter time.Duration,
) (*PoolConfig, error) {
	// Check if the data source name is empty
	if dataSourceName == "" {
		return nil, godatabases.ErrEmptyDataSourceName
	}

	return &PoolConfig{
		dataSourceName:              dataSourceName,
		maxOpenConnections:          maxOpenConnections,
		maxIdleConnections:          maxIdleConnections,
		connectionMaxIdleTime:       connectionMaxIdleTime,
		connectionMaxLifetime:       connectionMaxLifetime,
		healthCheckPeriod:           healthCheckPeriod,
		connectionMaxLifetimeJitter: connectionMaxLifetimeJitter,
	}, nil
}

// DataSourceName returns the data source name
//
// Returns:
//
//   - string: the data source name
func (p *PoolConfig) DataSourceName() string {
	if p == nil {
		return ""
	}
	return p.dataSourceName
}

// MaxOpenConnections returns the maximum open connections
//
// Returns:
//
//   - int: the maximum open connections
func (p *PoolConfig) MaxOpenConnections() int {
	if p == nil {
		return 0
	}
	return p.maxOpenConnections
}

// MaxIdleConnections returns the maximum idle connections
//
// Returns:
//
//   - int: the maximum idle connections
func (p *PoolConfig) MaxIdleConnections() int {
	if p == nil {
		return 0
	}
	return p.maxIdleConnections
}

// ConnectionMaxLifetime returns the connection max lifetime
//
// Returns:
//
//   - time.Duration: the connection max lifetime
func (p *PoolConfig) ConnectionMaxLifetime() time.Duration {
	if p == nil {
		return 0
	}
	return p.connectionMaxLifetime
}

// ConnectionMaxIdleTime returns the connection max idle time
//
// Returns:
//
//   - time.Duration: the connection max idle time
func (p *PoolConfig) ConnectionMaxIdleTime() time.Duration {
	if p == nil {
		return 0
	}
	return p.connectionMaxIdleTime
}

// HealthCheckPeriod returns the health check period
//
// Returns:
//
//   - time.Duration: the health check period
func (p *PoolConfig) HealthCheckPeriod() time.Duration {
	if p == nil {
		return 0
	}
	return p.healthCheckPeriod
}

// ConnectionMaxLifetimeJitter returns the connection max lifetime jitter
//
// Returns:
//
//   - time.Duration: the connection max lifetime jitter
func (p *PoolConfig) ConnectionMaxLifetimeJitter() time.Duration {
	if p == nil {
		return 0
	}
	return p.connectionMaxLifetimeJitter
}

// ParsedConfig returns the parsed configuration
//
// Returns:
//
//   - *pgxpool.Config: the parsed configuration
func (p *PoolConfig) ParsedConfig() (*pgxpool.Config, error) {
	if p == nil {
		return nil, godatabases.ErrNilPoolConfig
	}

	// Check if the parsed configuration is nil
	if p.parsedConfig != nil {
		return p.parsedConfig, nil
	}

	// Create a new parsed configuration
	parsedConfig, err := pgxpool.ParseConfig(p.DataSourceName())
	if err != nil {
		return nil, err
	}

	// Set the configuration values
	parsedConfig.MaxConnIdleTime = p.ConnectionMaxIdleTime()
	parsedConfig.MaxConnLifetime = p.ConnectionMaxLifetime()
	parsedConfig.MaxConns = int32(p.MaxOpenConnections())
	parsedConfig.MinConns = int32(p.MaxIdleConnections())
	parsedConfig.HealthCheckPeriod = p.HealthCheckPeriod()
	parsedConfig.MaxConnLifetimeJitter = p.ConnectionMaxLifetimeJitter()

	// Set the parsed configuration
	p.parsedConfig = parsedConfig

	return p.parsedConfig, nil
}
