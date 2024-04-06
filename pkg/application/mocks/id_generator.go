package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MockIDGenerator is a mock for the ProcessIDGenerator interface.
type MockIDGenerator struct {
	mock.Mock
}

// GenerateID provides a mock function to generate a process ID.
func (m *MockIDGenerator) GenerateID() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}
