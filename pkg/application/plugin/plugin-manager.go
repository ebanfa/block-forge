package plugin

import (
	"sync"

	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"plugin"

	"github.com/edward1christian/block-forge/pkg/application/appl"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
)

// PluginManagerModule is a module that manages plugins in the system.
type PluginManagerModule struct {
	id          string
	name        string
	description string
	plugins     map[string]Plugin
	mu          sync.RWMutex
	started     bool
}

// NewPluginManagerModule creates a new instance of PluginManagerModule.
func NewPluginManagerModule(id, name, description string) *PluginManagerModule {
	return &PluginManagerModule{
		id:          id,
		name:        name,
		description: description,
		plugins:     make(map[string]Plugin),
	}
}

// ID returns the unique identifier of the component.
func (m *PluginManagerModule) ID() string {
	return m.id
}

// Name returns the name of the component.
func (m *PluginManagerModule) Name() string {
	return m.name
}

// Description returns the description of the component.
func (m *PluginManagerModule) Description() string {
	return m.description
}

// Start starts the module and all its plugins.
func (m *PluginManagerModule) Start(ctx *context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.started {
		return nil
	}

	// Start all plugins
	for _, plugin := range m.plugins {
		if err := plugin.Start(ctx); err != nil {
			return err
		}
	}

	m.started = true
	return nil
}

// Stop stops the module and all its plugins.
func (m *PluginManagerModule) Stop(ctx *context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.started {
		return nil
	}

	// Stop all plugins
	for _, plugin := range m.plugins {
		if err := plugin.Stop(ctx); err != nil {
			return err
		}
	}

	m.started = false
	return nil
}

// Initialize initializes the module with the given context and application instance.
func (m *PluginManagerModule) Initialize(ctx *context.Context, app appl.ApplicationInterface) error {
	// Discover and load plugins
	plugins, err := m.DiscoverPlugins(ctx)
	if err != nil {
		return err
	}

	for _, plugin := range plugins {
		// Initialize the plugin
		if err := plugin.Initialize(ctx, app.System()); err != nil {
			return err
		}

		// Add the initialized plugin to the plugin manager
		if err := m.AddPlugin(plugin); err != nil {
			return err
		}
	}

	return nil
}

// AddPlugin adds a plugin to the plugin manager.
func (m *PluginManagerModule) AddPlugin(plugin Plugin) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.plugins[plugin.ID()]; exists {
		return errors.New("plugin already exists")
	}

	m.plugins[plugin.ID()] = plugin
	return nil
}

// RemovePlugin removes a plugin from the plugin manager.
func (m *PluginManagerModule) RemovePlugin(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.plugins[id]; !exists {
		return errors.New("plugin not found")
	}

	delete(m.plugins, id)
	return nil
}

// GetPlugin returns the plugin with the given id.
func (m *PluginManagerModule) GetPlugin(id string) (Plugin, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	plugin, exists := m.plugins[id]
	if !exists {
		return nil, errors.New("plugin not found")
	}

	return plugin, nil
}

// StartPlugins starts all plugins managed by the plugin manager.
func (m *PluginManagerModule) StartPlugins(ctx *context.Context) error {
	return m.Start(ctx)
}

// StopPlugins stops all plugins managed by the plugin manager.
func (m *PluginManagerModule) StopPlugins(ctx *context.Context) error {
	return m.Stop(ctx)
}

// PluginManagerModuleConfig represents the configuration for the PluginManagerModule.
type PluginManagerModuleConfig struct {
	PluginDir     string   `json:"plugin_dir"`
	RemotePlugins []string `json:"remote_plugins"`
}

func (m *PluginManagerModule) DiscoverPlugins(ctx *context.Context) ([]Plugin, error) {
	var plugins []Plugin

	// Get the module configuration
	config, err := m.getModuleConfig()
	if err != nil {
		return nil, err
	}

	pluginManagerConfig, ok := config.(PluginManagerModuleConfig)
	if !ok {
		return nil, errors.New("invalid module configuration type")
	}

	// Scan the plugin directory
	err = filepath.Walk(pluginManagerConfig.PluginDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".so" {
			plugin, err := m.loadPlugin(path)
			if err != nil {
				return err
			}
			plugins = append(plugins, plugin)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Load remote plugins
	for _, remoteURL := range pluginManagerConfig.RemotePlugins {
		plugin, err := m.LoadRemotePlugin(ctx, remoteURL)
		if err != nil {
			return nil, err
		}
		plugins = append(plugins, plugin)
	}

	return plugins, nil
}

func (m *PluginManagerModule) LoadRemotePlugin(ctx *context.Context, pluginURL string) (Plugin, error) {
	// Download the compressed plugin archive
	resp, err := http.Get(pluginURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Create a temporary file to store the downloaded archive
	tempFile, err := os.CreateTemp("", "plugin-*.zip")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tempFile.Name())

	// Write the downloaded content to the temporary file
	_, err = io.Copy(tempFile, resp.Body)
	if err != nil {
		return nil, err
	}

	// Extract the plugin from the archive
	plugin, err := m.loadPlugin(tempFile.Name())
	if err != nil {
		return nil, err
	}

	return plugin, nil
}

func (m *PluginManagerModule) getModuleConfig() (interface{}, error) {
	// Implement logic to retrieve the module configuration from the application configuration
	// This could involve reading a configuration file, querying a database, etc.
	return PluginManagerModuleConfig{
		PluginDir: "./",
		//RemotePlugins: []string{"https://example.com/plugin1.zip", "https://example.com/plugin2.zip"},
		RemotePlugins: []string{},
	}, nil
}

func (m *PluginManagerModule) loadPlugin(pluginPath string) (Plugin, error) {
	// Load the plugin using the Go plugin system
	plug, err := plugin.Open(pluginPath)
	if err != nil {
		return nil, err
	}

	// Lookup the symbol for the plugin factory function
	symPlugin, err := plug.Lookup("NewPlugin")
	if err != nil {
		return nil, err
	}

	// Create an instance of the plugin
	pluginFactory, ok := symPlugin.(func() (Plugin, error))
	if !ok {
		return nil, errors.New("invalid plugin factory function")
	}

	plugin, err := pluginFactory()
	if err != nil {
		return nil, err
	}

	return plugin, nil
}

func (m *PluginManagerModule) getConfiguredPlugins() []string {
	// Implement logic to retrieve the list of configured plugins
	// This could involve reading a configuration file, querying a database, etc.
	return []string{"path/to/plugin1.so", "path/to/plugin2.so"}
}
