package commands

import (
	"errors"
	"sync"
	"testing"

	"github.com/edward1christian/block-forge/nova/pkg/components/operations/commands"
	"github.com/edward1christian/block-forge/nova/pkg/mocks"
	"github.com/edward1christian/block-forge/nova/pkg/store"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"

	"github.com/stretchr/testify/assert"
)

// Define a mutex to protect access to the mockMetadataDB
var mutex sync.Mutex

// TestListConfigurationsOp_Execute_Success tests the Execute method of ListConfigurationsOp for success.
func TestListConfigurationsOp_Execute_Success(t *testing.T) {
	// Arrange
	mutex.Lock() // Lock before accessing shared resource
	mockMetadataDB := &mocks.MockMetadataStore{}

	mockEntries := []*store.MetadataEntry{
		{ProjectID: "project1"},
		{ProjectID: "project2"},
	}

	// Mock behavior
	mockMetadataDB.On("GetAllMetadata").Return(mockEntries, nil)
	mutex.Unlock() // Unlock after finishing accessing shared resource

	op := commands.NewListConfigurationsOp("id", "name", "description")

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
	mutex.Lock() // Lock before accessing shared resource
	mockMetadataDB := &mocks.MockMetadataStore{}
	expectedErr := errors.New("database error")

	// Mock behavior
	mockMetadataDB.On("GetAllMetadata").Return([]*store.MetadataEntry{}, expectedErr)
	mutex.Unlock() // Unlock after finishing accessing shared resource

	op := commands.NewListConfigurationsOp("id", "name", "description")

	// Act
	output, err := op.Execute(&context.Context{}, &system.SystemOperationInput{})

	// Assert
	assert.Error(t, err, "Execute should return an error")
	assert.Nil(t, output, "Output should be nil")
	assert.EqualError(t, err, expectedErr.Error(), "Error message should match")
}
