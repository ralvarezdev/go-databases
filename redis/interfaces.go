package redis

import (
	"github.com/go-redis/redis/v8"
)

type (
	// Handler interface
	Handler interface {
		Connect() (*redis.Client, error)
		Client() (*redis.Client, error)
		Disconnect() error
	}
)
