package common

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/edward1christian/block-forge/nova/pkg/components/common"
	"github.com/edward1christian/block-forge/nova/pkg/mocks"
	typesApi "github.com/edward1christian/block-forge/nova/pkg/types"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	mocksApi "github.com/edward1christian/block-forge/pkg/application/mocks"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
)

func TestPipeline_AddStage_Success(t *testing.T) {
	// Arrange
	mockStage := &mocks.MockStage{}
	pipeline := common.NewPipeline("id", "name", "description")

	// Act
	err := pipeline.AddStage("Stage1", mockStage)

	// Assert
	assert.NoError(t, err, "Adding stage should succeed")
}

func TestPipeline_AddStage_Error(t *testing.T) {
	// Arrange
	mockStage := &mocks.MockStage{}
	pipeline := common.NewPipeline("id", "name", "description")
	pipeline.AddStage("Stage1", mockStage) // Adding a stage to simulate existing stage

	// Act
	err := pipeline.AddStage("Stage1", mockStage)

	// Assert
	assert.Error(t, err, "Adding existing stage should return an error")
	assert.EqualError(t, err, "stage already exists", "Error message should indicate stage already exists")
}

func TestPipeline_GetStage_Success(t *testing.T) {
	// Arrange
	mockStage := &mocks.MockStage{}
	pipeline := common.NewPipeline("id", "name", "description")
	pipeline.AddStage("Stage1", mockStage)

	// Act
	stage, err := pipeline.GetStage("Stage1")

	// Assert
	assert.NoError(t, err, "Getting existing stage should succeed")
	assert.Equal(t, mockStage, stage, "Retrieved stage should match the added stage")
}

func TestPipeline_GetStage_Error(t *testing.T) {
	// Arrange
	pipeline := common.NewPipeline("id", "name", "description")

	// Act
	_, err := pipeline.GetStage("NonExistentStage")

	// Assert
	assert.Error(t, err, "Getting non-existent stage should return an error")
	assert.EqualError(t, err, "stage not found", "Error message should indicate stage not found")
}

func TestPipeline_GetStages(t *testing.T) {
	// Arrange
	mockStage := &mocks.MockStage{}
	pipeline := common.NewPipeline("id", "name", "description")
	pipeline.AddStage("Stage1", mockStage)

	// Act
	stages := pipeline.GetStages()

	// Assert
	assert.Len(t, stages, 1, "Number of stages retrieved should be 1")
	assert.Equal(t, mockStage, stages[0], "Retrieved stage should match the added stage")
}

func TestPipeline_Execute_Success(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	mockStage := &mocks.MockStage{}
	mockComponent := &mocks.MockTask{}
	mockSystem := &mocksApi.MockSystem{}
	mockOperationInput := &systemApi.SystemOperationInput{}

	pipeline := common.NewPipeline("id", "name", "description")
	pipeline.AddStage("Stage1", mockStage)

	// Mock behavior
	mockComponent.On("ID").Return("mockComponentId")
	mockStage.On("GetTasks").Return([]typesApi.TaskInterface{mockComponent})
	mockSystem.On("ExecuteOperation", ctx, mock.Anything, mock.Anything).Return(nil, nil)

	// Act
	err := pipeline.Initialize(ctx, mockSystem)
	assert.NoError(t, err)

	err = pipeline.Execute(ctx, mockOperationInput)

	// Assert
	assert.NoError(t, err, "Executing pipeline should succeed")
	mockStage.AssertExpectations(t)
}

func TestPipeline_Execute_Error(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	mockStage := &mocks.MockStage{}
	mockComponent := &mocks.MockTask{}
	mockSystem := &mocksApi.MockSystem{}
	mockOperationInput := &systemApi.SystemOperationInput{}

	pipeline := common.NewPipeline("id", "name", "description")
	pipeline.AddStage("Stage1", mockStage)

	// Mock behavior
	mockComponent.On("ID").Return("mockComponentId")
	mockStage.On("GetTasks").Return([]typesApi.TaskInterface{mockComponent})
	mockSystem.On("ExecuteOperation", ctx, mock.Anything, mock.Anything).Return(nil, errors.New("task execution failed"))

	// Act
	err := pipeline.Initialize(ctx, mockSystem)
	assert.NoError(t, err)

	err = pipeline.Execute(ctx, mockOperationInput)

	// Assert
	assert.Error(t, err, "Executing pipeline with failing task should return an error")
	assert.EqualError(t, err, "task execution failed", "Error message should indicate task execution failure")
	mockStage.AssertExpectations(t)
}
