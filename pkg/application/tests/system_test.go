package tests

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/components"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
)

// Mocking the logger, event bus, and component registrar is assumed to be done elsewhere

var (
	ctx                    *context.Context
	logger                 *mocks.MockLogger
	system                 systemApi.SystemInterface
	registrar              *mocks.MockComponentRegistrar
	serviceFactory         *mocks.MockComponentFactory
	operationFactory       *mocks.MockComponentFactory
	mockServiceComponent   *mocks.MockSystemService
	mockOperationComponent *mocks.MockOperation
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

	configuration := &components.Configuration{
		Services: []*components.ServiceConfiguration{
			{
				ComponentConfig: components.ComponentConfig{
					ID:          "Service1_ID",
					Name:        "Service1",
					FactoryName: "Service1Factory",
				},
			},
		},
		Operations: []*components.OperationConfiguration{
			{
				ComponentConfig: components.ComponentConfig{
					ID:          "Operation1_ID",
					Name:        "Operation1",
					FactoryName: "Operation1Factory",
				},
			},
		},
	}

	system = systemApi.NewSystem(logger, eventBus, configuration, registrar)

	// Mock the behavior of the component and factory
	mockServiceComponent.On("Type").Return(components.ServiceType)
	mockServiceComponent.On("Initialize", mock.Anything, mock.Anything).Return(nil)

	mockOperationComponent.On("Type").Return(components.OperationType)
	mockOperationComponent.On("Initialize", mock.Anything, mock.Anything).Return(nil)

	serviceFactory.On("CreateComponent", mock.Anything).Return(mockServiceComponent, nil)
	operationFactory.On("CreateComponent", mock.Anything).Return(mockOperationComponent, nil)

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
	err := system.Initialize(ctx)
	assert.NoError(t, err)
}

func TestSystemImpl_Initialize_Error(t *testing.T) {
	// Mock context
	ctx := &context.Context{}

	// Define different behaviors based on the arguments
	registrar.On("GetComponentFactory", "NonexistentFactory").Return(
		operationFactory, components.ErrComponentFactoryNil)

	// Mock configuration with invalid data
	configuration := &components.Configuration{
		Services: []*components.ServiceConfiguration{
			{
				ComponentConfig: components.ComponentConfig{
					ID:          "Operation1_ID",
					Name:        "Operation1",
					FactoryName: "NonexistentFactory",
				},
			},
		},
		Operations: []*components.OperationConfiguration{
			{
				ComponentConfig: components.ComponentConfig{
					ID:          "Service1_ID",
					Name:        "Service1",
					FactoryName: "NonexistentFactory",
				},
			},
		},
	}

	// Create a system instance
	sys := systemApi.NewSystem(nil, nil, configuration, registrar)

	// Test initialization with error
	err := sys.Initialize(ctx)
	assert.Error(t, err)
}

func TestSystemImpl_Start_Success(t *testing.T) {
	// Define different behaviors based on the arguments
	registrar.On("GetComponentFactory", "Service1Factory").Return(serviceFactory, nil)
	registrar.On("GetComponentFactory", "Operation1Factory").Return(operationFactory, nil)

	registrar.On("GetComponentByType", components.ServiceType).Return([]components.ComponentInterface{mockServiceComponent}, nil)
	registrar.On("GetComponentByType", components.OperationType).Return([]components.ComponentInterface{mockOperationComponent}, nil)

	// Mock the behavior of the component and factory
	mockServiceComponent.On("Start", ctx).Return(nil)

	// Test starting the system
	err := system.Start(ctx)
	assert.NoError(t, err)
}

func TestSystemImpl_Start_Error(t *testing.T) {

	registrar.On("GetComponentFactory", "Service1Factory").Return(serviceFactory, nil)
	registrar.On("GetComponentFactory", "Operation1Factory").Return(operationFactory, nil)

	registrar.On("GetComponentByType", components.ServiceType).Return([]components.ComponentInterface{mockServiceComponent}, nil)
	registrar.On("GetComponentByType", components.OperationType).Return([]components.ComponentInterface{mockOperationComponent}, nil)

	// Create a system instance
	sys := systemApi.NewSystem(nil, nil, &components.Configuration{}, registrar)

	// Test starting the system with error
	err := sys.Start(ctx)
	assert.Error(t, err)
}

