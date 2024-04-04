package test_ops

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/edward1christian/block-forge/pkg/etl"
	"github.com/edward1christian/block-forge/pkg/etl/mocks"
	etl_ops "github.com/edward1christian/block-forge/pkg/etl/ops"
)

func TestRollbackOperation_Execute_Success(t *testing.T) {
	// Mock context
	ctx := &context.Context{}

	// Mock input data
	process := &etl.ETLProcess{
		Components: map[string]etl.ETLComponent{
			"pipeline1": &mocks.MockPipeline{},
			"pipeline2": &mocks.MockPipeline{},
		},
	}
	opInput := &system.OperationInput{
		Data: process,
	}

	// Mock pipelines
	mockPipeline1 := process.Components["pipeline1"].(*mocks.MockPipeline)
	mockPipeline1.On("Stop", ctx).Return(nil)

	mockPipeline2 := process.Components["pipeline2"].(*mocks.MockPipeline)
	mockPipeline2.On("Stop", ctx).Return(nil)

	// Create operation
	operation := etl_ops.NewRollbackOperation()

	// Execute operation
	opOutput, err := operation.Execute(ctx, opInput)

	// Assertions
	assert.NoError(t, err, "Execute should not return an error")
	assert.Nil(t, opOutput, "Operation output should be nil")

	// Ensure Stop method is called for each component
	mockPipeline1.AssertCalled(t, "Stop", ctx)
	mockPipeline2.AssertCalled(t, "Stop", ctx)
}

func TestRollbackOperation_Execute_InvalidProcess(t *testing.T) {
	// Mock context
	ctx := &context.Context{}

	// Mock input data (invalid)
	data := "invalid data"
	opInput := &system.OperationInput{
		Data: data,
	}

	// Create operation
	operation := etl_ops.NewRollbackOperation()

	// Execute operation
	opOutput, err := operation.Execute(ctx, opInput)

	// Assertions
	assert.Error(t, err, "Execute should return an error for invalid process")
	assert.Nil(t, opOutput, "Operation output should be nil")
}
