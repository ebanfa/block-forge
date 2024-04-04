package components

import "github.com/edward1christian/block-forge/pkg/application/common/context"

// ComponentType represents the type of a component.
type ComponentType int

const (
	// BasicComponentType represents the type of a basic component.
	BasicComponentType ComponentType = iota

	// SystemComponentType represents the type of a system component.
	SystemComponentType

	// OperationType represents the type of an operation component.
	OperationType

	// ServiceType represents the type of a service component.
	ServiceType

	// ModuleType represents the type of a module component.
	ApplicationComponentType
)

// ComponentInterface represents a generic component in the system.
type ComponentInterface interface {
	// ID returns the unique identifier of the component.
	ID() string
	// Name returns the name of the component.
	Name() string
	// Type returns the type of the component.
	Type() ComponentType
	// Description returns the description of the component.
	Description() string
}

// ComponentFactoryInterface is responsible for creating components.
type ComponentFactoryInterface interface {
	// CreateComponent creates a new instance of the component.
	// Returns the created component and an error if the creation fails.
	CreateComponent(config *ComponentConfig) (ComponentInterface, error)
}

// BootableComponentInterface represents a component that can be initialized and started.
type BootableInterface interface {
	// Initialize initializes the component.
	// Returns an error if the initialization fails.
	Initialize(ctx *context.Context) error
}

// Startable defines the interface for instances that can be started and stopped.
type StartableInterface interface {
	// Start starts the component.
	// Returns an error if the start operation fails.
	Start(ctx *context.Context) error

	// Stop stops the component.
	// Returns an error if the stop operation fails.
	Stop(ctx *context.Context) error
}
