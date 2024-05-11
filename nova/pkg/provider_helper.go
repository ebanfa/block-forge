package provider

import (
	"errors"
	"os"

	"github.com/edward1christian/block-forge/nova/pkg/components/plugin"
	contextApi "github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/common/event"
	"github.com/edward1christian/block-forge/pkg/application/common/logger"
	"github.com/edward1christian/block-forge/pkg/application/component"
	"github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/edward1christian/block-forge/pkg/application/store"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
	"go.uber.org/fx"
)

type SystemProvider func(
	lc fx.Lifecycle,
	shutdowner fx.Shutdowner,
	logger logger.LoggerInterface,
	eventBus event.EventBusInterface,
	configuration *config.Configuration,
	pluginManager systemApi.PluginManagerInterface,
	registrar component.ComponentRegistrarInterface,
	multiStore store.MultiStore) systemApi.SystemInterface

// Returns FX provider function.
func SystemProviderFn(options *InitOptions) SystemProvider {
	return func(
		lc fx.Lifecycle,
		shutdowner fx.Shutdowner,
		logger logger.LoggerInterface,
		eventBus event.EventBusInterface,
		configuration *config.Configuration,
		pluginManager systemApi.PluginManagerInterface,
		registrar component.ComponentRegistrarInterface,
		multiStore store.MultiStore) systemApi.SystemInterface {

		// Create a new system instance with the provided dependencies.
		system := systemApi.NewSystem(
			logger, eventBus, configuration,
			pluginManager, registrar, multiStore)

		// Add lifecycle hooks to start and stop the system.
		lc.Append(fx.Hook{
			OnStart: OnStart(options, system, shutdowner),
			OnStop:  OnStop(options, system),
		})

		return system
	}
}

// InitializeSystem initializes the system with the provided context, system initializer, and operation executor.
func InitializeSystem(ctx *contextApi.Context, system systemApi.SystemInterface) error {
	// Initialize the system.
	if err := system.Initialize(ctx); err != nil {
		return errors.New("failed to initialize system: " + err.Error())
	}

	// Add the Nova plugin to the system.
	if err := AddNovaPlugin(ctx, system); err != nil {
		return errors.New("failed to add Nova plugin: " + err.Error())
	}

	// Get the home directory of the current user.
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.New("failed to get user's home directory: " + err.Error())
	}

	// Execute the operation to initialize directories.
	_, err = system.ExecuteOperation(ctx, "InitDirectoriesOperation", &systemApi.SystemOperationInput{
		Data: homeDir,
	})
	if err != nil {
		return errors.New("failed to execute operation: " + err.Error())
	}

	return nil
}

// AddNovaPlugin adds the Nova plugin to the system.
func AddNovaPlugin(ctx *contextApi.Context, system systemApi.SystemInterface) error {
	return AddPlugin(ctx, system, plugin.NewNovaPlugin())
}

// AddPlugin adds a plugin to the system using the provided context and system interface.
func AddPlugin(ctx *contextApi.Context, system systemApi.SystemInterface, p systemApi.PluginInterface) error {
	// Add the provided plugin to the plugin manager.
	if err := system.PluginManager().AddPlugin(ctx, p); err != nil {
		return err
	}
	return nil
}
