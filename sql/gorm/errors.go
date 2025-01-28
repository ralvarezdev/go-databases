package gorm

import (
	"errors"
)

var (
	ErrNilJoinField        = errors.New("join field cannot be nil")
	ErrNilModelConstraints = errors.New("model constraints cannot be nil")
)
