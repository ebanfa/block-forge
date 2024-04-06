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

// TestStartProcessOperation_Execute_Success verifies that StartProcessOperation.Execute() runs successfully.
func TestStartProcessOperation_Execute_Success(t *testing.T) {
	// Mocks
	ctx := &context.Context{}
	mockEtlComponent := &etlMocksApi.MockETLProcessComponent{}
	mockEtlComponent.On("Start", ctx).Return(nil)

	// Create StartProcessOperation instance
	startOp := operations.NewStartProcessOperation("1", "TestOperation", "Test Description")

	// Setup mock function calls
	process := &processApi.ETLProcess{
		ID: "processID",
		Components: map[string]etlComponentsApi.ETLProcessComponent{
			"AdapterID": mockEtlComponent,
		},
		Status: processApi.ETLProcessStatusInitialized,
	}

	// Call Execute method
	output, err := startOp.Execute(ctx, &system.OperationInput{Data: process})

	// Check if the Execute method returns no error
	assert.NoError(t, err)
	assert.NotNil(t, output)

	// Assert that the process status is updated to running
	assert.Equal(t, processApi.ETLProcessStatusRunning, process.Status)
}

// TestStartProcessOperation_Execute_Error_NotProcess verifies that StartProcessOperation.Execute() returns an error when input is not an ETLProcess.
func TestStartProcessOperation_Execute_Error_NotProcess(t *testing.T) {
	// Mocks
	ctx := &context.Context{}

	// Create StartProcessOperation instance
	startOp := operations.NewStartProcessOperation("1", "TestOperation", "Test Description")

	// Call Execute method with non-process input
	output, err := startOp.Execute(ctx, &system.OperationInput{Data: "not a process"})

	// Check if the Execute method returns an error
	assert.Error(t, err)
	assert.Nil(t, output)
}

// TestStartProcessOperation_Execute_Error_ComponentStart verifies that StartProcessOperation.Execute() returns an error when a component fails to start.
func TestStartProcessOperation_Execute_Error_ComponentStart(t *testing.T) {
	// Mocks
	ctx := &context.Context{}
	mockEtlComponent := &etlMocksApi.MockETLProcessComponent{}
	mockEtlComponent.On("Start", ctx).Return(errors.New("Component start error"))

	// Create StartProcessOperation instance
	startOp := operations.NewStartProcessOperation("1", "TestOperation", "Test Description")

	// Setup mock function calls
	process := &processApi.ETLProcess{
		ID: "processID",
		Components: map[string]etlComponentsApi.ETLProcessComponent{
			"AdapterID": mockEtlComponent,
		},
		Status: processApi.ETLProcessStatusInitialized,
	}

	// Call Execute method
	output, err := startOp.Execute(ctx, &system.OperationInput{Data: process})

	// Check if the Execute method returns an error
	assert.Error(t, err)
	assert.Nil(t, output)
}
