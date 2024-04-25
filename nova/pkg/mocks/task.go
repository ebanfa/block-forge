package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MockTask is a mock implementation of TaskInterface.
type MockTask struct {
	mock.Mock
}

// ID mocks the ID method of TaskInterface.
func (m *MockTask) ID() string {
	args := m.Called()
	return args.String(0)
}

// Name mocks the Name method of TaskInterface.
func (m *MockTask) Name() string {
	args := m.Called()
	return args.String(0)
}

// Description mocks the Description method of TaskInterface.
func (m *MockTask) Description() string {
	args := m.Called()
	return args.String(0)
}
