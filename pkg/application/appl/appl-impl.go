package appl

import (
	"fmt"

	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// ApplicationImpl represents the main application.
type ApplicationImpl struct {
	id            string
	name          string
	description   string
	moduleManager ModuleManager
	system        system.System
	started       bool
}

// NewApplication creates a new instance of the ApplicationImpl.
func NewApplication(id, name, description string, moduleManager ModuleManager, system system.System) *ApplicationImpl {
	return &ApplicationImpl{
		id:            id,
		name:          name,
		description:   description,
		moduleManager: moduleManager,
		system:        system,
	}
}

// ID returns the unique identifier of the application.
func (app *ApplicationImpl) ID() string {
	return app.id
}

// Name returns the name of the application.
func (app *ApplicationImpl) Name() string {
	return app.name
}

// Description returns the description of the application.
func (app *ApplicationImpl) Description() string {
	return app.description
}

// ModuleManager returns the module manager of the application.
func (app *ApplicationImpl) ModuleManager() ModuleManager {
	return app.moduleManager
}

// System returns the system of the application.
func (app *ApplicationImpl) System() system.System {
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
