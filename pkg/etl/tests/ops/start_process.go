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

func TestStartETLOperation_Execute_Success(t *testing.T) {
	// Mock context
	ctx := &context.Context{}

	// Mock input data
	process := &etl.ETLProcess{
		Components: map[string]etl.ETLComponent{
			"pipeline1": &mocks.MockPipeline{},
		},
	}
	opInput := &system.OperationInput{
		Data: process,
	}

	// Mock pipeline start
	mockPipeline := process.Components["pipeline1"].(*mocks.MockPipeline)
	mockPipeline.On("Start", ctx).Return(nil)

	// Create operation
	operation := etl_ops.NewStartETLOperation()

	// Execute operation
	opOutput, err := operation.Execute(ctx, opInput)

	// Assertions
	assert.NoError(t, err, "Execute should not return an error")
	assert.NotNil(t, opOutput, "Operation output should not be nil")
	assert.NotNil(t, opOutput.Data, "Operation output data should not be nil")
	returnedProcess, ok := opOutput.Data.(*etl.ETLProcess)
	assert.True(t, ok, "Operation output data should be of type *etl.ETLProcess")
	assert.Equal(t, etl.ETLStatusRunning, returnedProcess.Status, "Process status should be running")
	assert.NotNil(t, returnedProcess.Components, "Process components map should be initialized")
}

func TestStartETLOperation_Execute_ComponentStartError(t *testing.T) {
	// Mock context
	ctx := &context.Context{}

	// Mock input data
	process := &etl.ETLProcess{
		Components: map[string]etl.ETLComponent{
			"pipeline1": &mocks.MockPipeline{},
		},
	}
	opInput := &system.OperationInput{
		Data: process,
	}

	// Mock pipeline start with error
	mockPipeline := process.Components["pipeline1"].(*mocks.MockPipeline)
	mockPipeline.On("Start", ctx).Return(errors.New("start error"))

	// Create operation
	operation := etl_ops.NewStartETLOperation()

	// Execute operation
	opOutput, err := operation.Execute(ctx, opInput)

	// Assertions
	assert.Error(t, err, "Execute should return an error")
	assert.Nil(t, opOutput, "Operation output should be nil")
	assert.Equal(t, etl.ETLStatusInitialized, process.Status, "Process status should remain initialized")
}
