package mocks

import (
	dbm "github.com/cosmos/iavl/db"
	"github.com/stretchr/testify/mock"
)

// MockDB is a mock implementation of the DB interface
type MockDB struct {
	mock.Mock
}

// Get is a mocked method for fetching the value of the given key
func (m *MockDB) Get(key []byte) ([]byte, error) {
	args := m.Called(key)
	return args.Get(0).([]byte), args.Error(1)
}

// Has is a mocked method for checking if a key exists
func (m *MockDB) Has(key []byte) (bool, error) {
	args := m.Called(key)
	return args.Bool(0), args.Error(1)
}

// Iterator is a mocked method for creating an iterator
func (m *MockDB) Iterator(start, end []byte) (dbm.Iterator, error) {
	args := m.Called(start, end)
	return args.Get(0).(dbm.Iterator), args.Error(1)
}

// ReverseIterator is a mocked method for creating a reverse iterator
func (m *MockDB) ReverseIterator(start, end []byte) (dbm.Iterator, error) {
	args := m.Called(start, end)
	return args.Get(0).(dbm.Iterator), args.Error(1)
}

// Close is a mocked method for closing the database connection
func (m *MockDB) Close() error {
	args := m.Called()
	return args.Error(0)
}

// NewBatch is a mocked method for creating a new batch
func (m *MockDB) NewBatch() dbm.Batch {
	args := m.Called()
	return args.Get(0).(dbm.Batch)
}

// NewBatchWithSize is a mocked method for creating a new batch with pre-allocated size
func (m *MockDB) NewBatchWithSize(size int) dbm.Batch {
	args := m.Called(size)
	return args.Get(0).(dbm.Batch)
}

// MockIterator is a mock implementation of the Iterator interface
type MockIterator struct {
	mock.Mock
}

// Domain is a mocked method for getting the domain of the iterator
func (m *MockIterator) Domain() (start []byte, end []byte) {
	args := m.Called()
	return args.Get(0).([]byte), args.Get(1).([]byte)
}

// Valid is a mocked method for checking if the iterator is valid
func (m *MockIterator) Valid() bool {
	args := m.Called()
	return args.Bool(0)
}

// Next is a mocked method for moving to the next key in the database
func (m *MockIterator) Next() {
	m.Called()
}

// Key is a mocked method for getting the key at the current position
func (m *MockIterator) Key() []byte {
	args := m.Called()
	return args.Get(0).([]byte)
}

// Value is a mocked method for getting the value at the current position
func (m *MockIterator) Value() []byte {
	args := m.Called()
	return args.Get(0).([]byte)
}

// Error is a mocked method for getting the last error encountered by the iterator
func (m *MockIterator) Error() error {
	args := m.Called()
	return args.Error(0)
}

// Close is a mocked method for closing the iterator
func (m *MockIterator) Close() error {
	args := m.Called()
	return args.Error(0)
}

// MockBatch is a mock implementation of the Batch interface
type MockBatch struct {
	mock.Mock
}

// Set is a mocked method for setting a key/value pair in the batch
func (m *MockBatch) Set(key, value []byte) error {
	args := m.Called(key, value)
	return args.Error(0)
}

// Delete is a mocked method for deleting a key/value pair from the batch
func (m *MockBatch) Delete(key []byte) error {
	args := m.Called(key)
	return args.Error(0)
}

// Write is a mocked method for writing the batch
func (m *MockBatch) Write() error {
	args := m.Called()
	return args.Error(0)
}

// WriteSync is a mocked method for writing the batch and flushing it to disk
func (m *MockBatch) WriteSync() error {
	args := m.Called()
	return args.Error(0)
}

// Close is a mocked method for closing the batch
func (m *MockBatch) Close() error {
	args := m.Called()
	return args.Error(0)
}

// GetByteSize is a mocked method for getting the current size of the batch in bytes
func (m *MockBatch) GetByteSize() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}
