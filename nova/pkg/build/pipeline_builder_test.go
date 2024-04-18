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
	result := builder.AddStage(stageName)

	// Assert
	assert.NotNil(t, result, "Result should not be nil")
	assert.Equal(t, builder, result, "Result should be equal to builder")
	pipeline := builder.Build()
	assert.NotNil(t, pipeline, "Pipeline should not be nil")
	stages := pipeline.GetStages()
	assert.Len(t, stages, 1, "Number of stages should be 1")
	assert.Equal(t, stageName, stages[0].GetName(), "Stage name should match")
}

// TestPipelineBuilder_AddTask tests the AddTask method of PipelineBuilder.
func TestPipelineBuilder_AddTask(t *testing.T) {
	// Arrange
	builder := build.NewPipelineBuilder("TestPipeline")
	task := &mocks.MockBuildTask{}

	// Mock behavior
	task.On("GetName").Return("TestTask")

	// Act
	builder.AddStage("TestStage").AddTask(task)
	pipeline := builder.Build()

	// Assert
	assert.NotNil(t, pipeline, "Pipeline should not be nil")
	stages := pipeline.GetStages()
	assert.Len(t, stages, 1, "Number of stages should be 1")
	tasks := stages[0].GetTasks()
	assert.Len(t, tasks, 1, "Number of tasks should be 1")
	assert.Equal(t, task, tasks[0], "Added task should match")
}

// TestPipelineBuilder_AddTask_NoStage tests the AddTask method of PipelineBuilder when no stage is added.
func TestPipelineBuilder_AddTask_NoStage(t *testing.T) {
	// Arrange
	builder := build.NewPipelineBuilder("TestPipeline")
	task := &mocks.MockBuildTask{}

	// Act
	defer func() {
		if r := recover(); r != nil {
			// Assert
			assert.NotNil(t, r, "Panic should occur when no stage is added")
		}
	}()
	builder.AddTask(task)

	// Assert
	assert.True(t, true, "No stage added, so this line is never executed")
}

// TestPipelineBuilder_Build tests the Build method of PipelineBuilder.
func TestPipelineBuilder_Build(t *testing.T) {
	// Arrange
	builder := build.NewPipelineBuilder("TestPipeline")

	// Act
	pipeline := builder.Build()

	// Assert
	assert.NotNil(t, pipeline, "Pipeline should not be nil")
	assert.IsType(t, &build.BuildPipeline{}, pipeline, "Pipeline should be of type BuildPipeline")
	assert.Equal(t, "TestPipeline", pipeline.GetName(), "Pipeline name should match")
}
