package etl_ops

import (
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/edward1christian/block-forge/pkg/etl"
	"github.com/edward1christian/block-forge/pkg/etl/utils"
)

// CreateETLOperation represents an operation to create an ETL process.
type CreateETLOperation struct {
	idGenerator etl.ProcessIDGenerator
}

// NewCreateETLOperation creates a new instance of CreateETLOperation.
func NewCreateETLOperation(idGenerator etl.ProcessIDGenerator) *CreateETLOperation {
	return &CreateETLOperation{
		idGenerator: idGenerator,
	}
}

// Execute implements the Execute method of the Operation interface.
func (o *CreateETLOperation) Execute(ctx *context.Context, input *system.OperationInput) (*system.OperationOutput, error) {
	// Type assertion to extract ETLConfig from OperationInput
	config, ok := input.Data.(*etl.ETLConfig)

	if !ok || utils.IsEmptyConfig(config) {
		return nil, etl.ErrInvalidETLProcessConfig
	}

	// Generate the process ID using the injected generator
	processID, err := o.idGenerator.GenerateID()
	if err != nil {
		return nil, err
	}

	// Create the process
	process := &etl.ETLProcess{
		ID:         processID,
		Config:     config,
		Status:     etl.ETLStatusInitialized,
		Components: make(map[string]etl.ETLComponent),
	}

	// Return the created process
	return &system.OperationOutput{
		Data: process,
	}, nil
}
