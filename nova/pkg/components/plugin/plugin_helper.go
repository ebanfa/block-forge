package plugin

import (
	"fmt"

	"github.com/edward1christian/block-forge/nova/pkg/common"
	"github.com/edward1christian/block-forge/nova/pkg/components/operations"
	"github.com/edward1christian/block-forge/nova/pkg/components/services"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/config"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
)

func RegisterServices(ctx *context.Context, system systemApi.SystemInterface) error {
	registrar := system.ComponentRegistry()

	// Register build service
	err := registrar.RegisterFactory(ctx, common.BuildServiceFactory, &services.BuildServiceFactory{})
	if err != nil {
		return fmt.Errorf("failed to register build service: %w", err)
	}
	// Register API service// Register build service
	/* err = registrar.RegisterFactory(ctx, common.BuildServiceFactory, &services.BuildServiceFactory{})
	if err != nil {
		return fmt.Errorf("failed to register REST API service: %w", err)
	} */
	return nil
}

func RegisterOperations(ctx *context.Context, system systemApi.SystemInterface) error {
	// Register Build Operations
	if err := RegisterBuildOperations(ctx, system); err != nil {
		return fmt.Errorf("failed to register build operations: %w", err)
	}
	// Register Other Operations
	return nil
}

func RegisterBuildOperations(ctx *context.Context, system systemApi.SystemInterface) error {
	config := &config.ComponentConfig{
		ID:        common.CreateWorkspaceTask,
		FactoryID: common.CreateWorkspaceTaskFactory,
	}
	return systemApi.RegisterComponent(ctx, system, config, operations.NewCreateDirectoryTaskFactory())
}

func StartServices(ctx *context.Context, system systemApi.SystemInterface) error {
	// Start the build service
	if err := systemApi.StartService(ctx, system, getBuildServiceConfig()); err != nil {
		return fmt.Errorf("failed to start BuildService: %v", err)
	}
	// Start the RestFul API service
	/* if err := systemApi.StartService(ctx, system, getAPIServiceConfig()); err != nil {
		return fmt.Errorf("failed to start BuildService: %v", err)
	} */
	return nil
}

func StopServices(ctx *context.Context, system systemApi.SystemInterface) error {
	// Stop the build service
	if err := systemApi.StopService(ctx, system, common.BuildService); err != nil {
		return fmt.Errorf("failed to stop service: %s %v", common.BuildService, err)
	}
	// Stop the RestFul API service
	if err := systemApi.StopService(ctx, system, common.BuildService); err != nil {
		return fmt.Errorf("failed to stop service: %s %v", common.BuildService, err)
	}
	return nil
}

func getBuildServiceConfig() *config.ComponentConfig {
	return &config.ComponentConfig{
		ID:        common.BuildService,
		FactoryID: common.BuildServiceFactory,
	}
}

func getAPIServiceConfig() *config.ComponentConfig {
	return &config.ComponentConfig{
		ID:        common.APIService,
		FactoryID: common.APIServiceFactory,
	}
}
