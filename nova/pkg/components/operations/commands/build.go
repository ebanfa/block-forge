package system

import (
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/component"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// BaseComponent represents a concrete implementation of the SystemOperationInterface.
type BuildOperation struct {
	system.BaseSystemOperation // Embedding BaseComponent
}

// Type returns the type of the component.
func (bo *BuildOperation) Type() component.ComponentType {
	return component.BasicComponentType
}

func NewBuildOperation(id, name, description string) *BuildOperation {
	return &BuildOperation{
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
func (bo *BuildOperation) Execute(ctx *context.Context,
	input *system.SystemOperationInput) (*system.SystemOperationOutput, error) {
	// Perform operation logic here
	// For demonstration purposes, just return an error
	return nil, nil
}
