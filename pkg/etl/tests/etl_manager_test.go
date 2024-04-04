package tests

import (
	"os"
	"testing"

	"github.com/edward1christian/block-forge/pkg/application/context"
	applMocks "github.com/edward1christian/block-forge/pkg/application/mocks"
	"github.com/edward1christian/block-forge/pkg/etl"
	"github.com/edward1christian/block-forge/pkg/etl/mocks"
	etlSystem "github.com/edward1christian/block-forge/pkg/etl/system"
	"github.com/stretchr/testify/assert"
)

var (
	system              etl.ETLSystem
	ctx                 *context.Context
	etlConfig           *etl.ETLConfig
	etlComponentFactory etl.ETLComponentFactory
	etlProcess          *etl.ETLProcess
	mockETLComponent    *mocks.MockETLComponent
)

func TestMain(m *testing.M) {
	ctx = context.Background()
	eventBus := &applMocks.MockEventBus{}
	configuration := system.Configuration{}
	system = etlSystem.NewETLSystem(eventBus, nil, nil, configuration)

	mockETLComponent = &mocks.MockETLComponent{}
	mockETLComponent.On("Initialize", ctx, system).Return(nil)

	etlConfig = &etl.ETLConfig{
		Components: []*etl.ETLComponentConfig{
			{Name: "Component1", FactoryNm: "mockFactory"},
		},
	}

	etlProcess = &etl.ETLProcess{
		Config: etlConfig,
		Components: map[string]etl.ETLComponent{
			"pipeline1": mockETLComponent,
		},
	}

	// Mock factory
	etlComponentFactory = func(ctx *context.Context, config *etl.ETLComponentConfig) (etl.ETLComponent, error) {
		// Mock factory implementation
		return mockETLComponent, nil
	}

	// Run tests
	exitCode := m.Run()

	// Clean up any global resources here if needed

	// Exit with the proper exit code
	os.Exit(exitCode)
}
func TestETLSystemImpl_RegisterETLComponentFactory_Success(t *testing.T) {
	// Test
	err := system.RegisterETLComponentFactory("mockFactory", etlComponentFactory)

	// Assertions
	assert.NoError(t, err)
}

func TestETLSystemImpl_RegisterETLComponentFactory_Error(t *testing.T) {
	// Attempt to register same factory again
	err := system.RegisterETLComponentFactory("mockFactory", etlComponentFactory)

	// Assertions
	assert.Error(t, err)
}

func TestETLSystemImpl_GetETLComponentFactory_Success(t *testing.T) {
	// Test
	factory, ok := system.GetETLComponentFactory("mockFactory")

	// Assertions
	assert.True(t, ok)
	assert.NotNil(t, factory)
}

func TestETLSystemImpl_GetETLComponentFactory_NotFound(t *testing.T) {
	// Test
	factory, ok := system.GetETLComponentFactory("non_existing_factory")

	// Assertions
	assert.False(t, ok)
	assert.Nil(t, factory)
}

func TestETLSystemImpl_InitializeETLProcess_Success(t *testing.T) {
	// Test
	process, err := system.InitializeETLProcess(ctx, etlConfig)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, process)
}

func TestETLSystemImpl_InitializeETLProcess_Error(t *testing.T) {
	// Test
	process, err := system.InitializeETLProcess(ctx, &etl.ETLConfig{})

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, process)
}

func TestETLSystemImpl_StartETLProcess_Success(t *testing.T) {
	// Setup
	mockETLComponent.On("Start", ctx).Return(nil)

	process, err := system.InitializeETLProcess(ctx, etlConfig)
	assert.NoError(t, err)

	// Test
	err = system.StartETLProcess(ctx, process.ID)

	// Assertions
	assert.NoError(t, err)
}

func TestETLSystemImpl_StartETLProcess_Error(t *testing.T) {
	// Test with non-existing process ID
	err := system.StartETLProcess(ctx, "non_existing_process_id")

	// Assertions
	assert.Error(t, err)
}

func TestETLSystemImpl_StopETLProcess_Success(t *testing.T) {
	// Setup
	mockETLComponent.On("Start", ctx).Return(nil)
	mockETLComponent.On("Stop", ctx).Return(nil)

	process, err := system.InitializeETLProcess(ctx, etlConfig)
	assert.NoError(t, err)

	// Test
	err = system.StopETLProcess(ctx, process.ID)

	// Assertions
	assert.NoError(t, err)
}

func TestETLSystemImpl_StopETLProcess_Error(t *testing.T) {
	// Test with non-existing process ID
	err := system.StopETLProcess(ctx, "non_existing_process_id")

	// Assertions
	assert.Error(t, err)
}

func TestETLSystemImpl_StopETLProcess_ProcessNotFound(t *testing.T) {
	// Test
	err := system.StopETLProcess(ctx, "non_existing_process_id")

	// Assertions
	assert.Equal(t, etl.ErrProcessNotFound, err)
	// Add more assertions as needed
}

func TestETLSystemImpl_GetETLProcess_Success(t *testing.T) {
	// Setup
	process, err := system.InitializeETLProcess(ctx, etlConfig)
	assert.NoError(t, err)

	err = system.StartETLProcess(ctx, process.ID)
	assert.NoError(t, err)

	// Test
	_, err = system.GetETLProcess(process.ID)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, process)
	//assert.Equal(t, retrievedProcess, process)
	// Add more assertions as needed
}

func TestETLSystemImpl_GetETLProcess_NotFound(t *testing.T) {
	// Test
	process, err := system.GetETLProcess("non_existing_process_id")

	// Assertions
	assert.Equal(t, etl.ErrProcessNotFound, err)
	assert.Nil(t, process)
	// Add more assertions as needed
}

func TestETLSystemImpl_GetAllETLProcesses(t *testing.T) {
	// Test
	processes := system.GetAllETLProcesses()

	// Assertions
	assert.Len(t, processes, 4)
	// Add more assertions as needed
}
