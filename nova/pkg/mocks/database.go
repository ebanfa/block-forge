package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MockDatabase is a mock implementation of the Database interface.
type MockDatabase struct {
	mock.Mock
}

// Get mocks the Get method of Database.
func (m *MockDatabase) Get(key []byte) ([]byte, error) {
	args := m.Called(key)
	return args.Get(0).([]byte), args.Error(1)
}

// Set mocks the Set method of Database.
func (m *MockDatabase) Set(key, value []byte) error {
	args := m.Called(key, value)
	return args.Error(0)
}

// Delete mocks the Delete method of Database.
func (m *MockDatabase) Delete(key []byte) error {
	args := m.Called(key)
	return args.Error(0)
}

// Has mocks the Has method of Database.
func (m *MockDatabase) Has(key []byte) (bool, error) {
	args := m.Called(key)
	return args.Bool(0), args.Error(1)
}

// Iterate mocks the Iterate method of Database.
func (m *MockDatabase) Iterate(fn func(key, value []byte) bool) error {
	args := m.Called(fn)
	return args.Error(0)
}

// IterateRange mocks the IterateRange method of Database.
func (m *MockDatabase) IterateRange(start, end []byte, ascending bool, fn func(key, value []byte) bool) error {
	args := m.Called(start, end, ascending, fn)
	return args.Error(0)
}

// Hash mocks the Hash method of Database.
func (m *MockDatabase) Hash() []byte {
	args := m.Called()
	return args.Get(0).([]byte)
}

// Version mocks the Version method of Database.
func (m *MockDatabase) Version() int64 {
	args := m.Called()
	return args.Get(0).(int64)
}

// Load mocks the Load method of Database.
func (m *MockDatabase) Load() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

// SaveVersion mocks the SaveVersion method of Database.
func (m *MockDatabase) SaveVersion() ([]byte, int64, error) {
	args := m.Called()
	return args.Get(0).([]byte), args.Get(1).(int64), args.Error(2)
}

// Rollback mocks the Rollback method of Database.
func (m *MockDatabase) Rollback() {
	m.Called()
}

// Close mocks the Close method of Database.
func (m *MockDatabase) Close() error {
	args := m.Called()
	return args.Error(0)
}

// String mocks the String method of Database.
func (m *MockDatabase) String() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

// WorkingVersion mocks the WorkingVersion method of Database.
func (m *MockDatabase) WorkingVersion() int64 {
	args := m.Called()
	return args.Get(0).(int64)
}

// LoadVersion mocks the LoadVersion method of Database.
func (m *MockDatabase) LoadVersion(targetVersion int64) (int64, error) {
	args := m.Called(targetVersion)
	return args.Get(0).(int64), args.Error(1)
}

// WorkingHash mocks the WorkingHash method of Database.
func (m *MockDatabase) WorkingHash() []byte {
	args := m.Called()
	return args.Get(0).([]byte)
}

// AvailableVersions mocks the AvailableVersions method of Database.
func (m *MockDatabase) AvailableVersions() []int {
	args := m.Called()
	return args.Get(0).([]int)
}

// IsEmpty mocks the IsEmpty method of Database.
func (m *MockDatabase) IsEmpty() bool {
	args := m.Called()
	return args.Bool(0)
}
