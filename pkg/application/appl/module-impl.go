package appl

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"plugin"
	"strings"
	"sync"

	"github.com/edward1christian/block-forge/pkg/application/context"
)

// moduleEntry represents a module entry in the module manager.
type moduleEntry struct {
	module Module
	name   string
}

// ModuleManagerImpl implements the ModuleManager interface.
type ModuleManagerImpl struct {
	started     bool
	modules     map[string]moduleEntry
	mutex       sync.RWMutex
	stopChan    chan struct{}
	application Application
	id          string
	name        string
	description string
}

// NewModuleManager creates a new instance of ModuleManager.
func NewModuleManager(id string, name string, description string) ModuleManager {
	return &ModuleManagerImpl{
		modules:     make(map[string]moduleEntry),
		stopChan:    make(chan struct{}),
		id:          id,
		name:        name,
		description: description,
	}
}

// ID returns the unique identifier of the component.
func (mm *ModuleManagerImpl) ID() string {
	return mm.id
}

// Name returns the name of the component.
func (mm *ModuleManagerImpl) Name() string {
	return mm.name
}

// Description returns the description of the component.
func (mm *ModuleManagerImpl) Description() string {
	return mm.description
}

// Initialize initializes the module manager.
func (mm *ModuleManagerImpl) Initialize(ctx *context.Context, app Application) error {
	// Initialize any global setup or resources required by the module manager.
	mm.application = app
	return nil
}

// AddModule adds a module to the module manager.
func (mm *ModuleManagerImpl) AddModule(module Module) error {
	mm.mutex.Lock()
	defer mm.mutex.Unlock()

	if _, exists := mm.modules[module.ID()]; exists {
		return fmt.Errorf("module with ID '%s' already exists", module.ID())
	}

	mm.modules[module.ID()] = moduleEntry{
		module: module,
		name:   module.Name(),
	}
	return nil
}

// RemoveModule removes a module from the module manager.
func (mm *ModuleManagerImpl) RemoveModule(name string) error {
	mm.mutex.Lock()
	defer mm.mutex.Unlock()

	for id, entry := range mm.modules {
		if entry.name == name {
			delete(mm.modules, id)
			return nil
		}
	}
	return fmt.Errorf("module with name '%s' not found", name)
}

// GetModule returns the module with the given name.
func (mm *ModuleManagerImpl) GetModule(name string) (Module, error) {
	mm.mutex.RLock()
	defer mm.mutex.RUnlock()

	for _, entry := range mm.modules {
		if entry.name == name {
			return entry.module, nil
		}
	}
	return nil, fmt.Errorf("module with name '%s' not found", name)
}

// StartModules starts all modules managed by the module manager.
func (mm *ModuleManagerImpl) StartModules(ctx *context.Context) error {
	mm.mutex.RLock()
	defer mm.mutex.RUnlock()

	if mm.started {
		return errors.New("module manager already started")
	}

	// Start each module concurrently.
	var wg sync.WaitGroup
	for _, entry := range mm.modules {
		wg.Add(1)
		go func(module Module) {
			defer wg.Done()
			if err := module.Start(ctx); err != nil {
				fmt.Printf("Failed to start module %s: %v\n", module.Name(), err)
			}
		}(entry.module)
	}

	// Wait for all modules to start.
	wg.Wait()
	mm.started = true
	return nil
}

// StopModules stops all modules managed by the module manager.
func (mm *ModuleManagerImpl) StopModules(ctx *context.Context) error {
	mm.mutex.Lock()
	defer mm.mutex.Unlock()

	if !mm.started {
		return errors.New("module manager not started")
	}

	// Signal all modules to stop.
	close(mm.stopChan)

	// Stop each module concurrently.
	var wg sync.WaitGroup
	for _, entry := range mm.modules {
		wg.Add(1)
		go func(module Module) {
			defer wg.Done()
			if err := module.Stop(ctx); err != nil {
				fmt.Printf("Failed to stop module %s: %v\n", module.Name(), err)
			}
		}(entry.module)
	}

	// Wait for all modules to stop.
	wg.Wait()
	mm.started = false
	return nil
}

// DiscoverModules discovers available modules within the system.
func (mm *ModuleManagerImpl) DiscoverModules(ctx *context.Context) ([]Module, error) {
	// Assuming modules are located in a specific directory
	modulesDir := "./modules"

	// Create a slice to store discovered modules
	var discoveredModules []Module

	// Open the modules directory
	files, err := os.ReadDir(modulesDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read modules directory: %v", err)
	}

	for _, file := range files {
		// Check if the file is a Go source file
		if !strings.HasSuffix(file.Name(), ".go") {
			continue
		}

		// Load the module source file
		modulePath := filepath.Join(modulesDir, file.Name())
		module, err := mm.loadModule(ctx, modulePath)
		if err != nil {
			fmt.Printf("Failed to load module %s: %v\n", file.Name(), err)
			continue
		}

		// Add the loaded module to the slice
		discoveredModules = append(discoveredModules, module)
	}

	return discoveredModules, nil
}

// LoadRemoteModule loads a module from a remote source.
func (mm *ModuleManagerImpl) LoadRemoteModule(ctx *context.Context, moduleURL string) (Module, error) {
	// Download the remote module file
	resp, err := http.Get(moduleURL)
	if err != nil {
		return nil, fmt.Errorf("failed to download remote module: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response is successful
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download remote module: status code %d", resp.StatusCode)
	}

	// Create a temporary file to store the downloaded module
	tempFile, err := ioutil.TempFile("", "module-*.go")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write the downloaded content to the temporary file
	_, err = io.Copy(tempFile, resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to write temporary file: %v", err)
	}

	// Load the module from the temporary file
	module, err := mm.loadModule(ctx, tempFile.Name())
	if err != nil {
		return nil, fmt.Errorf("failed to load remote module: %v", err)
	}

	return module, nil
}

// loadModule loads a module from a given file path.
// loadModule loads a module from a pre-compiled package.
func (mm *ModuleManagerImpl) loadModule(ctx *context.Context, packagePath string) (Module, error) {
	// Import the package
	pkg, err := plugin.Open(packagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open module package: %v", err)
	}

	// Look up the module symbol
	moduleSymbol, err := pkg.Lookup("Module")
	if err != nil {
		return nil, fmt.Errorf("failed to find module symbol: %v", err)
	}

	// Cast the module symbol to the Module interface
	module, ok := moduleSymbol.(Module)
	if !ok {
		return nil, fmt.Errorf("failed to cast module symbol to Module interface")
	}

	return module, nil
}
