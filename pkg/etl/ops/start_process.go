package etl_ops

import (
	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/edward1christian/block-forge/pkg/etl"
)

// StartETLOperation represents an operation to start an ETL process.
type StartETLOperation struct{}

// NewStartETLOperation creates a new instance of StartETLOperation.
func NewStartETLOperation() *StartETLOperation {
	return &StartETLOperation{}
}

// Execute implements the Execute method of the Operation interface.
func (o *StartETLOperation) Execute(ctx *context.Context, input *system.OperationInput) (*system.OperationOutput, error) {
	process, ok := input.Data.(*etl.ETLProcess)
	if !ok {
		return nil, etl.ErrInvalidProcess
	}

	// Start components
	for _, component := range process.Components {
		if err := component.Start(ctx); err != nil {
			// Rollback: stop already started components
			for _, startedComponent := range process.Components {
				_ = startedComponent.Stop(ctx)
			}
			return nil, err
		}
	}

	// Update process status
	process.Status = etl.ETLStatusRunning
	return &system.OperationOutput{
		Data: process,
	}, nil
}
