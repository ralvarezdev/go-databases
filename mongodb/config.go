package mongodb

import (
	"time"
)

type (
	// Config interface
	Config interface {
		URI() string
		Timeout() time.Duration
	}

	// ConnConfig struct
	ConnConfig struct {
		uri     string
		timeout time.Duration
	}
)

// NewConnConfig creates a new MongoDB connection configuration
//
// Parameters:
//
// - uri: MongoDB URI
// - timeout: MongoDB connection timeout
//
// Returns:
//
// - *ConnConfig: MongoDB connection configuration
func NewConnConfig(uri string, timeout time.Duration) *ConnConfig {
	return &ConnConfig{
		uri,
		timeout,
	}
}

// URI returns the MongoDB URI
//
// Returns:
//
// - string: MongoDB URI
func (c ConnConfig) URI() string {
	return c.uri
}

// Timeout returns the MongoDB connection timeout
//
// Returns:
//
// - time.Duration: MongoDB connection timeout
func (c ConnConfig) Timeout() time.Duration {
	return c.timeout
}
