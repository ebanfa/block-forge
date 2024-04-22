package plugin

import (
	"errors"
	"fmt"

	"github.com/edward1christian/block-forge/nova/pkg/common"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/common/logger"
	"github.com/edward1christian/block-forge/pkg/application/component"
	"github.com/edward1christian/block-forge/pkg/application/config"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
)

// createAndStartBuilderService creates and starts the BuilderService component.
func StartBuildService(ctx *context.Context, system systemApi.SystemInterface) error {
	system.Logger().Log(logger.LevelInfo, "NovaPluginHelper: Creating build service")

	// Create and start the BuilderService component
	buildService, err := CreateBuilderService(ctx, system.ComponentRegistry())
	if err != nil {
		return err
	}
	system.Logger().Log(logger.LevelInfo, "NovaPluginHelper: Initializing service:"+buildService.ID())

	// Initialize the build service
	if err := buildService.Initialize(ctx, system); err != nil {
		return fmt.Errorf("failed to initialize BuilderService: %v", err)
	}

	system.Logger().Log(logger.LevelInfo, "NovaPluginHelper: Starting service:"+buildService.ID())
	return buildService.Start(ctx)
}

// Helper function to create the BuilderService component
func CreateBuilderService(ctx *context.Context, registrar component.ComponentRegistrarInterface) (systemApi.SystemServiceInterface, error) {
	component, err := registrar.CreateComponent(ctx, &component.ComponentCreationInfo{
		FactoryID: common.IgniteBuildServiceFactory,
		Config: &config.ComponentConfig{
			ID:        common.IgniteBuildService,
			FactoryID: common.IgniteBuildServiceFactory,
		},
	})
	if err != nil {
		return nil, err
	}
	// Ensure the component implements SystemServiceInterface
	builderService, ok := component.(systemApi.SystemServiceInterface)
	if !ok {
		return nil, errors.New("BuildService component does not implement SystemServiceInterface")
	}
	return builderService, nil
}
