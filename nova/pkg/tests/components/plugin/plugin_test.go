package plugin

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/edward1christian/block-forge/nova/pkg/common"
	"github.com/edward1christian/block-forge/nova/pkg/components/plugin"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
)

func TestNovaPlugin_Initialize(t *testing.T) {
	// Initialize NovaPlugin instance
	p := plugin.NewNovaPlugin()

	// Mock SystemInterface
	mockSystem := &mocks.MockSystem{}
	mockSystem.On("Logger").Return(&mocks.MockLogger{})

	// Mock Context
	ctx := &context.Context{}

	// Test initialization
	err := p.Initialize(ctx, mockSystem)
	assert.NoError(t, err)
}

func TestNovaPlugin_RegisterResources_Success(t *testing.T) {

	ctx := &context.Context{}
	// Initialize NovaPlugin instance
	p := plugin.NewNovaPlugin()

	// Mock SystemInterface
	mockSystem := &mocks.MockSystem{}
	mockSystem.On("Logger").Return(&mocks.MockLogger{})

	// Mock ComponentRegistry
	registrarMock := &mocks.MockComponentRegistrar{}

	// Expectations for registering resources
	registrarMock.On("RegisterFactory", ctx, mock.Anything, mock.Anything).Return(nil)
	registrarMock.On("CreateComponent", ctx, mock.Anything).Return(&mocks.MockOperation{}, nil)

	// Set the mocked ComponentRegistry
	mockSystem.On("ComponentRegistry").Return(registrarMock)

	// Initialize plugin
	err := p.Initialize(ctx, mockSystem)
	assert.NoError(t, err)

	// Test resource registration
	err = p.RegisterResources(ctx)
	assert.NoError(t, err)
}

