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

func TestBuilderService_Start_Success(t *testing.T) {
	// Initialize mock objects
	ctx := &context.Context{}
	factory := &mocks.MockBuilderFactory{}
	mockPipeline := &mocks.MockBuildPipeline{}
	mockBuilder := &mocks.MockPipelineBuilder{}

	mockPipeline.On("Execute", ctx).Return(nil)
	mockBuilder.On("Build").Return(mockPipeline, nil)
	factory.On("CreatePipelineBuilder", mock.Anything, mock.Anything).Return(mockBuilder, nil)

	// Create a new BuilderService instance
	builderService := build.NewBuilderService("id", "name", "description", factory)

	// Call Start method
	err := builderService.Start(ctx)

	// Assert that no error is returned
	assert.NoError(t, err)
}

func TestBuilderService_Start_FailedToCreateBuilder(t *testing.T) {
	// Initialize mock objects
	ctx := &context.Context{}
	factory := &mocks.MockBuilderFactory{}
	mockPipeline := &mocks.MockBuildPipeline{}
	mockBuilder := &mocks.MockPipelineBuilder{}

	mockPipeline.On("Execute", ctx).Return(nil)
	mockBuilder.On("Build").Return(mockPipeline, nil)
	factory.On("CreatePipelineBuilder", mock.Anything, mock.Anything).Return(
		mockBuilder, errors.New("failed to create pipeline builder"))

	// Create a new BuilderService instance
	builderService := build.NewBuilderService("id", "name", "description", factory)

	// Call Start method with a factory that returns nil
	err := builderService.Start(ctx)

	// Assert that an error is returned
	assert.Error(t, err)
	assert.Equal(t, "failed to create pipeline builder", err.Error())
}

func TestBuilderService_Start_FailedToBuildPipeline(t *testing.T) {
	// Initialize mock objects
	ctx := &context.Context{}
	factory := &mocks.MockBuilderFactory{}
	mockPipeline := &mocks.MockBuildPipeline{}
	mockBuilder := &mocks.MockPipelineBuilder{}

	mockBuilder.On("Build").Return(mockPipeline, errors.New("failed to build pipeline"))
	mockPipeline.On("Execute", ctx).Return(nil)
	factory.On("CreatePipelineBuilder", mock.Anything, mock.Anything).Return(mockBuilder, nil)

	// Create a new BuilderService instance
	builderService := build.NewBuilderService("id", "name", "description", factory)

	// Override the factory method to return a builder that returns nil when Build is called
	factory.On("ExecuteTasks", mock.Anything).Return(nil)

	// Call Start method
	err := builderService.Start(ctx)

	// Assert that an error is returned
	assert.Error(t, err)
	assert.Equal(t, "failed to build pipeline", err.Error())
}
