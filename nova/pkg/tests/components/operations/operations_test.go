package operations

import (
	"testing"

	"github.com/edward1christian/block-forge/nova/pkg/components/operations"
	"github.com/stretchr/testify/assert"
)

// TestGetOperationsToRegister_Success tests the GetOperationsToRegister function for success.
// It ensures that the function returns the list of operations to register without any errors.
func TestGetOperationsToRegister_Success(t *testing.T) {
	// Act
	operations := operations.GetOperationsToRegister()

	// Assert
	assert.NotEmpty(t, operations, "List of operations should not be empty")
}

// TestGetOperationsToRegister_FactoryIDs tests the GetOperationsToRegister function for correct factory IDs.
// It ensures that the factory IDs in the list of operations to register match the expected format.
func TestGetOperationsToRegister_FactoryIDs(t *testing.T) {
	// Act
	operations := operations.GetOperationsToRegister()

	// Assert
	for _, op := range operations {
		expectedFactoryID := op.ID + "Factory"
		assert.Equal(t, expectedFactoryID, op.FactoryID, "Factory ID should match the expected format")
	}
}
