package appl

import (
	"github.com/edward1christian/block-forge/pkg/application/component"
	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// Application represents the main application.
type Application interface {
	component.Startable

	// System returns the system instance
	System() system.System

	// ModuleManager returns the module manager instance
	ModuleManager() ModuleManager

	// Initialize initializes the application.
	Initialize(ctx *context.Context) error
}

// Module represents a module component in the system.
type Module interface {
	component.StartableComponent

	// Initialize starts the module with the given context and application instance
	Initialize(ctx *context.Context, app Application) error
}

// ModuleManager defines the interface for managing modules.
type ModuleManager interface {
	// Initialize starts the module with the given context and application instance
	Initialize(ctx *context.Context, app Application) error

	// AddModule adds a module to the module manager
	AddModule(module Module) error

	// RemoveModule removes a module from the module manager
	RemoveModule(name string) error

	// GetModule returns the module with the given name
	GetModule(name string) (Module, error)

	// StartModules starts all modules managed by the module manager
	StartModules(ctx *context.Context) error

	// StopModules stops all modules managed by the module manager
	StopModules(ctx *context.Context) error

	// DiscoverModules discovers available modules within the system
	DiscoverModules(ctx *context.Context) ([]Module, error)

	// LoadRemoteModule loads a module from a remote source
	LoadRemoteModule(ctx *context.Context, moduleURL string) (Module, error)
}
