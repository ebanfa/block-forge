package tests

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edward1christian/block-forge/pkg/application/appl"
	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
)

// TestNewModuleManager tests the creation of a new ModuleManagerImpl instance.
func TestNewModuleManager(t *testing.T) {
	moduleManager := appl.NewModuleManager("id", "name", "description")

	assert.NotNil(t, moduleManager)
	assert.Equal(t, "id", moduleManager.ID())
	assert.Equal(t, "name", moduleManager.Name())
	assert.Equal(t, "description", moduleManager.Description())
}

// TestModuleManagerImpl_Initialize tests initializing the module manager.
func TestModuleManagerImpl_Initialize(t *testing.T) {
	ctx := new(context.Context)
	app := new(mocks.MockApplication)

	moduleManager := appl.NewModuleManager("id", "name", "description")

	app.On("ID").Return("appID")

	moduleManager.Initialize(ctx, app)
	//err := moduleManager.Initialize(ctx, app)
	//assert.NoError(t, err)
	//app.AssertCalled(t, "ID")
}

// TestModuleManagerImpl_AddModule tests adding a module to the module manager.
func TestModuleManagerImpl_AddModule(t *testing.T) {
	module := new(mocks.MockModule)
	module.On("ID").Return("moduleID")

	moduleManager := appl.NewModuleManager("id", "name", "description")

	err := moduleManager.AddModule(module)

	assert.NoError(t, err)
	module.AssertCalled(t, "ID")
}

// TestModuleManagerImpl_AddModule_DuplicateID tests adding a module with a duplicate ID.
func TestModuleManagerImpl_AddModule_DuplicateID(t *testing.T) {
	module1 := new(mocks.MockModule)
	module1.On("ID").Return("moduleID")

	module2 := new(mocks.MockModule)
	module2.On("ID").Return("moduleID")

	moduleManager := appl.NewModuleManager("id", "name", "description")

	err := moduleManager.AddModule(module1)
	assert.NoError(t, err)

	err = moduleManager.AddModule(module2)
	assert.Error(t, err)
	assert.Equal(t, "module with ID 'moduleID' already exists", err.Error())
}

// TestModuleManagerImpl_RemoveModule tests removing a module from the module manager.
func TestModuleManagerImpl_RemoveModule(t *testing.T) {
	module := new(mocks.MockModule)
	module.On("ID").Return("moduleID")

	moduleManager := appl.NewModuleManager("id", "name", "description")
	moduleManager.AddModule(module)

	err := moduleManager.RemoveModule("moduleID")

	assert.NoError(t, err)
}

// TestModuleManagerImpl_RemoveModule_NotFound tests removing a module that does not exist.
func TestModuleManagerImpl_RemoveModule_NotFound(t *testing.T) {
	moduleManager := appl.NewModuleManager("id", "name", "description")

	err := moduleManager.RemoveModule("moduleID")

	assert.Error(t, err)
	assert.Equal(t, "module with ID 'moduleID' not found", err.Error())
}

// TestModuleManagerImpl_GetModule tests getting a module from the module manager.
func TestModuleManagerImpl_GetModule(t *testing.T) {
	module := new(mocks.MockModule)
	module.On("ID").Return("moduleID")

	moduleManager := appl.NewModuleManager("id", "name", "description")
	moduleManager.AddModule(module)

	resModule, err := moduleManager.GetModule("moduleID")

	assert.NoError(t, err)
	assert.Equal(t, module, resModule)
}

// TestModuleManagerImpl_GetModule_NotFound tests getting a module that does not exist.
func TestModuleManagerImpl_GetModule_NotFound(t *testing.T) {
	moduleManager := appl.NewModuleManager("id", "name", "description")

	_, err := moduleManager.GetModule("moduleID")

	assert.Error(t, err)
	assert.Equal(t, "module with ID 'moduleID' not found", err.Error())
}

// TestModuleManagerImpl_StartModules tests starting all modules managed by the module manager.
func TestModuleManagerImpl_StartModules(t *testing.T) {
	ctx := new(context.Context)
	module1 := new(mocks.MockModule)
	module1.On("ID").Return("moduleID1")
	module1.On("Start", ctx).Return(nil)

	moduleManager := appl.NewModuleManager("id", "name", "description")
	moduleManager.AddModule(module1)

	err := moduleManager.Start(ctx)

	assert.NoError(t, err)
	module1.AssertCalled(t, "Start", ctx)
}

// TestModuleManagerImpl_StartModules_AlreadyStartedError tests starting modules when module manager is already started.
func TestModuleManagerImpl_StartModules_AlreadyStartedError(t *testing.T) {
	ctx := new(context.Context)
	module := new(mocks.MockModule)
	module.On("ID").Return("moduleID")
	module.On("Start", ctx).Return(nil)

	moduleManager := appl.NewModuleManager("id", "name", "description")
	moduleManager.AddModule(module)
	moduleManager.Start(ctx)

	err := moduleManager.Start(ctx)

	assert.Error(t, err)
	assert.Equal(t, "module manager already started", err.Error())
}

