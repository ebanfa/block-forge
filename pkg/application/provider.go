package application

import (
	"fmt"

	"github.com/edward1christian/block-forge/pkg/application/appl"
	"github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/edward1christian/block-forge/pkg/application/event"
	"github.com/edward1christian/block-forge/pkg/application/logger"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"go.uber.org/fx"
)

// Init initializes the Fx application.
func Init(appConfigFile string, frameworkConfigFile string) {
	// Create an Fx application.
	app := fx.New(
		// Provide dependencies.
		fx.Provide(ProvideConfiguration(appConfigFile, frameworkConfigFile)),
		fx.Provide(ProvideEventBus),
		fx.Provide(ProvideModuleManager),
		fx.Provide(ProvideLogger),
		fx.Provide(ProvideSystem),
		fx.Provide(ProvideApplication),
	)
	// Run the application.
	app.Run()
}

// ProvideEventBus provides an event bus interface.
func ProvideEventBus() event.EventBusInterface {
	return event.NewSystemEventBus()
}

// ProvideModuleManager provides a module manager interface.
func ProvideModuleManager(id string, name string, description string) appl.ModuleManager {
	return appl.NewModuleManager(id, name, description)
}

// ProvideSystem provides a system interface.
func ProvideSystem(eventBus event.EventBusInterface, logger logger.LoggerInterface, configuration system.Configuration) system.System {
	return system.NewSystem(eventBus, logger, configuration)
}

// ProvideConfiguration loads and provides the application configuration.
func ProvideConfiguration(appConfigFile, frameworkConfigFile string) func() (*config.Configuration, error) {
	return func() (*config.Configuration, error) {
		var appConfig interface{}
		if appConfigFile != "" {
			if err := config.LoadConfigurationFromFile(appConfigFile, &appConfig); err != nil {
				return nil, fmt.Errorf("failed to load custom configuration: %v", err)
			}
		}

		var frameworkConfig config.ApplicationConfig
		if err := config.LoadConfigurationFromFile(frameworkConfigFile, &frameworkConfig); err != nil {
			return nil, fmt.Errorf("failed to load framework configuration: %v", err)
		}

		// Initialize the configuration struct
		configuration := &config.Configuration{
			ApplicationConfig: frameworkConfig,
			CustomConfig:      appConfig,
		}

		return configuration, nil
	}
}

// ProvideLogger provides a logger interface.
func ProvideLogger() logger.LoggerInterface {
	return logger.NewLogrusLogger()
}

// ProvideApplication provides the application interface.
func ProvideApplication(
	ID string,
	name string,
	description string,
	moduleManager appl.ModuleManager,
	system system.System) appl.Application {
	return appl.NewApplication(ID, name, description, moduleManager, system)
}
