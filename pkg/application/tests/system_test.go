package tests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// MockComponentFactory is a mock implementation of ComponentFactory for testing.

var (
	eventBus *mocks.MockEventBus
	logger   *mocks.MockLogger
	sys      *system.SystemImpl
)

// TestMain initializes common variables used across tests.
func TestMain(m *testing.M) {
	eventBus = eventBus
	logger = new(mocks.MockLogger)
	configuration := system.Configuration{
		Services: []*system.ServiceConfiguration{
			{
				ComponentConfig: system.ComponentConfig{
					ID:          "service1",
					FactoryName: "sysServiceFactory",
					Name:        "testService",
					Description: "Test Service",
				},
			},
		},
		Operations: []*system.OperationConfiguration{
			{
				ComponentConfig: system.ComponentConfig{
					ID:          "testOp1",
					FactoryName: "sysServiceFactory",
					Name:        "testOperation",
					Description: "Test Operation",
				},
			},
		},
	}

	sys = system.NewSystem(eventBus, logger, configuration)

	sysServiceFactory := func(ctx *context.Context, config *system.ComponentConfig) (system.Component, error) {
		return new(mocks.MockSystemService), nil
	}

	sysOpFactory := func(ctx *context.Context, config *system.ComponentConfig) (system.Component, error) {
		return new(mocks.MockOperation), nil
	}

	sys.RegisterComponentFactory("sysServiceFactory", sysServiceFactory)
	sys.RegisterComponentFactory("sysOpFactory", sysOpFactory)

	m.Run()
}

// Tests for NewSystem function

// TestNewSystem tests the creation of a new SystemImpl instance.
func TestNewSystem(t *testing.T) {
	assert.NotNil(t, sys)
	assert.Equal(t, "system", sys.ID())
	assert.Equal(t, "System", sys.Name())
	assert.Equal(t, "Core system in the application", sys.Description())
}

// TestSystem_RegisterComponentFactory_Success tests registering a component factory successfully.
func TestSystem_RegisterComponentFactory_Success(t *testing.T) {
	// Create a new system with mocks
	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})
	factoryName := "testRegisterComponentFactory_Success"
	factory := func(ctx *context.Context, config *system.ComponentConfig) (system.Component, error) {
		// Mock factory implementation
		return new(mocks.MockSystemService), nil
	}

	// Register a component factory
	err := sys.RegisterComponentFactory(factoryName, factory)

	// Assert that no error occurred
	assert.NoError(t, err)
}

// TestSystem_RegisterComponentFactory_Error tests registering a component factory when a factory with the same name already exists.
func TestSystem_RegisterComponentFactory_Error(t *testing.T) {
	// Create a new system with mocks
	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})
	factoryName := "testRegisterComponentFactory_Error"
	factory := func(ctx *context.Context, config *system.ComponentConfig) (system.Component, error) {
		// Mock factory implementation
		return new(mocks.MockSystemService), nil
	}

	// Register a component factory
	err := sys.RegisterComponentFactory(factoryName, factory)
	assert.NoError(t, err)

	// Attempt to register the same component factory again
	err = sys.RegisterComponentFactory(factoryName, factory)

	// Assert that an error occurred and it matches the expected error
	assert.Error(t, err)
	assert.Equal(t, system.ErrComponentFactoryAlreadyExists, err)
}

// TestSystem_GetComponentFactory_Success tests retrieving a component factory.
func TestSystem_GetComponentFactory_Success(t *testing.T) {
	// Create a new system with mocks
	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})

	factoryName := "testGetComponentFactory_Success"

	// Define a custom type for the component factory function
	type ComponentFactoryFunc func(ctx *context.Context, config *system.ComponentConfig) (system.Component, error)

	// Create a factory function
	factoryFunc := ComponentFactoryFunc(func(ctx *context.Context, config *system.ComponentConfig) (system.Component, error) {
		// Mock factory implementation
		return new(mocks.MockSystemService), nil
	})

	// Conversion function to convert ComponentFactoryFunc to system.ComponentFactory
	convertFunc := func(ctx *context.Context, config *system.ComponentConfig) (system.Component, error) {
		return factoryFunc(ctx, config)
	}

	// Register a component factory
	err := sys.RegisterComponentFactory(factoryName, convertFunc)
	assert.NoError(t, err)

	// Retrieve the registered component factory
	fetchedFactory, err := sys.GetComponentFactory(factoryName)

	// Assert that no error occurred and the fetched factory matches the registered factory
	assert.NoError(t, err)
	assert.NotNil(t, fetchedFactory)
}

