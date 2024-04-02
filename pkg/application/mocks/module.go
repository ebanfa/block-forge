package mocks

import (
	"github.com/edward1christian/block-forge/pkg/application/appl"
	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/stretchr/testify/mock"
)

// MockModule is a mock implementation of the Module interface for testing.
type MockModule struct {
	mock.Mock
}

func (m *MockModule) ID() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockModule) Name() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockModule) Description() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockModule) Initialize(ctx *context.Context, interfaces appl.Application) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockModule) Start(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockModule) Stop(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// MockModuleManager is a mock implementation of the MockModuleManager interface
type MockModuleManager struct {
	mock.Mock
}

func (m *MockModuleManager) ID() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockModuleManager) Name() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockModuleManager) Description() string {
	args := m.Called()
	return args.String(0)
}

// Initialize mocks the Initialize method of the MockModuleManager interface
func (m *MockModuleManager) Initialize(ctx *context.Context, app appl.Application) error {
	args := m.Called(ctx, app)
	return args.Error(0)
}

// AddModule mocks the AddModule method of the MockModuleManager interface
func (m *MockModuleManager) AddModule(module appl.Module) error {
	args := m.Called(module)
	return args.Error(0)
}

// RemoveModule mocks the RemoveModule method of the MockModuleManager interface
func (m *MockModuleManager) RemoveModule(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

// GetModule mocks the GetModule method of the MockModuleManager interface
func (m *MockModuleManager) GetModule(name string) (appl.Module, error) {
	args := m.Called(name)
	return args.Get(0).(appl.Module), args.Error(1)
}

// StartModules mocks the StartModules method of the MockModuleManager interface
func (m *MockModuleManager) StartModules(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// StopModules mocks the StopModules method of the MockModuleManager interface
func (m *MockModuleManager) StopModules(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// DiscoverModules mocks the DiscoverModules method of the MockModuleManager interface
func (m *MockModuleManager) DiscoverModules(ctx *context.Context) ([]appl.Module, error) {
	args := m.Called(ctx)
	return args.Get(0).([]appl.Module), args.Error(1)
}

// LoadRemoteModule mocks the LoadRemoteModule method of the MockModuleManager interface
func (m *MockModuleManager) LoadRemoteModule(ctx *context.Context, moduleURL string) (appl.Module, error) {
	args := m.Called(ctx, moduleURL)
	return args.Get(0).(appl.Module), args.Error(1)
}
