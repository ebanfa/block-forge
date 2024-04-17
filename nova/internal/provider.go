package internal

import (
	"context"
	"fmt"

	contextApi "github.com/edward1christian/block-forge/pkg/application/common/context"
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
	ConfigFilePath string
}

// Init initializes the Fx application.
func Init(options *InitOptions) {
	// Create an Fx application.
	app := fx.New(
		// Provide dependencies.
		fx.Provide(ProvideLogger),
		fx.Provide(ProvideEventBus),
		fx.Provide(ProvideConfiguration(options)),
		fx.Provide(ProvidComponentRegistrar),
		fx.Provide(ProvideSystem),
		fx.Invoke(func(system.SystemInterface) {}),
	)
	// Run the application.
	app.Run()
}

// ProvideConfiguration loads and provides the application configuration.
func ProvideConfiguration(options *InitOptions) func() (*config.Configuration, error) {
	return func() (*config.Configuration, error) {
		var appConfig interface{}
		if options.ConfigFilePath != "" {
			if err := config.LoadConfigurationFromFile(options.ConfigFilePath, &appConfig); err != nil {
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

// ProvidComponentRegistrar provides a component registrar interface.
func ProvidComponentRegistrar() components.ComponentRegistrar {
	return components.NewComponentRegistrar()
}

// ProvideLogger provides a logger interface.
func ProvideLogger() logger.LoggerInterface {
	return logger.NewLogrusLogger()
}

// ProvideSystem provides a system interface.
func ProvideSystem(
	lc fx.Lifecycle,
	logger logger.LoggerInterface,
	eventBus event.EventBusInterface,
	configuration *config.Configuration,
	registrar components.ComponentRegistrar) system.SystemInterface {

	sys := system.NewSystem(logger, eventBus, configuration, registrar)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			contx := contextApi.WithContext(ctx)
			err := sys.Initialize(contx)
			if err != nil {
				return err
			}

			return sys.Start(contx)
		},
		OnStop: func(ctx context.Context) error {
			contx := contextApi.WithContext(ctx)
			return sys.Stop(contx)
		},
	})
	return sys
}