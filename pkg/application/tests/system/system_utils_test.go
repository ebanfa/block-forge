package system

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// TestStartService_Success tests the system.StartService function for successful service start.
func TestStartService_Success(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	mockSystem := &mocks.MockSystem{}
	mockConfig := &config.ComponentConfig{}
	mockComponent := &mocks.MockSystemService{}
	mockRegistrar := &mocks.MockComponentRegistrar{}

	mockComponent.On("Start", ctx).Return(nil)
	mockComponent.On("ID").Return("mockService")
	mockComponent.On("Initialize", ctx, mockSystem).Return(nil)

	mockSystem.On("ComponentRegistry").Return(mockRegistrar)

	mockRegistrar.On("CreateComponent", ctx, mock.Anything).Return(mockComponent, nil)

	// Act
	err := system.StartService(ctx, mockSystem, mockConfig)

	// Assert
	assert.NoError(t, err, "Starting service should not return an error")
}

// TestStartService_CreateComponentError tests the system.StartService function for error when creating component.
func TestStartService_CreateComponentError(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	mockSystem := &mocks.MockSystem{}
	mockConfig := &config.ComponentConfig{}
	mockComponent := &mocks.MockSystemService{}
	mockRegistrar := &mocks.MockComponentRegistrar{}
	expectedErr := errors.New("create component error")

	mockSystem.On("ComponentRegistry").Return(mockRegistrar)

	mockRegistrar.On("CreateComponent", ctx, mock.Anything).Return(mockComponent, expectedErr)

	// Act
	err := system.StartService(ctx, mockSystem, mockConfig)

	// Assert
	assert.EqualError(t, err, fmt.Sprintf("failed to start service. Could not create component %s", mockConfig.ID), "Starting service with create component error should return an error")
}

// TestStartService_InitializeError tests the system.StartService function for error during service initialization.
func TestStartService_InitializeError(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	mockSystem := &mocks.MockSystem{}
	mockConfig := &config.ComponentConfig{}
	mockComponent := &mocks.MockSystemService{}
	mockRegistrar := &mocks.MockComponentRegistrar{}
	expectedErr := errors.New("initialize error")

	mockComponent.On("ID").Return("mockService")
	mockComponent.On("Initialize", ctx, mockSystem).Return(expectedErr)

	mockSystem.On("ComponentRegistry").Return(mockRegistrar)
	mockRegistrar.On("CreateComponent", ctx, mockConfig).Return(mockComponent, nil)

	// Act
	err := system.StartService(ctx, mockSystem, mockConfig)

	// Assert
	assert.EqualError(t, err, fmt.Sprintf("failed to initialize service: %s %v", mockComponent.ID(), expectedErr), "Starting service with initialization error should return an error")
}

// TestStartService_StartError tests the system.StartService function for error during service start.
func TestStartService_StartError(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	mockSystem := &mocks.MockSystem{}
	mockConfig := &config.ComponentConfig{}
	mockComponent := &mocks.MockSystemService{}
	mockRegistrar := &mocks.MockComponentRegistrar{}
	expectedErr := errors.New("start error")

	mockComponent.On("Start", ctx).Return(expectedErr)
	mockComponent.On("ID").Return("mockService")
	mockComponent.On("Initialize", ctx, mockSystem).Return(nil)

	mockSystem.On("ComponentRegistry").Return(mockRegistrar)
	mockRegistrar.On("CreateComponent", ctx, mock.Anything).Return(mockComponent, nil)

	// Act
	err := system.StartService(ctx, mockSystem, mockConfig)

	// Assert
	assert.EqualError(t, err, expectedErr.Error(), "Starting service with start error should return an error")
}

// TestStopService_Success tests the system.StopService function for successful service stop.
func TestStopService_Success(t *testing.T) {
	// Arrange
	mockID := "mockService"
	ctx := &context.Context{}
	mockSystem := &mocks.MockSystem{}
	mockComponent := &mocks.MockSystemService{}
	mockRegistrar := &mocks.MockComponentRegistrar{}

	mockComponent.On("Stop", ctx).Return(nil)
	mockComponent.On("Start", ctx).Return(nil)
	mockComponent.On("ID").Return("mockService")
	mockComponent.On("Initialize", ctx, mockSystem).Return(nil)

	mockSystem.On("ComponentRegistry").Return(mockRegistrar)
	mockRegistrar.On("GetComponent", mock.Anything).Return(mockComponent, nil)
	mockRegistrar.On("CreateComponent", ctx, mock.Anything).Return(mockComponent, nil)

	// Act
	err := system.StopService(ctx, mockSystem, mockID)

	// Assert
	assert.NoError(t, err, "Stopping service should not return an error")
}

