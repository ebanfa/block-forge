package application

import (
	"fmt"

	"github.com/edward1christian/block-forge/pkg/application/appl"
	"github.com/edward1christian/block-forge/pkg/application/common/event"
	"github.com/edward1christian/block-forge/pkg/application/common/logger"
	"github.com/edward1christian/block-forge/pkg/application/components"
	"github.com/edward1christian/block-forge/pkg/application/config"
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
func ProvideSystem(
	logger logger.LoggerInterface,
	eventBus event.EventBusInterface,
	configuration *components.Configuration,
	registrar components.ComponentRegistrar) system.SystemInterface {
	return system.NewSystem(logger, eventBus, configuration, registrar)
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
func ProvideApplication(moduleManager appl.ModuleManager, system system.SystemInterface) appl.ApplicationInterface {
	return appl.NewApplication(moduleManager, system)
}
