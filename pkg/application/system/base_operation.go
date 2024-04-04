package system

import (
	"errors"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/components"
)

// BaseComponent represents a concrete implementation of the OperationInterface.
type BaseSystemOperation struct {
	BaseSystemComponent // Embedding BaseComponent
}

// Type returns the type of the component.
func (bo *BaseSystemOperation) Type() components.ComponentType {
	return components.BasicComponentType
}

func NewBaseSystemOperation(id, name, description string) *BaseSystemOperation {
	return &BaseSystemOperation{
		BaseSystemComponent: BaseSystemComponent{
			BaseComponent: components.BaseComponent{
				Id:   id,
				Nm:   name,
				Desc: description,
			},
		},
	}
}

// Execute performs the operation with the given context and input parameters,
// and returns any output or error encountered.
func (bo *BaseSystemOperation) Execute(ctx *context.Context, input *OperationInput) (*OperationOutput, error) {
	// Perform operation logic here
	// For demonstration purposes, just return an error
	return nil, errors.New("operation not implemented")
}
