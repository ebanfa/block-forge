package system_test

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/component"
	configApi "github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
)

// Mocking the logger, event bus, and component registrar is assumed to be done elsewhere

var (
	ctx                    *context.Context
	logger                 *mocks.MockLogger
	sys                    systemApi.SystemInterface
	registrar              *mocks.MockComponentRegistrar
	serviceFactory         *mocks.MockComponentFactory
	operationFactory       *mocks.MockComponentFactory
	mockServiceComponent   *mocks.MockSystemService
	mockOperationComponent *mocks.MockOperation
	mockPluginManager      *mocks.MockPluginManager
	mockMultiStore         *mocks.MockMultiStore
)

func TestMain(m *testing.M) {
	ctx = context.Background()
	logger = &mocks.MockLogger{}
	eventBus := &mocks.MockEventBus{}
	registrar = &mocks.MockComponentRegistrar{}
	serviceFactory = &mocks.MockComponentFactory{}
	operationFactory = &mocks.MockComponentFactory{}
	mockServiceComponent = &mocks.MockSystemService{}
	mockOperationComponent = &mocks.MockOperation{}
	mockPluginManager = &mocks.MockPluginManager{}
	mockMultiStore = &mocks.MockMultiStore{}

	configuration := &configApi.Configuration{
		Services: []*configApi.ServiceConfiguration{
			{
				ComponentConfig: configApi.ComponentConfig{
					ID:   "Service1_ID",
					Name: "Service1",
				},
			},
		},
		Operations: []*configApi.OperationConfiguration{
			{
				ComponentConfig: configApi.ComponentConfig{
					ID:   "Operation1_ID",
					Name: "Operation1",
				},
			},
		},
	}

	sys = systemApi.NewSystem(logger, eventBus, configuration, mockPluginManager, registrar, mockMultiStore)

	// Mock the behavior of the component and factory
	mockServiceComponent.On("Type").Return(component.ServiceType)
	mockServiceComponent.On("Initialize", mock.Anything, mock.Anything).Return(nil)

	mockOperationComponent.On("Type").Return(component.OperationType)
	mockOperationComponent.On("Initialize", mock.Anything, mock.Anything).Return(nil)

	serviceFactory.On("CreateComponent", mock.Anything).Return(mockServiceComponent, nil)
	operationFactory.On("CreateComponent", mock.Anything).Return(mockOperationComponent, nil)

	mockPluginManager.On("Initialize", ctx, mock.Anything).Return(nil)
	mockPluginManager.On("StartPlugins", ctx).Return(nil)

	// Run tests
	exitCode := m.Run()

	// Clean up any global resources here if needed

	// Exit with the proper exit code
	os.Exit(exitCode)
}

func TestSystemImpl_Initialize_Success(t *testing.T) {
	// Define different behaviors based on the arguments
	registrar.On("GetComponentFactory", "Service1Factory").Return(serviceFactory, nil)
	registrar.On("GetComponentFactory", "Operation1Factory").Return(operationFactory, nil)

	// Test initialization
	err := sys.Initialize(ctx)
	assert.NoError(t, err)
}

func TestSystemImpl_Initialize_Error(t *testing.T) {
	// Mock context

	// Define different behaviors based on the arguments
	registrar.On("GetComponentFactory", "NonexistentFactory").Return(
		operationFactory, systemApi.ErrComponentFactoryNil)

	// Mock configuration with invalid data
	configuration := &configApi.Configuration{
		Services: []*configApi.ServiceConfiguration{
			{
				ComponentConfig: configApi.ComponentConfig{
					ID:   "Operation1_ID",
					Name: "Operation1",
				},
			},
		},
		Operations: []*configApi.OperationConfiguration{
			{
				ComponentConfig: configApi.ComponentConfig{
					ID:   "Service1_ID",
					Name: "Service1",
				},
			},
		},
	}

	// Create a sys instance
	sys := systemApi.NewSystem(nil, nil, configuration, mockPluginManager, registrar, mockMultiStore)

	// Test initialization with error
	_ = sys.Initialize(ctx)
	//assert.Error(t, err)
}

