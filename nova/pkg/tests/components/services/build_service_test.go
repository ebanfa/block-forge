package services

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/edward1christian/block-forge/nova/pkg/components/services"
	novaMocksApi "github.com/edward1christian/block-forge/nova/pkg/mocks"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
)

func TestBuildService_Start_Success(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	mockSystem := &mocks.MockSystem{}
	mockPipeline := &novaMocksApi.MockPipeline{}
	mockRegistrar := &mocks.MockComponentRegistrar{}

	bs := services.NewBuildService("id", "name", "description")

	mockPipeline.On("Execute", ctx, mock.Anything).Return(nil)

	mockSystem.On("ComponentRegistry").Return(mockRegistrar)
	mockRegistrar.On("RegisterFactory", ctx, mock.Anything).Return(nil)
	mockRegistrar.On("CreateComponent", ctx, mock.Anything).Return(mockPipeline, nil)

	// Act
	err := bs.Initialize(ctx, mockSystem)
	assert.NoError(t, err)

	err = bs.Start(ctx)

	// Assert
	assert.NoError(t, err, "Starting BuildService should not return an error")
	mockSystem.AssertExpectations(t)
	mockRegistrar.AssertExpectations(t)
	mockPipeline.AssertExpectations(t)
}

func TestBuildService_Start_Error_CreatePipelineFailed(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	mockSystem := &mocks.MockSystem{}
	mockRegistrar := &mocks.MockComponentRegistrar{}
	mockPipeline := &novaMocksApi.MockPipeline{}

	bs := services.NewBuildService("id", "name", "description")

	mockSystem.On("ComponentRegistry").Return(mockRegistrar)
	mockRegistrar.On("RegisterFactory", ctx, mock.Anything).Return(nil)
	mockRegistrar.On("CreateComponent", ctx, mock.Anything).Return(mockPipeline, fmt.Errorf("failed to create pipeline"))

	// Act
	err := bs.Initialize(ctx, mockSystem)
	assert.NoError(t, err)

	err = bs.Start(ctx)

	// Assert
	assert.Error(t, err, "Starting BuildService should return an error when creating pipeline fails")
	mockSystem.AssertExpectations(t)
	mockRegistrar.AssertExpectations(t)
}

func TestBuildService_Stop_Success(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	mockSystem := &mocks.MockSystem{}
	mockRegistrar := &mocks.MockComponentRegistrar{}

	bs := services.NewBuildService("id", "name", "description")

	mockSystem.On("ComponentRegistry").Return(mockRegistrar)
	mockRegistrar.On("RemoveComponent", ctx, "pipeline").Return(nil)
	mockRegistrar.On("UnregisterFactory", ctx, "factory").Return(nil)
	mockRegistrar.On("RegisterFactory", ctx, mock.Anything).Return(nil)

	// Act
	err := bs.Initialize(ctx, mockSystem)
	assert.NoError(t, err)

	err = bs.Stop(ctx)

	// Assert
	assert.NoError(t, err, "Stopping BuildService should not return an error")
	mockSystem.AssertExpectations(t)
	mockRegistrar.AssertExpectations(t)
}

func TestBuildService_Stop_Error_RemoveComponentFailed(t *testing.T) {
	// Arrange
	mockSystem := &mocks.MockSystem{}
	ctx := &context.Context{}
	mockRegistrar := &mocks.MockComponentRegistrar{}

	bs := services.NewBuildService("id", "name", "description")

	mockSystem.On("ComponentRegistry").Return(mockRegistrar)
	mockRegistrar.On("RegisterFactory", ctx, mock.Anything).Return(nil)
	mockRegistrar.On("RemoveComponent", ctx, "pipeline").Return(fmt.Errorf("failed to remove pipeline"))

	// Act
	err := bs.Initialize(ctx, mockSystem)
	assert.NoError(t, err)

	err = bs.Stop(ctx)

	// Assert
	assert.Error(t, err, "Stopping BuildService should return an error when removing pipeline component fails")
	mockSystem.AssertExpectations(t)
	mockRegistrar.AssertExpectations(t)
}

func TestBuildService_Stop_Error_UnregisterFactoryFailed(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	mockSystem := &mocks.MockSystem{}
	mockRegistrar := &mocks.MockComponentRegistrar{}

	bs := services.NewBuildService("id", "name", "description")

	mockSystem.On("ComponentRegistry").Return(mockRegistrar)
	mockRegistrar.On("RemoveComponent", ctx, "pipeline").Return(nil)
	mockRegistrar.On("RegisterFactory", ctx, mock.Anything).Return(nil)
	mockRegistrar.On("UnregisterFactory", ctx, "factory").Return(fmt.Errorf("failed to unregister factory"))

	// Act
	err := bs.Initialize(ctx, mockSystem)
	assert.NoError(t, err)

	err = bs.Stop(ctx)

	// Assert
	assert.Error(t, err, "Stopping BuildService should return an error when unregistering pipeline factory fails")
	mockSystem.AssertExpectations(t)
	mockRegistrar.AssertExpectations(t)
}
