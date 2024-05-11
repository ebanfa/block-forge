package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MockStore is a mock implementation of the Store interface for testing purposes.
type MockStore struct {
	mock.Mock
}

// Get mocks the Get method of Database.
func (m *MockStore) Get(key []byte) ([]byte, error) {
	args := m.Called(key)
	return args.Get(0).([]byte), args.Error(1)
}

// Set mocks the Set method of Database.
func (m *MockStore) Set(key, value []byte) error {
	args := m.Called(key, value)
	return args.Error(0)
}

// Delete mocks the Delete method of Database.
func (m *MockStore) Delete(key []byte) error {
	args := m.Called(key)
	return args.Error(0)
}

// Has mocks the Has method of Database.
func (m *MockStore) Has(key []byte) (bool, error) {
	args := m.Called(key)
	return args.Bool(0), args.Error(1)
}

// Iterate mocks the Iterate method of Database.
func (m *MockStore) Iterate(fn func(key, value []byte) bool) error {
	args := m.Called(fn)
	return args.Error(0)
}

// IterateRange mocks the IterateRange method of Database.
func (m *MockStore) IterateRange(start, end []byte, ascending bool, fn func(key, value []byte) bool) error {
	args := m.Called(start, end, ascending, fn)
	return args.Error(0)
}

// Hash mocks the Hash method of Database.
func (m *MockStore) Hash() []byte {
	args := m.Called()
	return args.Get(0).([]byte)
}

// Version mocks the Version method of Database.
func (m *MockStore) Version() int64 {
	args := m.Called()
	return args.Get(0).(int64)
}

// Load mocks the Load method of Database.
func (m *MockStore) Load() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

// SaveVersion mocks the SaveVersion method of Database.
func (m *MockStore) SaveVersion() ([]byte, int64, error) {
	args := m.Called()
	return args.Get(0).([]byte), args.Get(1).(int64), args.Error(2)
}

// Rollback mocks the Rollback method of Database.
func (m *MockStore) Rollback() {
	m.Called()
}

// Close mocks the Close method of Database.
func (m *MockStore) Close() error {
	args := m.Called()
	return args.Error(0)
}

// String mocks the String method of Database.
func (m *MockStore) String() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

// WorkingVersion mocks the WorkingVersion method of Database.
func (m *MockStore) WorkingVersion() int64 {
	args := m.Called()
	return args.Get(0).(int64)
}

// LoadVersion mocks the LoadVersion method of Database.
func (m *MockStore) LoadVersion(targetVersion int64) (int64, error) {
	args := m.Called(targetVersion)
	return args.Get(0).(int64), args.Error(1)
}

// WorkingHash mocks the WorkingHash method of Database.
func (m *MockStore) WorkingHash() []byte {
	args := m.Called()
	return args.Get(0).([]byte)
}

// AvailableVersions mocks the AvailableVersions method of Database.
func (m *MockStore) AvailableVersions() []int {
	args := m.Called()
	return args.Get(0).([]int)
}

// IsEmpty mocks the IsEmpty method of Database.
func (m *MockStore) IsEmpty() bool {
	args := m.Called()
	return args.Bool(0)
}

// Name implements the Name method of the Store interface.
func (m *MockStore) Name() string {
	args := m.Called()
	return args.String(0)
}

// Path implements the Path method of the Store interface.
func (m *MockStore) Path() string {
	args := m.Called()
	return args.String(0)
}
