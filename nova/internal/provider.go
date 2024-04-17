package application

import (
	"fmt"

	"github.com/edward1christian/block-forge/pkg/application/common/event"
	"github.com/edward1christian/block-forge/pkg/application/common/logger"
	"github.com/edward1christian/block-forge/pkg/application/components"
	"github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"go.uber.org/fx"
)

type InitOptions struct {
	Debug          bool
	Verbose        bool
	configFilePath string
}

// Init initializes the Fx application.
func Init(options *InitOptions) {
	// Create an Fx application.
	app := fx.New(
		// Provide dependencies.
		fx.Provide(ProvideConfiguration(options)),
		fx.Provide(ProvideEventBus),
		fx.Provide(ProvideLogger),
		fx.Provide(ProvideSystem),
	)
	// Run the application.
	app.Run()
}

// ProvideConfiguration loads and provides the application configuration.
func ProvideConfiguration(options *InitOptions) func() (*config.Configuration, error) {
	return func() (*config.Configuration, error) {
		var appConfig interface{}
		if options.configFilePath != "" {
			if err := config.LoadConfigurationFromFile(options.configFilePath, &appConfig); err != nil {
				return nil, fmt.Errorf("failed to load custom configuration: %v", err)
			}
		}

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
func ProvideLogger() logger.LoggerInterface {
	return logger.NewLogrusLogger()
}

// ProvideSystem provides a system interface.
func ProvideSystem(
	logger logger.LoggerInterface,
	eventBus event.EventBusInterface,
	configuration *components.Configuration,
	registrar components.ComponentRegistrar) system.SystemInterface {
	return system.NewSystem(logger, eventBus, configuration, registrar)
}
