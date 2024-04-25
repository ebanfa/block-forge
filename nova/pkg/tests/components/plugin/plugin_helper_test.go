package plugin

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/edward1christian/block-forge/nova/pkg/common"
	"github.com/edward1christian/block-forge/nova/pkg/components/plugin"
	"github.com/edward1christian/block-forge/nova/pkg/components/services"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
)

// TestRegisterServices_Success tests the RegisterServices function for successful service registration.
func TestRegisterServices_Success(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	mockSystem := &mocks.MockSystem{}

	mockRegistrar := &mocks.MockComponentRegistrar{}
	mockRegistrar.On("RegisterFactory", ctx, mock.Anything, mock.Anything).Return(nil)

	mockSystem.On("ComponentRegistry").Return(mockRegistrar)

	// Act
	err := plugin.RegisterServices(ctx, mockSystem)

	// Assert
	assert.NoError(t, err, "Registering services should not return an error")
}

// TestRegisterServices_BuildServiceFactoryError tests the RegisterServices function for error when registering build service factory.
func TestRegisterServices_BuildServiceFactoryError(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	mockSystem := &mocks.MockSystem{}
	expectedErr := errors.New("register build service factory error")

	mockRegistrar := &mocks.MockComponentRegistrar{}
	mockRegistrar.On("RegisterFactory",
		ctx, common.BuildServiceFactory, &services.BuildServiceFactory{}).Return(expectedErr)

	mockSystem.On("ComponentRegistry").Return(mockRegistrar)

	// Act
	err := plugin.RegisterServices(ctx, mockSystem)

	// Assert
	assert.ErrorContains(t, err, "failed to register build service")
}

// TestRegisterOperations_Success tests the RegisterOperations function for successful registration.
func TestRegisterOperations_Success(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	mockSystem := &mocks.MockSystem{}
	mockRegistrar := &mocks.MockComponentRegistrar{}

	mockSystem.On("ComponentRegistry").Return(mockRegistrar)
	mockRegistrar.On("RegisterFactory",
		ctx, common.BuildServiceFactory, &services.BuildServiceFactory{}).Return(nil)

	// Act
	err := plugin.RegisterOperations(ctx, mockSystem)

	// Assert
	assert.NoError(t, err, "Registering operations should not return an error")
}

// TestRegisterBuildOperations_Success tests the RegisterBuildOperations function for successful registration.
func TestRegisterBuildOperations_Success(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	mockSystem := &mocks.MockSystem{}

	mockRegistrar := &mocks.MockComponentRegistrar{}
	mockRegistrar.On("RegisterFactory",
		ctx, common.BuildServiceFactory, &services.BuildServiceFactory{}).Return(nil).Once()

	mockSystem.On("ComponentRegistry").Return(mockRegistrar)

	// Act
	err := plugin.RegisterBuildOperations(ctx, mockSystem)

	// Assert
	assert.NoError(t, err, "Registering build operations should not return an error")
}

// TestStartServices_Success tests the StartServices function for successful service start.
func TestStartServices_Success(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	mockSystem := &mocks.MockSystem{}
	mockComponent := &mocks.MockSystemService{}
	mockRegistrar := &mocks.MockComponentRegistrar{}

	mockComponent.On("Start", ctx).Return(nil)
	mockComponent.On("Initialize", ctx, mock.Anything).Return(nil)

	mockSystem.On("ComponentRegistry").Return(mockRegistrar)
	mockSystem.On("StartService", ctx, mock.Anything, mock.Anything).Return(nil).Twice()
	mockRegistrar.On("CreateComponent", ctx, mock.Anything).Return(mockComponent, nil)
	// Act
	err := plugin.StartServices(ctx, mockSystem)

	// Assert
	assert.NoError(t, err, "Starting services should not return an error")
}

// TestStopServices_Success tests the StopServices function for successful service stop.
func TestStopServices_Success(t *testing.T) {
	// Arrange
	// Arrange
	ctx := &context.Context{}
	mockSystem := &mocks.MockSystem{}
	mockComponent := &mocks.MockSystemService{}
	mockRegistrar := &mocks.MockComponentRegistrar{}

	mockComponent.On("Stop", ctx).Return(nil)

	mockSystem.On("ComponentRegistry").Return(mockRegistrar)
	mockSystem.On("StopService", ctx, mock.Anything, mock.Anything).Return(nil).Twice()

	mockRegistrar.On("GetComponent", mock.Anything).Return(mockComponent, nil)

	// Act
	err := plugin.StopServices(ctx, mockSystem)

	// Assert
	assert.NoError(t, err, "Stopping services should not return an error")
}
