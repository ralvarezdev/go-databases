package gorm

import (
	"errors"
)

var (
	ErrNilJoinField  = errors.New("join field cannot be nil")
	ErrNilConstraint = errors.New("constraint cannot be nil")
)