func TestNovaPlugin_RegisterResources_Failure_RegisterFactory(t *testing.T) {
	ctx := &context.Context{}
	// Initialize NovaPlugin instance
	p := plugin.NewNovaPlugin()

	// Mock SystemInterface
	mockSystem := &mocks.MockSystem{}
	mockSystem.On("Logger").Return(&mocks.MockLogger{})

	// Mock ComponentRegistry
	registrarMock := &mocks.MockComponentRegistrar{}

	// Error to be returned by RegisterFactory
	expectedErr := errors.New("failed to register service factory")

	// Expectations for registering resources
	registrarMock.On("RegisterFactory", ctx, mock.Anything, mock.Anything).Return(expectedErr)

	// Set the mocked ComponentRegistry
	mockSystem.On("ComponentRegistry").Return(registrarMock)

	// Initialize plugin
	err := p.Initialize(ctx, mockSystem)
	assert.NoError(t, err)

	// Test resource registration failure for RegisterFactory
	err = p.RegisterResources(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to register service factory")
}

func TestNovaPlugin_RegisterResources_Failure_CreateComponent(t *testing.T) {
	ctx := &context.Context{}
	// Initialize NovaPlugin instance
	p := plugin.NewNovaPlugin()

	// Mock SystemInterface
	mockSystem := &mocks.MockSystem{}
	mockSystem.On("Logger").Return(&mocks.MockLogger{})

	// Mock ComponentRegistry
	registrarMock := &mocks.MockComponentRegistrar{}

	// Error to be returned by CreateComponent
	expectedErr := errors.New("failed to create and register builder service")

	// Expectations for registering resources
	registrarMock.On("RegisterFactory", ctx, mock.Anything, mock.Anything).Return(expectedErr)

	// Set the mocked ComponentRegistry
	mockSystem.On("ComponentRegistry").Return(registrarMock)

	// Initialize plugin
	err := p.Initialize(ctx, mockSystem)
	assert.NoError(t, err)

	// Test resource registration failure for CreateComponent
	err = p.RegisterResources(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create and register builder service")
}

func TestNovaPlugin_Start_Success(t *testing.T) {
	// Initialize NovaPlugin instance
	p := plugin.NewNovaPlugin()

	// Mock SystemInterface
	mockSystem := &mocks.MockSystem{}
	mockSystem.On("Logger").Return(&mocks.MockLogger{})

	// Mock ComponentRegistry
	registrarMock := &mocks.MockComponentRegistrar{}

	// Mock BuildService component
	mockBuildService := &mocks.MockSystemService{}
	mockBuildService.On("ID").Return(common.BuildService)
	mockBuildService.On("Initialize", mock.Anything, mockSystem).Return(nil)

	// Expectations for creating BuildService component
	registrarMock.On(
		"CreateComponent", mock.Anything, mock.Anything).Return(mockBuildService, nil)

	// Expectations for starting BuildService
	mockBuildService.On("Start", mock.Anything).Return(nil)

	// Set the mocked ComponentRegistry
	mockSystem.On("ComponentRegistry").Return(registrarMock)

	ctx := &context.Context{}

	// Initialize plugin
	err := p.Initialize(ctx, mockSystem)
	assert.NoError(t, err)

	// Test starting the plugin
	err = p.Start(ctx)
	assert.NoError(t, err)
}

func TestNovaPlugin_Start_Failure_StartService(t *testing.T) {
	// Initialize NovaPlugin instance
	p := plugin.NewNovaPlugin()

	// Mock SystemInterface
	mockSystem := &mocks.MockSystem{}
	mockSystem.On("Logger").Return(&mocks.MockLogger{})

	// Mock ComponentRegistry
	registrarMock := &mocks.MockComponentRegistrar{}

	// Mock BuildService component
	mockBuildService := &mocks.MockSystemService{}
	mockBuildService.On("ID").Return(common.BuildService)
	mockBuildService.On("Initialize", mock.Anything, mockSystem).Return(nil)

	// Error to be returned by Start
	expectedErr := errors.New("failed to start BuildService")

	// Expectations for creating BuildService component
	registrarMock.On("CreateComponent", mock.Anything, mock.Anything).Return(&mocks.MockComponent{}, expectedErr)

	// Expectations for starting BuildService
	mockBuildService.On("Start", mock.AnythingOfType("*context.Context")).Return(expectedErr)

	// Set the mocked ComponentRegistry
	mockSystem.On("ComponentRegistry").Return(registrarMock)

	ctx := &context.Context{}

	// Initialize plugin
	err := p.Initialize(ctx, mockSystem)
	assert.NoError(t, err)

	// Test starting the plugin failure for Start service
	err = p.Start(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to start BuildService")
}

func TestNovaPlugin_Stop_Success(t *testing.T) {
	// Initialize NovaPlugin instance
	p := plugin.NewNovaPlugin()

	// Mock SystemInterface
	mockSystem := &mocks.MockSystem{}
	mockSystem.On("Logger").Return(&mocks.MockLogger{})

	// Mock ComponentRegistry
	registrarMock := &mocks.MockComponentRegistrar{}

	// Mock BuildService component
	mockBuildService := &mocks.MockSystemService{}
	mockBuildService.On("ID").Return(common.BuildService)

	// Expectations for retrieving BuildService component
	registrarMock.On("GetComponent", common.BuildService).Return(mockBuildService, nil)

	// Expectations for stopping BuildService
	mockBuildService.On("Stop", mock.AnythingOfType("*context.Context")).Return(nil)

	// Set the mocked ComponentRegistry
	mockSystem.On("ComponentRegistry").Return(registrarMock)

	ctx := &context.Context{}

	// Initialize plugin
	err := p.Initialize(ctx, mockSystem)
	assert.NoError(t, err)

	// Test stopping the plugin
	err = p.Stop(ctx)
	assert.NoError(t, err)
}

func TestNovaPlugin_Stop_Failure_GetComponent(t *testing.T) {
	// Initialize NovaPlugin instance
	p := plugin.NewNovaPlugin()

	// Mock SystemInterface
	mockSystem := &mocks.MockSystem{}
	mockSystem.On("Logger").Return(&mocks.MockLogger{})

	// Mock ComponentRegistry
	registrarMock := &mocks.MockComponentRegistrar{}

	// Error to be returned by GetComponent
	expectedErr := errors.New("failed to get BuildService component")

	// Expectations for retrieving BuildService component
	registrarMock.On("GetComponent",
		common.BuildService).Return(&mocks.MockComponent{}, expectedErr)

	// Set the mocked ComponentRegistry
	mockSystem.On("ComponentRegistry").Return(registrarMock)

	ctx := &context.Context{}

	// Initialize plugin
	err := p.Initialize(ctx, mockSystem)
	assert.NoError(t, err)

	// Test stopping the plugin failure for GetComponent
	err = p.Stop(ctx)
	assert.Error(t, err)
	assert.EqualError(t, err, "failed to get BuildService component: "+expectedErr.Error())
}

func TestNovaPlugin_Stop_Failure_CastComponent(t *testing.T) {
	// Initialize NovaPlugin instance
	p := plugin.NewNovaPlugin()

	// Mock SystemInterface
	mockSystem := &mocks.MockSystem{}
	mockSystem.On("Logger").Return(&mocks.MockLogger{})

	// Mock ComponentRegistry
	registrarMock := &mocks.MockComponentRegistrar{}

	// Mock invalid BuildService component
	invalidComponentMock := &mocks.MockComponent{}

	// Expectations for retrieving BuildService component
	registrarMock.On("GetComponent", common.BuildService).Return(invalidComponentMock, nil)

	// Set the mocked ComponentRegistry
	mockSystem.On("ComponentRegistry").Return(registrarMock)

	ctx := &context.Context{}

	// Initialize plugin
	err := p.Initialize(ctx, mockSystem)
	assert.NoError(t, err)

	// Test stopping the plugin failure for invalid component cast
	err = p.Stop(ctx)
	assert.Error(t, err)
	assert.EqualError(t, err, "BuildService component does not implement SystemServiceInterface")
}

func TestNovaPlugin_Stop_Failure_StopService(t *testing.T) {
	// Initialize NovaPlugin instance
	p := plugin.NewNovaPlugin()

	// Mock SystemInterface
	mockSystem := &mocks.MockSystem{}
	mockSystem.On("Logger").Return(&mocks.MockLogger{})

	// Mock ComponentRegistry
	registrarMock := &mocks.MockComponentRegistrar{}

	// Mock BuildService component
	mockBuildService := &mocks.MockSystemService{}
	mockBuildService.On("ID").Return(common.BuildService)

	// Error to be returned by Stop
	expectedErr := errors.New("failed to stop BuildService")

	// Expectations for retrieving BuildService component
	registrarMock.On("GetComponent", common.BuildService).Return(mockBuildService, nil)

	// Expectations for stopping BuildService
	mockBuildService.On("Stop", mock.AnythingOfType("*context.Context")).Return(expectedErr)

	// Set the mocked ComponentRegistry
	mockSystem.On("ComponentRegistry").Return(registrarMock)

	ctx := &context.Context{}

	// Initialize plugin
	err := p.Initialize(ctx, mockSystem)
	assert.NoError(t, err)

	// Test stopping the plugin failure for Stop service
	err = p.Stop(ctx)
	assert.Error(t, err)
	assert.EqualError(t, err, "failed to stop BuildService: "+expectedErr.Error())
}
