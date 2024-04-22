package build_test

import (
	"testing"

	"github.com/edward1christian/block-forge/nova/pkg/build"
	"github.com/edward1christian/block-forge/nova/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

// TestPipelineBuilder_AddStage tests the AddStage method of PipelineBuilder.
func TestPipelineBuilder_AddStage(t *testing.T) {
	// Arrange
	builder := build.NewPipelineBuilder("TestPipeline")
	stageName := "TestStage"

	// Act
	_, err := builder.AddStage(stageName)

	// Assert
	assert.NoError(t, err, "AddStage should not return an error")

	pipeline, err := builder.Build()
	assert.NoError(t, err, "Build should not return an error")
	assert.NotNil(t, pipeline, "Pipeline should not be nil")

	stages := pipeline.GetStages()
	assert.Len(t, stages, 1, "Number of stages should be 1")
}

// TestPipelineBuilder_AddStage_AddTask tests the AddStage and AddTask methods of PipelineBuilder.
func TestPipelineBuilder_AddStage_AddTask(t *testing.T) {
	// Arrange
	builder := build.NewPipelineBuilder("TestPipeline")
	stageName := "TestStage"
	task := &mocks.MockBuildTask{}
	task.On("GetName").Return("TestTask")

	// Act
	_, err := builder.AddStage(stageName)
	assert.NoError(t, err, "AddStage should not return an error")

	_, err = builder.AddTask(task)
	assert.NoError(t, err, "AddTask should not return an error")

	pipeline, err := builder.Build()

	// Assert
	assert.NoError(t, err, "Build should not return an error")
	assert.NotNil(t, pipeline, "Pipeline should not be nil")
}

// TestPipelineBuilder_AddTask_NoStage tests the AddTask method of PipelineBuilder when no stage is added.
func TestPipelineBuilder_AddTask_NoStage(t *testing.T) {
	// Arrange
	builder := build.NewPipelineBuilder("TestPipeline")
	task := &mocks.MockBuildTask{}

	// Act
	_, err := builder.AddTask(task)

	// Assert
	assert.Error(t, err, "AddTask should return an error when no stage is added")
}

// TestPipelineBuilder_Build tests the Build method of PipelineBuilder.
func TestPipelineBuilder_Build(t *testing.T) {
	// Arrange
	builder := build.NewPipelineBuilder("TestPipeline")
	stageName := "TestStage"
	task := &mocks.MockBuildTask{}
	task.On("GetName").Return("TestTask")

	// Act
	_, err := builder.AddStage(stageName)
	assert.NoError(t, err, "AddStage should not return an error")

	_, err = builder.AddTask(task)
	assert.NoError(t, err, "AddTask should not return an error")

	// Act
	pipeline, err := builder.Build()

	// Assert
	assert.NoError(t, err, "Build should not return an error")
	assert.NotNil(t, pipeline, "Pipeline should not be nil")
	assert.IsType(t, &build.BuildPipeline{}, pipeline, "Pipeline should be of type BuildPipeline")
	assert.Equal(t, "TestPipeline", pipeline.GetName(), "Pipeline name should match")
}
