package system

import (
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/component"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// BaseComponent represents a concrete implementation of the SystemOperationInterface.
type RunOperation struct {
	system.BaseSystemOperation // Embedding BaseComponent
}

// Type returns the type of the component.
func (ro *RunOperation) Type() component.ComponentType {
	return component.BasicComponentType
}

func NewRunOperation(id, name, description string) *RunOperation {
	return &RunOperation{
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
func (ro *RunOperation) Execute(ctx *context.Context,
	input *system.SystemOperationInput) (*system.SystemOperationOutput, error) {
	// Perform operation logic here
	// For demonstration purposes, just return an error
	return nil, nil
}
