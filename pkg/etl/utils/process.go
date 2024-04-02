package utils

import (
	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// ExecuteOperationWithConfig executes an operation with the provided configuration and returns the result.
func ExecuteSystemOp(
	ctx *context.Context,
	sys system.System,
	operationID string,
	data interface{}) (*system.OperationOutput, error) {

	opInput := &system.OperationInput{
		Data: data,
	}

	opOutput, err := sys.ExecuteOperation(ctx, operationID, opInput)
	if err != nil {
		return nil, err
	}

	return opOutput, nil
}
