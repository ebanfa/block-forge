package test_ops

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/edward1christian/block-forge/pkg/etl"
	"github.com/edward1christian/block-forge/pkg/etl/mocks"
	etl_ops "github.com/edward1christian/block-forge/pkg/etl/ops"
)

func TestCreateETLOperation_Execute_Success(t *testing.T) {
	// Mock context
	ctx := &context.Context{}

	// Mock input data
	config := &etl.ETLConfig{
		Components: []*etl.ETLComponentConfig{
			{Name: "Component1"},
			// Add more entries as needed
		},
	}
	opInput := &system.OperationInput{
		Data: config,
	}

	// Mock ID generator
	mockIDGenerator := mocks.NewMockProcessIDGenerator()
	mockIDGenerator.On("GenerateID").Return("mockProcessID", nil)

	// Create operation
	operation := etl_ops.NewCreateETLOperation(mockIDGenerator)

	// Execute operation
	opOutput, err := operation.Execute(ctx, opInput)

	// Assertions
	assert.NoError(t, err, "Execute should not return an error")
	assert.NotNil(t, opOutput, "Operation output should not be nil")
	assert.NotNil(t, opOutput.Data, "Operation output data should not be nil")
	process, ok := opOutput.Data.(*etl.ETLProcess)
	assert.True(t, ok, "Operation output data should be of type *etl.ETLProcess")
	assert.Equal(t, "mockProcessID", process.ID, "Process ID should match the mocked ID")
	assert.Equal(t, etl.ETLStatusInitialized, process.Status, "Process status should be initialized")
	assert.NotNil(t, process.Components, "Process components map should be initialized")
}

func TestCreateETLOperation_Execute_EmptyConfig(t *testing.T) {
	// Mock context
	ctx := &context.Context{}

	// Mock input data
	config := &etl.ETLConfig{} // Empty config
	opInput := &system.OperationInput{
		Data: config,
	}

	// Mock ID generator
	mockIDGenerator := mocks.NewMockProcessIDGenerator()

	// Create operation
	operation := etl_ops.NewCreateETLOperation(mockIDGenerator)

	// Execute operation
	opOutput, err := operation.Execute(ctx, opInput)

	// Assertions
	assert.Error(t, err, "Execute should return an error for empty config")
	assert.Nil(t, opOutput, "Operation output should be nil")
}

func TestCreateETLOperation_Execute_IDGenerationError(t *testing.T) {
	// Mock context
	ctx := &context.Context{}

	// Mock input data
	config := &etl.ETLConfig{
		Components: []*etl.ETLComponentConfig{
			{Name: "Component1"},
			// Add more entries as needed
		},
	}
	opInput := &system.OperationInput{
		Data: config,
	}

	// Mock ID generator with error
	mockIDGenerator := mocks.NewMockProcessIDGenerator()
	mockIDGenerator.On("GenerateID").Return("", errors.New("ID generation error"))

	// Create operation
	operation := etl_ops.NewCreateETLOperation(mockIDGenerator)

	// Execute operation
	opOutput, err := operation.Execute(ctx, opInput)

	// Assertions
	assert.Error(t, err, "Execute should return an error for ID generation error")
	assert.Nil(t, opOutput, "Operation output should be nil")
}
