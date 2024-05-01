package commands

import (
	"testing"

	"github.com/edward1christian/block-forge/nova/pkg/components/operations/commands"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/component"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/stretchr/testify/assert"
)

// TestNewCreateConfigurationOp tests the NewCreateConfigurationOp function.
func TestNewCreateConfigurationOp(t *testing.T) {
	// Arrange
	id := "test-id"
	name := "Test Config Op"
	description := "Test description"

	// Act
	op := commands.NewCreateConfigurationOp(id, name, description)

	// Assert
	assert.NotNil(t, op)
	assert.Equal(t, id, op.Id)
	assert.Equal(t, name, op.Nm)
	assert.Equal(t, description, op.Desc)
	assert.Equal(t, component.BasicComponentType, op.Type())
}

// TestCreateConfigurationOp_Execute_Success tests the successful execution of the CreateConfigurationOp.
func TestCreateConfigurationOp_Execute_Success(t *testing.T) {
	// Arrange
	op := commands.NewCreateConfigurationOp("test-id", "Test Config Op", "Test description")
	ctx := context.Background()
	input := &system.SystemOperationInput{
		Data: map[string]interface{}{
			"projectID":   "test-project-id",
			"projectName": "Test Project",
		},
	}

	// Act
	output, err := op.Execute(ctx, input)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, output)
}

// TestCreateConfigurationOp_Execute_InvalidInputData tests the case when the input data format is invalid.
func TestCreateConfigurationOp_Execute_InvalidInputData(t *testing.T) {
	// Arrange
	op := commands.NewCreateConfigurationOp("test-id", "Test Config Op", "Test description")
	ctx := context.Background()
	input := &system.SystemOperationInput{
		Data: "invalid-data",
	}

	// Act
	output, err := op.Execute(ctx, input)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, output)
	assert.EqualError(t, err, "invalid input data format")
}

// TestCreateConfigurationOp_Execute_InvalidProjectID tests the case when the project ID is not a string.
func TestCreateConfigurationOp_Execute_InvalidProjectID(t *testing.T) {
	// Arrange
	op := commands.NewCreateConfigurationOp("test-id", "Test Config Op", "Test description")
	ctx := context.Background()
	input := &system.SystemOperationInput{
		Data: map[string]interface{}{
			"projectID":   123,
			"projectName": "Test Project",
		},
	}

	// Act
	output, err := op.Execute(ctx, input)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, output)
	assert.EqualError(t, err, "project ID must be a string")
}

// TestCreateConfigurationOp_Execute_InvalidProjectName tests the case when the project name is not a string.
func TestCreateConfigurationOp_Execute_InvalidProjectName(t *testing.T) {
	// Arrange
	op := commands.NewCreateConfigurationOp("test-id", "Test Config Op", "Test description")
	ctx := context.Background()
	input := &system.SystemOperationInput{
		Data: map[string]interface{}{
			"projectID":   "test-project-id",
			"projectName": 123,
		},
	}

	// Act
	output, err := op.Execute(ctx, input)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, output)
	assert.EqualError(t, err, "project name must be a string")
}
