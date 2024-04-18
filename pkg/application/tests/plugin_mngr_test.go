package tests

import (
	"errors"
	"testing"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddPlugin_Success(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	pluginManager := systemApi.NewPluginManager()
	mockPlugin := new(mocks.MockPlugin)

	// Mock behavior
	mockPlugin.On("Initialize", mock.Anything, mock.Anything).Return(nil)
	mockPlugin.On("RegisterResources", ctx).Return(nil)
	mockPlugin.On("ID").Return("mock_plugin")

	// Act
	err := pluginManager.AddPlugin(ctx, mockPlugin)

	// Assert
	assert.NoError(t, err, "AddPlugin should not return an error")
}

func TestAddPlugin_Error_DuplicateID(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	pluginManager := systemApi.NewPluginManager()
	mockPlugin := new(mocks.MockPlugin)

	// Mock behavior
	mockPlugin.On("Initialize", mock.Anything, mock.Anything).Return(nil)
	mockPlugin.On("RegisterResources", ctx).Return(nil)
	mockPlugin.On("ID").Return("duplicate_plugin")

	// Add a plugin with the same ID first

	err := pluginManager.AddPlugin(ctx, mockPlugin)
	assert.NoError(t, err, "AddPlugin should not return an error")

	// Act
	err = pluginManager.AddPlugin(ctx, mockPlugin)

	// Assert
	assert.Error(t, err, "AddPlugin should return an error for duplicate plugin ID")
	assert.Contains(t, err.Error(), "already exists", "Error message should indicate duplicate ID")
}

func TestRemovePlugin_Success(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	pluginManager := systemApi.NewPluginManager()
	mockPlugin := new(mocks.MockPlugin)

	// Mock behavior
	mockPlugin.On("Initialize", mock.Anything, mock.Anything).Return(nil)
	mockPlugin.On("RegisterResources", ctx).Return(nil)
	mockPlugin.On("ID").Return("duplicate_plugin")

	// Add plugin to the manager
	err := pluginManager.AddPlugin(ctx, mockPlugin)
	assert.NoError(t, err, "AddPlugin should not return an error")

	// Act
	err = pluginManager.RemovePlugin(mockPlugin)

	// Assert
	assert.NoError(t, err, "RemovePlugin should not return an error")
}

func TestRemovePlugin_Error_NotFound(t *testing.T) {
	// Arrange
	pluginManager := systemApi.NewPluginManager()
	mockPlugin := new(mocks.MockPlugin)

	// Mock behavior
	mockPlugin.On("ID").Return("super_plugin")

	// Act
	err := pluginManager.RemovePlugin(mockPlugin)

	// Assert
	assert.Error(t, err, "RemovePlugin should return an error for non-existent plugin")
	assert.Contains(t, err.Error(), "not found", "Error message should indicate plugin not found")
}

func TestGetPlugin_Success(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	pluginManager := systemApi.NewPluginManager()
	mockPlugin := new(mocks.MockPlugin)

	// Mock behavior
	mockPlugin.On("Initialize", mock.Anything, mock.Anything).Return(nil)
	mockPlugin.On("RegisterResources", ctx).Return(nil)
	mockPlugin.On("ID").Return("super_plugin")

	// Add plugin to the manager
	err := pluginManager.AddPlugin(ctx, mockPlugin)
	assert.NoError(t, err, "AddPlugin should not return an error")

	// Act
	plugin, err := pluginManager.GetPlugin(mockPlugin.ID())

	// Assert
	assert.NoError(t, err, "GetPlugin should not return an error")
	assert.Equal(t, mockPlugin, plugin, "Retrieved plugin should match the added plugin")
}

func TestGetPlugin_Error_NotFound(t *testing.T) {
	// Arrange
	pluginManager := systemApi.NewPluginManager()

	// Act
	_, err := pluginManager.GetPlugin("non_existent_plugin")

	// Assert
	assert.Error(t, err, "GetPlugin should return an error for non-existent plugin")
	assert.Contains(t, err.Error(), "not found", "Error message should indicate plugin not found")
}

func TestStartPlugins_Success(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	pluginManager := systemApi.NewPluginManager()
	mockPlugin1 := new(mocks.MockPlugin)
	mockPlugin2 := new(mocks.MockPlugin)

	// Mock behavior
	mockPlugin1.On("Initialize", mock.Anything, mock.Anything).Return(nil)
	mockPlugin2.On("Initialize", mock.Anything, mock.Anything).Return(nil)

	mockPlugin1.On("RegisterResources", ctx).Return(nil)
	mockPlugin2.On("RegisterResources", ctx).Return(nil)

	mockPlugin1.On("ID").Return("super_plugin1")
	mockPlugin2.On("ID").Return("super_plugin2")

	// Mock behavior
	mockPlugin1.On("Start", mock.Anything).Return(nil)
	mockPlugin2.On("Start", mock.Anything).Return(nil)

	// Add plugins to the manager
	pluginManager.AddPlugin(ctx, mockPlugin1)
	pluginManager.AddPlugin(ctx, mockPlugin2)

	// Act
	err := pluginManager.StartPlugins(&context.Context{})

	// Assert
	assert.NoError(t, err, "StartPlugins should not return an error")
	mockPlugin1.AssertCalled(t, "Start", mock.Anything)
	mockPlugin2.AssertCalled(t, "Start", mock.Anything)
}

func TestStartPlugins_Error(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	pluginManager := systemApi.NewPluginManager()
	mockPlugin1 := new(mocks.MockPlugin)
	mockPlugin2 := new(mocks.MockPlugin)

	// Mock behavior
	mockPlugin1.On("Initialize", ctx, mock.Anything).Return(nil)
	mockPlugin2.On("Initialize", ctx, mock.Anything).Return(nil)

	mockPlugin1.On("RegisterResources", ctx).Return(nil)
	mockPlugin2.On("RegisterResources", ctx).Return(nil)

	mockPlugin1.On("ID").Return("super_plugin1")
	mockPlugin2.On("ID").Return("super_plugin2")

	mockPlugin1.On("Start", mock.Anything).Return(errors.New("start error"))
	mockPlugin2.On("Start", mock.Anything).Return(nil)

	// Add plugins to the manager
	pluginManager.AddPlugin(ctx, mockPlugin1)
	pluginManager.AddPlugin(ctx, mockPlugin2)

	// Act
	err := pluginManager.StartPlugins(&context.Context{})

	// Assert
	assert.Error(t, err, "StartPlugins should return an error")
	assert.Contains(t, err.Error(), "start error", "Error message should indicate start error")
	mockPlugin1.AssertCalled(t, "Start", mock.Anything)
}

func TestStopPlugins_Success(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	pluginManager := systemApi.NewPluginManager()
	mockPlugin1 := new(mocks.MockPlugin)
	mockPlugin2 := new(mocks.MockPlugin)

	// Mock behavior
	mockPlugin1.On("Initialize", ctx, mock.Anything).Return(nil)
	mockPlugin2.On("Initialize", ctx, mock.Anything).Return(nil)

	mockPlugin1.On("RegisterResources", ctx).Return(nil)
	mockPlugin2.On("RegisterResources", ctx).Return(nil)

	mockPlugin1.On("ID").Return("super_plugin1")
	mockPlugin2.On("ID").Return("super_plugin2")

	mockPlugin1.On("Stop", ctx).Return(nil)
	mockPlugin2.On("Stop", ctx).Return(nil)

	// Add plugins to the manager
	pluginManager.AddPlugin(ctx, mockPlugin1)
	pluginManager.AddPlugin(ctx, mockPlugin2)

	// Act
	err := pluginManager.StopPlugins(&context.Context{})

	// Assert
	assert.NoError(t, err, "StopPlugins should not return an error")
}

func TestStopPlugins_Error(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	pluginManager := systemApi.NewPluginManager()
	mockPlugin1 := new(mocks.MockPlugin)
	mockPlugin2 := new(mocks.MockPlugin)

	// Mock behavior
	mockPlugin1.On("Initialize", ctx, mock.Anything).Return(nil)
	mockPlugin2.On("Initialize", ctx, mock.Anything).Return(nil)

	mockPlugin1.On("RegisterResources", ctx).Return(nil)
	mockPlugin2.On("RegisterResources", ctx).Return(nil)

	mockPlugin1.On("ID").Return("super_plugin1")
	mockPlugin2.On("ID").Return("super_plugin2")

	mockPlugin1.On("Start", mock.Anything).Return(nil)
	mockPlugin2.On("Start", mock.Anything).Return(nil)

	mockPlugin1.On("Stop", ctx).Return(errors.New("stop error"))
	mockPlugin2.On("Stop", ctx).Return(nil)

	// Add plugins to the manager
	pluginManager.AddPlugin(ctx, mockPlugin1)
	pluginManager.AddPlugin(ctx, mockPlugin2)

	// Act
	err := pluginManager.StartPlugins(&context.Context{})
	assert.NoError(t, err, "StartPlugins should not return an error")

	err = pluginManager.StopPlugins(ctx)

	// Assert
	assert.Error(t, err, "StopPlugins should return an error")
	assert.Contains(t, err.Error(), "stop error", "Error message should indicate stop error")
	mockPlugin1.AssertCalled(t, "Stop", ctx)
}