// TestSystem_GetComponentFactory_Error tests retrieving a component factory that does not exist.
func TestSystem_GetComponentFactory_Error(t *testing.T) {
	// Create a new system with mocks
	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})

	// Attempt to retrieve a component factory that does not exist
	_, err := sys.GetComponentFactory("nonexistentFactory")

	// Assert that the expected error occurred
	assert.Equal(t, system.ErrComponentNotFound, err)
}

// TestSystem_RegisterService_Success tests registering a service successfully.
func TestSystem_RegisterService_Success(t *testing.T) {
	// Create a new system with mocks
	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})
	serviceID := "testRegisterService_Success"
	service := new(mocks.MockSystemService)

	// Register a service
	err := sys.RegisterService(serviceID, service)

	// Assert that no error occurred during registration
	assert.NoError(t, err)
}

// TestSystem_RegisterService_Error tests registering a service that already exists.
func TestSystem_RegisterService_Error(t *testing.T) {
	// Create a new system with mocks
	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})
	serviceID := "testRegisterService_Error"
	service := new(mocks.MockSystemService)

	// Register a service
	err := sys.RegisterService(serviceID, service)
	assert.NoError(t, err)

	// Attempt to register the same service again
	err = sys.RegisterService(serviceID, service)

	// Assert that the expected error occurred
	assert.Error(t, err)
	assert.Equal(t, system.ErrServiceAlreadyExists, err)
}

// TestSystem_UnregisterService_Success tests unregistering a service successfully.
func TestSystem_UnregisterService_Success(t *testing.T) {
	// Create a new system with mocks
	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})
	serviceID := "testUnregisterService_Success"
	service := new(mocks.MockSystemService)

	// Register a service
	err := sys.RegisterService(serviceID, service)
	assert.NoError(t, err)

	// Unregister the service
	err = sys.UnregisterService(serviceID)

	// Assert that no error occurred during unregistering
	assert.NoError(t, err)
}

// TestSystem_UnregisterService_NotFound tests unregistering a service that does not exist.
func TestSystem_UnregisterService_NotFound(t *testing.T) {
	// Create a new system with mocks
	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})
	serviceID := "testUnregisterService_NotFound"
	service := new(mocks.MockSystemService)

	// Register a service
	err := sys.RegisterService(serviceID, service)
	assert.NoError(t, err)

	// Unregister the service
	err = sys.UnregisterService(serviceID)
	assert.NoError(t, err)

	// Attempt to unregister a nonexistent service
	err = sys.UnregisterService("nonexistentService")

	// Assert that the expected error occurred
	assert.Error(t, err)
	assert.Equal(t, system.ErrServiceNotRegistered, err)
}

// TestSystem_StartService_Success tests starting a service successfully.
func TestSystem_StartService_Success(t *testing.T) {
	// Create a new system with mocks
	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})
	serviceID := "testStartService_Success"
	service := new(mocks.MockSystemService)
	ctx := new(context.Context)

	// Register a service
	err := sys.RegisterService(serviceID, service)
	assert.NoError(t, err)

	// Mock the service's Start method to return nil
	service.On("Start", ctx).Return(nil)

	// Start the service
	err = sys.StartService(serviceID, ctx)

	// Assert that no error occurred during service start
	assert.NoError(t, err)
}

// TestSystem_StartService_NotRegistered tests starting a service that is not registered.
func TestSystem_StartService_NotRegistered(t *testing.T) {
	// Create a new system with mocks
	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})
	serviceID := "testStartService_NotRegistered"
	service := new(mocks.MockSystemService)
	ctx := new(context.Context)

	// Register a service
	err := sys.RegisterService(serviceID, service)
	assert.NoError(t, err)

	// Attempt to start a nonexistent service
	err = sys.StartService("nonexistentService", ctx)

	// Assert that the expected error occurred
	assert.Error(t, err)
	assert.Equal(t, system.ErrServiceNotRegistered, err)
}

