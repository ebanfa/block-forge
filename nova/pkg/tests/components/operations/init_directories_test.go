package operations

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/edward1christian/block-forge/nova/pkg/components/operations"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/stretchr/testify/assert"
)

// Mock function for creating directories
var createDirectory func(string) error

func TestInitDirectoriesOperation_Execute_Success(t *testing.T) {
	// Arrange
	op := operations.NewInitDirectoriesOperation("testID", "testName", "testDescription")
	ctx := &context.Context{}
	// User home directory
	homeDir := os.TempDir()

	// Create SystemOperationInput with project ID and home directory
	inputData := &system.SystemOperationInput{
		Data: homeDir,
	}
	// Act
	output, err := op.Execute(ctx, inputData)

	// Assert
	assert.NoError(t, err, "Initialization should succeed without errors")
	assert.Nil(t, output, "Output should be nil")

	novaDir := filepath.Join(homeDir, ".nova")
	subdirectories := []string{"databases", "configs", "logs", "cache"}
	for _, subdir := range subdirectories {
		dirPath := filepath.Join(novaDir, subdir)
		_, err := os.Stat(dirPath)
		assert.NoError(t, err, fmt.Sprintf("Directory %s should be created", dirPath))
	}
}

func TestInitDirectoriesOperation_Execute_Failure(t *testing.T) {
	// Arrange
	op := operations.NewInitDirectoriesOperation("testID", "testName", "testDescription")
	ctx := &context.Context{}
	input := &system.SystemOperationInput{}

	// Mock the createDirectory function to simulate failure
	createDirectory = func(path string) error {
		fmt.Println("Mock createDirectory function called with path:", path)
		return fmt.Errorf("failed to create directory: %s", path)
	}
	defer func() { createDirectory = osMkdirAllWrapper }()

	// Act
	output, err := op.Execute(ctx, input)

	// Assert
	assert.Error(t, err, "Initialization should fail with an error")
	assert.Nil(t, output, "Output should be nil")
}

// osMkdirAllWrapper is a wrapper around os.MkdirAll with default permissions
func osMkdirAllWrapper(path string) error {
	fmt.Println("osMkdirAllWrapper called with path:", path)
	return os.MkdirAll(path, 0755) // Using default permissions 0755
}
