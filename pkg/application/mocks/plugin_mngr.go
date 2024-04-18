package mocks

import (
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/stretchr/testify/mock"
)

// MockPluginManager is a mock implementation of PluginManagerInterface.
type MockPluginManager struct {
	mock.Mock
}

// AddPlugin mocks the AddPlugin method.
func (m *MockPluginManager) AddPlugin(ctx *context.Context, plugin system.PluginInterface) error {
	args := m.Called(plugin)
	return args.Error(0)
}

// RemovePlugin mocks the RemovePlugin method.
func (m *MockPluginManager) RemovePlugin(plugin system.PluginInterface) error {
	args := m.Called(plugin)
	return args.Error(0)
}

// GetPlugin mocks the GetPlugin method.
func (m *MockPluginManager) GetPlugin(name string) (system.PluginInterface, error) {
	args := m.Called(name)
	return args.Get(0).(system.PluginInterface), args.Error(1)
}

// StartPlugins mocks the StartPlugins method.
func (m *MockPluginManager) StartPlugins(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// StopPlugins mocks the StopPlugins method.
func (m *MockPluginManager) StopPlugins(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// DiscoverPlugins mocks the DiscoverPlugins method.
func (m *MockPluginManager) DiscoverPlugins(ctx *context.Context) ([]system.PluginInterface, error) {
	args := m.Called(ctx)
	return args.Get(0).([]system.PluginInterface), args.Error(1)
}

// LoadRemotePlugin mocks the LoadRemotePlugin method.
func (m *MockPluginManager) LoadRemotePlugin(ctx *context.Context, pluginURL string) (system.PluginInterface, error) {
	args := m.Called(ctx, pluginURL)
	return args.Get(0).(system.PluginInterface), args.Error(1)
}
