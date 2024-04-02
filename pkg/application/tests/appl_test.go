package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/edward1christian/block-forge/pkg/application/appl"
	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
)

// TestNewApplication tests the creation of a new ApplicationImpl instance.
func TestNewApplication(t *testing.T) {
	moduleManager := &mocks.MockModuleManager{}
	sys := new(mocks.MockSystem)

	app := appl.NewApplication("appID", "appName", "appDescription", moduleManager, sys)

	assert.NotNil(t, app)
	assert.Equal(t, "appID", app.ID())
	assert.Equal(t, "appName", app.Name())
	assert.Equal(t, "appDescription", app.Description())
	assert.Equal(t, moduleManager, app.ModuleManager())
	assert.Equal(t, sys, app.System())
}

// TestApplication_Initialize tests initializing the application.
func TestApplication_Initialize(t *testing.T) {
	ctx := new(context.Context)
	moduleManager := &mocks.MockModuleManager{}
	sys := new(mocks.MockSystem)

	moduleManager.On("Initialize", ctx, mock.AnythingOfType("*appl.ApplicationImpl")).Return(nil)

	app := appl.NewApplication("appID", "appName", "appDescription", moduleManager, sys)

	err := app.Initialize(ctx)

	assert.NoError(t, err)
	moduleManager.AssertCalled(t, "Initialize", ctx, mock.AnythingOfType("*appl.ApplicationImpl"))
}

// TestApplication_Initialize_Error tests initializing the application with an error.
func TestApplication_Initialize_Error(t *testing.T) {
	ctx := new(context.Context)
	moduleManager := &mocks.MockModuleManager{}
	sys := new(mocks.MockSystem)

	expectedErr := fmt.Errorf("module manager initialization error")
	moduleManager.On("Initialize", ctx, mock.AnythingOfType("*appl.ApplicationImpl")).Return(expectedErr)

	app := appl.NewApplication("appID", "appName", "appDescription", moduleManager, sys)

	err := app.Initialize(ctx)

	assert.Error(t, err)

	assert.Equal(t, "failed to initialize module manager: module manager initialization error", err.Error())
	moduleManager.AssertCalled(t, "Initialize", ctx, mock.AnythingOfType("*appl.ApplicationImpl"))
}

// TestApplication_Start tests starting the application.
func TestApplication_Start(t *testing.T) {
	ctx := new(context.Context)
	moduleManager := &mocks.MockModuleManager{}
	sys := new(mocks.MockSystem)

	moduleManager.On("Start", ctx).Return(nil)

	app := appl.NewApplication("appID", "appName", "appDescription", moduleManager, sys)

	err := app.Start(ctx)

	assert.NoError(t, err)
	moduleManager.AssertCalled(t, "Start", ctx)
}

// TestApplication_Start_AlreadyStartedError tests starting the application when it's already started.
func TestApplication_Start_AlreadyStartedError(t *testing.T) {
	ctx := new(context.Context)
	sys := new(mocks.MockSystem)

	moduleManager := &mocks.MockModuleManager{}
	moduleManager.On("Start", ctx).Return(nil)

	app := appl.NewApplication("appID", "appName", "appDescription", moduleManager, sys)
	app.Start(ctx)

	err := app.Start(ctx)

	assert.Error(t, err)
	assert.Equal(t, "application already started", err.Error())
}

// TestApplication_Start_Error tests starting the application with an error.
func TestApplication_Start_Error(t *testing.T) {
	ctx := new(context.Context)
	moduleManager := &mocks.MockModuleManager{}
	sys := new(mocks.MockSystem)

	expectedErr := fmt.Errorf("module manager start error")
	moduleManager.On("Start", ctx).Return(expectedErr)

	app := appl.NewApplication("appID", "appName", "appDescription", moduleManager, sys)

	err := app.Start(ctx)

	assert.Error(t, err)
	assert.Equal(t, "failed to start modules: module manager start error", err.Error())
	moduleManager.AssertCalled(t, "Start", ctx)
}

// TestApplication_Stop tests stopping the application.
func TestApplication_Stop(t *testing.T) {
	ctx := new(context.Context)
	moduleManager := &mocks.MockModuleManager{}
	sys := new(mocks.MockSystem)

	moduleManager.On("Start", ctx).Return(nil)
	moduleManager.On("Stop", ctx).Return(nil)

	app := appl.NewApplication("appID", "appName", "appDescription", moduleManager, sys)
	app.Start(ctx)

	err := app.Stop(ctx)

	assert.NoError(t, err)
	moduleManager.AssertCalled(t, "Stop", ctx)
}

// TestApplication_Stop_NotStartedError tests stopping the application when it's not started.
func TestApplication_Stop_NotStartedError(t *testing.T) {
	ctx := new(context.Context)
	sys := new(mocks.MockSystem)

	moduleManager := &mocks.MockModuleManager{}
	moduleManager.On("Start", ctx).Return(nil)
	moduleManager.On("Stop", ctx).Return(nil)

	app := appl.NewApplication("appID", "appName", "appDescription", moduleManager, sys)

	err := app.Stop(ctx)

	assert.Error(t, err)
	assert.Equal(t, "application not started", err.Error())
}

// TestApplication_Stop_Error tests stopping the application with an error.
func TestApplication_Stop_Error(t *testing.T) {
	ctx := new(context.Context)
	moduleManager := &mocks.MockModuleManager{}
	sys := new(mocks.MockSystem)

	expectedErr := fmt.Errorf("module manager stop error")
	moduleManager.On("Start", ctx).Return(nil)
	moduleManager.On("Stop", ctx).Return(expectedErr)

	app := appl.NewApplication("appID", "appName", "appDescription", moduleManager, sys)
	app.Start(ctx)

	err := app.Stop(ctx)

	assert.Error(t, err)
	assert.Equal(t, "failed to stop modules: module manager stop error", err.Error())
	moduleManager.AssertCalled(t, "Stop", ctx)
}
