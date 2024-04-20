package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/component"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// MockOperationInput is a mock structure for OperationInput.
type MockOperationInput struct {
	mock.Mock
}

// MockOperationOutput is a mock structure for OperationOutput.
type MockOperationOutput struct {
	mock.Mock
}

// NewMockOperationInput creates a new instance of MockOperationInput.
func NewMockOperationInput(data interface{}) *MockOperationInput {
	return &MockOperationInput{}
}

// NewMockOperationOutput creates a new instance of MockOperationOutput.
func NewMockOperationOutput(data interface{}) *MockOperationOutput {
	return &MockOperationOutput{}
}

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
func (m *MockOperation) Type() component.ComponentType {
	args := m.Called()
	return args.Get(0).(component.ComponentType)
}

// Start mocks the Start method of the SystemService interface.
func (m *MockOperation) Initialize(ctx *context.Context, system *MockSystem) error {
	args := m.Called(ctx, system)
	return args.Error(0)
}

// Execute mocks the Execute method of the Operation interface.
func (m *MockOperation) Execute(ctx *context.Context, input *system.SystemOperationInput) (*system.SystemOperationOutput, error) {
	args := m.Called(ctx, input)
	return args.Get(0).(*system.SystemOperationOutput), args.Error(1)
}
