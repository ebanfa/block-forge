package build_test

import (
	"errors"
	"testing"

	"github.com/edward1christian/block-forge/nova/pkg/build"
	"github.com/edward1christian/block-forge/nova/pkg/mocks"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestBuildStage_GetName tests the GetName method of BuildStageInterface.
func TestBuildStage_GetName(t *testing.T) {
	// Arrange
	expectedName := "TestStage"
	stage := build.NewBuildStage(expectedName)

	// Act
	name := stage.GetName()

	// Assert
	assert.Equal(t, expectedName, name, "Expected and actual names should match")
}

// TestBuildStage_GetTasks tests the GetTasks method of BuildStageInterface.
func TestBuildStage_GetTasks(t *testing.T) {
	// Arrange
	stage := build.NewBuildStage("TestStage")
	mockTask := &mocks.MockBuildTask{}

	// Mock behavior
	mockTask.On("GetName").Return("TestTask")
	stage.AddTask(mockTask)

	// Act
	tasks := stage.GetTasks()

	// Assert
	assert.NotNil(t, tasks, "Tasks should not be nil")
	assert.Len(t, tasks, 1, "Number of tasks should be 1")
}

// TestBuildStage_ExecuteTasks_Success tests the ExecuteTasks method of BuildStageInterface for success.
func TestBuildStage_ExecuteTasks_Success(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	stage := build.NewBuildStage("TestStage")
	mockTask := &mocks.MockBuildTask{}

	// Mock behavior
	mockTask.On("GetName").Return("TestTask")
	mockTask.On("Execute", ctx, mock.Anything).Return(&system.OperationOutput{}, nil)

	stage.AddTask(mockTask)

	// Act
	err := stage.ExecuteTasks(ctx)

	// Assert
	assert.NoError(t, err, "Executing tasks should not return an error")
	mockTask.AssertCalled(t, "Execute", mock.Anything, mock.Anything)
}

// TestBuildStage_ExecuteTasks_Error tests the ExecuteTasks method of BuildStageInterface for error.
func TestBuildStage_ExecuteTasks_Error(t *testing.T) {
	// Arrange
	stage := build.NewBuildStage("TestStage")
	mockTask := &mocks.MockBuildTask{}

	// Mock behavior
	mockTask.On("GetName").Return("TestTask")
	mockTask.On("Execute", mock.Anything, mock.Anything).Return(&system.OperationOutput{}, errors.New("execution error"))

	stage.AddTask(mockTask)

	// Act
	err := stage.ExecuteTasks(&context.Context{})

	// Assert
	assert.Error(t, err, "Executing tasks should return an error")
	mockTask.AssertCalled(t, "Execute", mock.Anything, mock.Anything)
}

// TestBuildStage_GetTaskByName_Success tests the GetTaskByName method of BuildStageInterface for success.
func TestBuildStage_GetTaskByName_Success(t *testing.T) {
	// Arrange
	stage := build.NewBuildStage("TestStage")
	mockTask := &mocks.MockBuildTask{}

	// Mock behavior
	mockTask.On("GetName").Return("TestTask")

	stage.AddTask(mockTask)

	// Act
	task, err := stage.GetTaskByName("TestTask")

	// Assert
	assert.NoError(t, err, "Getting task by name should not return an error")
	assert.Equal(t, mockTask, task, "Retrieved task should match the added task")
}

// TestBuildStage_GetTaskByName_Error tests the GetTaskByName method of BuildStageInterface for error.
func TestBuildStage_GetTaskByName_Error(t *testing.T) {
	// Arrange
	stage := build.NewBuildStage("TestStage")
	mockTask := &mocks.MockBuildTask{}

	// Mock behavior
	mockTask.On("GetName").Return("TestTask")
	stage.AddTask(mockTask)

	// Act
	_, err := stage.GetTaskByName("NonExistentTask")

	// Assert
	assert.Error(t, err, "Getting non-existent task should return an error")
}
