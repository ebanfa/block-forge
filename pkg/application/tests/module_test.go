package tests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/edward1christian/block-forge/pkg/application/appl"
	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
	"github.com/stretchr/testify/assert"
)

func TestModuleManager_Initialize(t *testing.T) {
	// Create a new ModuleManager
	mm := appl.NewModuleManager("MMID", "ModuleManager", "Module Manager")
	ctx := &context.Context{}

	// Create a mock applications
	mockApplication := &mocks.MockApplication{}

	// Add the mock module to the ModuleManager
	err := mm.Initialize(ctx, mockApplication)

	// Verify that no error occurred
	assert.NoError(t, err)
}

func TestModuleManager_AddModule(t *testing.T) {
	// Create a new ModuleManager
	mm := appl.NewModuleManager("MMID", "ModuleManager", "Module Manager")

	// Create a mock module
	mockModule := &mocks.MockModule{}

	// Expect ID() method call on mockModule
	mockModule.On("ID").Return("mockID")

	// Expect Name() method call on mockModule
	mockModule.On("Name").Return("mockName")

	// Add the mock module to the ModuleManager
	err := mm.AddModule(mockModule)

	// Verify that no error occurred
	assert.NoError(t, err)
}

func TestModuleManager_AddModule_AlreadyExists(t *testing.T) {
	mm := appl.NewModuleManager("MMID", "ModuleManager", "Module Manager")
	mockModule := &mocks.MockModule{}

	// Expect ID() method call on mockModule
	mockModule.On("ID").Return("mockID")

	// Expect Name() method call on mockModule
	mockModule.On("Name").Return("Mock")

	err := mm.AddModule(mockModule)
	assert.NoError(t, err)

	// Attempt to add the same module again
	err = mm.AddModule(mockModule)

	assert.Error(t, err)
	assert.EqualError(t, err, fmt.Sprintf("module with ID '%s' already exists", mockModule.ID()))
}

func TestModuleManager_RemoveModule(t *testing.T) {
	mm := appl.NewModuleManager("MMID", "ModuleManager", "Module Manager")
	mockModule := &mocks.MockModule{}

	// Expect ID() method call on mockModule
	mockModule.On("ID").Return("mockID")

	// Expect Name() method call on mockModule
	mockModule.On("Name").Return("Mock")

	err := mm.AddModule(mockModule)
	assert.NoError(t, err)

	err = mm.RemoveModule(mockModule.Name())

	assert.NoError(t, err)
}

func TestModuleManager_RemoveModule_NotFound(t *testing.T) {
	mm := appl.NewModuleManager("MMID", "ModuleManager", "Module Manager")
	mockModule := &mocks.MockModule{}

	// Expect ID() method call on mockModule
	mockModule.On("ID").Return("mockID")

	// Expect Name() method call on mockModule
	mockModule.On("Name").Return("Mock")

	err := mm.RemoveModule(mockModule.Name())

	assert.Error(t, err)
	assert.EqualError(t, err, fmt.Sprintf("module with name '%s' not found", mockModule.Name()))
}

func TestModuleManager_GetModule(t *testing.T) {
	mm := appl.NewModuleManager("MMID", "ModuleManager", "Module Manager")
	mockModule := &mocks.MockModule{}

	// Expect ID() method call on mockModule
	mockModule.On("ID").Return("mockID")

	// Expect Name() method call on mockModule
	mockModule.On("Name").Return("Mock")

	err := mm.AddModule(mockModule)
	assert.NoError(t, err)

	module, err := mm.GetModule(mockModule.Name())

	assert.NoError(t, err)
	assert.NotNil(t, module)
	assert.Equal(t, mockModule, module)
}

func TestModuleManager_StartModules(t *testing.T) {
	mm := appl.NewModuleManager("MMID", "ModuleManager", "Module Manager")
	mockModule := &mocks.MockModule{}

	// Expect ID() method call on mockModule
	mockModule.On("ID").Return("mockID")

	// Expect Name() method call on mockModule
	mockModule.On("Name").Return("Mock")

	err := mm.AddModule(mockModule)
	assert.NoError(t, err)

	// Prepare the context
	ctx := &context.Context{}

	// Mock the Start method of the module
	mockModule.On("Start", ctx).Return(nil)

	// Test StartModules
	err = mm.StartModules(ctx)

	assert.NoError(t, err)
}

