package etl_ops

import (
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/edward1christian/block-forge/pkg/etl"
)

// StopETLOperation represents an operation to stop an ETL process.
type StopETLOperation struct{}

// NewStopETLOperation creates a new instance of StopETLOperation.
func NewStopETLOperation() *StopETLOperation {
	return &StopETLOperation{}
}

// Execute implements the Execute method of the Operation interface.
func (o *StopETLOperation) Execute(ctx *context.Context, input *system.OperationInput) (*system.OperationOutput, error) {
	process, ok := input.Data.(*etl.ETLProcess)
	if !ok {
		return nil, etl.ErrInvalidProcess
	}

	// Stop components in reverse order
	for i := len(process.Components) - 1; i >= 0; i-- {
		componentName := process.Config.Components[i].Name
		if component, ok := process.Components[componentName]; ok {
			if err := component.Stop(ctx); err != nil {
				return nil, err
			}
		}
	}

	// Update process status
	process.Status = etl.ETLStatusStopped
	return &system.OperationOutput{
		Data: process,
	}, nil
}
