package store

import (
	"errors"
	"sync"

	"github.com/edward1christian/block-forge/pkg/application/db"
)

// StoreOptions contains options for configuring a store.
type StoreOptions struct {
	// Initial state of the store
	InitialHeight int64
	Path          string
}

// MultiStore is a multi-store interface that manages multiple key-value stores.
type MultiStore interface {
	Store

	// GetStore returns the store with the given namespace. If the store doesn't exist, it creates and initializes
	// a new store using the provided options.
	GetStore(namespace []byte, options StoreOptions) (Store, error)

	// CreateStore creates and initializes a new store with the given namespace and options. If a store with the same
	// namespace already exists, it returns an error.
	CreateStore(namespace []byte, options StoreOptions) (Store, error)

	// GetStoreCount returns the total number of stores in the multistore.
	GetStoreCount() int
}

// MultiStoreImpl is a concrete implementation of the MultiStore interface.
type MultiStoreImpl struct {
	db.Database // Embedding Database to satisfy the Database interface
	stores      map[string]*StoreImpl
	mutex       sync.RWMutex
}

// NewMultiStore creates a new instance of MultiStoreImpl.
func NewMultiStore(tree db.Database) *MultiStoreImpl {
	return &MultiStoreImpl{
		Database: tree,
		stores:   make(map[string]*StoreImpl),
		mutex:    sync.RWMutex{},
	}
}

// GetStore returns the store with the given namespace.
// If the store doesn't exist, it returns an error.
func (ms *MultiStoreImpl) GetStore(namespace []byte, options StoreOptions) (Store, error) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	ns := string(namespace)
	store, ok := ms.stores[ns]
	if !ok {
		return nil, errors.New("store does not exist")
	}
	return store, nil
}

// CreateStore creates and initializes a new store with the given namespace and options.
// If a store with the same namespace already exists, it returns an error.
func (ms *MultiStoreImpl) CreateStore(namespace []byte, options StoreOptions, database db.Database) (Store, error) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	ns := string(namespace)

	if _, exists := ms.stores[ns]; exists {
		return nil, errors.New("store already exists")
	}

	// Create a new StoreImpl instance with the provided database
	store, err := NewStoreImpl(database)
	if err != nil {
		return nil, err
	}

	ms.stores[ns] = store

	return store, nil
}

// GetStoreCount returns the total number of stores in the multistore.
func (ms *MultiStoreImpl) GetStoreCount() int {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()

	return len(ms.stores)
}
