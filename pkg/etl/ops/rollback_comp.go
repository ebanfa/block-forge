package etl_ops

import (
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/edward1christian/block-forge/pkg/etl"
)

// RollbackOperation represents an operation to rollback components of an ETL process.
type RollbackOperation struct{}

// NewRollbackOperation creates a new instance of RollbackOperation.
func NewRollbackOperation() *RollbackOperation {
	return &RollbackOperation{}
}

// Execute implements the Execute method of the Operation interface.
func (o *RollbackOperation) Execute(ctx *context.Context, input *system.OperationInput) (*system.OperationOutput, error) {
	process, ok := input.Data.(*etl.ETLProcess)
	if !ok {
		return nil, etl.ErrInvalidProcess
	}

	// Rollback components
	for _, component := range process.Components {
		_ = component.Stop(ctx)
	}

	return nil, nil
}
