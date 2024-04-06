package operations_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/components"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
	etlOpsApi "github.com/edward1christian/block-forge/pkg/etl/components/operations"
	"github.com/edward1christian/block-forge/pkg/etl/process"

	"github.com/edward1christian/block-forge/pkg/application/mocks"
	etlMocksApi "github.com/edward1christian/block-forge/pkg/etl/mocks"
)

var (
	component1 = &components.ComponentConfig{
		ID:           "AdapterID",
		Name:         "Adapter",
		Description:  "Extracts data from source",
		FactoryName:  "AdaptorFactory",
		CustomConfig: map[string]interface{}{"param1": "value1", "param2": 123},
	}

	component2 = &components.ComponentConfig{
		ID:           "TransformerID",
		Name:         "Transformer",
		Description:  "Transforms extracted data",
		FactoryName:  "TransformerFactory",
		CustomConfig: map[string]interface{}{"param3": "value3", "param4": 456},
	}
)

// Demo data for ETLProcessConfig
var config = &process.ETLProcessConfig{
	Components: []*components.ComponentConfig{component1, component2},
}

// createMocks creates and returns mock objects required for testing.
func createMocks() (
	*mocks.MockSystem, *mocks.MockIDGenerator,
	*mocks.MockComponentFactory, *mocks.MockComponentRegistrar,
	*etlMocksApi.MockETLProcessComponent, *etlMocksApi.MockETLProcessComponent) {

	mockSystem := &mocks.MockSystem{}
	mockIDGenerator := &mocks.MockIDGenerator{}
	mockFactory := &mocks.MockComponentFactory{}
	mockRegistrar := &mocks.MockComponentRegistrar{}

	adapterComponent := &etlMocksApi.MockETLProcessComponent{}
	transformerComponent := &etlMocksApi.MockETLProcessComponent{}

	adapterComponent.On("ID").Return("TransformerID", nil)
	transformerComponent.On("ID").Return("TransformerID", nil)

	// Setup mock function calls
	mockIDGenerator.On("GenerateID").Return("123456", nil)
	mockSystem.On("ComponentRegistry").Return(mockRegistrar)

	mockRegistrar.On("GetComponentFactory", "AdaptorFactory").Return(mockFactory, nil)
	mockRegistrar.On("GetComponentFactory", "TransformerFactory").Return(mockFactory, nil)

	return mockSystem, mockIDGenerator, mockFactory, mockRegistrar, adapterComponent, transformerComponent
}

// TestInitializeProcessOperation_Initialize tests the Initialize method of InitializeETLProcessOperation.
func TestInitializeProcessOperation_Initialize(t *testing.T) {
	// Mocks
	ctx := &context.Context{}
	mockSystem, mockIDGenerator, _, _, _, _ := createMocks()

	// Create InitializeETLProcessOperation instance
	ie := etlOpsApi.NewInitializeETLProcessOperation(
		"1", "TestOperation", "Test Description", mockIDGenerator)

	// Call Initialize method
	err := ie.Initialize(ctx, mockSystem)

	// Check if the Initialize method returns no error
	assert.NoError(t, err)
}

// TestInitializeProcessOperation_InitializeProcess_Success tests the InitializeProcess method of InitializeETLProcessOperation with successful execution.
func TestInitializeProcessOperation_InitializeProcess_Success(t *testing.T) {
	// Mocks
	ctx := &context.Context{}
	mockSystem, mockIDGenerator, mockFactory, mockRegistrar, adapterComponent, transformerComponent := createMocks()

	mockFactory.On("CreateComponent", mock.Anything).Return(adapterComponent, nil)
	mockFactory.On("CreateComponent", mock.Anything).Return(transformerComponent, nil)

	// Create and initialize the InitializeETLProcessOperation instance
	ie := etlOpsApi.NewInitializeETLProcessOperation(
		"123456", "TestOperation", "Test Description", mockIDGenerator)

	ie.Initialize(ctx, mockSystem)

	// Call InitializeProcess method
	etlProcess, err := ie.InitializeProcess(ctx, config)

	// Check if the InitializeProcess method returns no error
	assert.NoError(t, err)

	// Assert that the ETL process is initialized correctly
	assert.NotNil(t, etlProcess)
	assert.Equal(t, "123456", etlProcess.ID)

	// Assert that mock functions were called
	mockSystem.AssertExpectations(t)
	mockRegistrar.AssertExpectations(t)
	mockFactory.AssertExpectations(t)
}

