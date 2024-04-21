package plugin

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/edward1christian/block-forge/nova/pkg/common"
	"github.com/edward1christian/block-forge/nova/pkg/components/factories"
	"github.com/edward1christian/block-forge/nova/pkg/components/plugin"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
)

func TestNovaPlugin_Initialize(t *testing.T) {
	// Initialize NovaPlugin instance
	p := plugin.NewNovaPlugin()

	// Mock SystemInterface
	mockSystem := &mocks.MockSystem{}

	// Mock Context
	ctx := &context.Context{}

	// Test initialization
	err := p.Initialize(ctx, mockSystem)
	assert.NoError(t, err)
}

func TestNovaPlugin_RegisterResources_Success(t *testing.T) {
	// Initialize NovaPlugin instance
	p := plugin.NewNovaPlugin()

	// Mock SystemInterface
	mockSystem := &mocks.MockSystem{}

	// Mock ComponentRegistry
	registrarMock := &mocks.MockComponentRegistrar{}

	// Expectations for registering resources
	registrarMock.On("RegisterFactory", common.IgniteBuildServiceFactory, &factories.BuilderServiceFactory{}).Return(nil)
	registrarMock.On("CreateComponent", &config.ComponentConfig{
		ID:        common.IgniteBuildService,
		FactoryID: common.IgniteBuildServiceFactory,
	}).Return(&mocks.MockComponent{}, nil)

	// Set the mocked ComponentRegistry
	mockSystem.On("ComponentRegistry").Return(registrarMock)

	ctx := &context.Context{}

	// Initialize plugin
	err := p.Initialize(ctx, mockSystem)

	// Test resource registration
	err = p.RegisterResources(ctx)
	assert.NoError(t, err)
}

func TestNovaPlugin_RegisterResources_Failure_RegisterFactory(t *testing.T) {
	// Initialize NovaPlugin instance
	p := plugin.NewNovaPlugin()

	// Mock SystemInterface
	mockSystem := &mocks.MockSystem{}

	// Mock ComponentRegistry
	registrarMock := &mocks.MockComponentRegistrar{}

	// Error to be returned by RegisterFactory
	expectedErr := errors.New("failed to register service factory")

	// Expectations for registering resources
	registrarMock.On("RegisterFactory", common.IgniteBuildServiceFactory, &factories.BuilderServiceFactory{}).Return(expectedErr)

	// Set the mocked ComponentRegistry
	mockSystem.On("ComponentRegistry").Return(registrarMock)

	ctx := &context.Context{}

	// Initialize plugin
	err := p.Initialize(ctx, mockSystem)

	// Test resource registration failure for RegisterFactory
	err = p.RegisterResources(ctx)
	assert.Error(t, err)
	assert.EqualError(t, err, "failed to register service factory: "+expectedErr.Error())
}

func TestNovaPlugin_RegisterResources_Failure_CreateComponent(t *testing.T) {
	// Initialize NovaPlugin instance
	p := plugin.NewNovaPlugin()

	// Mock SystemInterface
	mockSystem := &mocks.MockSystem{}

	// Mock ComponentRegistry
	registrarMock := &mocks.MockComponentRegistrar{}

	// Error to be returned by CreateComponent
	expectedErr := errors.New("failed to create and register builder service")

	// Expectations for registering resources
	registrarMock.On("RegisterFactory", common.IgniteBuildServiceFactory, &factories.BuilderServiceFactory{}).Return(nil)
	registrarMock.On("CreateComponent", &config.ComponentConfig{
		ID:        common.IgniteBuildService,
		FactoryID: common.IgniteBuildServiceFactory,
	}).Return(&mocks.MockComponent{}, expectedErr)

	// Set the mocked ComponentRegistry
	mockSystem.On("ComponentRegistry").Return(registrarMock)

	ctx := &context.Context{}

	// Initialize plugin
	err := p.Initialize(ctx, mockSystem)

	// Test resource registration failure for CreateComponent
	err = p.RegisterResources(ctx)
	assert.Error(t, err)
	assert.EqualError(t, err, "failed to create and register builder service: "+expectedErr.Error())
}

