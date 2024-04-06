package plugin

import (
	"github.com/edward1christian/block-forge/pkg/application/appl"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// Plugin represents a plugin in the system.
type PluginInterface interface {
	system.SystemServiceInterface

	// RegisterEventHandlers registers event handlers provided by the plugin.
	RegisterEventHandlers() error

	// RegisterComponents registers additional operations provided by the plugin.
	RegisterComponents() error
}

// PluginManager represents functionality for managing plugins.
type PluginManagerInterface interface {
	// ModuleInterface includes methods for managing modules.
	appl.Module

	// AddPlugin adds a plugin to the plugin manager.
	AddPlugin(plugin PluginInterface) error

	// RemovePlugin removes a plugin from the plugin manager.
	RemovePlugin(name string) error

	// GetPlugin returns the plugin with the given name.
	GetPlugin(name string) (PluginInterface, error)

	// StartPlugins starts all plugins managed by the plugin manager.
	StartPlugins(ctx context.Context) error

	// StopPlugins stops all plugins managed by the plugin manager.
	StopPlugins(ctx context.Context) error

	// DiscoverPlugins discovers available plugins within the system.
	DiscoverPlugins(ctx context.Context) ([]PluginInterface, error)

	// LoadRemotePlugin loads a plugin from a remote source.
	LoadRemotePlugin(ctx context.Context, pluginURL string) (PluginInterface, error)
}