// TestInitializeProcessOperation_InitializeETLProcess_Error tests the InitializeProcess method of InitializeETLProcessOperation with an error returned during initialization.
func TestInitializeProcessOperation_InitializeETLProcess_Error(t *testing.T) {
	// Mocks
	ctx := &context.Context{}
	mockSystem, mockIDGenerator, mockFactory, mockRegistrar, adapterComponent, transformerComponent := createMocks()

	mockFactory.On("CreateComponent", component1).Return(adapterComponent, nil)
	mockFactory.On("CreateComponent", component2).Return(
		transformerComponent, errors.New("error creating component"))

	// Create InitializeETLProcessOperation instance
	ie := etlOpsApi.NewInitializeETLProcessOperation(
		"1", "TestOperation", "Test Description", mockIDGenerator)

	ie.Initialize(ctx, mockSystem)

	// Call InitializeProcess method
	etlProcess, err := ie.InitializeProcess(ctx, config)

	// Check if the InitializeProcess method returns the expected error
	assert.Error(t, err)
	assert.Nil(t, etlProcess)

	// Assert that the ComponentRegistry and CreateComponent methods are called
	mockSystem.AssertExpectations(t)
	mockRegistrar.AssertExpectations(t)
	mockFactory.AssertExpectations(t)
}

// TestInitializeProcessOperation_Execute tests the Execute method of InitializeETLProcessOperation.
func TestInitializeProcessOperation_Execute_Success(t *testing.T) {
	// Mocks
	ctx := &context.Context{}
	mockSystem, mockIDGenerator, mockFactory, _, adapterComponent, transformerComponent := createMocks()

	mockFactory.On("CreateComponent", mock.Anything).Return(adapterComponent, nil)
	mockFactory.On("CreateComponent", mock.Anything).Return(transformerComponent, nil)

	// Create InitializeETLProcessOperation instance
	ie := etlOpsApi.NewInitializeETLProcessOperation(
		"1", "TestOperation", "Test Description", mockIDGenerator)

	ie.Initialize(ctx, mockSystem)

	// Call Execute method
	output, err := ie.Execute(ctx, &systemApi.OperationInput{
		Data: []*process.ETLProcessConfig{config},
	})

	// Check if the Execute method returns no error
	assert.NoError(t, err)
	assert.Nil(t, output)

}

// TestInitializeProcessOperation_Execute_Error tests the Execute method of InitializeETLProcessOperation with an error returned during execution.
func TestInitializeProcessOperation_Execute_Error(t *testing.T) {
	// Mocks
	ctx := &context.Context{}
	mockSystem, mockIDGenerator, mockFactory, _, adapterComponent, transformerComponent := createMocks()

	mockFactory.On("CreateComponent", component1).Return(adapterComponent, nil)
	mockFactory.On("CreateComponent", component2).Return(transformerComponent,
		errors.New("error creating component"))

	// Simulate error during Execute
	adapterComponent.On("Start", ctx).Return(errors.New("error starting adapter"))

	// Create InitializeETLProcessOperation instance
	ie := etlOpsApi.NewInitializeETLProcessOperation(
		"1", "TestOperation", "Test Description", mockIDGenerator)

	ie.Initialize(ctx, mockSystem)

	// Call Execute method
	output, err := ie.Execute(ctx, &systemApi.OperationInput{
		Data: []*process.ETLProcessConfig{config},
	})

	// Check if the Execute method returns an error
	assert.Error(t, err)
	assert.Nil(t, output)
}
