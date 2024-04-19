package mocks

import (
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/component"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/stretchr/testify/mock"
)

// MockBuildTask is a mock implementation of the BuildTaskInterface.
type MockBuildTask struct {
	mock.Mock
}

// ID returns the unique identifier of the component.
func (m *MockBuildTask) ID() string {
	args := m.Called()
	return args.String(0)
}

// Name returns the name of the component.
func (m *MockBuildTask) Name() string {
	args := m.Called()
	return args.String(0)
}

// Description returns the description of the component.
func (m *MockBuildTask) Description() string {
	args := m.Called()
	return args.String(0)
}

// Description returns the description of the component.
func (m *MockBuildTask) Type() component.ComponentType {
	args := m.Called()
	return args.Get(0).(component.ComponentType)
}

// GetName implements the GetName method of the BuildTaskInterface.
func (m *MockBuildTask) GetName() string {
	args := m.Called()
	return args.String(0)
}

// Start mocks the Start method of the SystemService interface.
func (m *MockBuildTask) Initialize(ctx *context.Context, system system.SystemInterface) error {
	args := m.Called(ctx, system)
	return args.Error(0)
}

// Execute mocks the Execute method of the SystemOperation interface.
func (m *MockBuildTask) Execute(ctx *context.Context, input *system.SystemOperationInput) (*system.SystemOperationOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*system.SystemOperationOutput), args.Error(1)
}
