package redis

import (
	"errors"
)

var (
	ErrNilClient = errors.New("redis client cannot be nil")
	ErrNilConfig = errors.New("redis config cannot be nil")
)
