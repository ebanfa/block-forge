package mocks

import (
	"github.com/edward1christian/block-forge/pkg/application/store"
	"github.com/stretchr/testify/mock"
)

// MockMultiStore is a mock implementation of the MultiStore interface.
type MockMultiStore struct {
	mock.Mock
}

// Get retrieves the value associated with the given key from the database.
func (m *MockMultiStore) Get(key []byte) ([]byte, error) {
	args := m.Called(key)
	return args.Get(0).([]byte), args.Error(1)
}

// Has checks if a key exists in the database.
func (m *MockMultiStore) Has(key []byte) (bool, error) {
	args := m.Called(key)
	return args.Bool(0), args.Error(1)
}

// Iterate iterates over all key-value pairs in the database and calls the given function for each pair.
// Iteration stops if the function returns true.
func (m *MockMultiStore) Iterate(fn func(key, value []byte) bool) error {
	args := m.Called(fn)
	return args.Error(0)
}

// IterateRange iterates over key-value pairs with keys in the specified range
// and calls the given function for each pair. Iteration stops if the function returns true.
func (m *MockMultiStore) IterateRange(start, end []byte, ascending bool, fn func(key, value []byte) bool) error {
	args := m.Called(start, end, ascending, fn)
	return args.Error(0)
}

// Hash returns the hash of the database.
func (m *MockMultiStore) Hash() []byte {
	args := m.Called()
	return args.Get(0).([]byte)
}

// Version returns the version of the database.
func (m *MockMultiStore) Version() int64 {
	args := m.Called()
	return args.Get(0).(int64)
}

// String returns a string representation of the database.
func (m *MockMultiStore) String() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

// WorkingVersion returns the current working version of the database.
func (m *MockMultiStore) WorkingVersion() int64 {
	args := m.Called()
	return args.Get(0).(int64)
}

// WorkingHash returns the hash of the current working version of the database.
func (m *MockMultiStore) WorkingHash() []byte {
	args := m.Called()
	return args.Get(0).([]byte)
}

// AvailableVersions returns a list of available versions.
func (m *MockMultiStore) AvailableVersions() []int {
	args := m.Called()
	return args.Get(0).([]int)
}

// IsEmpty checks if the database is empty.
func (m *MockMultiStore) IsEmpty() bool {
	args := m.Called()
	return args.Bool(0)
}

// Set stores the key-value pair in the database. If the key already exists, its value will be updated.
func (m *MockMultiStore) Set(key, value []byte) error {
	args := m.Called(key, value)
	return args.Error(0)
}

// Delete removes the key-value pair from the database.
func (m *MockMultiStore) Delete(key []byte) error {
	args := m.Called(key)
	return args.Error(0)
}

// Load loads the latest versioned database from disk.
func (m *MockMultiStore) Load() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

// LoadVersion loads a specific version of the database from disk.
func (m *MockMultiStore) LoadVersion(targetVersion int64) (int64, error) {
	args := m.Called(targetVersion)
	return args.Get(0).(int64), args.Error(1)
}

// SaveVersion saves a new version of the database to disk.
func (m *MockMultiStore) SaveVersion() ([]byte, int64, error) {
	args := m.Called()
	return args.Get(0).([]byte), args.Get(1).(int64), args.Error(2)
}

// Rollback resets the working database to the latest saved version, discarding any unsaved modifications.
func (m *MockMultiStore) Rollback() {
	m.Called()
}

// Close closes the database.
func (m *MockMultiStore) Close() error {
	args := m.Called()
	return args.Error(0)
}

// GetStore returns the store with the given namespace. If the store doesn't exist, it creates and initializes
// a new store using the provided options.
func (m *MockMultiStore) GetStore(namespace []byte, options store.StoreOptions) (store.Store, error) {
	args := m.Called(namespace, options)
	return args.Get(0).(store.Store), args.Error(1)
}

// CreateStore creates and initializes a new store with the given namespace and options. If a store with the same
// namespace already exists, it returns an error.
func (m *MockMultiStore) CreateStore(namespace []byte, options store.StoreOptions) (store.Store, error) {
	args := m.Called(namespace, options)
	return args.Get(0).(store.Store), args.Error(1)
}

// GetStoreCount returns the total number of stores in the multistore.
func (m *MockMultiStore) GetStoreCount() int {
	args := m.Called()
	return args.Int(0)
}
