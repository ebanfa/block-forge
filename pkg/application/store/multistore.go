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
	Name          string
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
	dbFactory db.DatabaseFactory // Factory for creating databases
	stores    map[string]Store
	mutex     sync.RWMutex
	database  db.Database
}

// NewMultiStore creates a new instance of MultiStoreImpl with the provided database factory.
func NewMultiStore(dbFactory db.DatabaseFactory) (MultiStore, error) {
	database, err := dbFactory.CreateDatabase("default", "default/path") // Assuming default values
	if err != nil {
		return nil, err
	}
	return &MultiStoreImpl{
		database:  database,
		dbFactory: dbFactory,
		stores:    make(map[string]Store),
		mutex:     sync.RWMutex{},
	}, nil
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

// GetStoreCount returns the total number of stores in the multistore.
func (ms *MultiStoreImpl) GetStoreCount() int {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()

	return len(ms.stores)
}

// CreateStore creates and initializes a new store with the given namespace and options.
// If a store with the same namespace already exists, it returns an error.
func (ms *MultiStoreImpl) CreateStore(namespace []byte, options StoreOptions) (Store, error) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()

	ns := string(namespace)
	// Ensure store does not exist
	if _, exists := ms.stores[ns]; exists {
		return nil, errors.New("store already exists")
	}

	// Create the underlying database using the factory
	database, err := ms.dbFactory.CreateDatabase(options.Name, options.Path)
	if err != nil {
		return nil, err
	}

	// Create a new StoreImpl instance with the provided database
	store, err := NewStoreImpl(database)
	if err != nil {
		return nil, err
	}

	ms.stores[ns] = store

	return store, nil
}

// Get retrieves the value associated with the given key from the database.
func (ms *MultiStoreImpl) Get(key []byte) ([]byte, error) {
	return ms.database.Get(key)
}

// Has checks if a key exists in the database.
func (ms *MultiStoreImpl) Has(key []byte) (bool, error) {
	return ms.database.Has(key)
}

// Iterate iterates over all key-value pairs in the database and calls the given function for each pair.
// Iteration stops if the function returns true.
func (ms *MultiStoreImpl) Iterate(fn func(key, value []byte) bool) error {
	return ms.database.Iterate(fn)
}

// IterateRange iterates over key-value pairs with keys in the specified range
// and calls the given function for each pair. Iteration stops if the function returns true.
func (ms *MultiStoreImpl) IterateRange(start, end []byte, ascending bool, fn func(key, value []byte) bool) error {
	return ms.database.IterateRange(start, end, ascending, fn)
}

// Hash returns the hash of the database.
func (ms *MultiStoreImpl) Hash() []byte {
	return ms.database.Hash()
}

// Version returns the version of the database.
func (ms *MultiStoreImpl) Version() int64 {
	return ms.database.Version()
}

// String returns a string representation of the database.
func (ms *MultiStoreImpl) String() (string, error) {
	return ms.database.String()
}

// WorkingVersion returns the current working version of the database.
func (ms *MultiStoreImpl) WorkingVersion() int64 {
	return ms.database.WorkingVersion()
}

// WorkingHash returns the hash of the current working version of the database.
func (ms *MultiStoreImpl) WorkingHash() []byte {
	return ms.database.WorkingHash()
}

// AvailableVersions returns a list of available versions.
func (ms *MultiStoreImpl) AvailableVersions() []int {
	return ms.database.AvailableVersions()
}

// IsEmpty checks if the database is empty.
func (ms *MultiStoreImpl) IsEmpty() bool {
	return ms.database.IsEmpty()
}

// Set stores the key-value pair in the database. If the key already exists, its value will be updated.
func (ms *MultiStoreImpl) Set(key, value []byte) error {
	return ms.database.Set(key, value)
}

// Delete removes the key-value pair from the database.
func (ms *MultiStoreImpl) Delete(key []byte) error {
	return ms.database.Delete(key)
}

// Load loads the latest versioned database from disk.
func (ms *MultiStoreImpl) Load() (int64, error) {
	return ms.database.Load()
}

// LoadVersion loads a specific version of the database from disk.
func (ms *MultiStoreImpl) LoadVersion(targetVersion int64) (int64, error) {
	return ms.database.LoadVersion(targetVersion)
}

// SaveVersion saves a new version of the database to disk.
func (ms *MultiStoreImpl) SaveVersion() ([]byte, int64, error) {
	return ms.database.SaveVersion()
}

// Rollback resets the working database to the latest saved version, discarding any unsaved modifications.
func (ms *MultiStoreImpl) Rollback() {
	ms.database.Rollback()
}

// Close closes the database.
func (ms *MultiStoreImpl) Close() error {
	return ms.database.Close()
}
