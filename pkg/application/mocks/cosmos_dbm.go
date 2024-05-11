package mocks

import (
	dbm "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/mock"
)

// MockDBM is a mock implementation of the DB interface for testing purposes.
type MockDBM struct {
	mock.Mock
}

// Get mocks the Get method of the DB interface.
func (m *MockDBM) Get(key []byte) ([]byte, error) {
	args := m.Called(key)
	return args.Get(0).([]byte), args.Error(1)
}

// Has mocks the Has method of the DB interface.
func (m *MockDBM) Has(key []byte) (bool, error) {
	args := m.Called(key)
	return args.Bool(0), args.Error(1)
}

// Set mocks the Set method of the DB interface.
func (m *MockDBM) Set(key, value []byte) error {
	args := m.Called(key, value)
	return args.Error(0)
}

// SetSync mocks the SetSync method of the DB interface.
func (m *MockDBM) SetSync(key, value []byte) error {
	args := m.Called(key, value)
	return args.Error(0)
}

// Delete mocks the Delete method of the DB interface.
func (m *MockDBM) Delete(key []byte) error {
	args := m.Called(key)
	return args.Error(0)
}

// DeleteSync mocks the DeleteSync method of the DB interface.
func (m *MockDBM) DeleteSync(key []byte) error {
	args := m.Called(key)
	return args.Error(0)
}

// Iterator mocks the Iterator method of the DB interface.
func (m *MockDBM) Iterator(start, end []byte) (dbm.Iterator, error) {
	args := m.Called(start, end)
	return args.Get(0).(dbm.Iterator), args.Error(1)
}

// ReverseIterator mocks the ReverseIterator method of the DB interface.
func (m *MockDBM) ReverseIterator(start, end []byte) (dbm.Iterator, error) {
	args := m.Called(start, end)
	return args.Get(0).(dbm.Iterator), args.Error(1)
}

// Close mocks the Close method of the DB interface.
func (m *MockDBM) Close() error {
	args := m.Called()
	return args.Error(0)
}

// NewBatch mocks the NewBatch method of the DB interface.
func (m *MockDBM) NewBatch() dbm.Batch {
	args := m.Called()
	return args.Get(0).(dbm.Batch)
}

// NewBatchWithSize mocks the NewBatchWithSize method of the DB interface.
func (m *MockDBM) NewBatchWithSize(size int) dbm.Batch {
	args := m.Called(size)
	return args.Get(0).(dbm.Batch)
}

// Print mocks the Print method of the DB interface.
func (m *MockDBM) Print() error {
	args := m.Called()
	return args.Error(0)
}

// Stats mocks the Stats method of the DB interface.
func (m *MockDBM) Stats() map[string]string {
	args := m.Called()
	return args.Get(0).(map[string]string)
}
