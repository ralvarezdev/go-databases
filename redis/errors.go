package redis

import (
	"errors"
)

var (
	ErrNilConfig = errors.New("redis config cannot be nil")
)
