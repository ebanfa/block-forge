package mocks

import (
	dbm "github.com/cosmos/iavl/db"
	"github.com/stretchr/testify/mock"
)

// MockDBM is a mock implementation of the DB interface for testing.
type MockDBM struct {
	mock.Mock
}

// Get is a mock implementation of the Get method.
func (m *MockDBM) Get(key []byte) ([]byte, error) {
	args := m.Called(key)
	return args.Get(0).([]byte), args.Error(1)
}

// Has is a mock implementation of the Has method.
func (m *MockDBM) Has(key []byte) (bool, error) {
	args := m.Called(key)
	return args.Bool(0), args.Error(1)
}

// Iterator is a mock implementation of the Iterator method.
func (m *MockDBM) Iterator(start, end []byte) (dbm.Iterator, error) {
	args := m.Called(start, end)
	return args.Get(0).(dbm.Iterator), args.Error(1)
}

// ReverseIterator is a mock implementation of the ReverseIterator method.
func (m *MockDBM) ReverseIterator(start, end []byte) (dbm.Iterator, error) {
	args := m.Called(start, end)
	return args.Get(0).(dbm.Iterator), args.Error(1)
}

// Close is a mock implementation of the Close method.
func (m *MockDBM) Close() error {
	args := m.Called()
	return args.Error(0)
}

// NewBatch is a mock implementation of the NewBatch method.
func (m *MockDBM) NewBatch() dbm.Batch {
	args := m.Called()
	return args.Get(0).(dbm.Batch)
}

// NewBatchWithSize is a mock implementation of the NewBatchWithSize method.
func (m *MockDBM) NewBatchWithSize(size int) dbm.Batch {
	args := m.Called(size)
	return args.Get(0).(dbm.Batch)
}
