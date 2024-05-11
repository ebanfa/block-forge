package store

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/edward1christian/block-forge/pkg/application/db"
)

// StoreOptions contains options for configuring a store.
type StoreOptions struct {
	DatabaseFactory db.DatabaseFactory `valid:"-"`
	Name            string             `valid:"required"`
	Path            string             `valid:"required"`
}

// NewStoreOptions creates a new instance of StoreOptions with the provided parameters.
func NewStoreOptions(databaseFactory db.DatabaseFactory, name, path string) *StoreOptions {
	return &StoreOptions{
		DatabaseFactory: databaseFactory,
		Name:            name,
		Path:            path,
	}
}

// Validate checks the validity of the StoreOptions struct.
func (so *StoreOptions) Validate() bool {
	// Using govalidator.ValidateStruct to perform validation
	result, _ := govalidator.ValidateStruct(so)
	return result
}

// Store represents a database store.
type Store interface {
	db.Database

	// Name returns the name of the store.
	Name() string

	// Path returns the path of the store.
	Path() string
}

// StoreImpl is a concrete implementation of the Store interface.
type StoreImpl struct {
	db.Database        // Embedding db.Database to satisfy the Database interface
	name        string // Name of the store
	path        string // Path of the store
}

// NewStoreImpl creates a new instance of StoreImpl with the provided StoreOptions object.
func NewStoreImpl(name, path string, database db.Database) (*StoreImpl, error) {
	if database == nil {
		return nil, errors.New("cannot create Store from nil database")
	}
	return &StoreImpl{
		Database: database,
		name:     name,
		path:     path,
	}, nil
}

// Name returns the name of the store.
func (s *StoreImpl) Name() string {
	return s.name
}

// Path returns the path of the store.
func (s *StoreImpl) Path() string {
	return s.path
}
