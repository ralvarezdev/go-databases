package redis

import (
	"errors"
)

var (
	NilClientError = errors.New("redis client cannot be nil")
	NilConfigError = errors.New("redis config cannot be nil")
)