// TestModuleManagerImpl_StartModules_ModuleStartError tests starting modules with a module failing to start.
func TestModuleManagerImpl_StartModules_ModuleStartError(t *testing.T) {
	ctx := new(context.Context)
	module1 := new(mocks.MockModule)

	module1.On("ID").Return("moduleID")
	module1.On("Name").Return("moduleID")
	module1.On("Start", ctx).Return(errors.New("module start error"))

	moduleManager := appl.NewModuleManager("id", "name", "description")
	moduleManager.AddModule(module1)

	/* err := moduleManager.Start(ctx)

	assert.Error(t, err) */
	moduleManager.Start(ctx)

	module1.AssertCalled(t, "Start", ctx)
}

// TestModuleManagerImpl_StopModules tests stopping all modules managed by the module manager.
func TestModuleManagerImpl_StopModules(t *testing.T) {
	ctx := new(context.Context)
	module1 := new(mocks.MockModule)
	module1.On("ID").Return("moduleID1")
	module1.On("Name").Return("moduleID1")
	module1.On("Start", ctx).Return(nil)

	moduleManager := appl.NewModuleManager("id", "name", "description")
	moduleManager.AddModule(module1)
	moduleManager.Start(ctx)

	module1.On("Stop", ctx).Return(nil)

	err := moduleManager.Stop(ctx)

	assert.NoError(t, err)
	module1.AssertCalled(t, "Stop", ctx)
}

// TestModuleManagerImpl_StopModules_NotStartedError tests stopping modules when module manager is not started.
func TestModuleManagerImpl_StopModules_NotStartedError(t *testing.T) {
	ctx := new(context.Context)
	module := new(mocks.MockModule)
	module.On("ID").Return("moduleID1")
	module.On("Name").Return("moduleID1")
	module.On("Start", ctx).Return(nil)

	moduleManager := appl.NewModuleManager("id", "name", "description")
	moduleManager.AddModule(module)

	err := moduleManager.Stop(ctx)

	assert.Error(t, err)
	assert.Equal(t, "module manager not started", err.Error())
}

// TestModuleManagerImpl_StopModules_ModuleStopError tests stopping modules with a module failing to stop.
func TestModuleManagerImpl_StopModules_ModuleStopError(t *testing.T) {
	ctx := new(context.Context)
	module1 := new(mocks.MockModule)
	module1.On("ID").Return("moduleID1")
	module1.On("Name").Return("moduleID1")
	module1.On("Start", ctx).Return(nil)

	moduleManager := appl.NewModuleManager("id", "name", "description")
	moduleManager.AddModule(module1)
	moduleManager.Start(ctx)

	module1.On("Stop", ctx).Return(errors.New("module stop error"))

	/* err := moduleManager.Stop(ctx)

	assert.Error(t, err)
	assert.Equal(t, "Failed to stop module name: module stop error", err.Error()) */
	moduleManager.Stop(ctx)
	module1.AssertCalled(t, "Stop", ctx)
}

// TestModuleManagerImpl_StartModule tests starting a single module managed by the module manager.
func TestModuleManagerImpl_StartModule(t *testing.T) {
	ctx := new(context.Context)
	module := new(mocks.MockModule)
	module.On("ID").Return("moduleID1")
	module.On("Name").Return("moduleID1")
	module.On("Start", ctx).Return(nil)

	moduleManager := appl.NewModuleManager("id", "name", "description")
	moduleManager.AddModule(module)

	err := moduleManager.StartModule(ctx, "moduleID1")

	assert.NoError(t, err)
	module.AssertCalled(t, "Start", ctx)
}

// TestModuleManagerImpl_StartModule_ModuleNotFound tests starting a module that does not exist.
func TestModuleManagerImpl_StartModule_ModuleNotFound(t *testing.T) {
	ctx := new(context.Context)

	moduleManager := appl.NewModuleManager("id", "name", "description")

	err := moduleManager.StartModule(ctx, "moduleID")

	assert.Error(t, err)
	assert.Equal(t, "module with name 'moduleID' not found", err.Error())
}

// TestModuleManagerImpl_StartModule_ModuleStartError tests starting a module with a module failing to start.
func TestModuleManagerImpl_StartModule_ModuleStartError(t *testing.T) {
	ctx := new(context.Context)
	module := new(mocks.MockModule)
	module.On("ID").Return("moduleID")
	module.On("Name").Return("moduleID")
	module.On("Start", ctx).Return(errors.New("module start error"))

	moduleManager := appl.NewModuleManager("id", "name", "description")
	moduleManager.AddModule(module)

	moduleManager.StartModule(ctx, "moduleID")

	module.AssertCalled(t, "Start", ctx)
}

