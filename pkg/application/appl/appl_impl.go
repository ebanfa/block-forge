package appl

import (
	"fmt"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/components"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// ApplicationImpl represents the main application.
type ApplicationImpl struct {
	components.BootableInterface
	components.StartableInterface
	moduleManager ModuleManager
	system        system.SystemInterface
	started       bool
}

// NewApplication creates a new instance of the ApplicationImpl.
func NewApplication(moduleManager ModuleManager, system system.SystemInterface) *ApplicationImpl {
	return &ApplicationImpl{
		moduleManager: moduleManager,
		system:        system,
	}
}

// ModuleManager returns the module manager of the application.
func (app *ApplicationImpl) ModuleManager() ModuleManager {
	return app.moduleManager
}

// System returns the system of the application.
func (app *ApplicationImpl) System() system.SystemInterface {
	return app.system
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
	if err := app.moduleManager.Start(ctx); err != nil {
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
	if err := app.moduleManager.Stop(ctx); err != nil {
		return fmt.Errorf("failed to stop modules: %w", err)
	}

	// Mark the application as stopped
	app.started = false
	return nil
}
