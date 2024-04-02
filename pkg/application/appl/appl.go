package appl

import (
	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// Application represents the main application.
type Application interface {
	system.BootableComponent

	// System returns the system instance
	System() system.System

	// ModuleManager returns the module manager instance
	ModuleManager() ModuleManager
}

// Component represents a generic component in the system.
type ApplicationComponent interface {
	system.StartableComponent

	// Initialize starts the module with the given context and application instance
	Initialize(ctx *context.Context, app Application) error
}

// Module represents a module component in the system.
type Module interface {
	ApplicationComponent
}

// ModuleManager defines the interface for managing modules.
type ModuleManager interface {
	ApplicationComponent

	// AddModule adds a module to the module manager
	AddModule(module Module) error

	// RemoveModule removes a module from the module manager
	RemoveModule(name string) error

	// GetModule returns the module with the given name
	GetModule(name string) (Module, error)

	// StartModules starts all modules managed by the module manager
	StartModule(ctx *context.Context, name string) error

	// StopModules stops all modules managed by the module manager
	StopModule(ctx *context.Context, name string) error

	// DiscoverModules discovers available modules within the system
	DiscoverAndLoadModules(ctx *context.Context) error

	// LoadRemoteModule loads a module from a remote source
	LoadRemoteModule(ctx *context.Context, moduleURL string) (Module, error)
}
