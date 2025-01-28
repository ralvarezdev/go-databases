package gorm

import (
	godatabases "github.com/ralvarezdev/go-databases"
	"gorm.io/gorm"
)

type (
	// JoinField struct
	JoinField struct {
		model     interface{}
		field     string
		joinTable interface{}
	}
)

// NewJoinField creates a new join field
func NewJoinField(
	model interface{},
	field string,
	joinTable interface{},
) *JoinField {
	return &JoinField{
		model:     model,
		field:     field,
		joinTable: joinTable,
	}
}

// SetupJoinTable setups the join table
func SetupJoinTable(
	database *gorm.DB,
	joinField *JoinField,
) error {
	// Check if the database or the join field is nil
	if database == nil {
		return godatabases.ErrNilConnection
	}
	if joinField == nil {
		return ErrNilJoinField
	}

	return database.SetupJoinTable(
		joinField.model,
		joinField.field,
		joinField.joinTable,
	)
}

// SetupJoinTables setups the join tables
func SetupJoinTables(
	database *gorm.DB,
	joinFields []*JoinField,
) (err error) {
	for _, joinField := range joinFields {
		if err = SetupJoinTable(database, joinField); err != nil {
			return err
		}
	}

	return nil
}
