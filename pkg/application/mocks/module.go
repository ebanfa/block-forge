package mocks

import (
	"github.com/edward1christian/block-forge/pkg/application/appl"
	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/stretchr/testify/mock"
)

// MockModule is a mock implementation of ApplicationComponent for testing purposes.
type MockModule struct {
	mock.Mock
}

// Start mocks the Start method of ApplicationComponent.
func (m *MockModule) Start(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// Stop mocks the Stop method of ApplicationComponent.
func (m *MockModule) Stop(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// Initialize mocks the Initialize method of ApplicationComponent.
func (m *MockModule) Initialize(ctx *context.Context, app appl.Application) error {
	args := m.Called(ctx, app)
	return args.Error(0)
}

// ID mocks the ID method of ApplicationComponent.
func (m *MockModule) ID() string {
	args := m.Called()
	return args.String(0)
}

// Name mocks the Name method of ApplicationComponent.
func (m *MockModule) Name() string {
	args := m.Called()
	return args.String(0)
}

// Description mocks the Description method of ApplicationComponent.
func (m *MockModule) Description() string {
	args := m.Called()
	return args.String(0)
}

// MockModuleManager is a mock implementation of ModuleManagerImpl for testing purposes.
type MockModuleManager struct {
	mock.Mock
}

// ID mocks the ID method of ModuleManagerImpl.
func (m *MockModuleManager) ID() string {
	args := m.Called()
	return args.String(0)
}

// Name mocks the Name method of ModuleManagerImpl.
func (m *MockModuleManager) Name() string {
	args := m.Called()
	return args.String(0)
}

// Description mocks the Description method of ModuleManagerImpl.
func (m *MockModuleManager) Description() string {
	args := m.Called()
	return args.String(0)
}

// Initialize mocks the Initialize method of ModuleManagerImpl.
func (m *MockModuleManager) Initialize(ctx *context.Context, app appl.Application) error {
	args := m.Called(ctx, app)
	return args.Error(0)
}

// AddModule mocks the AddModule method of ModuleManagerImpl.
func (m *MockModuleManager) AddModule(module appl.Module) error {
	args := m.Called(module)
	return args.Error(0)
}

// RemoveModule mocks the RemoveModule method of ModuleManagerImpl.
func (m *MockModuleManager) RemoveModule(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// GetModule mocks the GetModule method of ModuleManagerImpl.
func (m *MockModuleManager) GetModule(id string) (appl.Module, error) {
	args := m.Called(id)
	return args.Get(0).(appl.Module), args.Error(1)
}

// Start mocks the Start method of ModuleManagerImpl.
func (m *MockModuleManager) Start(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// Stop mocks the Stop method of ModuleManagerImpl.
func (m *MockModuleManager) Stop(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// StartModule mocks the StartModule method of ModuleManagerImpl.
func (m *MockModuleManager) StartModule(ctx *context.Context, name string) error {
	args := m.Called(ctx, name)
	return args.Error(0)
}

// StopModule mocks the StopModule method of ModuleManagerImpl.
func (m *MockModuleManager) StopModule(ctx *context.Context, name string) error {
	args := m.Called(ctx, name)
	return args.Error(0)
}

// DiscoverAndLoadModules mocks the DiscoverAndLoadModules method of ModuleManagerImpl.
func (m *MockModuleManager) DiscoverAndLoadModules(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// LoadRemoteModule mocks the LoadRemoteModule method of ModuleManagerImpl.
func (m *MockModuleManager) LoadRemoteModule(ctx *context.Context, moduleURL string) (appl.Module, error) {
	args := m.Called(ctx, moduleURL)
	return args.Get(0).(appl.Module), args.Error(1)
}
