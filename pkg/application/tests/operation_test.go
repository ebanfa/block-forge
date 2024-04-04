package tests

import (
	"testing"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/components"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/stretchr/testify/assert"
)

// TestNewBaseSystemOperation tests the NewBaseSystemOperation function.
func TestNewBaseSystemOperation(t *testing.T) {
	// Call the NewBaseSystemOperation function to create a new BaseSystemOperation instance
	operation := systemApi.NewBaseSystemOperation("1", "Operation1", "Description1")

	// Check if the instance is not nil
	assert.NotNil(t, operation)

	// Check if the ID of the created instance matches the expected value
	assert.Equal(t, "1", operation.ID())

	// Check if the Name of the created instance matches the expected value
	assert.Equal(t, "Operation1", operation.Name())

	// Check if the Description of the created instance matches the expected value
	assert.Equal(t, "Description1", operation.Description())
}

// TestBaseSystemOperation_Type tests the Type method of BaseSystemOperation.
func TestBaseSystemOperation_Type(t *testing.T) {
	// Create a new BaseSystemOperation instance
	operation := systemApi.NewBaseSystemOperation("1", "Operation1", "Description1")

	// Call the Type method to get the component type
	componentType := operation.Type()

	// Check if the returned component type matches the expected value
	assert.Equal(t, components.BasicComponentType, componentType)
}

// TestBaseSystemOperation_ImplementingInterface tests if BaseSystemOperation implements the OperationInterface.
func TestBaseSystemOperation_ImplementingInterface(t *testing.T) {
	// Ensure that BaseSystemOperation implements the OperationInterface
	var _ systemApi.OperationInterface = (*systemApi.BaseSystemOperation)(nil)
}

// TestBaseSystemOperation_Execute tests the Execute method of BaseSystemOperation.
func TestBaseSystemOperation_Execute(t *testing.T) {
	// Create a new BaseSystemOperation instance
	operation := systemApi.NewBaseSystemOperation("1", "Operation1", "Description1")

	// Create a mock context
	mockContext := &context.Context{}

	// Call the Execute method
	output, err := operation.Execute(mockContext, nil)

	// Check if an error is returned
	assert.Error(t, err)

	// Check if the error message matches the expected value
	assert.EqualError(t, err, "operation not implemented")

	// Check if the output is nil
	assert.Nil(t, output)
}

// TestBaseSystemOperation_Initialize tests the Initialize method of BaseSystemOperation.
func TestBaseSystemOperation_Initialize(t *testing.T) {
	// Create a new BaseSystemOperation instance
	operation := systemApi.NewBaseSystemOperation("1", "Operation1", "Description1")

	// Create a mock context
	mockContext := &context.Context{}

	// Call the Initialize method
	err := operation.Initialize(mockContext, nil)

	// Check if an error is returned
	assert.Error(t, err)

	// Check if the error message matches the expected value
	assert.EqualError(t, err, "initialize not implemented")
}
