package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/components"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// MockOperation represents a mock for the Operation interface.
type MockOperation struct {
	mock.Mock
}

// ID returns the unique identifier of the component.
func (m *MockOperation) ID() string {
	args := m.Called()
	return args.String(0)
}

// Name returns the name of the component.
func (m *MockOperation) Name() string {
	args := m.Called()
	return args.String(0)
}

// Description returns the description of the component.
func (m *MockOperation) Description() string {
	args := m.Called()
	return args.String(0)
}

// Description returns the description of the component.
func (m *MockOperation) Type() components.ComponentType {
	args := m.Called()
	return args.Get(0).(components.ComponentType)
}

// Start mocks the Start method of the SystemService interface.
func (m *MockOperation) Initialize(ctx *context.Context, system system.SystemInterface) error {
	args := m.Called(ctx, system)
	return args.Error(0)
}

// Execute mocks the Execute method of the Operation interface.
func (m *MockOperation) Execute(ctx *context.Context, input *system.OperationInput) (*system.OperationOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*system.OperationOutput), args.Error(1)
}