// TestModuleManagerImpl_StopModule tests stopping a single module managed by the module manager.
func TestModuleManagerImpl_StopModule(t *testing.T) {
	ctx := new(context.Context)
	module := new(mocks.MockModule)
	module.On("ID").Return("moduleID")
	module.On("Name").Return("moduleID")
	module.On("Start", ctx).Return(nil)
	module.On("Stop", ctx).Return(nil)

	moduleManager := appl.NewModuleManager("id", "name", "description")
	moduleManager.AddModule(module)
	moduleManager.Start(ctx)

	err := moduleManager.StopModule(ctx, "moduleID")

	assert.NoError(t, err)
	module.AssertCalled(t, "Stop", ctx)
}

// TestModuleManagerImpl_StopModule_ModuleNotFound tests stopping a module that does not exist.
func TestModuleManagerImpl_StopModule_ModuleNotFound(t *testing.T) {
	ctx := new(context.Context)

	moduleManager := appl.NewModuleManager("id", "name", "description")

	err := moduleManager.StopModule(ctx, "moduleID")

	assert.Error(t, err)
	assert.Equal(t, "module with name 'moduleID' not found", err.Error())
}

// TestModuleManagerImpl_StopModule_ModuleStopError tests stopping a module with a module failing to stop.
func TestModuleManagerImpl_StopModule_ModuleStopError(t *testing.T) {
	ctx := new(context.Context)
	module := new(mocks.MockModule)
	module.On("ID").Return("moduleID")
	module.On("Name").Return("moduleID")
	module.On("Start", ctx).Return(nil)

	moduleManager := appl.NewModuleManager("id", "name", "description")
	moduleManager.AddModule(module)
	moduleManager.Start(ctx)

	module.On("Stop", ctx).Return(errors.New("module stop error"))

	err := moduleManager.StopModule(ctx, "moduleID")

	assert.Error(t, err)
	assert.Equal(t, "failed to stop module moduleID: module stop error", err.Error())
	module.AssertCalled(t, "Stop", ctx)
}

// TestDiscoverAndLoadModules_Valid tests discovering and loading valid modules.
func TestDiscoverAndLoadModules_Valid(t *testing.T) {
	/* // Create a temporary directory to simulate modules directory
	tmpDir, err := ioutil.TempDir("", "test-modules")
	assert.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Create a mock module zip file
	mockZipFilePath := filepath.Join(tmpDir, "mock_module.zip")
	err = ioutil.WriteFile(mockZipFilePath, []byte("mock module content"), 0644)
	assert.NoError(t, err)

	// Create a module manager
	moduleManager := appl.NewModuleManager("id", "name", "description")

	// Mock context
	ctx := new(context.Context)

	// Discover and load modules
	err = moduleManager.DiscoverAndLoadModules(ctx)
	assert.NoError(t, err)

	// Check if the module is loaded
	module, err := moduleManager.GetModule("mock_module")
	assert.NoError(t, err)
	assert.NotNil(t, module) */
}

// TestDiscoverAndLoadModules_Invalid tests discovering invalid modules.
func TestDiscoverAndLoadModules_Invalid(t *testing.T) {
	/* // Create a temporary directory to simulate modules directory
	tmpDir, err := ioutil.TempDir("", "test-modules")
	assert.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	// Create an invalid module file (not a zip file)
	invalidModuleFilePath := filepath.Join(tmpDir, "invalid_module.txt")
	err = ioutil.WriteFile(invalidModuleFilePath, []byte("invalid module content"), 0644)
	assert.NoError(t, err)

	// Create a module manager
	moduleManager := appl.NewModuleManager("id", "name", "description")

	// Mock context
	ctx := new(context.Context)

	// Discover and load modules
	err = moduleManager.DiscoverAndLoadModules(ctx)
	assert.NoError(t, err)

	// Check if the invalid module is not loaded
	_, err = moduleManager.GetModule("invalid_module")
	assert.Error(t, err) */
}

// TestLoadRemoteModule_Valid tests loading a valid module from a remote source.
func TestLoadRemoteModule_Valid(t *testing.T) {
	/* // Provide a valid remote module URL (you can mock a remote server if needed)
	remoteModuleURL := "https://example.com/module.zip"

	// Create a module manager
	moduleManager := appl.NewModuleManager("id", "name", "description")

	// Mock context
	ctx := new(context.Context)

	// Load the remote module
	module, err := moduleManager.LoadRemoteModule(ctx, remoteModuleURL)
	assert.NoError(t, err)
	assert.NotNil(t, module) */
}

// TestLoadRemoteModule_InvalidURL tests loading a module from an invalid remote URL.
func TestLoadRemoteModule_InvalidURL(t *testing.T) {
	/* // Provide an invalid remote module URL
	invalidRemoteModuleURL := "https://invalid-url.com/module.zip"

	// Create a module manager
	moduleManager := appl.NewModuleManager("id", "name", "description")

	// Mock context
	ctx := new(context.Context)

	// Load the remote module
	_, err := moduleManager.LoadRemoteModule(ctx, invalidRemoteModuleURL)
	assert.Error(t, err) */
}
