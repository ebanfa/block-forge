package commands

import (
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/component"
	configApi "github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// BuildServiceFactory is responsible for creating instances of BuildService.
type RunProjectOpFactory struct {
}

// CreateComponent creates a new instance of the BuildService.
func (bf *RunProjectOpFactory) CreateComponent(config *configApi.ComponentConfig) (component.ComponentInterface, error) {
	// Construct the service
	return NewRunProjectOp(config.ID, config.Name, config.Description), nil
}

// BaseComponent represents a concrete implementation of the SystemOperationInterface.
type RunProjectOp struct {
	system.BaseSystemOperation // Embedding BaseComponent
}

// Type returns the type of the component.
func (ro *RunProjectOp) Type() component.ComponentType {
	return component.BasicComponentType
}

func NewRunProjectOp(id, name, description string) *RunProjectOp {
	return &RunProjectOp{
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
func (ro *RunProjectOp) Execute(ctx *context.Context,
	input *system.SystemOperationInput) (*system.SystemOperationOutput, error) {
	// Perform operation logic here
	// For demonstration purposes, just return an error
	return nil, nil
}
