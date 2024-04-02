package config

import "time"

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
	Logger            LoggerConfiguration   // Logger configuration
	EventBus          EventBusConfiguration // Event bus configuration
	ApplicationConfig ApplicationConfig
	Services          []ServiceConfiguration   // Service configurations
	Operations        []OperationConfiguration // Operation configurations
	CustomConfig      interface{}
}

// Configuration represents the system configuration.
type ApplicationConfig struct {
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	ModuleLoadPath string   `json:"module_load_path"` // Path to load modules from
	Modules        []Module `json:"modules"`
}

// Module represents a module configuration in the configuration file.
type Module struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Path        string      `json:"path,omitempty"`
	Type        string      `json:"type,omitempty"`
	Config      interface{} `json:"config,omitempty"` // Module-specific configuration
}

// ServiceConfiguration represents the configuration for a service.
type ServiceConfiguration struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	FactoryName   string        // Name of the factory to use for creating the service
	MaxRetries    int           // Maximum number of retries for the service
	RetryInterval time.Duration // Interval between retries
	// Other service-specific configuration optionsetl.ErrScheduledProcessNotFound
	CustomConfig interface{} // Custom configuration
}

// OperationConfiguration represents the configuration for an operation.
type OperationConfiguration struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	FactoryName   string        // Name of the factory to use for creating the operation
	Timeout       time.Duration // Timeout for the operation
	RetryStrategy string        // Retry strategy for the operation (e.g., exponential backoff)
	// Other operation-specific configuration optionsetl.ErrScheduledProcessNotFound
	CustomConfig interface{} // Custom configuration
}
