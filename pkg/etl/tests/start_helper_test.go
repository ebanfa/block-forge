package tests

import (
	"errors"
	"testing"

	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/etl"
	"github.com/edward1christian/block-forge/pkg/etl/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewETLProcessStartHelper(t *testing.T) {
	helper := etl.NewETLProcessStartHelper(nil)
	assert.NotNil(t, helper, "NewETLProcessStartHelper should return a non-nil object")
}

func TestStartETLProcess_Success(t *testing.T) {
	// Create a mocks
	ctx := &context.Context{}
	mockPipeline1 := &mocks.MockPipeline{}
	mockPipeline1.On("Start", ctx).Return(nil)

	process := &etl.ETLProcess{
		Components: map[string]etl.ETLComponent{
			"pipeline1": mockPipeline1,
		},
	}
	// Create the helper
	helper := etl.NewETLProcessStartHelper(nil)

	// Start the process
	err := helper.StartETLProcess(ctx, process)

	// Assertions
	assert.NoError(t, err, "StartETLProcess should not return an error")
	assert.Equal(t, etl.ETLStatusRunning, process.Status, "Process should be in running status")
}

func TestStartETLProcess_Error(t *testing.T) {
	// Create a mocks
	ctx := &context.Context{}
	mockPipeline1 := &mocks.MockPipeline{}

	mockPipeline1.On("Start", ctx).Return(errors.New(""))
	mockPipeline1.On("Stop", ctx).Return(nil)

	process := &etl.ETLProcess{
		Components: map[string]etl.ETLComponent{
			"pipeline1": mockPipeline1,
		},
		Status: etl.ETLStatusInitialized,
	}

	// Create the helper
	helper := etl.NewETLProcessStartHelper(nil)

	// Start the process
	err := helper.StartETLProcess(ctx, process)

	// Assertions
	assert.Error(t, err, "StartETLProcess should return an error")
	assert.Equal(t, etl.ETLStatusInitialized, process.Status, "Process should remain in initialized status")
}
