package appl

import (
	"archive/zip"
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

// ModuleManagerImpl implements the ModuleManager interface.
type ModuleManagerImpl struct {
	id          string
	name        string
	description string
	started     bool
	modules     map[string]Module
	mutex       sync.RWMutex
	stopChan    chan struct{}
	application Application
}

// NewModuleManager creates a new instance of ModuleManager.
func NewModuleManager(id, name, description string) ModuleManager {
	return &ModuleManagerImpl{
		id:          id,
		name:        name,
		description: description,
		modules:     make(map[string]Module),
		stopChan:    make(chan struct{}),
	}
}

// ID returns the unique identifier of the module manager.
func (mm *ModuleManagerImpl) ID() string {
	return mm.id
}

// Name returns the name of the module manager.
func (mm *ModuleManagerImpl) Name() string {
	return mm.name
}

// Description returns the description of the module manager.
func (mm *ModuleManagerImpl) Description() string {
	return mm.description
}

// Initialize initializes the module manager.
func (mm *ModuleManagerImpl) Initialize(ctx *context.Context, app Application) error {
	// Initialize any global setup or resources required by the module manager.
	mm.application = app

	// Discover and load modules
	// You may need to adjust this based on your actual implementation
	if err := mm.DiscoverAndLoadModules(ctx); err != nil {
		return err
	}

	return nil
}

// AddModule adds a module to the module manager.
func (mm *ModuleManagerImpl) AddModule(module Module) error {
	mm.mutex.Lock()
	defer mm.mutex.Unlock()

	if _, exists := mm.modules[module.ID()]; exists {
		return fmt.Errorf("module with ID '%s' already exists", module.ID())
	}

	mm.modules[module.ID()] = module
	return nil
}

// RemoveModule removes a module from the module manager.
func (mm *ModuleManagerImpl) RemoveModule(id string) error {
	mm.mutex.Lock()
	defer mm.mutex.Unlock()

	if _, exists := mm.modules[id]; !exists {
		return fmt.Errorf("module with ID '%s' not found", id)
	}

	delete(mm.modules, id)
	return nil
}

// GetModule returns the module with the given ID.
func (mm *ModuleManagerImpl) GetModule(id string) (Module, error) {
	mm.mutex.RLock()
	defer mm.mutex.RUnlock()

	module, exists := mm.modules[id]
	if !exists {
		return nil, fmt.Errorf("module with ID '%s' not found", id)
	}

	return module, nil
}

// StartModules starts all modules managed by the module manager.
func (mm *ModuleManagerImpl) Start(ctx *context.Context) error {
	mm.mutex.RLock()
	defer mm.mutex.RUnlock()

	if mm.started {
		return errors.New("module manager already started")
	}

	// Start each module concurrently.
	var wg sync.WaitGroup
	for _, module := range mm.modules {
		wg.Add(1)
		go func(module Module) {
			defer wg.Done()
			if err := module.Start(ctx); err != nil {
				fmt.Printf("Failed to start module %s: %v\n", module.Name(), err)
			}
		}(module)
	}

	// Wait for all modules to start.
	wg.Wait()
	mm.started = true
	return nil
}

// StopModules stops all modules managed by the module manager.
func (mm *ModuleManagerImpl) Stop(ctx *context.Context) error {
	mm.mutex.Lock()
	defer mm.mutex.Unlock()

	if !mm.started {
		return errors.New("module manager not started")
	}

	// Signal all modules to stop.
	close(mm.stopChan)

	// Stop each module concurrently.
	var wg sync.WaitGroup
	for _, module := range mm.modules {
		wg.Add(1)
		go func(module Module) {
			defer wg.Done()
			if err := module.Stop(ctx); err != nil {
				fmt.Printf("Failed to stop module %s: %v\n", module.Name(), err)
			}
		}(module)
	}

	// Wait for all modules to stop.
	wg.Wait()
	mm.started = false
	return nil
}

// StartModule starts the module with the given name.
func (mm *ModuleManagerImpl) StartModule(ctx *context.Context, name string) error {
	mm.mutex.RLock()
	defer mm.mutex.RUnlock()

	module, exists := mm.modules[name]
	if !exists {
		return fmt.Errorf("module with name '%s' not found", name)
	}

	if err := module.Start(ctx); err != nil {
		return fmt.Errorf("failed to start module %s: %v", name, err)
	}

	return nil
}

// StopModule stops the module with the given name.
func (mm *ModuleManagerImpl) StopModule(ctx *context.Context, name string) error {
	mm.mutex.RLock()
	defer mm.mutex.RUnlock()

	module, exists := mm.modules[name]
	if !exists {
		return fmt.Errorf("module with name '%s' not found", name)
	}

	if err := module.Stop(ctx); err != nil {
		return fmt.Errorf("failed to stop module %s: %v", name, err)
	}

	return nil
}

// DiscoverAndLoadModules discovers and loads available modules within the system.
func (mm *ModuleManagerImpl) DiscoverAndLoadModules(ctx *context.Context) error {
	// Implement module discovery and loading logic based on your requirements
	// Example implementation using filepath.Walk to find modules in a directory
	modulesDir := "./modules"

	files, err := ioutil.ReadDir(modulesDir)
	if err != nil {
		return fmt.Errorf("failed to read modules directory: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".zip") {
			zipPath := filepath.Join(modulesDir, file.Name())

			// Decompress the module zip file
			moduleDir, err := ioutil.TempDir("", "module-dir-*")
			if err != nil {
				return fmt.Errorf("failed to create temporary directory for module: %v", err)
			}
			defer os.RemoveAll(moduleDir)

			// Open the module zip file
			zipReader, err := zip.OpenReader(zipPath)
			if err != nil {
				return fmt.Errorf("failed to open module zip file: %v", err)
			}
			defer zipReader.Close()

			// Extract all files from the zip archive
			for _, zipFile := range zipReader.File {
				// Open each file in the zip archive
				zipFileReader, err := zipFile.Open()
				if err != nil {
					return fmt.Errorf("failed to open file in module zip: %v", err)
				}
				defer zipFileReader.Close()

				// Create the file on disk
				extractedFilePath := filepath.Join(moduleDir, zipFile.Name)
				extractedFile, err := os.Create(extractedFilePath)
				if err != nil {
					return fmt.Errorf("failed to create extracted file: %v", err)
				}
				defer extractedFile.Close()

				// Copy the file contents
				_, err = io.Copy(extractedFile, zipFileReader)
				if err != nil {
					return fmt.Errorf("failed to extract file: %v", err)
				}
			}

			// Load the module from the extracted directory
			module, err := mm.loadModule(moduleDir)
			if err != nil {
				fmt.Printf("Failed to load module %s: %v\n", file.Name(), err)
				continue
			}

			if err := mm.AddModule(module); err != nil {
				fmt.Printf("Failed to add module %s: %v\n", module.Name(), err)
				continue
			}
		}
	}

	return nil
}

// LoadRemoteModule loads a module from a remote source.
func (mm *ModuleManagerImpl) LoadRemoteModule(ctx *context.Context, moduleURL string) (Module, error) {
	// Download the module zip file
	resp, err := http.Get(moduleURL)
	if err != nil {
		return nil, fmt.Errorf("failed to download remote module: %v", err)
	}
	defer resp.Body.Close()

	// Create a temporary file to store the downloaded module zip
	tempZipFile, err := ioutil.TempFile("", "module-*.zip")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary zip file: %v", err)
	}
	defer os.Remove(tempZipFile.Name())

	// Write the downloaded content to the temporary zip file
	_, err = io.Copy(tempZipFile, resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to write temporary zip file: %v", err)
	}

	// Decompress the module zip file
	moduleDir, err := ioutil.TempDir("", "module-dir-*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary directory for module: %v", err)
	}
	defer os.RemoveAll(moduleDir)

	// Open the module zip file
	zipReader, err := zip.OpenReader(tempZipFile.Name())
	if err != nil {
		return nil, fmt.Errorf("failed to open module zip file: %v", err)
	}
	defer zipReader.Close()

	// Extract all files from the zip archive
	for _, file := range zipReader.File {
		// Open each file in the zip archive
		fileReader, err := file.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file in module zip: %v", err)
		}
		defer fileReader.Close()

		// Create the file on disk
		extractedFilePath := filepath.Join(moduleDir, file.Name)
		extractedFile, err := os.Create(extractedFilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to create extracted file: %v", err)
		}
		defer extractedFile.Close()

		// Copy the file contents
		_, err = io.Copy(extractedFile, fileReader)
		if err != nil {
			return nil, fmt.Errorf("failed to extract file: %v", err)
		}
	}

	// Load the module from the extracted directory
	return mm.loadModule(moduleDir)
}

// loadModule loads a module from a given file path.
func (mm *ModuleManagerImpl) loadModule(modulePath string) (Module, error) {
	plug, err := plugin.Open(modulePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open module: %v", err)
	}

	sym, err := plug.Lookup("NewModule")
	if err != nil {
		return nil, fmt.Errorf("failed to find NewModule function: %v", err)
	}

	newModuleFunc, ok := sym.(func() Module)
	if !ok {
		return nil, errors.New("NewModule function has invalid type")
	}

	module := newModuleFunc()
	return module, nil
}
