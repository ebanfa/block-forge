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
	"github.com/edward1christian/block-forge/pkg/application/db"
	"github.com/edward1christian/block-forge/pkg/application/store"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
	"go.uber.org/fx"
)

const (
	DaemonMode  = "daemon"  // Indicates the system should run continuously as a daemon.
	CommandMode = "command" // Indicates the system should execute a single command and exit.
)

// InitOptions represents system initialization options.
type InitOptions struct {
	Debug    bool                            // Debug mode flag
	Verbose  bool                            // Verbose mode flag
	Command  string                          // Command to execute during initialization
	InitMode string                          // InitMode specifies the operational mode of the system, determining whether it runs continuously as a daemon or executes a single command and exits.
	Data     *systemApi.SystemOperationInput // Data for system initialization
}

// Init initializes the Fx application with the provided options.
func Init(options *InitOptions) {
	// Create an Fx application.
	app := fx.New(
		fx.NopLogger,
		// Provide dependencies.
		fx.Provide(ProvideConfiguration(options)),
		fx.Provide(ProvideLogger(options)),
		fx.Provide(ProvideEventBus),
		fx.Provide(ProvidComponentRegistrar),
		fx.Provide(ProvidPluginManager),
		fx.Provide(ProvideMultiStore(options)),
		fx.Provide(ProvideSystem(options)),
		fx.Invoke(func(systemApi.SystemInterface) {}),
	)
	// Run the application.
	app.Run()
}

// ProvideConfiguration provides a function to load and provide the application configuration.
func ProvideConfiguration(options *InitOptions) func() (*config.Configuration, error) {
	return func() (*config.Configuration, error) {
		var appConfig interface{}

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

// ProvideLogger provides a logger interface based on the initialization options.
func ProvideLogger(options *InitOptions) func() logger.LoggerInterface {
	// Return a function that creates a logger interface based on the provided options.
	return func() logger.LoggerInterface {
		// Determine the log level based on the debug option.
		level := logger.LevelInfo
		if options.Debug {
			level = logger.LevelDebug
		}

		// Create a new logger with the determined log level.
		return logger.NewLogrusLogger(level)
	}
}

func ProvideMultiStore(options *InitOptions) func() store.MultiStore {
	// Return a function that creates a MultiStore interface based on the provided options.
	return func() store.MultiStore {
		// Create the underlying database
		database, err := db.CreateIAVLDatabase(options.Name, options.Path)
		if err != nil {
			return nil
		}

		// Create a new logger with the determined log level.
		return store.NewMultiStore(database)
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
func ProvideSystem(options *InitOptions) func(
	lc fx.Lifecycle,
	shutdowner fx.Shutdowner,
	logger logger.LoggerInterface,
	eventBus event.EventBusInterface,
	configuration *config.Configuration,
	pluginManager systemApi.PluginManagerInterface,
	registrar component.ComponentRegistrarInterface,
	multiStore store.MultiStore) systemApi.SystemInterface {
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
		system := systemApi.NewSystem(logger, eventBus, configuration, pluginManager, registrar, multiStore)

		// Add lifecycle hooks to start and stop the system.
		lc.Append(fx.Hook{
			OnStart: OnStart(options, system, shutdowner),
			OnStop:  OnStop(options, system),
		})

		return system
	}
}

// OnStart returns a function to initialize the system and execute a command on system start.
func OnStart(options *InitOptions, system systemApi.SystemInterface, shutdowner fx.Shutdowner) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		contx := contextApi.WithContext(ctx)

		// Initialize the system and execute the command.
		if err := InitializeSystem(contx, system); err != nil {
			return err
		}

		// Execute the command.
		ExecuteCommand(contx, options, system)

		// If the system is not running in daemon mode, perform shutdown.
		return ShutdownIfNotDaemon(options, shutdowner)
	}
}

// OnStop returns a function to stop the system on system shutdown.
func OnStop(options *InitOptions, system systemApi.SystemInterface) func(ctx context.Context) error {
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
	options := &InitOptions{
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
func ExecuteCommand(ctx *contextApi.Context, options *InitOptions, system systemApi.SystemInterface) error {
	// Execute the operation using the provided system interface.
	if _, err := system.ExecuteOperation(ctx, options.Command, options.Data); err != nil {
		return fmt.Errorf("failed to execute operation: %w", err)
	}
	return nil
}

// ShutdownIfNotDaemon shuts down the system if it's not running in daemon mode.
// If the system is not running as a daemon, it invokes the shutdowner.Shutdown() function to gracefully shut down the system.
// If the system is running as a daemon, it returns nil indicating no action is needed.
func ShutdownIfNotDaemon(options *InitOptions, shutdowner fx.Shutdowner) error {
	// Check if the system is running as a daemon
	if options.InitMode != DaemonMode {
		// If not running as a daemon, perform shutdown
		return shutdowner.Shutdown()
	}
	// If running as a daemon, return nil indicating no action is needed
	return nil
}
