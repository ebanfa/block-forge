// Package provider provides functionality to initialize and configure the system components
// using dependency injection with the Fx framework.
package provider

import (
	"context"
	"errors"
	"fmt"

	valid "github.com/asaskevich/govalidator"

	novaConfigApi "github.com/edward1christian/block-forge/nova/pkg/config"

	dbm "github.com/cosmos/iavl/db"
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
	MultiDbName = ""
	MultiDbPath = ""
)

// InitOptions represents system initialization options.
type InitOptions struct {
	Daemon  bool        `valid:"type(bool),optional"` //Run in daemon mode
	Debug   bool        `valid:"type(bool),optional"` // Debug mode flag
	Verbose bool        `valid:"type(bool),optional"` // Verbose mode flag
	Command string      `valid:"alpha,optional"`      // Command to execute during initialization
	Data    interface{} `valid:"-"`                   // Data for system initialization
}

// Init initializes the Fx application with the provided options.
func Init(options *InitOptions) {
	// Validate the user struct
	if _, err := valid.ValidateStruct(options); err != nil {
		// Handle validation errors
		fmt.Println("Validation errors:")
		for _, e := range err.(valid.Errors) {
			fmt.Println(e.Error())
		}
		return
	}

	// Create an Fx application.
	app := fx.New(
		//fx.NopLogger,
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

		configuration, err := novaConfigApi.GetDefaultConfig()
		if err != nil {
			return nil, err
		}

		return &config.Configuration{
			Debug:        options.Debug,
			Verbose:      options.Verbose,
			CustomConfig: configuration,
		}, nil
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

// ProvideMultiStore creates a function that generates a MultiStore interface based on the provided options.
func ProvideMultiStore(options *InitOptions) func(configuration *config.Configuration) (store.MultiStore, error) {
	return func(configuration *config.Configuration) (store.MultiStore, error) {
		// Get the configuration
		novaConfig, ok := configuration.CustomConfig.(novaConfigApi.NovaConfig)
		if !ok {
			return nil, errors.New("invalid configuration for creating multistore")
		}

		// Create the underlying data store factory
		storeFactory := store.NewStoreFactory(
			novaConfig.DatabasesDir, db.NewIAVLDatabaseFactory(dbm.NewDB))

		// Create the MultiStore
		multiStore, err := store.CreateMultiStore(
			novaConfig.MultiStoreDbName, novaConfig.DatabasesDir, storeFactory)

		if err != nil {
			return nil, err
		}

		return multiStore, nil
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
func ProvideSystem(options *InitOptions) SystemProvider {
	return SystemProviderFn(options)
}

// OnStart returns a function to initialize the system and execute a command on system start.
func OnStart(options *InitOptions, system systemApi.SystemInterface, shutdowner fx.Shutdowner) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		contx := contextApi.WithContext(ctx)

		// Initialize the system and execute the command.
		if err := InitializeSystem(contx, system); err != nil {
			return err
		}

		// If a command is provided, execute it.
		if options.Command != "" {
			// Execute the specified command.
			result, err := system.ExecuteOperation(contx, options.Command, &systemApi.SystemOperationInput{
				Data: options.Data,
			})
			if err != nil {
				// If an error occurs during execution, return a formatted error message.
				return fmt.Errorf("failed to execute operation '%s': %w", options.Command, err)
			}

			// Optionally handle the result of the operation, if needed.
			_ = result // You may use the result here if necessary.
		}

		// Check if the system is running as a daemon
		if !options.Daemon {
			// If not running as a daemon, perform shutdown
			return shutdowner.Shutdown()
		}

		return nil
	}
}

// OnStop returns a function to stop the system on system shutdown.
func OnStop(options *InitOptions, system systemApi.SystemInterface) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		contx := contextApi.WithContext(ctx)

		// Check if the system is running as a daemon
		if options.Daemon {
			// If not running as a daemon, perform shutdown
			return system.Stop(contx)
		}
		return nil
	}
}
