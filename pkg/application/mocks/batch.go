package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MockBatch is a mock implementation of the Batch interface for testing.
type MockBatch struct {
	mock.Mock
}

// Set is a mock implementation of the Set method.
func (m *MockBatch) Set(key, value []byte) error {
	args := m.Called(key, value)
	return args.Error(0)
}

// Delete is a mock implementation of the Delete method.
func (m *MockBatch) Delete(key []byte) error {
	args := m.Called(key)
	return args.Error(0)
}

// Write is a mock implementation of the Write method.
func (m *MockBatch) Write() error {
	args := m.Called()
	return args.Error(0)
}

// WriteSync is a mock implementation of the WriteSync method.
func (m *MockBatch) WriteSync() error {
	args := m.Called()
	return args.Error(0)
}

// Close is a mock implementation of the Close method.
func (m *MockBatch) Close() error {
	args := m.Called()
	return args.Error(0)
}

// GetByteSize is a mock implementation of the GetByteSize method.
func (m *MockBatch) GetByteSize() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}
