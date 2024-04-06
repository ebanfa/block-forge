package operations

import (
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/components"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/edward1christian/block-forge/pkg/etl"
	processApi "github.com/edward1christian/block-forge/pkg/etl/process"
)

// StopProcessOperation represents an operation to stop an ETL process.
type StopProcessOperation struct {
	system.BaseSystemOperation // Embedding BaseComponent
}

// NewStopProcessOperation creates a new StopProcessOperation instance.
func NewStopProcessOperation(id, name, description string) *StopProcessOperation {
	return &StopProcessOperation{
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

// Execute performs the operation to stop the ETL process.
func (so *StopProcessOperation) Execute(ctx *context.Context, input *system.OperationInput) (*system.OperationOutput, error) {
	process, ok := input.Data.(*processApi.ETLProcess)
	if !ok {
		return nil, etl.ErrNotProcess
	}

	// Stop each component of the ETL process
	for _, component := range process.Components {
		// Check if the component is stoppable
		stoppable, ok := component.(components.StartableInterface)
		if !ok {
			continue
		}
		// Stop the component
		if err := stoppable.Stop(ctx); err != nil {
			return nil, err
		}
	}

	// Update the process status
	process.Status = processApi.ETLProcessStatusStopped

	return &system.OperationOutput{Data: process}, nil
}
