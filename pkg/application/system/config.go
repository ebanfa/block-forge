package system

import "time"

// ComponentConfig represents the configuration for a component.
type ComponentConfig struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	FactoryName  string      // Name of the factory to use for creating the service
	CustomConfig interface{} // Custom configuration
}

type ServiceConfiguration struct {
	ComponentConfig
	RetryInterval time.Duration // Interval between retries
	// Other service-specific configuration optionsetl.ErrScheduledProcessNotFound
	CustomConfig interface{} // Custom configuration
}

// OperationConfiguration represents the configuration for an operation.
type OperationConfiguration struct {
	ComponentConfig
	// Other operation-specific configuration optionsetl.ErrScheduledProcessNotFound
}

// LoggerConfiguration represents the configuration for the logger.
type LoggerConfiguration struct {
	Level  string // Log level: debug, info, warn, error, etc.
	Format string // Log format: text, json, etc.
	// Other logger configuration optionsetl.ErrScheduledProcessNotFound
}

// EventBusConfiguration represents the configuration for the event bus.
type EventBusConfiguration struct {
}

// Configuration represents the system configuration.
type Configuration struct {
	Logger       LoggerConfiguration       // Logger configuration
	EventBus     EventBusConfiguration     // Event bus configuration
	Services     []*ServiceConfiguration   // Service configurations
	Operations   []*OperationConfiguration // Operation configurations
	CustomConfig interface{}
}
