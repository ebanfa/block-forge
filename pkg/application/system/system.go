package system

import (
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/common/event"
	"github.com/edward1christian/block-forge/pkg/application/common/logger"
	"github.com/edward1christian/block-forge/pkg/application/components"
)

// ComponentType represents the type of a component.
type SystemStatusType int

const (
	// OperationType represents the type of an operation component.
	SystemInitializedType SystemStatusType = iota

	// ServiceType represents the type of a service component.
	SystemStartedType

	// ModuleType represents the type of a module component.
	SystemStoppedType
)

// SystemComponentInterface represents a component in the system.
type SystemComponentInterface interface {
	components.ComponentInterface

	// Initialize initializes the module.
	// Returns an error if the initialization fails.
	Initialize(ctx *context.Context, system SystemInterface) error
}

// SystemServiceInterface represents a service within the system.
type SystemServiceInterface interface {
	components.StartableInterface
	SystemComponentInterface
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
type OperationInterface interface {
	SystemComponentInterface

	// Execute performs the operation with the given context and input parameters,
	// and returns any output or error encountered.
	Execute(ctx *context.Context, input *OperationInput) (*OperationOutput, error)
}

// SystemInterface represents the core system in the application.
type SystemInterface interface {
	components.BootableInterface
	components.StartableInterface

	// Logger returns the system logger.
	Logger() logger.LoggerInterface

	// EventBus returns the system event bus.
	EventBus() event.EventBusInterface

	// Configuration returns the system configuration.
	Configuration() components.Configuration

	// ComponentRegistry returns the component registry
	ComponentRegistry() components.ComponentRegistrar

	// ExecuteOperation executes the operation with the given ID and input data.
	// Returns the output of the operation and an error if the operation is not found or if execution fails.
	ExecuteOperation(ctx *context.Context, operationID string, data *OperationInput) (*OperationOutput, error)

	// StartService starts the service with the given ID.
	// Returns an error if the service ID is not found or other error.
	StartService(ctx *context.Context, serviceID string) error

	// StopService stops the service with the given ID.
	// Returns an error if the service ID is not found or other error.
	StopService(ctx *context.Context, serviceID string) error

	// RestartService restarts the service with the given ID.
	// Returns an error if the service ID is not found or other error.
	RestartService(ctx *context.Context, serviceID string) error
}
