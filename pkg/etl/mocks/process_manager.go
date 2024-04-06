package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/etl/process"
)

// MockProcessManager is a mock implementation of the ETLManagerService interface.
type MockProcessManager struct {
	mock.Mock
}

// InitializeProcess initializes the ETLManagerService.
func (m *MockProcessManager) InitializeProcess(ctx *context.Context, config *process.ETLProcessConfig) (*process.ETLProcess, error) {
	args := m.Called(ctx, config)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*process.ETLProcess), args.Error(1)
}

// StartProcess starts an ETL process with the given ID.
func (m *MockProcessManager) StartProcess(ctx *context.Context, processID string) error {
	args := m.Called(ctx, processID)
	return args.Error(0)
}

// StopProcess stops an ETL process with the given ID.
func (m *MockProcessManager) StopProcess(ctx *context.Context, processID string) error {
	args := m.Called(ctx, processID)
	return args.Error(0)
}

// RestartProcess restarts an ETL process with the given ID.
func (m *MockProcessManager) RestartProcess(ctx *context.Context, processID string) error {
	args := m.Called(ctx, processID)
	return args.Error(0)
}

// GetProcess retrieves an ETL process by its ID.
func (m *MockProcessManager) GetProcess(processID string) (*process.ETLProcess, error) {
	args := m.Called(processID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*process.ETLProcess), args.Error(1)
}

// GetAllProcesses retrieves all ETL processes.
func (m *MockProcessManager) GetAllProcesses() []*process.ETLProcess {
	args := m.Called()
	return args.Get(0).([]*process.ETLProcess)
}

// RemoveProcess removes an ETL process with the given ID.
func (m *MockProcessManager) RemoveProcess(processID string) error {
	args := m.Called(processID)
	return args.Error(0)
}
