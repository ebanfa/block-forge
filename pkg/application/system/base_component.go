package system

import (
	"errors"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/components"
)

// BaseComponent represents a concrete implementation of the OperationInterface.
type BaseSystemComponent struct {
	components.BaseComponent // Embedding BaseComponent
}

// Type returns the type of the component.
func (bo *BaseSystemComponent) Type() components.ComponentType {
	return components.SystemComponentType
}

func NewBaseSystemComponent(id, name, description string) *BaseSystemComponent {
	return &BaseSystemComponent{
		BaseComponent: components.BaseComponent{
			Id:   id,
			Nm:   name,
			Desc: description,
		},
	}
}

// Initialize initializes the module.
// Returns an error if the initialization fails.
func (bo *BaseSystemComponent) Initialize(ctx *context.Context, system SystemInterface) error {
	// Perform initialization tasks specific to operation component if needed.
	return errors.New("initialize not implemented")
}