func TestNovaPlugin_Start_Success(t *testing.T) {
	// Initialize NovaPlugin instance
	p := plugin.NewNovaPlugin()

	// Mock SystemInterface
	mockSystem := &mocks.MockSystem{}

	// Mock ComponentRegistry
	registrarMock := &mocks.MockComponentRegistrar{}

	// Mock BuilderService component
	builderServiceMock := &mocks.MockSystemService{}

	// Expectations for retrieving BuilderService component
	registrarMock.On("GetComponent", common.IgniteBuildService).Return(builderServiceMock, nil)

	// Expectations for starting BuilderService
	builderServiceMock.On("Start", mock.AnythingOfType("*context.Context")).Return(nil)

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

func TestNovaPlugin_Start_Failure_GetComponent(t *testing.T) {
	// Initialize NovaPlugin instance
	p := plugin.NewNovaPlugin()

	// Mock SystemInterface
	mockSystem := &mocks.MockSystem{}

	// Mock ComponentRegistry
	registrarMock := &mocks.MockComponentRegistrar{}

	// Error to be returned by GetComponent
	expectedErr := errors.New("failed to get BuilderService component")

	// Expectations for retrieving BuilderService component
	registrarMock.On("GetComponent",
		common.IgniteBuildService).Return(&mocks.MockComponent{}, expectedErr)

	// Set the mocked ComponentRegistry
	mockSystem.On("ComponentRegistry").Return(registrarMock)

	ctx := &context.Context{}

	// Initialize plugin
	err := p.Initialize(ctx, mockSystem)
	assert.NoError(t, err)

	// Test starting the plugin failure for GetComponent
	err = p.Start(ctx)
	assert.Error(t, err)
	assert.EqualError(t, err, "failed to get BuilderService component: "+expectedErr.Error())
}

func TestNovaPlugin_Start_Failure_CastComponent(t *testing.T) {
	// Initialize NovaPlugin instance
	p := plugin.NewNovaPlugin()

	// Mock SystemInterface
	mockSystem := &mocks.MockSystem{}

	// Mock ComponentRegistry
	registrarMock := &mocks.MockComponentRegistrar{}

	// Mock invalid BuilderService component
	invalidComponentMock := &mocks.MockComponent{}

	// Expectations for retrieving BuilderService component
	registrarMock.On("GetComponent", common.IgniteBuildService).Return(invalidComponentMock, nil)

	// Set the mocked ComponentRegistry
	mockSystem.On("ComponentRegistry").Return(registrarMock)

	ctx := &context.Context{}

	// Initialize plugin
	err := p.Initialize(ctx, mockSystem)
	assert.NoError(t, err)

	// Test starting the plugin failure for invalid component cast
	err = p.Start(ctx)
	assert.Error(t, err)
	assert.EqualError(t, err, "BuilderService component does not implement SystemServiceInterface")
}

func TestNovaPlugin_Start_Failure_StartService(t *testing.T) {
	// Initialize NovaPlugin instance
	p := plugin.NewNovaPlugin()

	// Mock SystemInterface
	mockSystem := &mocks.MockSystem{}

	// Mock ComponentRegistry
	registrarMock := &mocks.MockComponentRegistrar{}

	// Mock BuilderService component
	builderServiceMock := &mocks.MockSystemService{}

	// Error to be returned by Start
	expectedErr := errors.New("failed to start BuilderService")

	// Expectations for retrieving BuilderService component
	registrarMock.On("GetComponent", common.IgniteBuildService).Return(builderServiceMock, nil)

	// Expectations for starting BuilderService
	builderServiceMock.On("Start", mock.AnythingOfType("*context.Context")).Return(expectedErr)

	// Set the mocked ComponentRegistry
	mockSystem.On("ComponentRegistry").Return(registrarMock)

	ctx := &context.Context{}

	// Initialize plugin
	err := p.Initialize(ctx, mockSystem)
	assert.NoError(t, err)

	// Test starting the plugin failure for Start service
	err = p.Start(ctx)
	assert.Error(t, err)
	assert.EqualError(t, err, "failed to start BuilderService: "+expectedErr.Error())
}

func TestNovaPlugin_Stop_Success(t *testing.T) {
	// Initialize NovaPlugin instance
	p := plugin.NewNovaPlugin()

	// Mock SystemInterface
	mockSystem := &mocks.MockSystem{}

	// Mock ComponentRegistry
	registrarMock := &mocks.MockComponentRegistrar{}

	// Mock BuilderService component
	builderServiceMock := &mocks.MockSystemService{}

	// Expectations for retrieving BuilderService component
	registrarMock.On("GetComponent", common.IgniteBuildService).Return(builderServiceMock, nil)

	// Expectations for stopping BuilderService
	builderServiceMock.On("Stop", mock.AnythingOfType("*context.Context")).Return(nil)

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

	// Mock ComponentRegistry
	registrarMock := &mocks.MockComponentRegistrar{}

	// Error to be returned by GetComponent
	expectedErr := errors.New("failed to get BuilderService component")

	// Expectations for retrieving BuilderService component
	registrarMock.On("GetComponent",
		common.IgniteBuildService).Return(&mocks.MockComponent{}, expectedErr)

	// Set the mocked ComponentRegistry
	mockSystem.On("ComponentRegistry").Return(registrarMock)

	ctx := &context.Context{}

	// Initialize plugin
	err := p.Initialize(ctx, mockSystem)
	assert.NoError(t, err)

	// Test stopping the plugin failure for GetComponent
	err = p.Stop(ctx)
	assert.Error(t, err)
	assert.EqualError(t, err, "failed to get BuilderService component: "+expectedErr.Error())
}

func TestNovaPlugin_Stop_Failure_CastComponent(t *testing.T) {
	// Initialize NovaPlugin instance
	p := plugin.NewNovaPlugin()

	// Mock SystemInterface
	mockSystem := &mocks.MockSystem{}

	// Mock ComponentRegistry
	registrarMock := &mocks.MockComponentRegistrar{}

	// Mock invalid BuilderService component
	invalidComponentMock := &mocks.MockComponent{}

	// Expectations for retrieving BuilderService component
	registrarMock.On("GetComponent", common.IgniteBuildService).Return(invalidComponentMock, nil)

	// Set the mocked ComponentRegistry
	mockSystem.On("ComponentRegistry").Return(registrarMock)

	ctx := &context.Context{}

	// Initialize plugin
	err := p.Initialize(ctx, mockSystem)
	assert.NoError(t, err)

	// Test stopping the plugin failure for invalid component cast
	err = p.Stop(ctx)
	assert.Error(t, err)
	assert.EqualError(t, err, "BuilderService component does not implement SystemServiceInterface")
}

func TestNovaPlugin_Stop_Failure_StopService(t *testing.T) {
	// Initialize NovaPlugin instance
	p := plugin.NewNovaPlugin()

	// Mock SystemInterface
	mockSystem := &mocks.MockSystem{}

	// Mock ComponentRegistry
	registrarMock := &mocks.MockComponentRegistrar{}

	// Mock BuilderService component
	builderServiceMock := &mocks.MockSystemService{}

	// Error to be returned by Stop
	expectedErr := errors.New("failed to stop BuilderService")

	// Expectations for retrieving BuilderService component
	registrarMock.On("GetComponent", common.IgniteBuildService).Return(builderServiceMock, nil)

	// Expectations for stopping BuilderService
	builderServiceMock.On("Stop", mock.AnythingOfType("*context.Context")).Return(expectedErr)

	// Set the mocked ComponentRegistry
	mockSystem.On("ComponentRegistry").Return(registrarMock)

	ctx := &context.Context{}

	// Initialize plugin
	err := p.Initialize(ctx, mockSystem)
	assert.NoError(t, err)

	// Test stopping the plugin failure for Stop service
	err = p.Stop(ctx)
	assert.Error(t, err)
	assert.EqualError(t, err, "failed to stop BuilderService: "+expectedErr.Error())
}
