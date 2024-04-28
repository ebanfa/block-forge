package plugin

import (
	"errors"
	"fmt"

	"github.com/edward1christian/block-forge/nova/pkg/common"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/common/logger"
	"github.com/edward1christian/block-forge/pkg/application/component"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
)

// NovaPlugin represents a plugin in the system.
type NovaPlugin struct {
	systemApi.BaseSystemComponent
}

// NewNovaPlugin creates a new instance of NovaPlugin.
func NewNovaPlugin() systemApi.PluginInterface {
	return &NovaPlugin{
		BaseSystemComponent: systemApi.BaseSystemComponent{
			BaseComponent: component.BaseComponent{
				Id: "NovaPlugin",
			},
		},
	}
}

// Initialize initializes the module.
// Returns an error if the initialization fails.
func (p *NovaPlugin) Initialize(ctx *context.Context, system systemApi.SystemInterface) error {
	// Initialization logic
	p.System = system
	// Initialize the
	system.Logger().Log(logger.LevelInfo, "NovaPlugin: Initializing plugin")
	return nil
}

// RegisterResources registers resources into the system.
// Returns an error if resource registration fails.
func (p *NovaPlugin) RegisterResources(ctx *context.Context) error {
	p.System.Logger().Log(logger.LevelInfo, "NovaPlugin: Registering resources")

	// Register components
	if err := RegisterComponents(ctx, p.System); err != nil {
		return fmt.Errorf("failed to register plugin components %w", err)
	}
	// Register system services
	err := RegisterServices(ctx, p.System)
	if err != nil {
		return fmt.Errorf("failed to register services %w", err)
	}
	p.System.Logger().Log(logger.LevelInfo, "NovaPlugin: Registered resources")
	return nil
}

// Start starts the component.
// Returns an error if the start operation fails.
func (p *NovaPlugin) Start(ctx *context.Context) error {
	p.System.Logger().Log(logger.LevelInfo, "NovaPlugin: Starting plugin")

	// Start the build service
	if err := StartServices(ctx, p.System); err != nil {
		return fmt.Errorf("failed to start BuildService: %v", err)
	}
	p.System.Logger().Log(logger.LevelInfo, "NovaPlugin: Started plugin")

	return nil
}

// Stop stops the component.
// Returns an error if the stop operation fails.
func (p *NovaPlugin) Stop(ctx *context.Context) error {
	// Retrieve the BuildService component from the ComponentRegistry
	component, err := p.System.ComponentRegistry().GetComponent(common.BuildService)
	if err != nil {
		return fmt.Errorf("failed to get BuildService component: %v", err)
	}

	// Check if the retrieved component implements the SystemServiceInterface
	buildService, ok := component.(systemApi.SystemServiceInterface)
	if !ok {
		return errors.New("BuildService component does not implement SystemServiceInterface")
	}

	// Stop the BuildService
	if err := buildService.Stop(ctx); err != nil {
		return fmt.Errorf("failed to stop BuildService: %v", err)
	}

	fmt.Println("NovaPlugin stopped")
	return nil
}
