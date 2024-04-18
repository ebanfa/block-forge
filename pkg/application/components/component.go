package components

import (
	"errors"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	configApi "github.com/edward1christian/block-forge/pkg/application/config"
)

// Custom errors
var (
	ErrComponentNil                  = errors.New("component is nil")
	ErrComponentAlreadyExist         = errors.New("component already exists")
	ErrFactoryNotFound               = errors.New("factory not found")
	ErrComponentFactoryNil           = errors.New("component factory is nil")
	ErrComponentFactoryAlreadyExists = errors.New("component factory already exists")
	ErrComponentNotFound             = errors.New("component not found")
	ErrComponentTypeNotFound         = errors.New("component type not found")
)

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
	CreateComponent(config *configApi.ComponentConfig) (ComponentInterface, error)
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

// BaseComponent represents a concrete implementation of the OperationInterface.
type BaseComponent struct {
	Id   string
	Nm   string
	Desc string
}

func NewComponentImpl(Id, Nm, Desc string) *BaseComponent {
	return &BaseComponent{Id: Id, Nm: Nm, Desc: Desc}
}

// ID returns the unique identifier of the component.
func (bc *BaseComponent) ID() string {
	return bc.Id
}

// Name returns the Nm of the component.
func (bc *BaseComponent) Name() string {
	return bc.Nm
}

// Type returns the type of the component.
func (bc *BaseComponent) Type() ComponentType {
	return BasicComponentType
}

// Description returns the Desc of the component.
func (bc *BaseComponent) Description() string {
	return bc.Desc
}
