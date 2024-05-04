package commands

import (
	"errors"

	"github.com/edward1christian/block-forge/nova/pkg/store"
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
func (bo *CreateConfigurationOp) Execute(ctx *context.Context, input *system.SystemOperationInput) (*system.SystemOperationOutput, error) {
	// Extract project information from input data
	projectID, _, err := extractProjectInfo(input)
	if err != nil {
		return nil, err
	}

	// Get the project-specific database path
	projectDbPath, err := store.GetDefaultDatabasePath(projectID)
	if err != nil {
		return nil, err
	}

	// Insert metadata entry into MetadataDatabase
	err = insertMetadataEntry(projectID, projectDbPath)
	if err != nil {
		return nil, err
	}

	// Return success response
	return &system.SystemOperationOutput{}, nil
}

// extractProjectInfo extracts project ID and project name from the input data.
func extractProjectInfo(input *system.SystemOperationInput) (string, string, error) {
	// Check if input data is in the expected format
	data, ok := input.Data.(map[string]interface{})
	if !ok {
		return "", "", errors.New("invalid input data format")
	}

	// Extract project ID from input data
	projectID, ok := data["projectID"].(string)
	if !ok {
		return "", "", errors.New("project ID must be a string")
	}

	// Extract project name from input data
	projectName, ok := data["projectName"].(string)
	if !ok {
		return "", "", errors.New("project name must be a string")
	}

	return projectID, projectName, nil
}

// insertMetadataEntry inserts a metadata entry into the MetadataDatabase for the specified project.
func insertMetadataEntry(projectID, projectDbPath string) error {
	// Get the default database path for the metadata database
	dbPath, err := store.GetDefaultDatabasePath(store.MetadataDatabaseID)
	if err != nil {
		return err
	}

	// Get an instance of the MetadataDatabase
	metaDB, err := store.GetMetadataStoreInstance(store.MetadataDatabaseID, dbPath)
	if err != nil {
		return err
	}

	// Create a new metadata entry
	entry := &store.MetadataEntry{
		ProjectID:    projectID,
		DatabaseName: "default",
		DatabasePath: projectDbPath,
	}

	// Insert the metadata entry into the database
	return metaDB.InsertMetadata(entry)
}
