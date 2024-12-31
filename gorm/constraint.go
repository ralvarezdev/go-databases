package gorm

import (
	godatabases "github.com/ralvarezdev/go-databases"
	"gorm.io/gorm"
)

type (
	// Constraint struct
	Constraint struct {
		model interface{}
		field string
	}
)

// NewConstraint creates a new constraint
func NewConstraint(model interface{}, field string) *Constraint {
	return &Constraint{
		model: model,
		field: field,
	}
}

// CreateConstraint creates a new constraint
func CreateConstraint(database *gorm.DB, constraint *Constraint) error {
	// Check if the database or the constraint is nil
	if database == nil {
		return godatabases.ErrNilDatabase
	}
	if constraint == nil {
		return ErrNilConstraint
	}

	// Create the constraint
	return database.Migrator().CreateConstraint(
		constraint.model,
		constraint.field,
	)
}

// CreateConstraints creates new constraints
func CreateConstraints(database *gorm.DB, constraints []*Constraint) error {
	for _, constraint := range constraints {
		if err := CreateConstraint(database, constraint); err != nil {
			return err
		}
	}
	return nil
}
