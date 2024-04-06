package process_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/components"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/edward1christian/block-forge/pkg/etl"
	"github.com/edward1christian/block-forge/pkg/etl/common"
	"github.com/edward1christian/block-forge/pkg/etl/process"
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

func TestProcessManager_InitializeProcess_Success(t *testing.T) {
	ctx := &context.Context{}
	mockSystem := new(mocks.MockSystem)
	expectedOutput := &process.ETLProcess{ID: "123"}

	mockSystem.On("ExecuteOperation", ctx, common.ProcessOpInitializeETL, mock.Anything).
		Return(&systemApi.OperationOutput{Data: expectedOutput}, nil)

	manager := process.NewETLManagerService("TestMangerID", "TestMangerID", "TestManger ID")
	manager.Initialize(ctx, mockSystem)

	etlProcess, err := manager.InitializeProcess(ctx, config)

	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, etlProcess)
	mockSystem.AssertExpectations(t)
}

func TestProcessManager_InitializeProcess_Failure(t *testing.T) {
	ctx := &context.Context{}
	mockSystem := new(mocks.MockSystem)
	expectedOutput := &process.ETLProcess{ID: "123"}

	mockSystem.On("ExecuteOperation", ctx, common.ProcessOpInitializeETL, mock.Anything).
		Return(&systemApi.OperationOutput{Data: expectedOutput}, errors.New("error executing system operation"))

	manager := process.NewETLManagerService("TestMangerID", "TestMangerID", "TestManger ID")
	manager.Initialize(ctx, mockSystem)

	etlProcess, err := manager.InitializeProcess(ctx, config)

	assert.Error(t, err)
	assert.Nil(t, etlProcess)
	mockSystem.AssertExpectations(t)
}

func TestProcessManager_StartProcess_Success(t *testing.T) {
	ctx := &context.Context{}
	mockSystem := new(mocks.MockSystem)
	expectedOutput := &process.ETLProcess{ID: "123"}

	mockSystem.On("ExecuteOperation", ctx, common.ProcessOpInitializeETL, mock.Anything).
		Return(&systemApi.OperationOutput{Data: expectedOutput}, nil)

	mockSystem.On("ExecuteOperation", ctx, common.ProcessOpStartETL, mock.Anything).
		Return(&systemApi.OperationOutput{Data: expectedOutput}, nil)

	manager := process.NewETLManagerService("TestMangerID", "TestMangerID", "TestManger ID")

	manager.Initialize(ctx, mockSystem)

	etlProcess, err := manager.InitializeProcess(ctx, config)
	assert.NoError(t, err)

	err = manager.StartProcess(ctx, etlProcess.ID)

	assert.NoError(t, err)
	mockSystem.AssertExpectations(t)
}

func TestProcessManager_StartProcess_Failure(t *testing.T) {
	ctx := &context.Context{}
	mockSystem := new(mocks.MockSystem)
	expectedOutput := &process.ETLProcess{ID: "123"}

	mockSystem.On("ExecuteOperation", ctx, common.ProcessOpInitializeETL, mock.Anything).
		Return(&systemApi.OperationOutput{Data: expectedOutput}, nil)

	mockSystem.On("ExecuteOperation", ctx, common.ProcessOpStartETL, mock.Anything).
		Return(&systemApi.OperationOutput{Data: expectedOutput}, errors.New("error executing system operation"))

	manager := process.NewETLManagerService("TestMangerID", "TestMangerID", "TestManger ID")
	manager.Initialize(ctx, mockSystem)

	etlProcess, err := manager.InitializeProcess(ctx, config)
	assert.NoError(t, err)

	err = manager.StartProcess(ctx, etlProcess.ID)

	assert.Error(t, err)
	mockSystem.AssertExpectations(t)
}

func TestProcessManager_StopProcess_Success(t *testing.T) {
	ctx := &context.Context{}
	mockSystem := new(mocks.MockSystem)
	expectedOutput := &process.ETLProcess{ID: "123"}

	mockSystem.On("ExecuteOperation", ctx, common.ProcessOpInitializeETL, mock.Anything).
		Return(&systemApi.OperationOutput{Data: expectedOutput}, nil)

	mockSystem.On("ExecuteOperation", ctx, common.ProcessOpStartETL, mock.Anything).
		Return(&systemApi.OperationOutput{}, nil)

	mockSystem.On("ExecuteOperation", ctx, common.ProcessOpStopETL, mock.Anything).
		Return(&systemApi.OperationOutput{}, nil)

	manager := process.NewETLManagerService("TestMangerID", "TestMangerID", "TestManger ID")
	manager.Initialize(ctx, mockSystem)

	etlProcess, err := manager.InitializeProcess(ctx, config)
	assert.NoError(t, err)

	err = manager.StartProcess(ctx, etlProcess.ID)
	assert.NoError(t, err)

	err = manager.StopProcess(ctx, etlProcess.ID)

	assert.NoError(t, err)
	mockSystem.AssertExpectations(t)
}

func TestProcessManager_StopProcess_Failure(t *testing.T) {
	processID := "123"
	ctx := &context.Context{}
	mockSystem := new(mocks.MockSystem)

	manager := process.NewETLManagerService("TestMangerID", "TestMangerID", "TestManger ID")
	manager.Initialize(ctx, mockSystem)

	err := manager.StopProcess(ctx, processID)

	assert.Error(t, err)
	mockSystem.AssertExpectations(t)
}