// TestSystem_StopService_Success tests stopping a registered service successfully.
func TestSystem_StopService_Success(t *testing.T) {
	// Initialize system, service, and context
	serviceID := "testStopService_Success"
	service := new(mocks.MockSystemService)
	ctx := new(context.Context)
	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})

	// Register the service
	err := sys.RegisterService(serviceID, service)
	assert.NoError(t, err)

	// Mock the service's Stop method to return no error
	service.On("Stop", ctx).Return(nil)

	// Stop the service
	err = sys.StopService(serviceID, ctx)
	assert.NoError(t, err)
}

// TestSystem_StopService_NotRegistered tests stopping a service that is not registered.
func TestSystem_StopService_NotRegistered(t *testing.T) {
	// Initialize system and context
	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})
	ctx := new(context.Context)

	// Attempt to stop a nonexistent service
	err := sys.StopService("nonexistentService", ctx)

	// Assert error occurred and it's the correct error
	assert.Error(t, err)
	assert.Equal(t, system.ErrServiceNotRegistered, err)
}

// TestSystem_StopService_Error tests stopping a service that returns an error.
func TestSystem_StopService_Error(t *testing.T) {
	// Initialize system, service, and context
	serviceID := "testStopService_Error"
	service := new(mocks.MockSystemService)
	ctx := new(context.Context)
	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})

	// Register the service
	err := sys.RegisterService(serviceID, service)
	assert.NoError(t, err)

	// Define the expected error
	expectedErr := errors.New("service stop error")

	// Mock the service's Stop method to return an error
	service.On("Stop", ctx).Return(expectedErr)

	// Attempt to stop the service
	err = sys.StopService(serviceID, ctx)

	// Assert error occurred and it's the expected error
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
}

// TestSystem_RegisterOperation_Success tests registering an operation successfully.
func TestSystem_RegisterOperation_Success(t *testing.T) {
	// Initialize system and mocks
	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})
	operationID := "testRegisterOperation_Success"
	operation := new(mocks.MockOperation)

	// Register the operation
	err := sys.RegisterOperation(operationID, operation)

	// Assert no error occurred
	assert.NoError(t, err)
}

// TestSystem_RegisterOperation_AlreadyExists tests registering an operation that already exists.
func TestSystem_RegisterOperation_AlreadyExists(t *testing.T) {
	// Initialize system and mocks
	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})
	operationID := "testRegisterOperation_AlreadyExists"
	operation := new(mocks.MockOperation)

	// Register the operation twice to simulate the already existing scenario
	err := sys.RegisterOperation(operationID, operation)
	assert.NoError(t, err)

	// Attempt to register the same operation again
	err = sys.RegisterOperation(operationID, operation)

	// Assert error occurred and it's the correct error
	assert.Error(t, err)
	assert.Equal(t, system.ErrOperationAlreadyExists, err)
}

// TestSystem_UnregisterOperation_Success tests unregistering a registered operation successfully.
func TestSystem_UnregisterOperation_Success(t *testing.T) {
	// Initialize system and operation
	operationID := "testUnregisterOperation_Success"
	operation := new(mocks.MockOperation)
	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})

	// Register the operation
	err := sys.RegisterOperation(operationID, operation)
	assert.NoError(t, err)

	// Unregister the operation
	err = sys.UnregisterOperation(operationID)
	assert.NoError(t, err)
}

// TestSystem_UnregisterOperation_NotRegistered tests unregistering an operation that is not registered.
func TestSystem_UnregisterOperation_NotRegistered(t *testing.T) {
	// Initialize system
	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})

	// Attempt to unregister a nonexistent operation
	err := sys.UnregisterOperation("nonexistentOperation")

	// Assert error occurred and it's the correct error
	assert.Error(t, err)
	assert.Equal(t, system.ErrOperationNotRegistered, err)
}

// TestSystem_ExecuteOperation_Success tests executing a registered operation successfully.
func TestSystem_ExecuteOperation_Success(t *testing.T) {
	// Initialize system, operation, context, input, and output
	operationID := "testExecuteOperation_Success"
	operation := new(mocks.MockOperation)
	ctx := new(context.Context)
	input := &system.OperationInput{}
	output := &system.OperationOutput{}

	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})

	// Register the operation
	err := sys.RegisterOperation(operationID, operation)
	assert.NoError(t, err)

	// Mock the Execute method of the operation
	operation.On("Execute", ctx, input).Return(output, nil)

	// Execute the operation
	resOutput, err := sys.ExecuteOperation(ctx, operationID, input)

	// Assert no error occurred and the output matches the expected output
	assert.NoError(t, err)
	assert.Equal(t, output, resOutput)
}

