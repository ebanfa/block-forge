package operations

import (
	"errors"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/components"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// BaseComponent represents a concrete implementation of the OperationInterface.
type DemoOperation struct {
	system.BaseSystemOperation // Embedding BaseComponent
}

func NewDemoOperation(id, name, description string) *DemoOperation {
	return &DemoOperation{
		BaseSystemOperation: system.BaseSystemOperation{
			BaseSystemComponent: system.BaseSystemComponent{
				BaseComponent: components.BaseComponent{
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
func (bo *DemoOperation) Execute(ctx *context.Context, input *system.OperationInput) (*system.OperationOutput, error) {
	// Perform operation logic here
	// For demonstration purposes, just return an error
	return nil, errors.New("operation not implemented")
}
