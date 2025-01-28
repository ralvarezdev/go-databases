package pgxpool

import (
	"github.com/jackc/pgx/v5/pgxpool"
	godatabases "github.com/ralvarezdev/go-databases"
	"time"
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
func (p *PoolConfig) DataSourceName() string {
	return p.dataSourceName
}

// MaxOpenConnections returns the maximum open connections
func (p *PoolConfig) MaxOpenConnections() int {
	return p.maxOpenConnections
}

// MaxIdleConnections returns the maximum idle connections
func (p *PoolConfig) MaxIdleConnections() int {
	return p.maxIdleConnections
}

// ConnectionMaxLifetime returns the connection max lifetime
func (p *PoolConfig) ConnectionMaxLifetime() time.Duration {
	return p.connectionMaxLifetime
}

// ConnectionMaxIdleTime returns the connection max idle time
func (p *PoolConfig) ConnectionMaxIdleTime() time.Duration {
	return p.connectionMaxIdleTime
}

// HealthCheckPeriod returns the health check period
func (p *PoolConfig) HealthCheckPeriod() time.Duration {
	return p.healthCheckPeriod
}

// ConnectionMaxLifetimeJitter returns the connection max lifetime jitter
func (p *PoolConfig) ConnectionMaxLifetimeJitter() time.Duration {
	return p.connectionMaxLifetimeJitter
}

// ParsedConfig returns the parsed configuration
func (p *PoolConfig) ParsedConfig() (*pgxpool.Config, error) {
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
