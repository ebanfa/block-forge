package build_test

import (
	"errors"
	"testing"

	"github.com/edward1christian/block-forge/nova/pkg/build"
	"github.com/edward1christian/block-forge/nova/pkg/mocks"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestBuildPipeline_GetName tests the GetName method of BuildPipelineInterface.
func TestBuildPipeline_GetName(t *testing.T) {
	// Arrange
	expectedName := "TestPipeline"
	pipeline := build.NewBuildPipeline(expectedName)

	// Act
	name := pipeline.GetName()

	// Assert
	assert.Equal(t, expectedName, name, "Expected and actual names should match")
}

// TestBuildPipeline_AddStage_Success tests the AddStage method of BuildPipelineInterface for success.
func TestBuildPipeline_AddStage_Success(t *testing.T) {
	// Arrange
	mockPipeline := build.NewBuildPipeline("testPipeline")
	mockStage := &mocks.MockBuildStage{}

	// Mock behavior
	mockPipeline.AddStage("Stage1", mockStage)

	// Act
	err := mockPipeline.AddStage("Stage2", mockStage)

	// Assert
	assert.NoError(t, err, "Adding stage should not return an error")
}

// TestBuildPipeline_AddStage_Error tests the AddStage method of BuildPipelineInterface for error when adding duplicate stage.
func TestBuildPipeline_AddStage_Error(t *testing.T) {
	// Arrange
	mockPipeline := build.NewBuildPipeline("testPipeline")
	mockStage := &mocks.MockBuildStage{}

	// Mock behavior
	mockPipeline.AddStage("Stage1", mockStage)

	// Act
	err := mockPipeline.AddStage("Stage1", mockStage)

	// Assert
	assert.Error(t, err, "Adding stage with same name should return an error")
}

// TestBuildPipeline_GetStage_Success tests the GetStage method of BuildPipelineInterface for success.
func TestBuildPipeline_GetStage_Success(t *testing.T) {
	// Arrange
	mockPipeline := build.NewBuildPipeline("testPipeline")
	mockStage := &mocks.MockBuildStage{}

	// Mock behavior
	mockPipeline.AddStage("Stage1", mockStage)

	// Act
	stage, err := mockPipeline.GetStage("Stage1")

	// Assert
	assert.NoError(t, err, "Getting stage should not return an error")
	assert.NotNil(t, stage, "Stage should not be nil")
}

// TestBuildPipeline_GetStage_Error tests the GetStage method of BuildPipelineInterface for error when getting non-existent stage.
func TestBuildPipeline_GetStage_Error(t *testing.T) {
	// Arrange
	mockPipeline := build.NewBuildPipeline("testPipeline")
	mockStage := &mocks.MockBuildStage{}

	// Mock behavior
	mockPipeline.AddStage("Stage1", mockStage)

	// Act
	_, err := mockPipeline.GetStage("NonExistentStage")

	// Assert
	assert.Error(t, err, "Getting non-existent stage should return an error")
}

// TestBuildPipeline_GetStages tests the GetStages method of BuildPipelineInterface.
func TestBuildPipeline_GetStages(t *testing.T) {
	// Arrange
	mockPipeline := build.NewBuildPipeline("testPipeline")
	mockStage := &mocks.MockBuildStage{}

	// Mock behavior
	mockPipeline.AddStage("Stage1", mockStage)

	// Act
	stages := mockPipeline.GetStages()

	// Assert
	assert.NotNil(t, stages, "Stages should not be nil")
	assert.Len(t, stages, 1, "Number of stages should be 1")
}

// TestBuildPipeline_Execute_Success tests the Execute method of BuildPipelineInterface for success.
func TestBuildPipeline_Execute_Success(t *testing.T) {
	// Arrange
	mockPipeline := build.NewBuildPipeline("testPipeline")
	mockStage := &mocks.MockBuildStage{}

	// Mock behavior
	mockPipeline.AddStage("Stage1", mockStage)
	mockStage.On("ExecuteTasks", mock.Anything).Return(nil)

	// Act
	err := mockPipeline.Execute(&context.Context{})

	// Assert
	assert.NoError(t, err, "Executing pipeline should not return an error")
	mockStage.AssertCalled(t, "ExecuteTasks", mock.Anything)
}

// TestBuildPipeline_Execute_Error tests the Execute method of BuildPipelineInterface for error.
func TestBuildPipeline_Execute_Error(t *testing.T) {
	// Arrange
	mockPipeline := build.NewBuildPipeline("testPipeline")
	mockStage := &mocks.MockBuildStage{}

	// Mock behavior
	mockPipeline.AddStage("Stage1", mockStage)
	mockStage.On("ExecuteTasks", mock.Anything).Return(errors.New("execution error"))

	// Act
	err := mockPipeline.Execute(&context.Context{})

	// Assert
	assert.Error(t, err, "Executing pipeline should return an error")
	mockStage.AssertCalled(t, "ExecuteTasks", mock.Anything)
}
