package mocks

import (
	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/stretchr/testify/mock"
)

// Mock implementations for testing
type MockOperation struct {
	mock.Mock
}

// ID implements application.Operations.
func (m *MockOperation) ID() string {
	args := m.Called()
	return args.String(0)
}

// Name implements application.Operations.
func (m *MockOperation) Name() string {
	args := m.Called()
	return args.String(0)
}

// Description implements application.Operations.
func (m *MockOperation) Description() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockOperation) Initialize(ctx *context.Context, system system.System) error {
	args := m.Called(ctx, system)
	return args.Error(0)
}

func (m *MockOperation) Execute(ctx *context.Context, data system.OperationInput) (system.OperationOutput, error) {
	args := m.Called(ctx, data)
	return args.Get(0).(system.OperationOutput), args.Error(1)
}

// OperationsMock is a mock for the Operations interface.
type MockOperations struct {
	mock.Mock
}

// ID implements application.Operations.
func (m *MockOperations) ID() string {
	args := m.Called()
	return args.String(0)
}

// Name implements application.Operations.
func (m *MockOperations) Name() string {
	args := m.Called()
	return args.String(0)
}

// Description implements application.Operations.
func (m *MockOperations) Description() string {
	args := m.Called()
	return args.String(0)
}

// Initialize implements application.Operations.
func (m *MockOperations) Initialize(ctx *context.Context, system system.System) error {
	args := m.Called(ctx, system)
	return args.Error(0)
}

// RegisterOperation mocks the RegisterOperation method of the Operations interface.
func (m *MockOperations) RegisterOperation(operationID string, operation system.Operation) error {
	args := m.Called(operationID, operation)
	return args.Error(0)
}

// ExecuteOperation mocks the ExecuteOperation method of the Operations interface.
func (m *MockOperations) ExecuteOperation(
	ctx *context.Context,
	operationID string,
	data system.OperationInput) (system.OperationOutput, error) {

	args := m.Called(ctx, operationID, data)
	return args.Get(0).(system.OperationOutput), args.Error(1)
}