// TestStopService_ComponentNotFoundError tests the system.StopService function for error when component not found.
func TestStopService_ComponentNotFoundError(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	mockSystem := &mocks.MockSystem{}
	mockComponent := &mocks.MockSystemService{}
	mockRegistrar := &mocks.MockComponentRegistrar{}
	mockID := "nonExistentService"
	expectedErr := fmt.Errorf("failed to stop build service. Service not found: %v", errors.New("component not found"))

	mockSystem.On("ComponentRegistry").Return(mockRegistrar)
	mockRegistrar.On("GetComponent", mockID).Return(mockComponent, errors.New("component not found"))

	// Act
	err := system.StopService(ctx, mockSystem, mockID)

	// Assert
	assert.EqualError(t, err, expectedErr.Error(), "Stopping service with component not found error should return an error")
}

// TestStopService_ServiceInterfaceError tests the system.StopService function for error when component does not implement SystemServiceInterface.
func TestStopService_ServiceInterfaceError(t *testing.T) {
	// Arrange
	mockID := "mockService"
	ctx := &context.Context{}
	mockSystem := &mocks.MockSystem{}
	mockComponent := &mocks.MockComponent{}
	mockRegistrar := &mocks.MockComponentRegistrar{}
	expectedErr := errors.New("failed to stop service. Service component is not a system service")

	mockComponent.On("ID").Return("mockService")
	mockComponent.On("Initialize", ctx, mockSystem).Return(nil)

	mockSystem.On("ComponentRegistry").Return(mockRegistrar)
	mockRegistrar.On("GetComponent", mockID).Return(mockComponent, nil)
	mockRegistrar.On("CreateComponent", ctx, mock.Anything).Return(mockComponent, nil)

	// Act
	err := system.StopService(ctx, mockSystem, mockID)

	// Assert
	assert.EqualError(t, err, expectedErr.Error(), "Stopping service with non-system service component should return an error")
}

// TestRegisterComponent_Success tests the RegisterComponent function for success.
func TestRegisterComponent_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockLogger := &mocks.MockLogger{}
	mockSystem := &mocks.MockSystem{}
	mockComponent := &mocks.MockComponent{}
	mockFactory := &mocks.MockComponentFactory{}
	mockRegistry := &mocks.MockComponentRegistrar{}
	mockConfig := &config.ComponentConfig{
		ID:        "TestComponent",
		FactoryID: "TestFactory",
	}

	// Mock behavior
	mockSystem.On("Logger").Return(mockLogger)
	mockSystem.On("ComponentRegistry").Return(mockRegistry)
	mockRegistry.On("RegisterFactory", ctx, mockConfig.FactoryID, mockFactory).Return(nil)
	mockRegistry.On("CreateComponent", ctx, mockConfig).Return(mockComponent, nil)

	// Act
	err := system.RegisterComponent(ctx, mockSystem, mockConfig, mockFactory)

	// Assert
	assert.NoError(t, err, "RegisterComponent should not return an error")
}

// TestRegisterComponent_Error_RegisterFactory tests the RegisterComponent function for error during factory registration.
func TestRegisterComponent_Error_RegisterFactory(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockLogger := &mocks.MockLogger{}
	mockSystem := &mocks.MockSystem{}
	mockFactory := &mocks.MockComponentFactory{}
	mockRegistry := &mocks.MockComponentRegistrar{}
	mockConfig := &config.ComponentConfig{
		ID:        "TestComponent",
		FactoryID: "TestFactory",
	}

	// Mock behavior - RegisterFactory returns an error
	mockSystem.On("Logger").Return(mockLogger)
	mockSystem.On("ComponentRegistry").Return(mockRegistry)
	mockRegistry.On("RegisterFactory", ctx, mockConfig.FactoryID, mockFactory).Return(errors.New("error registering factory"))

	// Act
	err := system.RegisterComponent(ctx, mockSystem, mockConfig, mockFactory)

	// Assert
	assert.Error(t, err, "RegisterComponent should return an error")
	assert.Contains(t, err.Error(), "failed to register component factory", "Error message should indicate failure in registering factory")
}

// TestRegisterComponent_Error_CreateComponent tests the RegisterComponent function for error during component creation.
func TestRegisterComponent_Error_CreateComponent(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockLogger := &mocks.MockLogger{}
	mockSystem := &mocks.MockSystem{}
	mockComponent := &mocks.MockComponent{}
	mockFactory := &mocks.MockComponentFactory{}
	mockRegistry := &mocks.MockComponentRegistrar{}
	mockConfig := &config.ComponentConfig{
		ID:        "TestComponent",
		FactoryID: "TestFactory",
	}

	// Mock behavior - CreateComponent returns an error
	mockSystem.On("Logger").Return(mockLogger)
	mockSystem.On("ComponentRegistry").Return(mockRegistry)
	mockRegistry.On("RegisterFactory", ctx, mockConfig.FactoryID, mockFactory).Return(nil)
	mockRegistry.On("CreateComponent", ctx, mockConfig).Return(mockComponent, errors.New("error creating component"))

	// Act
	err := system.RegisterComponent(ctx, mockSystem, mockConfig, mockFactory)

	// Assert
	assert.Error(t, err, "RegisterComponent should return an error")
	assert.Contains(t, err.Error(), "failed to create component", "Error message should indicate failure in creating component")
}
