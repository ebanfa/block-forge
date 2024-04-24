package plugin

import (
	"errors"
	"fmt"

	"github.com/edward1christian/block-forge/nova/pkg/common"
	"github.com/edward1christian/block-forge/nova/pkg/components/services"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/common/logger"
	"github.com/edward1christian/block-forge/pkg/application/component"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
)

// NovaPlugin represents a plugin in the system.
type NovaPlugin struct {
	systemApi.BaseSystemComponent
	system       systemApi.SystemInterface
	buildService systemApi.SystemServiceInterface
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
	// Initialize the
	system.Logger().Log(logger.LevelInfo, "NovaPlugin: Initializing plugin")
	return nil
}

// RegisterResources registers resources into the system.
// Returns an error if resource registration fails.
func (p *NovaPlugin) RegisterResources(ctx *context.Context) error {
	p.system.Logger().Log(logger.LevelInfo, "NovaPlugin: Registering resources")
	// Register and create the build service component
	if err := p.registerBuildService(ctx); err != nil {
		return fmt.Errorf("failed to register build service: %w", err)
	}

	return nil
}

// Helper function to register and create the build service component
func (p *NovaPlugin) registerBuildService(ctx *context.Context) error {
	p.system.Logger().Log(logger.LevelInfo, "NovaPlugin: registering factory "+common.IgniteBuildServiceFactory)

	registrar := p.system.ComponentRegistry()
	err := registrar.RegisterFactory(ctx, &component.FactoryRegistrationInfo{
		ID:      common.IgniteBuildServiceFactory,
		Factory: &services.BuildServiceFactory{},
	})
	if err != nil {
		return fmt.Errorf("failed to register factory %s: %w", common.IgniteBuildServiceFactory, err)
	}
	return nil
}

// Start starts the component.
// Returns an error if the start operation fails.
func (p *NovaPlugin) Start(ctx *context.Context) error {
	p.system.Logger().Log(logger.LevelInfo, "NovaPlugin: Starting plugin")

	// Start the build service
	if err := StartBuildService(ctx, p.system); err != nil {
		return fmt.Errorf("failed to register BuildService: %v", err)
	}

	return nil
}

// Stop stops the component.
// Returns an error if the stop operation fails.
func (p *NovaPlugin) Stop(ctx *context.Context) error {
	// Retrieve the BuildService component from the ComponentRegistry
	component, err := p.system.ComponentRegistry().GetComponent(common.IgniteBuildService)
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