func TestSystemImpl_InitializeOperation_Success(t *testing.T) {
	// Mock operation configuration
	operationConfig := &components.OperationConfiguration{
		ComponentConfig: components.ComponentConfig{
			ID:          "operation_id",
			Name:        "testOperation",
			FactoryName: "testFactory",
		},
	}
	registrar.On("GetComponentFactory", "testFactory").Return(operationFactory, nil)
	registrar.On("GetComponentByType", components.OperationType).Return([]components.ComponentInterface{mockOperationComponent}, nil)

	operationFactory.On("CreateComponent", mock.Anything).Return(mockOperationComponent, nil)

	// Create a system instance
	sys := systemApi.NewSystem(nil, nil, &components.Configuration{}, registrar)

	// Test initializing an operation
	err := sys.InitializeOperation(ctx, operationConfig)
	assert.NoError(t, err)
}

func TestSystemImpl_InitializeOperation_Error(t *testing.T) {
	// Mock operation configuration
	operationConfig := &components.OperationConfiguration{
		// Missing FactoryName
		ComponentConfig: components.ComponentConfig{
			ID:          "operation_id",
			Name:        "testOperation",
			FactoryName: "",
		},
	}
	registrar.On("GetComponentFactory", "").Return(operationFactory, components.ErrFactoryNotFound)
	registrar.On("GetComponentByType", components.OperationType).Return([]components.ComponentInterface{mockOperationComponent}, nil)

	operationFactory.On("CreateComponent", mock.Anything).Return(mockOperationComponent, nil)

	// Create a system instance
	sys := systemApi.NewSystem(nil, nil, &components.Configuration{}, registrar)

	// Test initializing an operation with error
	err := sys.InitializeOperation(ctx, operationConfig)
	assert.Error(t, err)
}

func TestSystemImpl_InitializeService_Success(t *testing.T) {

	// Mock service configuration
	serviceConfig := &components.ServiceConfiguration{
		ComponentConfig: components.ComponentConfig{
			ID:          "service_id",
			Name:        "testInitializeServiceSuccess",
			FactoryName: "testFactoryInitializeServiceSuccess",
		},
	}

	registrar.On("GetComponentFactory", "testFactoryInitializeServiceSuccess").Return(serviceFactory, nil)
	registrar.On("GetComponentByType", components.ServiceType).Return([]components.ComponentInterface{mockServiceComponent}, nil)

	// Create a system instance
	sys := systemApi.NewSystem(nil, nil, &components.Configuration{}, registrar)

	// Test initializing a service
	err := sys.InitializeService(ctx, serviceConfig)
	assert.NoError(t, err)
}

func TestSystemImpl_InitializeService_Error(t *testing.T) {
	// Mock context
	ctx := &context.Context{}

	// Mock service configuration with missing factory name
	serviceConfig := &components.ServiceConfiguration{
		// Missing FactoryName
		ComponentConfig: components.ComponentConfig{
			ID:          "service_id",
			Name:        "testService",
			FactoryName: "",
		},
	}

	registrar.On("GetComponentFactory", "").Return(operationFactory, components.ErrFactoryNotFound)
	registrar.On("GetComponentByType", components.OperationType).Return([]components.ComponentInterface{mockOperationComponent}, nil)

	operationFactory.On("CreateComponent", mock.Anything).Return(mockOperationComponent, nil)

	// Create a system instance
	sys := systemApi.NewSystem(nil, nil, &components.Configuration{}, registrar)

	// Test initializing a service with error
	err := sys.InitializeService(ctx, serviceConfig)
	assert.Error(t, err)
}

func TestSystemImpl_Stop_Success(t *testing.T) {
	// Mock service configuration
	registrar.On("GetComponentFactory", "testFactoryInitializeServiceSuccess").Return(serviceFactory, nil)
	registrar.On("GetComponentByType", components.ServiceType).Return([]components.ComponentInterface{}, nil)
	mockServiceComponent.On("Stop", ctx).Return(nil)
	// Create a system instance
	sys := systemApi.NewSystem(nil, nil, &components.Configuration{}, registrar)

	err := sys.Initialize(ctx)
	assert.NoError(t, err)

	err = sys.Start(ctx)
	assert.NoError(t, err)

	// Test stopping the system
	err = sys.Stop(ctx)
	assert.NoError(t, err)
}

func TestSystemImpl_Stop_Error(t *testing.T) {
	// Mock context
	ctx := &context.Context{}

	// Mock component registrar with error
	componentReg := &mocks.MockComponentRegistrar{}

	// Create a system instance
	sys := systemApi.NewSystem(nil, nil, &components.Configuration{}, componentReg)

	// Test stopping the system with error
	err := sys.Stop(ctx)
	assert.Error(t, err)
}

