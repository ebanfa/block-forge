package mocks

import (
	"github.com/edward1christian/block-forge/pkg/application/component"
	"github.com/stretchr/testify/mock"
)

type MockComponentType int

// MockComponent is a mock implementation of the ComponentInterface interface.
type MockComponent struct {
	mock.Mock
}

// ID returns the unique identifier of the component.
func (m *MockComponent) ID() string {
	args := m.Called()
	return args.String(0)
}

// Name returns the name of the component.
func (m *MockComponent) Name() string {
	args := m.Called()
	return args.String(0)
}

// Type returns the type of the component.
func (m *MockComponent) Type() component.ComponentType {
	args := m.Called()
	return args.Get(0).(component.ComponentType)
}

// Description returns the description of the component.
func (m *MockComponent) Description() string {
	args := m.Called()
	return args.String(0)
}