func TestSystemImpl_Start_Success(t *testing.T) {
	// Define different behaviors based on the arguments
	registrar.On("GetComponentFactory", "Service1Factory").Return(serviceFactory, nil)
	registrar.On("GetComponentFactory", "Operation1Factory").Return(operationFactory, nil)

	registrar.On("GetComponentsByType", component.ServiceType).Return([]component.ComponentInterface{mockServiceComponent}, nil)
	registrar.On("GetComponentsByType", component.OperationType).Return([]component.ComponentInterface{mockOperationComponent}, nil)

	// Mock the behavior of the component and factory
	mockServiceComponent.On("Start", ctx).Return(nil)

	// Test starting the sys
	err := sys.Start(ctx)
	assert.NoError(t, err)
}

func TestSystemImpl_Start_Error(t *testing.T) {

	registrar.On("GetComponentFactory", "Service1Factory").Return(serviceFactory, nil)
	registrar.On("GetComponentFactory", "Operation1Factory").Return(operationFactory, nil)

	registrar.On("GetComponentByType", component.ServiceType).Return([]component.ComponentInterface{mockServiceComponent}, nil)
	registrar.On("GetComponentByType", component.OperationType).Return([]component.ComponentInterface{mockOperationComponent}, nil)

	// Create a sys instance
	sys := systemApi.NewSystem(nil, nil, &configApi.Configuration{}, mockPluginManager, registrar, mockMultiStore)

	// Test starting the sys with error
	err := sys.Start(ctx)
	assert.Error(t, err)
}

func TestSystemImpl_Stop_Success(t *testing.T) {
	// Mock service configuration
	registrar.On("GetComponentFactory", "testFactoryInitializeServiceSuccess").Return(serviceFactory, nil)
	registrar.On("GetComponentByType", component.ServiceType).Return([]component.ComponentInterface{}, nil)
	mockServiceComponent.On("Stop", ctx).Return(nil)
	// Create a sys instance
	sys := systemApi.NewSystem(nil, nil, &configApi.Configuration{}, mockPluginManager, registrar, mockMultiStore)

	err := sys.Initialize(ctx)
	assert.NoError(t, err)

	err = sys.Start(ctx)
	assert.NoError(t, err)

	// Test stopping the sys
	err = sys.Stop(ctx)
	assert.NoError(t, err)
}

func TestSystemImpl_Stop_Error(t *testing.T) {
	// Mock context
	ctx := &context.Context{}

	// Mock component registrar with error
	componentReg := &mocks.MockComponentRegistrar{}

	// Create a sys instance
	sys := systemApi.NewSystem(nil, nil, &configApi.Configuration{}, mockPluginManager, componentReg, mockMultiStore)

	// Test stopping the sys with error
	err := sys.Stop(ctx)
	assert.Error(t, err)
}

func TestSystemImpl_ExecuteOperation_Success(t *testing.T) {
	// Mocks//
	mockOperation := &mocks.MockOperation{}
	operationInput := &systemApi.SystemOperationInput{}
	expectedOutput := &systemApi.SystemOperationOutput{}

	registrar.On("GetComponent", "Operation1_ID").Return(mockOperation, nil)
	mockOperation.On("Execute", ctx, operationInput).Return(expectedOutput, nil)

	// Test executing an operation
	//_, err := sys.ExecuteOperation(ctx, "Operation1_ID", operationInput)
	//assert.NoError(t, err)
}

func TestSystemImpl_ExecuteOperation_Error_ComponentNotFound(t *testing.T) {
	// Mock context
	ctx := &context.Context{}

	// Mock operation input
	operationInput := &systemApi.SystemOperationInput{}

	componentReg := &mocks.MockComponentRegistrar{}
	componentReg.On("GetComponent", "operation_id").Return(mockOperationComponent, systemApi.ErrComponentNotFound)

	sys := systemApi.NewSystem(nil, nil, &configApi.Configuration{}, mockPluginManager, componentReg, mockMultiStore)

	// Test executing an operation with component not found error
	output, err := sys.ExecuteOperation(ctx, "operation_id", operationInput)
	assert.Error(t, err)
	assert.Nil(t, output)
}

func TestSystemImpl_ExecuteOperation_Error_ComponentNotOperation(t *testing.T) {
	// Mock context
	ctx := &context.Context{}

	// Mock operation input
	operationInput := &systemApi.SystemOperationInput{}

	// Mock component registrar returning a non-operation component
	componentReg := &mocks.MockComponentRegistrar{}
	componentReg.On("GetComponent", "operation_id").Return(mockServiceComponent, nil)

	// Create a sys instance
	sys := systemApi.NewSystem(nil, nil, &configApi.Configuration{}, mockPluginManager, componentReg, mockMultiStore)

	// Test executing an operation with component not an operation error
	output, err := sys.ExecuteOperation(ctx, "operation_id", operationInput)
	assert.Error(t, err)
	assert.Nil(t, output)
}

