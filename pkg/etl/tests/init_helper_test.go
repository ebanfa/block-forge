package tests

import (
	"testing"

	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/etl"
	"github.com/edward1christian/block-forge/pkg/etl/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewETLProcessInitHelper(t *testing.T) {
	helper := etl.NewETLProcessInitHelper(nil)
	assert.NotNil(t, helper, "NewETLProcessInitHelper should return a non-nil object")
}

func TestCreateETLProcess_Success(t *testing.T) {
	// Mock context
	ctx := &context.Context{}

	// Create a mock ETLConfig
	config := &etl.ETLConfig{
		Components: []*etl.ETLComponentConfig{
			{Name: "Component1"},
			// Add more entries as needed
		},
	}

	// Create the helper
	generator := etl.NewProcessIDGenerator("etl")
	helper := etl.NewETLProcessInitHelper(generator)

	// Create the process
	process, err := helper.CreateETLProcess(ctx, config)

	// Assertions
	assert.NoError(t, err, "CreateETLProcess should not return an error")
	assert.NotNil(t, process, "CreateETLProcess should return a non-nil process object")
	assert.NotEmpty(t, process.ID, "Process ID should not be empty")
	assert.Equal(t, etl.ETLStatusInitialized, process.Status, "Process should be in initialized status")
	assert.NotNil(t, process.Components, "Process components map should be initialized")
}

func TestCreateETLProcess_Error(t *testing.T) {
	// Mock context
	ctx := &context.Context{}

	// Create a mock ETLConfig
	config := &etl.ETLConfig{}

	// Create the helper with a mock GenerateProcessID function that returns an error
	generator := etl.NewProcessIDGenerator("etl")
	helper := etl.NewETLProcessInitHelper(generator)

	// Create the process
	process, err := helper.CreateETLProcess(ctx, config)

	// Assertions
	assert.Error(t, err, "CreateETLProcess should return an error")
	assert.Nil(t, process, "CreateETLProcess should return nil for the process object")
}

func TestRollbackComponents(t *testing.T) {
	// Create mocks
	ctx := &context.Context{}
	mockPipeline1 := &mocks.MockPipeline{}
	mockPipeline1.On("Stop", ctx).Return(nil)

	// Create a mock ETLProcess with components
	process := &etl.ETLProcess{
		Components: map[string]etl.ETLComponent{
			"pipeline1": mockPipeline1,
		},
	}

	// Create the helper
	generator := etl.NewProcessIDGenerator("etl")
	helper := etl.NewETLProcessInitHelper(generator)

	// Rollback components
	helper.RollbackComponents(ctx, process)

	// Assertions: Ensure Stop() is called on each component
	for _, component := range process.Components {
		_, ok := component.(*mocks.MockPipeline)
		assert.True(t, ok, "Component should be a MockPipeline")
		mockPipeline1.AssertExpectations(t)
	}
}
