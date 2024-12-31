package gorm

import (
	godatabases "github.com/ralvarezdev/go-databases"
	"gorm.io/gorm"
)

type (
	// ModelConstraints struct
	ModelConstraints struct {
		model interface{}
		names []string
	}
)

// NewModelConstraints creates a new model constraints
func NewModelConstraints(model interface{}, names ...string) *ModelConstraints {
	return &ModelConstraints{
		model: model,
		names: names,
	}
}

// HasConstraint checks if a constraint exists
func HasConstraint(database *gorm.DB, model interface{}, name string) bool {
	// Check if the database or the model is nil
	if database == nil || model == nil {
		return false
	}

	// Check if the constraint exists
	return database.Migrator().HasConstraint(
		model,
		name,
	)
}

// CreateModelConstraints creates model constraints
func CreateModelConstraints(
	database *gorm.DB,
	modelConstraints *ModelConstraints,
) (err error) {
	// Check if the database or the constraint is nil
	if database == nil {
		return godatabases.ErrNilDatabase
	}
	if modelConstraints == nil {
		return ErrNilModelConstraints
	}

	for _, name := range modelConstraints.names {
		// Check if the constraint exists
		if HasConstraint(database, modelConstraints.model, name) {
			return nil
		}

		// Create the constraint
		if err = database.Migrator().CreateConstraint(
			modelConstraints.model,
			name,
		); err != nil {
			return err
		}
	}
	return nil
}

// CreateModelsConstraints creates models constraints
func CreateModelsConstraints(
	database *gorm.DB,
	modelsConstraints []*ModelConstraints,
) error {
	for _, modelConstraint := range modelsConstraints {
		if err := CreateModelConstraints(
			database,
			modelConstraint,
		); err != nil {
			return err
		}
	}
	return nil
}
