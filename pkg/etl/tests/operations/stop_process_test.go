package operations_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
	etlComponentsApi "github.com/edward1christian/block-forge/pkg/etl/components"
	"github.com/edward1christian/block-forge/pkg/etl/components/operations"
	etlMocksApi "github.com/edward1christian/block-forge/pkg/etl/mocks"
	processApi "github.com/edward1christian/block-forge/pkg/etl/process"
)

// TestStopProcessOperation_Execute_Success verifies that StopProcessOperation.Execute() runs successfully.
func TestStopProcessOperation_Execute_Success(t *testing.T) {
	// Mocks
	ctx := &context.Context{}
	mockEtlComponent := &etlMocksApi.MockETLProcessComponent{}
	mockEtlComponent.On("Stop", ctx).Return(nil)

	// Create StopProcessOperation instance
	stopOp := operations.NewStopProcessOperation("1", "TestOperation", "Test Description")

	// Setup mock function calls
	process := &processApi.ETLProcess{
		ID: "processID",
		Components: map[string]etlComponentsApi.ETLProcessComponent{
			"AdapterID": mockEtlComponent,
		},
		Status: processApi.ETLProcessStatusRunning,
	}

	// Call Execute method
	output, err := stopOp.Execute(ctx, &system.OperationInput{Data: process})

	// Check if the Execute method returns no error
	assert.NoError(t, err)
	assert.NotNil(t, output)

	// Assert that the process status is updated to stopped
	assert.Equal(t, processApi.ETLProcessStatusStopped, process.Status)
}

// TestStopProcessOperation_Execute_Error_NotProcess verifies that StopProcessOperation.Execute() returns an error when input is not an ETLProcess.
func TestStopProcessOperation_Execute_Error_NotProcess(t *testing.T) {
	// Mocks
	ctx := &context.Context{}

	// Create StopProcessOperation instance
	stopOp := operations.NewStopProcessOperation("1", "TestOperation", "Test Description")

	// Call Execute method with non-process input
	output, err := stopOp.Execute(ctx, &system.OperationInput{Data: "not a process"})

	// Check if the Execute method returns an error
	assert.Error(t, err)
	assert.Nil(t, output)
}

// TestStopProcessOperation_Execute_Error_ComponentStop verifies that StopProcessOperation.Execute() returns an error when a component fails to stop.
func TestStopProcessOperation_Execute_Error_ComponentStop(t *testing.T) {
	// Mocks
	ctx := &context.Context{}
	mockEtlComponent := &etlMocksApi.MockETLProcessComponent{}
	mockEtlComponent.On("Stop", ctx).Return(errors.New("Component stop error"))

	// Create StopProcessOperation instance
	stopOp := operations.NewStopProcessOperation("1", "TestOperation", "Test Description")

	// Setup mock function calls
	process := &processApi.ETLProcess{
		ID: "processID",
		Components: map[string]etlComponentsApi.ETLProcessComponent{
			"AdapterID": mockEtlComponent,
		},
		Status: processApi.ETLProcessStatusRunning,
	}

	// Call Execute method
	output, err := stopOp.Execute(ctx, &system.OperationInput{Data: process})

	// Check if the Execute method returns an error
	assert.Error(t, err)
	assert.Nil(t, output)
}
