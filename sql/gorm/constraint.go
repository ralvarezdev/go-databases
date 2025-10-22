package gorm

import (
	"gorm.io/gorm"

	godatabases "github.com/ralvarezdev/go-databases"
)

type (
	// ModelConstraints struct
	ModelConstraints struct {
		model any
		names []string
	}
)

// NewModelConstraints creates a new model constraints
//
// Parameters:
//
//   - model: The model to create the constraints for
//   - names: The names of the constraints to create
//
// Returns:
//
//   - *ModelConstraints: The model constraints
func NewModelConstraints(model any, names ...string) *ModelConstraints {
	return &ModelConstraints{
		model,
		names,
	}
}

// HasConstraint checks if a constraint exists
//
// Parameters:
//
//   - database: The GORM database connection
//   - model: The model to check the constraint
//   - name: The name of the constraint to check
//
// Returns:
//
//   - bool: True if the constraint exists, false otherwise
func HasConstraint(database *gorm.DB, model any, name string) bool {
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
//
// Parameters:
//
//   - database: The GORM database connection
//   - modelConstraints: The model constraints to create
//
// Returns:
//
//   - error: An error if something went wrong, or nil if everything went fine
func CreateModelConstraints(
	database *gorm.DB,
	modelConstraints *ModelConstraints,
) error {
	// Check if the database or the constraint is nil
	if database == nil {
		return godatabases.ErrNilConnection
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
		if err := database.Migrator().CreateConstraint(
			modelConstraints.model,
			name,
		); err != nil {
			return err
		}
	}
	return nil
}

// CreateModelsConstraints creates models constraints
//
// Parameters:
//
//   - database: The GORM database connection
//   - modelsConstraints: The models constraints to create
//
// Returns:
//
//   - error: An error if something went wrong, or nil if everything went fine
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
