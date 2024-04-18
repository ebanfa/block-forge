package internal

import (
	"fmt"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
)

// NovaPlugin represents a plugin in the system.
type NovaPlugin struct {
	systemApi.BaseSystemComponent
	system systemApi.SystemInterface
}

// NewNovaPlugin creates a new instance of NovaPlugin.
func NewNovaPlugin() systemApi.PluginInterface {
	return &NovaPlugin{}
}

// Initialize initializes the module.
// Returns an error if the initialization fails.
func (p *NovaPlugin) Initialize(ctx *context.Context, system systemApi.SystemInterface) error {
	// Initialization logic
	p.system = system
	fmt.Println("NovaPlugin initialized")
	return nil
}

// RegisterResources registers resources into the system.
// Returns an error if resource registration fails.
func (p *NovaPlugin) RegisterResources(ctx *context.Context) error {
	// Register resources using p.system
	fmt.Println("Resources registered for NovaPlugin")
	return nil
}

// Start starts the component.
// Returns an error if the start operation fails.
func (p *NovaPlugin) Start(ctx *context.Context) error {
	// Start logic
	fmt.Println("NovaPlugin started")
	return nil
}

// Stop stops the component.
// Returns an error if the stop operation fails.
func (p *NovaPlugin) Stop(ctx *context.Context) error {
	// Stop logic
	fmt.Println("NovaPlugin stopped")
	return nil
}
