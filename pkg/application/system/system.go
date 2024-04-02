package system

import (
	"github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/application/event"
	"github.com/edward1christian/block-forge/pkg/application/logger"
)

// Operation represents a unit of work that can be executed.
type Operation interface {

	// ID returns the unique identifier of the component.
	ID() string

	// Name returns the name of the component.
	Name() string

	// Description returns the description of the component.
	Description() string

	// Initialize initializes the module.
	Initialize(ctx *context.Context, system System) error

	// Execute performs the operation with the given context and input parameters, and returns any output or error encountered.
	Execute(ctx *context.Context, input OperationInput) (OperationOutput, error)
}

// Operations represents the interface for managing operations.
type Operations interface {

	// ID returns the unique identifier of the component.
	ID() string

	// Name returns the name of the component.
	Name() string

	// Description returns the description of the component.
	Description() string

	// Initialize initializes the module.
	Initialize(ctx *context.Context, system System) error

	// RegisterOperation registers an operation with the given ID.
	// Returns an error if the operation ID is already registered or if the operation is nil.
	RegisterOperation(operationID string, operation Operation) error

	// ExecuteOperation executes the operation with the given ID and input data.
	// Returns the output of the operation and an error if the operation is not found or if execution fails.
	ExecuteOperation(ctx *context.Context, operationID string, data OperationInput) (OperationOutput, error)
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

// System represents the core system in the application.
type System interface {
	EventBus() event.EventBusInterface
	Operations() Operations
	Logger() logger.LoggerInterface

	// Configuration returns the system configuration.
	Configuration() config.Configuration
}
