package component

import (
	"github.com/edward1christian/block-forge/pkg/application/context"
)

// Component represents a generic component in the system.
type Component interface {
	// ID returns the unique identifier of the component.
	ID() string

	// Name returns the name of the component.
	Name() string

	// Description returns the description of the component.
	Description() string
}

// Startable defines the interface for components that can be started and stopped.
type Startable interface {
	// Start starts the component.
	Start(ctx *context.Context) error

	// Stop stops the component.
	Stop(ctx *context.Context) error
}

// Startable defines the interface for components that can be started and stopped.
type StartableComponent interface {
	Component
	Startable
}
