// Package provider provides functionality to initialize and configure the system components
// using dependency injection with the Fx framework.
package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/edward1christian/block-forge/nova/pkg/components/plugin"
	contextApi "github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/common/event"
	"github.com/edward1christian/block-forge/pkg/application/common/logger"
	"github.com/edward1christian/block-forge/pkg/application/component"
	"github.com/edward1christian/block-forge/pkg/application/config"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
	"go.uber.org/fx"
)

// InitOptions contains options for initializing the system.
type InitOptions struct {
	Debug          bool
	Verbose        bool
	ConfigFilePath string
}

// CommandOptions contains options for executing a command.
type CommandOptions struct {
	Debug   bool
	Verbose bool
	Command string
	Data    *systemApi.SystemOperationInput
}

// Init initializes the Fx application with the provided options.
func Init(options *CommandOptions) {
	// Create an Fx application.
	app := fx.New(
		fx.NopLogger,
		// Provide dependencies.
		fx.Provide(ProvideConfiguration(options)),
		fx.Provide(ProvideLogger(options)),
		fx.Provide(ProvideEventBus),
		fx.Provide(ProvidComponentRegistrar),
		fx.Provide(ProvidPluginManager),
		fx.Provide(ProvideSystem(options)),
		fx.Invoke(func(systemApi.SystemInterface) {}),
	)
	// Run the application.
	app.Run()
}

// ProvideConfiguration provides a function to load and provide the application configuration.
func ProvideConfiguration(options *CommandOptions) func() (*config.Configuration, error) {
	return func() (*config.Configuration, error) {
		var appConfig interface{}
		// Load custom configuration if provided.
		/* if options.ConfigFilePath != "" {
			if err := config.LoadConfigurationFromFile(options.ConfigFilePath, &appConfig); err != nil {
				return nil, fmt.Errorf("failed to load custom configuration: %v", err)
			}
		} */

		configuration := &config.Configuration{
			Debug:        options.Debug,
			Verbose:      options.Verbose,
			CustomConfig: appConfig,
		}

		return configuration, nil
	}
}

// ProvideEventBus provides an event bus interface.
func ProvideEventBus() event.EventBusInterface {
	return event.NewSystemEventBus()
}

// ProvideLogger provides a logger interface.
func ProvideLogger(options *CommandOptions) func() logger.LoggerInterface {
	return func() logger.LoggerInterface {
		var level logger.Level
		if options.Debug {
			level = logger.LevelDebug
		} else {
			level = logger.LevelInfo
		}
		return logger.NewLogrusLogger(level)
	}
}

// ProvidComponentRegistrar provides a component registrar interface.
func ProvidComponentRegistrar() component.ComponentRegistrarInterface {
	return component.NewComponentRegistrar()
}

// ProvidPluginManager provides a plugin manager interface.
func ProvidPluginManager() systemApi.PluginManagerInterface {
	return systemApi.NewPluginManager()
}

// ProvideSystem provides a function to configure and provide a system instance.
func ProvideSystem(options *CommandOptions) func(
	lc fx.Lifecycle,
	logger logger.LoggerInterface,
	eventBus event.EventBusInterface,
	configuration *config.Configuration,
	pluginManager systemApi.PluginManagerInterface,
	registrar component.ComponentRegistrarInterface) systemApi.SystemInterface {
	return func(
		lc fx.Lifecycle,
		logger logger.LoggerInterface,
		eventBus event.EventBusInterface,
		configuration *config.Configuration,
		pluginManager systemApi.PluginManagerInterface,
		registrar component.ComponentRegistrarInterface) systemApi.SystemInterface {

		// Create a new system instance with the provided dependencies.
		system := systemApi.NewSystem(logger, eventBus, configuration, pluginManager, registrar)

		// Add lifecycle hooks to start and stop the system.
		lc.Append(fx.Hook{
			OnStart: OnStart(options, system),
			OnStop:  OnStop(options, system),
		})

		return system
	}
}

// OnStart returns a function to initialize the system and execute a command on system start.
func OnStart(options *CommandOptions, system systemApi.SystemInterface) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		contx := contextApi.WithContext(ctx)

		// Initialize the system and execute the command.
		if err := InitializeSystem(contx, system); err != nil {
			return err
		}

		// Execute the command.
		return ExecuteCommand(contx, options, system)
		//return nil
	}
}

// OnStop returns a function to stop the system on system shutdown.
func OnStop(options *CommandOptions, system systemApi.SystemInterface) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		contx := contextApi.WithContext(ctx)
		return system.Stop(contx)
	}
}

// InitializeSystem initializes the system with the provided context and system interface.
func InitializeSystem(ctx *contextApi.Context, system systemApi.SystemInterface) error {
	contx := contextApi.WithContext(ctx)

	// Initialize the system.
	if err := system.Initialize(contx); err != nil {
		return fmt.Errorf("failed to initialize system: %w", err)
	}

	// Add the Nova plugin to the system.
	if err := AddPlugin(contx, system, plugin.NewNovaPlugin()); err != nil {
		return fmt.Errorf("failed to add plugin: %w", err)
	}

	// Get the home directory of the current user
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user's home directory: %w", err)
	}

	// Create input data containing the home directory
	input := &systemApi.SystemOperationInput{
		Data: homeDir,
	}

	// Define the command options
	options := &CommandOptions{
		Command: "InitDirectoriesOperation",
		Data:    input,
	}

	// Execute the command
	if err := ExecuteCommand(ctx, options, system); err != nil {
		return fmt.Errorf("failed to initialize directories: %w", err)
	}

	return nil
}

// AddPlugin adds a plugin to the system using the provided context and system interface.
func AddPlugin(ctx *contextApi.Context, system systemApi.SystemInterface, p systemApi.PluginInterface) error {
	// Add the provided plugin to the plugin manager.
	if err := system.PluginManager().AddPlugin(ctx, p); err != nil {
		return err
	}
	return nil
}

// ExecuteCommand executes a command on the provided system using the provided context.
func ExecuteCommand(ctx *contextApi.Context, options *CommandOptions, system systemApi.SystemInterface) error {
	// Execute the operation using the provided system interface.
	if _, err := system.ExecuteOperation(ctx, options.Command, options.Data); err != nil {
		return fmt.Errorf("failed to execute operation: %w", err)
	}
	return nil
}
