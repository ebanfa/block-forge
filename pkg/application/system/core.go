package system

import (
	"github.com/edward1christian/block-forge/pkg/application/common/context"
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