func TestSystemImpl_ExecuteOperation_Success(t *testing.T) {
	// Mocks
	operationInput := &systemApi.OperationInput{}
	expectedOutput := &systemApi.OperationOutput{}

	registrar.On("GetComponent", "Operation1_ID").Return(mockOperationComponent, nil)
	mockOperationComponent.On("Execute", ctx, operationInput).Return(expectedOutput, nil)

	// Test executing an operation
	output, err := system.ExecuteOperation(ctx, "Operation1_ID", operationInput)
	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, output)
}

func TestSystemImpl_ExecuteOperation_Error_ComponentNotFound(t *testing.T) {
	// Mock context
	ctx := &context.Context{}

	// Mock operation input
	operationInput := &systemApi.OperationInput{}

	componentReg := &mocks.MockComponentRegistrar{}
	componentReg.On("GetComponent", "operation_id").Return(mockOperationComponent, components.ErrComponentNotFound)

	sys := systemApi.NewSystem(nil, nil, &components.Configuration{}, componentReg)

	// Test executing an operation with component not found error
	output, err := sys.ExecuteOperation(ctx, "operation_id", operationInput)
	assert.Error(t, err)
	assert.Nil(t, output)
}

func TestSystemImpl_ExecuteOperation_Error_ComponentNotOperation(t *testing.T) {
	// Mock context
	ctx := &context.Context{}

	// Mock operation input
	operationInput := &systemApi.OperationInput{}

	// Mock component registrar returning a non-operation component
	componentReg := &mocks.MockComponentRegistrar{}
	componentReg.On("GetComponent", "operation_id").Return(mockServiceComponent, nil)

	// Create a system instance
	sys := systemApi.NewSystem(nil, nil, &components.Configuration{}, componentReg)

	// Test executing an operation with component not an operation error
	output, err := sys.ExecuteOperation(ctx, "operation_id", operationInput)
	assert.Error(t, err)
	assert.Nil(t, output)
}

func TestSystemImpl_StartService_Success(t *testing.T) {
	// Mock component registrar
	componentReg := &mocks.MockComponentRegistrar{}
	componentReg.On("GetComponent", "service_id").Return(mockServiceComponent, nil)

	// Create a system instance
	sys := systemApi.NewSystem(nil, nil, &components.Configuration{}, componentReg)

	// Test starting a service
	err := sys.StartService(ctx, "service_id")
	assert.NoError(t, err)
}

func TestSystemImpl_StartService_Error_ComponentNotFound(t *testing.T) {

	// Mock component registrar with error
	componentReg := &mocks.MockComponentRegistrar{}
	componentReg.On("GetComponent", "service_id").Return(mockServiceComponent, components.ErrComponentNotFound)

	// Create a system instance
	sys := systemApi.NewSystem(nil, nil, &components.Configuration{}, componentReg)

	// Test starting a service with component not found error
	err := sys.StartService(ctx, "service_id")
	assert.Error(t, err)
}

func TestSystemImpl_StopService_Success(t *testing.T) {
	// Mock component registrar
	componentReg := &mocks.MockComponentRegistrar{}
	componentReg.On("GetComponent", "service_id").Return(mockServiceComponent, nil)

	// Create a system instance
	sys := systemApi.NewSystem(nil, nil, &components.Configuration{}, componentReg)

	// Test stopping a service
	err := sys.StopService(ctx, "service_id")
	assert.NoError(t, err)
}

func TestSystemImpl_StopService_Error_ComponentNotFound(t *testing.T) {
	// Mock component registrar with error
	componentReg := &mocks.MockComponentRegistrar{}
	componentReg.On("GetComponent", "service_id").Return(mockServiceComponent, components.ErrComponentNotFound)

	// Create a system instance
	sys := systemApi.NewSystem(nil, nil, &components.Configuration{}, componentReg)

	// Test stopping a service with component not found error
	err := sys.StopService(ctx, "service_id")
	assert.Error(t, err)
}

func TestSystemImpl_RestartService_Success(t *testing.T) {
	// Mock component registrar
	componentReg := &mocks.MockComponentRegistrar{}
	componentReg.On("GetComponent", "service_id").Return(mockServiceComponent, nil)

	// Create a system instance
	sys := systemApi.NewSystem(nil, nil, &components.Configuration{}, componentReg)

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

	// Create a system instance
	sys := systemApi.NewSystem(nil, nil, &components.Configuration{}, componentReg)

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

	// Create a system instance
	sys := systemApi.NewSystem(nil, nil, &components.Configuration{}, componentReg)

	// Test restarting a service with error while starting service
	err := sys.RestartService(ctx, "service_id")
	assert.Error(t, err)
}
