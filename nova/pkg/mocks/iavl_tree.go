package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MockMutableTree is a mock implementation of the iavl.MutableTree interface.
type MockMutableTree struct {
	mock.Mock
}

// Get provides a mock function to get the value associated with the given key from the tree.
func (m *MockMutableTree) Get(key []byte) ([]byte, error) {
	args := m.Called(key)
	return args.Get(0).([]byte), args.Error(1)
}

// Set provides a mock function to store the key-value pair in the tree.
func (m *MockMutableTree) Set(key, value []byte) ([]byte, error) {
	args := m.Called(key, value)
	return args.Get(0).([]byte), args.Error(1)
}

// Remove provides a mock function to remove the key-value pair from the tree.
func (m *MockMutableTree) Remove(key []byte) ([]byte, []byte, error) {
	args := m.Called(key)
	return args.Get(0).([]byte), args.Get(1).([]byte), args.Error(2)
}

// Has provides a mock function to check if the key exists in the tree.
func (m *MockMutableTree) Has(key []byte) (bool, error) {
	args := m.Called(key)
	return args.Bool(0), args.Error(1)
}

// Iterate provides a mock function to iterate over all keys of the tree.
func (m *MockMutableTree) Iterate(fn func(key []byte, value []byte) bool) (bool, error) {
	args := m.Called(fn)
	return args.Bool(0), args.Error(1)
}

// IterateRange provides a mock function to iterate over all key-value pairs with keys in the specified range.
func (m *MockMutableTree) IterateRange(start, end []byte, ascending bool, fn func(key []byte, value []byte) bool) error {
	args := m.Called(start, end, ascending, fn)
	return args.Error(0)
}

// Hash provides a mock function to get the root hash of the tree.
func (m *MockMutableTree) Hash() []byte {
	args := m.Called()
	return args.Get(0).([]byte)
}

// WorkingVersion provides a mock function to get the version of the tree.
func (m *MockMutableTree) WorkingVersion() int64 {
	args := m.Called()
	return args.Get(0).(int64)
}

// Load provides a mock function to load the latest versioned tree from disk.
func (m *MockMutableTree) Load() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

// SaveVersion provides a mock function to save a new tree version to disk.
func (m *MockMutableTree) SaveVersion() ([]byte, int64, error) {
	args := m.Called()
	return args.Get(0).([]byte), args.Get(1).(int64), args.Error(2)
}

// Rollback provides a mock function to reset the working tree to the latest saved version.
func (m *MockMutableTree) Rollback() {
	m.Called()
}

// Close provides a mock function to close the tree.
func (m *MockMutableTree) Close() error {
	args := m.Called()
	return args.Error(0)
}

// String provides a mock function to get a string representation of the tree.
func (m *MockMutableTree) String() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}
