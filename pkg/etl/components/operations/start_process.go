package operations

import (
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/components"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/edward1christian/block-forge/pkg/etl"
	processApi "github.com/edward1christian/block-forge/pkg/etl/process"
)

// BaseComponent represents a concrete implementation of the OperationInterface.
type StartProcessOperation struct {
	system.BaseSystemOperation // Embedding BaseComponent
}

func NewStartProcessOperation(id, name, description string) *StartProcessOperation {
	return &StartProcessOperation{
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
func (bo *StartProcessOperation) Execute(ctx *context.Context, input *system.OperationInput) (*system.OperationOutput, error) {
	process, ok := input.Data.(*processApi.ETLProcess)
	if !ok {
		return nil, etl.ErrNotProcess
	}
	// Perform operation logic here
	// For demonstration purposes, just return an error
	// Start each component of the ETL process
	for _, component := range process.Components {
		// Check if the component is startable
		startable, ok := component.(components.StartableInterface)
		if !ok {
			continue
		}
		// Start the component
		if err := startable.Start(ctx); err != nil {
			return nil, err
		}
	}

	// Update the process status
	process.Status = processApi.ETLProcessStatusRunning

	return &system.OperationOutput{Data: process}, nil
}
