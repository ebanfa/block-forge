package tests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/edward1christian/block-forge/pkg/application/mocks"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// TestNewSystemOperations tests the creation of a new SystemOperations instance.
func TestNewSystemOperations(t *testing.T) {
	ops := system.NewSystemOperations("system_operations", "System Operations", "Manages and executes operations within the system")
	require.NotNil(t, ops)
	assert.Equal(t, "system_operations", ops.ID())
	assert.Equal(t, "System Operations", ops.Name())
	assert.Equal(t, "Manages and executes operations within the system", ops.Description())
}

// TestSystemOperations_Initialize tests the initialization of the SystemOperations instance.
func TestSystemOperations_Initialize(t *testing.T) {
	ops := system.NewSystemOperations("system_operations", "System Operations", "Manages and executes operations within the system")
	system := &mocks.MockSystem{}
	err := ops.Initialize(nil, system)
	require.NoError(t, err)
}

// TestSystemOperations_RegisterOperation tests the registration of an operation.
func TestSystemOperations_RegisterOperation(t *testing.T) {
	ops := system.NewSystemOperations("system_operations", "System Operations", "Manages and executes operations within the system")
	op := new(mocks.MockOperation)
	err := ops.RegisterOperation("mock_op", op)
	assert.NoError(t, err)
}

// TestSystemOperations_ExecuteOperation tests the execution of a registered operation.
// TestSystemOperations_ExecuteOperation tests the execution of a registered operation.
func TestSystemOperations_ExecuteOperation(t *testing.T) {
	ops := system.NewSystemOperations("system_operations", "System Operations", "Manages and executes operations within the system")
	// Create a mock system using testify/mock
	mockSystem := &mocks.MockSystem{}
	ops.Initialize(nil, mockSystem)

	mockOp := new(mocks.MockOperation)
	err := ops.RegisterOperation("mock_op", mockOp)
	require.NoError(t, err)

	input := system.OperationInput{Data: "test_data"}
	expectedOutput := system.OperationOutput{Data: "test_data_result"}

	// Configure mockOp expectations
	mockOp.On("Initialize", mock.Anything, mockSystem).
		Return(nil).
		Once()

	// Configure mockOp expectations
	mockOp.On("Execute", mock.Anything, input).
		Return(expectedOutput, nil).
		Once()

	output, err := ops.ExecuteOperation(nil, "mock_op", input)
	require.NoError(t, err)
	assert.Equal(t, expectedOutput, output)

	// Assert that the mock operation's Execute method was called with the expected parameters
	mockOp.AssertExpectations(t)
}

// TestSystemOperations_ExecuteOperation_InitializeFailure tests the execution when the operation initialization fails.
func TestSystemOperations_ExecuteOperation_InitializeFailure(t *testing.T) {
	ops := system.NewSystemOperations("system_operations", "System Operations", "Manages and executes operations within the system")
	mockSystem := &mocks.MockSystem{}
	ops.Initialize(nil, mockSystem)

	mockOp := &mocks.MockOperation{}
	err := ops.RegisterOperation("mock_op", mockOp)
	require.NoError(t, err)

	expectedError := errors.New("initialization failed")

	// Configure mockOp expectations for initialization failure
	mockOp.On("Initialize", mock.Anything, mock.Anything).
		Return(expectedError).
		Once()

	_, err = ops.ExecuteOperation(nil, "mock_op", system.OperationInput{})
	require.Error(t, err)
	assert.EqualError(t, err, "failed to initialize operation: initialization failed")

	// Assert that the mock operation's Initialize method was called with the expected parameters
	mockOp.AssertExpectations(t)
}

// TestSystemOperations_ExecuteOperation_NotFound tests the execution of a non-registered operation.
func TestSystemOperations_ExecuteOperation_NotFound(t *testing.T) {
	ops := system.NewSystemOperations("system_operations", "System Operations", "Manages and executes operations within the system")
	mockSystem := &mocks.MockSystem{}
	ops.Initialize(nil, mockSystem)

	_, err := ops.ExecuteOperation(nil, "unknown_op", system.OperationInput{})
	assert.Error(t, err)
	assert.EqualError(t, err, "operation with ID unknown_op not found")
}
