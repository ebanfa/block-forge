package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MockProcessIDGenerator is a mock implementation of ProcessIDGenerator interface.
type MockProcessIDGenerator struct {
	mock.Mock
}

// GenerateID generates a unique process ID.
func (m *MockProcessIDGenerator) GenerateID() (string, error) {
	// Mocking the behavior of GenerateID method
	return "mock-process-id", nil
}

// NewMockProcessIDGenerator creates a new instance of MockProcessIDGenerator.
func NewMockProcessIDGenerator() *MockProcessIDGenerator {
	return &MockProcessIDGenerator{}
}
