package store

import (
	"errors"

	"github.com/edward1christian/block-forge/pkg/application/db"
)

type Store interface {
	db.Database
}

// StoreImpl is a concrete implementation of the Store interface.
type StoreImpl struct {
	db.Database     // Embedding db.Database to satisfy the Database interface
	index       int // Index of the store
}

// NewStoreImpl creates a new instance of StoreImpl with the provided database.
func NewStoreImpl(database db.Database) (*StoreImpl, error) {
	if database == nil {
		return nil, errors.New("database cannot be nil")
	}

	return &StoreImpl{
		Database: database,
	}, nil
}
