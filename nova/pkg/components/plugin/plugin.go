package plugin

import (
	"errors"
	"fmt"

	"github.com/edward1christian/block-forge/nova/pkg/common"
	"github.com/edward1christian/block-forge/nova/pkg/components/factories"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/config"
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
	// Initialize the
	return nil
}

// RegisterResources registers resources into the system.
// Returns an error if resource registration fails.
func (p *NovaPlugin) RegisterResources(ctx *context.Context) error {
	registrar := p.system.ComponentRegistry()

	// Register the service factory
	err := registrar.RegisterFactory(
		common.IgniteBuildServiceFactory, &factories.BuilderServiceFactory{})
	if err != nil {
		return fmt.Errorf("failed to register service factory: %w", err)
	}

	// Create and register the builder service
	_, err = registrar.CreateComponent(&config.ComponentConfig{
		ID:        common.IgniteBuildService,
		FactoryID: common.IgniteBuildServiceFactory,
	})

	if err != nil {
		return fmt.Errorf("failed to create and register builder service: %w", err)
	}

	return nil
}

// Start starts the component.
// Returns an error if the start operation fails.
func (p *NovaPlugin) Start(ctx *context.Context) error {
	// Retrieve the BuilderService component from the ComponentRegistry
	component, err := p.system.ComponentRegistry().GetComponent(common.IgniteBuildService)
	if err != nil {
		return fmt.Errorf("failed to get BuilderService component: %v", err)
	}

	// Check if the retrieved component implements the SystemServiceInterface
	builderService, ok := component.(systemApi.SystemServiceInterface)
	if !ok {
		return errors.New("BuilderService component does not implement SystemServiceInterface")
	}

	// Start the BuilderService
	if err := builderService.Start(ctx); err != nil {
		return fmt.Errorf("failed to start BuilderService: %v", err)
	}

	return nil
}

// Stop stops the component.
// Returns an error if the stop operation fails.
func (p *NovaPlugin) Stop(ctx *context.Context) error {
	// Retrieve the BuilderService component from the ComponentRegistry
	component, err := p.system.ComponentRegistry().GetComponent(common.IgniteBuildService)
	if err != nil {
		return fmt.Errorf("failed to get BuilderService component: %v", err)
	}

	// Check if the retrieved component implements the SystemServiceInterface
	builderService, ok := component.(systemApi.SystemServiceInterface)
	if !ok {
		return errors.New("BuilderService component does not implement SystemServiceInterface")
	}

	// Stop the BuilderService
	if err := builderService.Stop(ctx); err != nil {
		return fmt.Errorf("failed to stop BuilderService: %v", err)
	}

	fmt.Println("NovaPlugin stopped")
	return nil
}
