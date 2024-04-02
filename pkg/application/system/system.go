package system

import (
	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/application/event"
	"github.com/edward1christian/block-forge/pkg/application/logger"
)

// Component represents a generic component.
type Component interface {
	// ID returns the unique identifier of the component.
	ID() string

	// Name returns the name of the component.
	Name() string

	// Description returns the description of the component.
	Description() string
}

// ComponentFactory defines the function signature for creating a component.
type ComponentFactory func(ctx *context.Context, config *ComponentConfig) (Component, error)

// Startable defines the interface for instances that can be started and stopped.
type Startable interface {
	// Start starts the component.
	Start(ctx *context.Context) error

	// Stop stops the component.
	Stop(ctx *context.Context) error
}

// StartableComponent defines the interface for components that can be started and stopped.
type StartableComponent interface {
	Component
	Startable
}

// BootableComponent represents a component that can be initialized and started.
type BootableComponent interface {
	StartableComponent // Embedding StartableComponent interface

	// Initialize initializes the component.
	Initialize(ctx *context.Context) error
}

// SystemComponent represents a component in the system.
type SystemComponent interface {
	Component

	// Initialize initializes the module.
	Initialize(ctx *context.Context, system System) error
}

// OperationInput represents the input data for an operation.
type OperationInput struct {
	// Data is the input data for the operation.
	Data interface{}
}

// OperationOutput represents the response data from an operation.
type OperationOutput struct {
	// Data is the response data from the operation.
	Data interface{}
}

// Operation represents a unit of work that can be executed.
type Operation interface {
	SystemComponent

	// Execute performs the operation with the given context and input parameters,
	// and returns any output or error encountered.
	Execute(ctx *context.Context, input *OperationInput) (*OperationOutput, error)
}

// SystemService represents a service within the system.
type SystemService interface {
	Startable
	SystemComponent
}

// System represents the core system in the application.
type System interface {
	BootableComponent

	// Logger returns the system logger.
	Logger() logger.LoggerInterface

	// EventBus returns the system event bus.
	EventBus() event.EventBusInterface

	// Configuration returns the system configuration.
	Configuration() Configuration

	// RegisterOperation registers an operation with the given ID.
	// Returns an error if the operation ID is already registered or if the operation is nil.
	RegisterOperation(operationID string, operation Operation) error

	// UnregisterOperation unregisters the operation with the given ID.
	// Returns an error if the operation ID is not found.
	UnregisterOperation(operationID string) error

	// ExecuteOperation executes the operation with the given ID and input data.
	// Returns the output of the operation and an error if the operation is not found or if execution fails.
	ExecuteOperation(ctx *context.Context, operationID string, data *OperationInput) (*OperationOutput, error)

	// RegisterService registers a SystemService with the given ID.
	// Returns an error if the service ID is already registered or if the service is nil.
	RegisterService(serviceID string, service SystemService) error

	// UnregisterService unregisters a SystemService with the given ID.
	// Returns an error if the service ID is not found.
	UnregisterService(serviceID string) error

	// StartService starts the service with the given ID.
	// Returns an error if the service ID is not found or other error
	StartService(serviceID string, ctx *context.Context) error

	// StopService stops the service with the given ID.
	// Returns an error if the service ID is not found or other error.
	StopService(serviceID string, ctx *context.Context) error

	// RegisterComponentFactory registers an component factory with the system.
	// Returns an error if the component name is already registered or if the component is nil.
	RegisterComponentFactory(name string, factory ComponentFactory) error

	// GetComponentFactory retrieves the factory for creating components by name.
	// Returns an error if the component name is not found or other error.
	GetComponentFactory(name string) (ComponentFactory, error)
}
