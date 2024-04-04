package test_ops

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/edward1christian/block-forge/pkg/etl"
	"github.com/edward1christian/block-forge/pkg/etl/mocks"
	etl_ops "github.com/edward1christian/block-forge/pkg/etl/ops"
)

func TestStopETLOperation_Execute_Success(t *testing.T) {
	// Mock context
	ctx := &context.Context{}

	// Mock input data
	process := &etl.ETLProcess{
		Config: &etl.ETLConfig{
			Components: []*etl.ETLComponentConfig{
				{Name: "pipeline1"},
			},
		},
		Components: map[string]etl.ETLComponent{
			"pipeline1": &mocks.MockPipeline{},
		},
	}
	opInput := &system.OperationInput{
		Data: process,
	}

	// Mock pipeline stop
	mockPipeline := process.Components["pipeline1"].(*mocks.MockPipeline)
	mockPipeline.On("Stop", ctx).Return(nil)

	// Create operation
	operation := etl_ops.NewStopETLOperation()

	// Execute operation
	opOutput, err := operation.Execute(ctx, opInput)

	// Assertions
	assert.NoError(t, err, "Execute should not return an error")
	assert.NotNil(t, opOutput, "Operation output should not be nil")
	assert.NotNil(t, opOutput.Data, "Operation output data should not be nil")
	returnedProcess, ok := opOutput.Data.(*etl.ETLProcess)
	assert.True(t, ok, "Operation output data should be of type *etl.ETLProcess")
	assert.Equal(t, etl.ETLStatusStopped, returnedProcess.Status, "Process status should be stopped")
	assert.NotNil(t, returnedProcess.Components, "Process components map should be initialized")
}

func TestStopETLOperation_Execute_ComponentStopError(t *testing.T) {
	// Mock context
	ctx := &context.Context{}

	// Mock input data
	process := &etl.ETLProcess{
		Config: &etl.ETLConfig{
			Components: []*etl.ETLComponentConfig{
				{Name: "pipeline1"},
			},
		},
		Components: map[string]etl.ETLComponent{
			"pipeline1": &mocks.MockPipeline{},
		},
	}
	opInput := &system.OperationInput{
		Data: process,
	}

	// Mock pipeline stop with error
	mockPipeline := process.Components["pipeline1"].(*mocks.MockPipeline)
	mockPipeline.On("Stop", ctx).Return(errors.New("stop error"))

	// Create operation
	operation := etl_ops.NewStopETLOperation()

	// Execute operation
	opOutput, err := operation.Execute(ctx, opInput)

	// Assertions
	assert.Error(t, err, "Execute should return an error")
	assert.Nil(t, opOutput, "Operation output should be nil")
	assert.Equal(t, etl.ETLStatusRunning, process.Status, "Process status should remain running")
}
