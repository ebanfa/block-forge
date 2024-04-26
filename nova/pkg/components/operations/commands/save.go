package system

import (
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/component"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// BaseComponent represents a concrete implementation of the SystemOperationInterface.
type SaveOperation struct {
	system.BaseSystemOperation // Embedding BaseComponent
}

// Type returns the type of the component.
func (bo *SaveOperation) Type() component.ComponentType {
	return component.BasicComponentType
}

func NewSaveOperation(id, name, description string) *SaveOperation {
	return &SaveOperation{
		BaseSystemOperation: system.BaseSystemOperation{
			BaseSystemComponent: system.BaseSystemComponent{
				BaseComponent: component.BaseComponent{
					Id:   id,
					Nm:   name,
					Desc: description,
				},
			},
		},
	}
}

// Execute performs the operation with the given context and input parameters,
// and returns any output or error encountered.
func (bo *SaveOperation) Execute(ctx *context.Context,
	input *system.SystemOperationInput) (*system.SystemOperationOutput, error) {
	// Perform operation logic here
	// For demonstration purposes, just return an error
	return nil, nil
}
