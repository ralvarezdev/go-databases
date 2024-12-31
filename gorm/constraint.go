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

// HasConstraint checks if a constraint exists
func HasConstraint(database *gorm.DB, constraint *Constraint) bool {
	// Check if the database or the constraint is nil
	if database == nil || constraint == nil {
		return false
	}

	// Check if the constraint exists
	return database.Migrator().HasConstraint(
		constraint.model,
		constraint.field,
	)
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

	// Check if the constraint exists
	if HasConstraint(database, constraint) {
		return nil
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
