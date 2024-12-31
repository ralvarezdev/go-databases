package gorm

import (
	godatabases "github.com/ralvarezdev/go-databases"
	"gorm.io/gorm"
)

type (
	// ForeignKey struct
	ForeignKey struct {
		model     interface{}
		field     string
		joinTable interface{}
	}
)

// NewForeignKey creates a new foreign key
func NewForeignKey(
	model interface{},
	field string,
	joinTable interface{},
) *ForeignKey {
	return &ForeignKey{
		model:     model,
		field:     field,
		joinTable: joinTable,
	}
}

// SetupJoinTable create foreign key
func SetupJoinTable(
	database *gorm.DB,
	foreignKey *ForeignKey,
) error {
	// Check if the database or the foreign key is nil
	if database == nil {
		return godatabases.ErrNilDatabase
	}
	if foreignKey == nil {
		return godatabases.ErrNilForeignKey
	}

	return database.SetupJoinTable(
		foreignKey.model,
		foreignKey.field,
		foreignKey.joinTable,
	)
}

// SetupJoinTables create foreign keys
func SetupJoinTables(
	database *gorm.DB,
	foreignKeys []*ForeignKey,
) error {
	for _, foreignKey := range foreignKeys {
		err := database.SetupJoinTable(
			foreignKey.model,
			foreignKey.field,
			foreignKey.joinTable,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
