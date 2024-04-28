package application

import (
	"context"
	"fmt"

	contextApi "github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/common/event"
	"github.com/edward1christian/block-forge/pkg/application/common/logger"
	"github.com/edward1christian/block-forge/pkg/application/component"
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
		fx.Provide(ProvideLogger(options)),
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
func ProvideLogger(options *InitOptions) func() logger.LoggerInterface {
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

// ProvideSystem provides a system interface.
func ProvideSystem(
	lc fx.Lifecycle,
	logger logger.LoggerInterface,
	eventBus event.EventBusInterface,
	configuration *config.Configuration,
	pluginManager system.PluginManagerInterface,
	registrar component.ComponentRegistrarInterface) system.SystemInterface {

	sys := system.NewSystem(logger, eventBus, configuration, pluginManager, registrar)

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