// TestSystem_ExecuteOperation_NotRegistered tests executing an operation that is not registered.
func TestSystem_ExecuteOperation_NotRegistered(t *testing.T) {
	// Initialize system, context, and input
	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})
	ctx := new(context.Context)
	input := &system.OperationInput{}

	// Execute an operation that is not registered
	_, err := sys.ExecuteOperation(ctx, "nonexistentOperation", input)

	// Assert error occurred and it's the correct error
	assert.Error(t, err)
	assert.Equal(t, system.ErrOperationNotRegistered, err)
}

// TestSystem_ExecuteOperation_Error tests executing a registered operation that returns an error.
func TestSystem_ExecuteOperation_Error(t *testing.T) {
	// Initialize system, operation, context, and input
	operationID := "testExecuteOperation_Error"
	operation := new(mocks.MockOperation)
	ctx := new(context.Context)
	input := &system.OperationInput{}
	output := &system.OperationOutput{}
	expectedErr := errors.New("operation execution error")

	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})

	// Register the operation
	err := sys.RegisterOperation(operationID, operation)
	assert.NoError(t, err)

	// Mock the Execute method of the operation to return an error
	operation.On("Execute", ctx, input).Return(output, expectedErr)

	// Execute the operation
	_, err = sys.ExecuteOperation(ctx, operationID, input)

	// Assert an error occurred and it's the expected error
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
}

// TestSystem_Initialize_Success tests initializing the system successfully.
func TestSystem_Initialize_Success(t *testing.T) {
	// Initialize system and context
	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})
	ctx := new(context.Context)

	// Attempt to initialize the system
	err := sys.Initialize(ctx)

	// Assert no error occurred
	assert.NoError(t, err)
}

// TestSystem_Initialize_Error tests initializing the system when an error occurs.
func TestSystem_Initialize_Error(t *testing.T) {
	// Initialize system and context
	configuration := system.Configuration{
		Services: []*system.ServiceConfiguration{
			{
				ComponentConfig: system.ComponentConfig{
					ID:          "service1",
					FactoryName: "sysServiceFactory",
					Name:        "testService",
					Description: "Test Service",
				},
			},
		},
	}
	sys := system.NewSystem(eventBus, new(mocks.MockLogger), configuration)
	ctx := new(context.Context)
	// Simulate an error during initialization
	expectedErr := errors.New("failed to get component factory: component not found")

	// Attempt to initialize the system
	err := sys.Initialize(ctx)

	// Assert that the expected error occurred
	assert.EqualError(t, err, expectedErr.Error())
}

// TestSystem_Start_Success tests starting the system successfully.
func TestSystem_Start_Success(t *testing.T) {
	// Initialize system, service, and context
	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})
	service := new(mocks.MockSystemService)
	ctx := new(context.Context)

	// Register service
	sys.RegisterService("serviceStart_Success", service)

	// Mock service start function
	service.On("Start", ctx).Return(nil)

	// Attempt to start the system
	err := sys.Start(ctx)

	// Assert no error occurred
	assert.NoError(t, err)
}

// TestSystem_Start_Error tests starting the system when an error occurs.
func TestSystem_Start_Error(t *testing.T) {
	// Initialize system, service, and context
	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})
	service := new(mocks.MockSystemService)
	ctx := new(context.Context)

	// Register service
	sys.RegisterService("serviceStart_Error", service)

	// Simulate an error during service start
	expectedErr := errors.New("error during service start")
	service.On("Start", ctx).Return(expectedErr)

	// Attempt to start the system
	err := sys.Start(ctx)

	// Assert that the expected error occurred
	assert.Equal(t, err.Error(), "failed to start service: error during service start")
}

// TestSystem_Stop_Success tests stopping the system successfully.
func TestSystem_Stop_Success(t *testing.T) {
	// Initialize system, service, and context
	sys := system.NewSystem(eventBus, new(mocks.MockLogger), system.Configuration{})
	service := new(mocks.MockSystemService)
	ctx := new(context.Context)

	// Register service
	sys.RegisterService("serviceStop_Success", service)

	// Mock service stop function
	service.On("Stop", ctx).Return(nil)

	// Attempt to stop the system
	err := sys.Stop(ctx)

	// Assert no error occurred
	assert.NoError(t, err)
}