func TestProcessManager_GetProcess_Success(t *testing.T) {
	processID := "123"
	ctx := &context.Context{}
	mockSystem := new(mocks.MockSystem)
	expectedOutput := &process.ETLProcess{ID: processID}

	mockSystem.On("ExecuteOperation", ctx, common.ProcessOpInitializeETL, mock.Anything).
		Return(&systemApi.OperationOutput{Data: expectedOutput}, nil)

	manager := process.NewETLManagerService("TestMangerID", "TestMangerID", "TestManger ID")
	manager.Initialize(ctx, mockSystem)

	etlProcess, err := manager.InitializeProcess(ctx, config)
	assert.NoError(t, err)

	retrievedProcess, err := manager.GetProcess(etlProcess.ID)
	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, retrievedProcess)
}

func TestProcessManager_GetProcess_NotFound(t *testing.T) {
	processID := "123"
	ctx := &context.Context{}
	mockSystem := new(mocks.MockSystem)

	manager := process.NewETLManagerService("TestMangerID", "TestMangerID", "TestManger ID")
	manager.Initialize(ctx, mockSystem)

	retrievedProcess, err := manager.GetProcess(processID)

	assert.Error(t, err)
	assert.Nil(t, retrievedProcess)
}

func TestProcessManager_GetAllProcesses(t *testing.T) {
	processID := "123"
	ctx := &context.Context{}
	mockSystem := new(mocks.MockSystem)
	expectedProcess := &process.ETLProcess{ID: processID}
	expectedProcesses := []*process.ETLProcess{expectedProcess}

	mockSystem.On("ExecuteOperation", ctx, common.ProcessOpInitializeETL, mock.Anything).
		Return(&systemApi.OperationOutput{Data: expectedProcess}, nil)

	manager := process.NewETLManagerService("TestMangerID", "TestMangerID", "TestManger ID")
	manager.Initialize(ctx, mockSystem)

	etlProcess, err := manager.InitializeProcess(ctx, config)
	assert.NoError(t, err)
	assert.NotNil(t, etlProcess)

	retrievedProcesses := manager.GetAllProcesses()

	assert.ElementsMatch(t, expectedProcesses, retrievedProcesses)
}

func TestProcessManager_RestartProcess_Success(t *testing.T) {
	ctx := &context.Context{}
	mockSystem := new(mocks.MockSystem)
	expectedOutput := &process.ETLProcess{ID: "123"}

	mockSystem.On("ExecuteOperation", ctx, common.ProcessOpInitializeETL, mock.Anything).
		Return(&systemApi.OperationOutput{Data: expectedOutput}, nil)

	mockSystem.On("ExecuteOperation", ctx, common.ProcessOpStartETL, mock.Anything).
		Return(&systemApi.OperationOutput{}, nil)

	mockSystem.On("ExecuteOperation", ctx, common.ProcessOpStopETL, mock.Anything).
		Return(&systemApi.OperationOutput{}, nil)

	manager := process.NewETLManagerService("TestMangerID", "TestMangerID", "TestManger ID")
	manager.Initialize(ctx, mockSystem)

	etlProcess, err := manager.InitializeProcess(ctx, config)
	assert.NoError(t, err)

	err = manager.StartProcess(ctx, etlProcess.ID)
	assert.NoError(t, err)

	err = manager.RestartProcess(ctx, etlProcess.ID)
	assert.NoError(t, err)
	mockSystem.AssertExpectations(t)
}

func TestProcessManager_RestartProcess_NotFound(t *testing.T) {
	processID := "123"
	ctx := &context.Context{}
	mockSystem := new(mocks.MockSystem)

	manager := process.NewETLManagerService("TestMangerID", "TestMangerID", "TestManger ID")
	manager.Initialize(ctx, mockSystem)

	err := manager.RestartProcess(ctx, processID)

	assert.Error(t, err)
	assert.Equal(t, etl.ErrNotProcessNotFound, err)
}

func TestProcessManager_RemoveProcess_Success(t *testing.T) {
	processID := "123"
	ctx := &context.Context{}
	mockSystem := new(mocks.MockSystem)
	expectedOutput := &process.ETLProcess{ID: processID}

	mockSystem.On("ExecuteOperation", ctx, common.ProcessOpInitializeETL, mock.Anything).
		Return(&systemApi.OperationOutput{Data: expectedOutput}, nil)

	manager := process.NewETLManagerService("TestMangerID", "TestMangerID", "TestManger ID")
	manager.Initialize(ctx, mockSystem)

	etlProcess, err := manager.InitializeProcess(ctx, config)

	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, etlProcess)

	err = manager.RemoveProcess(processID)

	assert.NoError(t, err)
}

func TestProcessManager_RemoveProcess_NotFound(t *testing.T) {
	processID := "123"
	ctx := &context.Context{}
	mockSystem := new(mocks.MockSystem)

	manager := process.NewETLManagerService("TestMangerID", "TestMangerID", "TestManger ID")
	manager.Initialize(ctx, mockSystem)

	err := manager.RemoveProcess(processID)

	assert.Error(t, err)
	assert.Equal(t, etl.ErrNotProcessNotFound, err)
}
