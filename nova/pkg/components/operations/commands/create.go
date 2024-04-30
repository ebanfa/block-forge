package commands

import (
	"errors"
	"fmt"
	"path/filepath"
	"time"

	"github.com/edward1christian/block-forge/nova/pkg/database"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/component"
	configApi "github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// BuildServiceFactory is responsible for creating instances of BuildService.
type CreateConfigurationOpFactory struct {
}

// CreateComponent creates a new instance of the BuildService.
func (bf *CreateConfigurationOpFactory) CreateComponent(config *configApi.ComponentConfig) (component.ComponentInterface, error) {
	// Construct the service
	return NewCreateConfigurationOp(config.ID, config.Name, config.Description), nil
}

// BaseComponent represents a concrete implementation of the SystemOperationInterface.
type CreateConfigurationOp struct {
	system.BaseSystemOperation // Embedding BaseComponent
}

// Type returns the type of the component.
func (bo *CreateConfigurationOp) Type() component.ComponentType {
	return component.BasicComponentType
}

func NewCreateConfigurationOp(id, name, description string) *CreateConfigurationOp {
	return &CreateConfigurationOp{
		BaseSystemOperation: system.BaseSystemOperation{
			BaseSystemComponent: system.BaseSystemComponent{
				BaseComponent: component.BaseComponent{
					Id:   id,
					Nm:   name,
					Desc: description,
				},
			},
		},
	}
}

// Execute performs the operation with the given context and input parameters,
// and returns any output or error encountered.
func (bo *CreateConfigurationOp) Execute(ctx *context.Context,
	input *system.SystemOperationInput) (*system.SystemOperationOutput, error) {
	// Extract project ID and home directory from input.Data
	data, ok := input.Data.(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid input data format")
	}

	projectID, ok := data["projectID"].(string)
	if !ok {
		return nil, errors.New("invalid input data format: project ID must be a string")
	}

	homeDir, ok := data["homeDir"].(string)
	if !ok {
		return nil, errors.New("invalid input data format: home directory must be a string")
	}

	// Define the database path within the .nova directory
	dbPath := filepath.Join(homeDir, ".nova", "databases", projectID+".db")

	// Ensure a single instance of MetadataDatabase is used throughout the application
	metaDB, err := database.GetMetadataDBInstance(projectID, dbPath)
	if err != nil {
		return nil, err
	}

	// Create a new metadata entry
	entry := &database.MetadataEntry{
		ProjectID:    projectID,
		DatabaseName: "default", // Set default database name
		DatabasePath: dbPath,    // Set database path based on home directory and project ID
		CreationDate: time.Now(),
		LastUpdated:  time.Now(),
	}

	// Insert the metadata entry into the database
	err = metaDB.Insert(entry)
	if err != nil {
		return nil, fmt.Errorf("error inserting metadata entry: %v", err)
	}

	// Return success response
	return &system.SystemOperationOutput{
		Data: map[string]interface{}{"message": "Configuration created successfully"},
	}, nil
}
