package appl

import (
	"fmt"

	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// application represents the main application.
type ApplicationImpl struct {
	moduleManager ModuleManager
	system        system.System
	started       bool
}

// ModuleManager implements Application.
func (app *ApplicationImpl) ModuleManager() ModuleManager {
	return app.moduleManager
}

// System implements Application.
func (app *ApplicationImpl) System() system.System {
	return app.system
}

// NewApplication creates a new instance of the ApplicationImpl.
func NewApplication(moduleManager ModuleManager, system system.System) *ApplicationImpl {
	return &ApplicationImpl{
		moduleManager: moduleManager,
		system:        system,
	}
}

// Initialize initializes the application.
func (app *ApplicationImpl) Initialize(ctx *context.Context) error {
	// Initialize the system and module manager
	if err := app.moduleManager.Initialize(ctx, app); err != nil {
		return fmt.Errorf("failed to initialize module manager: %w", err)
	}
	return nil
}

// Start starts the application.
func (app *ApplicationImpl) Start(ctx *context.Context) error {
	if app.started {
		return fmt.Errorf("application already started")
	}

	// Start the module manager
	if err := app.moduleManager.StartModules(ctx); err != nil {
		return fmt.Errorf("failed to start modules: %w", err)
	}

	// Mark the application as started
	app.started = true
	return nil
}

// Stop stops the application.
func (app *ApplicationImpl) Stop(ctx *context.Context) error {
	if !app.started {
		return fmt.Errorf("application not started")
	}

	// Stop the module manager
	if err := app.moduleManager.StopModules(ctx); err != nil {
		return fmt.Errorf("failed to stop modules: %w", err)
	}

	// Mark the application as stopped
	app.started = false
	return nil
}
