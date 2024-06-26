package system

import (
	"errors"
	"fmt"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/common/logger"
	compApi "github.com/edward1christian/block-forge/pkg/application/component"
	"github.com/edward1christian/block-forge/pkg/application/config"
)

func StartService(
	ctx *context.Context,
	system SystemInterface,
	config *config.ComponentConfig) error {

	registrar := system.ComponentRegistry()

	component, err := registrar.CreateComponent(ctx, config)
	if err != nil {
		return fmt.Errorf("failed to start service. Could not create component %s", config.ID)
	}

	service, ok := component.(SystemServiceInterface)
	if !ok {
		return fmt.Errorf("failed to start service. Component %s is not a system service", component.ID())
	}

	// Initialize the service
	if err := service.Initialize(ctx, system); err != nil {
		return fmt.Errorf("failed to initialize service: %s %v", component.ID(), err)
	}

	return service.Start(ctx)
}

func StopService(ctx *context.Context, system SystemInterface, id string) error {
	// Retrieve the BuildService component from the ComponentRegistry
	component, err := system.ComponentRegistry().GetComponent(id)
	if err != nil {
		return fmt.Errorf("failed to stop build service. Service not found: %v", err)
	}

	// Check if the retrieved component implements the SystemServiceInterface
	service, ok := component.(SystemServiceInterface)
	if !ok {
		return errors.New("failed to stop service. Service component is not a system service")
	}

	return service.Stop(ctx)
}

func RegisterComponent(ctx *context.Context, system SystemInterface, config *config.ComponentConfig, factory compApi.ComponentFactoryInterface) error {
	registrar := system.ComponentRegistry()
	system.Logger().Logf(logger.LevelDebug, "Registering component %s with factory ID %s", config.ID, config.FactoryID)
	// Register the factory
	err := registrar.RegisterFactory(ctx, config.FactoryID, factory)
	if err != nil {
		return fmt.Errorf("failed to register component factory: %w", err)
	}

	// Create the component
	component, err := registrar.CreateComponent(ctx, config)
	if err != nil {
		return fmt.Errorf("failed to create component: %s", config.ID)
	}

	// Initial system services and operations
	switch v := component.(type) {
	case SystemOperationInterface:
		v.Initialize(ctx, system)
	case SystemServiceInterface:
		v.Initialize(ctx, system)
	default:

	}

	return nil
}