func TestSystemImpl_StartService_Success(t *testing.T) {
	// Mock component registrar
	componentReg := &mocks.MockComponentRegistrar{}
	componentReg.On("GetComponent", "service_id").Return(mockServiceComponent, nil)

	// Create a sys instance
	sys := systemApi.NewSystem(nil, nil, &configApi.Configuration{}, mockPluginManager, componentReg, mockMultiStore)

	// Test starting a service
	err := sys.StartService(ctx, "service_id")
	assert.NoError(t, err)
}

func TestSystemImpl_StartService_Error_ComponentNotFound(t *testing.T) {

	// Mock component registrar with error
	componentReg := &mocks.MockComponentRegistrar{}
	componentReg.On("GetComponent", "service_id").Return(mockServiceComponent, systemApi.ErrComponentNotFound)

	// Create a sys instance
	sys := systemApi.NewSystem(nil, nil, &configApi.Configuration{}, mockPluginManager, componentReg, mockMultiStore)

	// Test starting a service with component not found error
	err := sys.StartService(ctx, "service_id")
	assert.Error(t, err)
}

func TestSystemImpl_StopService_Success(t *testing.T) {
	// Mock component registrar
	componentReg := &mocks.MockComponentRegistrar{}
	componentReg.On("GetComponent", "service_id").Return(mockServiceComponent, nil)

	// Create a sys instance
	sys := systemApi.NewSystem(nil, nil, &configApi.Configuration{}, mockPluginManager, componentReg, mockMultiStore)

	// Test stopping a service
	err := sys.StopService(ctx, "service_id")
	assert.NoError(t, err)
}

func TestSystemImpl_StopService_Error_ComponentNotFound(t *testing.T) {
	// Mock component registrar with error
	componentReg := &mocks.MockComponentRegistrar{}
	componentReg.On("GetComponent", "service_id").Return(mockServiceComponent, systemApi.ErrComponentNotFound)

	// Create a sys instance
	sys := systemApi.NewSystem(nil, nil, &configApi.Configuration{}, mockPluginManager, componentReg, mockMultiStore)

	// Test stopping a service with component not found error
	err := sys.StopService(ctx, "service_id")
	assert.Error(t, err)
}

func TestSystemImpl_RestartService_Success(t *testing.T) {
	// Mock component registrar
	componentReg := &mocks.MockComponentRegistrar{}
	componentReg.On("GetComponent", "service_id").Return(mockServiceComponent, nil)

	// Create a sys instance
	sys := systemApi.NewSystem(nil, nil, &configApi.Configuration{}, mockPluginManager, componentReg, mockMultiStore)

	// Test restarting a service
	err := sys.RestartService(ctx, "service_id")
	assert.NoError(t, err)
}

func TestSystemImpl_RestartService_Error_StopService(t *testing.T) {
	// Mock context
	ctx := &context.Context{}

	// Mock component registrar
	componentReg := &mocks.MockComponentRegistrar{}
	componentReg.On("GetComponent", "service_id").Return(mockServiceComponent, nil)
	mockServiceComponent.On("Stop", ctx).Return(errors.New("Error stopping service"))

	// Create a sys instance
	sys := systemApi.NewSystem(nil, nil, &configApi.Configuration{}, mockPluginManager, componentReg, mockMultiStore)

	// Test restarting a service with error while stopping service
	err := sys.RestartService(ctx, "service_id")
	assert.Error(t, err)
}

func TestSystemImpl_RestartService_Error_StartService(t *testing.T) {
	// Mock context
	ctx := &context.Context{}

	// Mock component registrar
	componentReg := &mocks.MockComponentRegistrar{}
	componentReg.On("GetComponent", "service_id").Return(mockServiceComponent, nil)
	mockServiceComponent.On("Stop", ctx).Return(nil)
	mockServiceComponent.On("Start", ctx).Return(errors.New("Error starting service"))

	// Create a sys instance
	sys := systemApi.NewSystem(nil, nil, &configApi.Configuration{}, mockPluginManager, componentReg, mockMultiStore)

	// Test restarting a service with error while starting service
	err := sys.RestartService(ctx, "service_id")
	assert.Error(t, err)
}
