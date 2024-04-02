package plugin

import (
	"github.com/edward1christian/block-forge/pkg/application/appl"
	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// Plugin represents a plugin in the system.
type Plugin interface {
	system.Startable
	system.SystemComponent

	// RegisterEventHandlers registers event handlers provided by the plugin.
	RegisterEventHandlers() error

	// RegisterOperations registers additional operations provided by the plugin.
	RegisterOperations() error

	// Configure configures the plugin with the provided configuration.
	Configure(config interface{}) error

	// Dependencies returns a list of dependencies required by the plugin.
	Dependencies() []string

	// OnInitialize is called during plugin initialization.
	OnInitialize() error

	// OnStart is called when the plugin starts.
	OnStart() error

	// OnStop is called when the plugin stops.
	OnStop() error

	// LogError logs an error that occurred during plugin execution.
	LogError(err error)
}

// PluginManager represents functionality for managing plugins.
type PluginManager interface {
	// ModuleInterface includes methods for managing modules.
	appl.Module

	// AddPlugin adds a plugin to the plugin manager.
	AddPlugin(plugin Plugin) error

	// RemovePlugin removes a plugin from the plugin manager.
	RemovePlugin(name string) error

	// GetPlugin returns the plugin with the given name.
	GetPlugin(name string) (Plugin, error)

	// StartPlugins starts all plugins managed by the plugin manager.
	StartPlugins(ctx context.Context) error

	// StopPlugins stops all plugins managed by the plugin manager.
	StopPlugins(ctx context.Context) error

	// DiscoverPlugins discovers available plugins within the system.
	DiscoverPlugins(ctx context.Context) ([]Plugin, error)

	// LoadRemotePlugin loads a plugin from a remote source.
	LoadRemotePlugin(ctx context.Context, pluginURL string) (Plugin, error)
}