func TestModuleManager_StopModules(t *testing.T) {
	mm := appl.NewModuleManager("MMID", "ModuleManager", "Module Manager")
	mockModule := &mocks.MockModule{}

	// Expect ID() method call on mockModule
	mockModule.On("ID").Return("mockID")

	// Expect Name() method call on mockModule
	mockModule.On("Name").Return("Mock")

	err := mm.AddModule(mockModule)
	assert.NoError(t, err)

	// Prepare the context
	ctx := &context.Context{}

	// Mock the Start method of the module
	mockModule.On("Start", ctx).Return(nil)

	// Mock the Stop method of the module
	mockModule.On("Stop", ctx).Return(nil)

	// Test StartModules
	mm.StartModules(ctx)

	// Test StopModules
	err = mm.StopModules(ctx)

	assert.NoError(t, err)
}

func TestModuleManager_StopModules_Failure(t *testing.T) {
	mm := appl.NewModuleManager("MMID", "ModuleManager", "Module Manager")
	mockModule := &mocks.MockModule{}

	// Expect ID() method call on mockModule
	mockModule.On("ID").Return("mockID")

	// Expect Name() method call on mockModule
	mockModule.On("Name").Return("Mock")

	err := mm.AddModule(mockModule)
	assert.NoError(t, err)

	// Prepare the context
	ctx := &context.Context{}

	// Mock the Stop method of the module
	mockModule.On("Stop", ctx).Return(nil)

	// Test StopModules
	err = mm.StopModules(ctx)

	assert.Error(t, err)
	assert.EqualError(t, err, "module manager not started")
}

func TestModuleManager_DiscoverModules(t *testing.T) {
	t.Skip("Pending, unit implementation incomplete")
	// Create a temporary directory for test modules
	tempDir, err := ioutil.TempDir("", "test-modules")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create a test Go source file
	testModulePath := filepath.Join(tempDir, "test_module.go")
	testModuleFile, err := os.Create(testModulePath)
	assert.NoError(t, err)
	defer testModuleFile.Close()

	// Write a dummy Go module code to the test file
	_, err = testModuleFile.WriteString(`
        package main

        import "fmt"

        type TestModule struct{}

        func (m *TestModule) ID() string {
            return "test_module"
        }

        func (m *TestModule) Name() string {
            return "Test Module"
        }

        func (m *TestModule) Start(ctx *context.Context) error {
            fmt.Println("Test Module started")
            return nil
        }

        func (m *TestModule) Stop(ctx *context.Context) error {
            fmt.Println("Test Module stopped")
            return nil
        }
    `)
	assert.NoError(t, err)

	// Create a new ModuleManager
	mm := appl.NewModuleManager("MMID", "ModuleManager", "Module Manager")

	// Prepare the context
	ctx := &context.Context{}

	// Discover modules
	discoveredModules, err := mm.DiscoverModules(ctx)

	// Ensure no errors
	assert.NoError(t, err)
	// Ensure at least one module is discovered
	assert.NotEmpty(t, discoveredModules)
}

func TestModuleManager_DiscoverModules_Failure(t *testing.T) {
	t.Skip("Pending, unit implementation incomplete")
	// Create a temporary directory without any Go source files
	tempDir, err := ioutil.TempDir("", "empty-directory")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Create a new ModuleManager
	mm := appl.NewModuleManager("MMID", "ModuleManager", "Module Manager")

	// Prepare the context
	ctx := &context.Context{}

	// Discover modules from an empty directory
	discoveredModules, err := mm.DiscoverModules(ctx)

	// Ensure an error occurred
	assert.Error(t, err)
	// Ensure no modules are discovered
	assert.Empty(t, discoveredModules)
}

func TestModuleManager_LoadRemoteModule(t *testing.T) {
	t.Skip("Pending, unit implementation incomplete")
	// Create a temporary HTTP server to serve a dummy module
	http.HandleFunc("/test_module.go", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`
            package main

            import "fmt"

            type TestModule struct{}

            func (m *TestModule) ID() string {
                return "test_module"
            }

            func (m *TestModule) Name() string {
                return "Test Module"
            }

            func (m *TestModule) Start(ctx *context.Context) error {
                fmt.Println("Test Module started")
                return nil
            }

            func (m *TestModule) Stop(ctx *context.Context) error {
                fmt.Println("Test Module stopped")
                return nil
            }
        `))
	})

	server := &http.Server{Addr: ":8080"}
	defer server.Close()

	go func() {
		_ = server.ListenAndServe()
	}()

	// Create a new ModuleManager
	mm := appl.NewModuleManager("MMID", "ModuleManager", "Module Manager")

	// Prepare the context
	ctx := &context.Context{}

	// Load the remote module
	remoteModuleURL := "http://localhost:8080/test_module.go"
	module, err := mm.LoadRemoteModule(ctx, remoteModuleURL)

	// Ensure no errors
	assert.NoError(t, err)
	// Ensure module is not nil
	assert.NotNil(t, module)
	// Ensure module has correct ID
	//assert.Equal(t, "test_module", appl.ID())
	// Ensure module has correct Name
	//assert.Equal(t, "Test Module", appl.Name())
}
