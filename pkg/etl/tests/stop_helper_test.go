package tests

import (
	"errors"
	"testing"

	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/etl"
	"github.com/edward1christian/block-forge/pkg/etl/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewETLProcessStopHelper(t *testing.T) {
	helper := etl.NewETLProcessStopHelper()
	assert.NotNil(t, helper, "NewETLProcessStopHelper should return a non-nil object")
}

func TestStopETLProcess_Success(t *testing.T) {
	// Create a mocks
	ctx := &context.Context{}
	mockPipeline1 := &mocks.MockPipeline{}
	mockPipeline1.On("Start", ctx).Return(nil)
	mockPipeline1.On("Stop", ctx).Return(nil)

	process := &etl.ETLProcess{
		Config: &etl.ETLConfig{
			Components: []*etl.ETLComponentConfig{
				{Name: "pipeline1"},
			},
		},
		Components: map[string]etl.ETLComponent{
			"pipeline1": mockPipeline1,
		},
	}

	// Create the helper
	helper := etl.NewETLProcessStopHelper()

	// Stop the process
	err := helper.StopETLProcess(ctx, process)

	// Assertions
	assert.NoError(t, err, "StopETLProcess should not return an error")
	assert.Equal(t, etl.ETLStatusStopped, process.Status, "Process should be in stopped status")
}

func TestStopETLProcess_Error(t *testing.T) {
	// Create a mocks
	ctx := &context.Context{}
	mockPipeline1 := &mocks.MockPipeline{}
	mockPipeline1.On("Stop", ctx).Return(errors.New(""))

	process := &etl.ETLProcess{
		Config: &etl.ETLConfig{
			Components: []*etl.ETLComponentConfig{
				{Name: "pipeline1"},
			},
		},
		Components: map[string]etl.ETLComponent{
			"pipeline1": mockPipeline1,
		},
		Status: etl.ETLStatusRunning,
	}
	// Create the helper
	helper := etl.NewETLProcessStopHelper()

	// Stop the process
	err := helper.StopETLProcess(ctx, process)

	// Assertions
	assert.Error(t, err, "StopETLProcess should return an error")
	assert.Equal(t, etl.ETLStatusRunning, process.Status, "Process should remain in running status")
}
