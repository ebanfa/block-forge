package operations

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/component"
	configApi "github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// InitDirectoriesOperationFactory is responsible for creating instances of InitDirectoriesOperation.
type InitDirectoriesOperationFactory struct {
}

// NewInitDirectoriesOperationFactory is a constructor function that creates and returns a new instance of InitDirectoriesOperationFactory.
func NewInitDirectoriesOperationFactory() *InitDirectoriesOperationFactory {
	return &InitDirectoriesOperationFactory{}
}

// CreateComponent creates a new instance of the InitDirectoriesOperation.
func (f *InitDirectoriesOperationFactory) CreateComponent(config *configApi.ComponentConfig) (component.ComponentInterface, error) {
	return NewInitDirectoriesOperation(config.ID, config.Name, config.Description), nil
}

// InitDirectoriesOperation represents an operation to initialize required application directories.
type InitDirectoriesOperation struct {
	system.BaseSystemComponent // Embedding BaseComponent
}

// Type returns the type of the component.
func (op *InitDirectoriesOperation) Type() component.ComponentType {
	return component.BasicComponentType
}

// NewInitDirectoriesOperation creates a new instance of InitDirectoriesOperation.
func NewInitDirectoriesOperation(id, name, description string) *InitDirectoriesOperation {
	return &InitDirectoriesOperation{
		BaseSystemComponent: system.BaseSystemComponent{
			BaseComponent: component.BaseComponent{
				Id:   id,
				Nm:   name,
				Desc: description,
			},
		},
	}
}

// Execute performs the operation to initialize directories.
func (op *InitDirectoriesOperation) Execute(ctx *context.Context, input *system.SystemOperationInput) (*system.SystemOperationOutput, error) {
	// Extract the home directory from the input
	homeDir, ok := input.Data.(string)
	if !ok {
		return nil, errors.New("invalid home directory provided in input data")
	}

	// Define the .nova directory path
	novaDir := filepath.Join(homeDir, ".nova")

	// Create the .nova directory if it does not exist
	if err := createDirectory(novaDir); err != nil {
		return nil, fmt.Errorf("failed to create .nova directory: %v", err)
	}

	// Create subdirectories within the .nova directory
	subdirectories := []string{"databases", "configs", "logs", "cache"}
	for _, subdir := range subdirectories {
		dirPath := filepath.Join(novaDir, subdir)
		if err := createDirectory(dirPath); err != nil {
			return nil, fmt.Errorf("failed to create directory %s: %v", dirPath, err)
		}
	}

	return nil, nil
}

// createDirectory creates a directory if it does not already exist
func createDirectory(path string) error {
	// Check if the directory already exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Create the directory with read-write permissions for the current user
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
