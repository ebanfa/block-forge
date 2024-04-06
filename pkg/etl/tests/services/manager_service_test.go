package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
	"github.com/edward1christian/block-forge/pkg/etl"
	"github.com/edward1christian/block-forge/pkg/etl/components/services"
	etlMocks "github.com/edward1christian/block-forge/pkg/etl/mocks"
	"github.com/edward1christian/block-forge/pkg/etl/process"
	"github.com/edward1christian/block-forge/pkg/etl/tests"
)

func createMocks() (mockSystem *mocks.MockSystem, mockProcessManager *etlMocks.MockProcessManager) {
	mockSystem = &mocks.MockSystem{}
	mockProcessManager = &etlMocks.MockProcessManager{}
	config := tests.DummySystemConfiguration()

	mockSystem.On("Configuration").Return(config)

	return mockSystem, mockProcessManager
}
func TestProcessManagerService_Initialize(t *testing.T) {
	ctx := &context.Context{}
	mockSystem, mockProcessManager := createMocks()
	expectedOutput := &process.ETLProcess{ID: "123"}
	service := services.NewProcessManagerService("id", "name", "description", mockProcessManager)

	mockProcessManager.On("InitializeProcess", ctx, mock.Anything).Return(expectedOutput, nil)

	err := service.Initialize(ctx, mockSystem)
	assert.NoError(t, err)
}

func TestProcessManagerService_Initialize_Error(t *testing.T) {
	ctx := &context.Context{}
	mockSystem, mockProcessManager := createMocks()
	service := services.NewProcessManagerService("id", "name", "description", mockProcessManager)

	mockProcessManager.On(
		"InitializeProcess", ctx, mock.Anything).Return(nil, etl.ErrInvalidProcessesConfig)

	err := service.Initialize(ctx, mockSystem)
	assert.Error(t, err)
	assert.Equal(
		t, "failed to initialize ETL process: invalid processes configuration provided", err.Error())
}

func TestProcessManagerService_Start(t *testing.T) {
	ctx := &context.Context{}
	_, mockProcessManager := createMocks()
	expectedOutput := &process.ETLProcess{ID: "123"}
	service := services.NewProcessManagerService("id", "name", "description", mockProcessManager)

	mockProcessManager.On("StartProcess", ctx, expectedOutput.ID).Return(nil)
	mockProcessManager.On("GetAllProcesses").Return([]*process.ETLProcess{expectedOutput}, nil)

	err := service.Start(ctx)
	assert.NoError(t, err)
}

func TestProcessManagerService_Start_Error(t *testing.T) {
	ctx := &context.Context{}
	_, mockProcessManager := createMocks()
	expectedOutput := &process.ETLProcess{ID: "123"}
	service := services.NewProcessManagerService("id", "name", "description", mockProcessManager)

	mockProcessManager.On("StartProcess", ctx, expectedOutput.ID).Return(etl.ErrNotProcessNotFound)
	mockProcessManager.On("GetAllProcesses").Return([]*process.ETLProcess{expectedOutput}, nil)

	err := service.Start(ctx)
	assert.Error(t, err)
	assert.Equal(
		t, "failed to start ETL process: etl process not found", err.Error())
}

func TestProcessManagerService_Stop(t *testing.T) {
	ctx := &context.Context{}
	_, mockProcessManager := createMocks()
	expectedOutput := &process.ETLProcess{ID: "123"}
	service := services.NewProcessManagerService("id", "name", "description", mockProcessManager)

	mockProcessManager.On("StopProcess", ctx, expectedOutput.ID).Return(nil)
	mockProcessManager.On("GetAllProcesses").Return([]*process.ETLProcess{expectedOutput}, nil)

	err := service.Stop(ctx)
	assert.NoError(t, err)
}

func TestProcessManagerService_Stop_Error(t *testing.T) {
	ctx := &context.Context{}
	_, mockProcessManager := createMocks()
	expectedOutput := &process.ETLProcess{ID: "123"}
	service := services.NewProcessManagerService("id", "name", "description", mockProcessManager)

	mockProcessManager.On("StopProcess", ctx, expectedOutput.ID).Return(etl.ErrNotProcessNotFound)
	mockProcessManager.On("GetAllProcesses").Return([]*process.ETLProcess{expectedOutput}, nil)

	err := service.Stop(ctx)
	assert.Error(t, err)
	assert.EqualError(t, err, "failed to stop ETL process: etl process not found")
}
