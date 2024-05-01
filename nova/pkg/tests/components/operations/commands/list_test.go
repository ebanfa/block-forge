package commands

import (
	"errors"
	"testing"

	"github.com/edward1christian/block-forge/nova/pkg/components/operations/commands"
	"github.com/edward1christian/block-forge/nova/pkg/database"
	"github.com/edward1christian/block-forge/nova/pkg/mocks"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"

	"github.com/stretchr/testify/assert"
)

// TestListConfigurationsOp_Execute_Success tests the Execute method of ListConfigurationsOp for success.
func TestListConfigurationsOp_Execute_Success(t *testing.T) {
	// Arrange
	mockMetadataDB := &mocks.MockMetadataDatabase{}

	mockEntries := []*database.MetadataEntry{
		{ProjectID: "project1"},
		{ProjectID: "project2"},
	}

	// Mock behavior
	mockMetadataDB.On("GetAll").Return(mockEntries, nil)

	op := commands.NewListConfigurationsOp("id", "name", "description", mockMetadataDB)

	// Act
	output, err := op.Execute(&context.Context{}, &system.SystemOperationInput{})

	// Assert
	assert.NoError(t, err, "Execute should not return an error")
	assert.NotNil(t, output, "Output should not be nil")
	assert.Equal(t, len(mockEntries), len(output.Data.([]string)), "Number of entries should match")
}

// TestListConfigurationsOp_Execute_Error tests the Execute method of ListConfigurationsOp for error.
func TestListConfigurationsOp_Execute_Error(t *testing.T) {
	// Arrange
	mockMetadataDB := &mocks.MockMetadataDatabase{}
	expectedErr := errors.New("database error")

	// Mock behavior
	mockMetadataDB.On("GetAll").Return([]*database.MetadataEntry{}, expectedErr)

	op := commands.NewListConfigurationsOp("id", "name", "description", mockMetadataDB)

	// Act
	output, err := op.Execute(&context.Context{}, &system.SystemOperationInput{})

	// Assert
	assert.Error(t, err, "Execute should return an error")
	assert.Nil(t, output, "Output should be nil")
	assert.EqualError(t, err, expectedErr.Error(), "Error message should match")
}
