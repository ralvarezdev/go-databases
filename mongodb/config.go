package mongodb

import (
	"time"
)

type (
	// Config struct
	Config struct {
		URI     string
		Timeout time.Duration
	}
)

// NewConfig creates a new MongoDB connection configuration
//
// Parameters:
//
// - uri: MongoDB URI
// - timeout: MongoDB connection timeout
//
// Returns:
//
// - *Config: MongoDB connection configuration
func NewConfig(uri string, timeout time.Duration) *Config {
	return &Config{
		uri,
		timeout,
	}
}
