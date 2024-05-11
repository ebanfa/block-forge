package commands

import (
	"errors"
	"fmt"

	novaConfigApi "github.com/edward1christian/block-forge/nova/pkg/config"
	"github.com/edward1christian/block-forge/nova/pkg/store"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/component"
	configApi "github.com/edward1christian/block-forge/pkg/application/config"
	storeApi "github.com/edward1christian/block-forge/pkg/application/store"
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

	// Check if input data is in the expected format
	projectName, ok := input.Data.(string)
	if !ok {
		return nil, errors.New("failed to create project. Invalid input data")
	}

	multiStore := bo.System.MultiStore()
	configuration := bo.System.Configuration()

	// Validate the configuration
	novaConfig, ok := configuration.CustomConfig.(novaConfigApi.NovaConfig)
	if !ok {
		return nil, errors.New("failed to create project. Invalid configuration")
	}

	_, err := bo.CreateProjectMetadataEntry(projectName, novaConfig, multiStore)
	if err != nil {
		return nil, fmt.Errorf("failed to create project metadata entry for project %s. %w", projectName, err)
	}

	// Load the latest version of the MultiStore database from disk
	_, err = multiStore.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load multistore for project %s. %w", projectName, err)
	}

	// Save the latest version of the MultiStore database to disk
	_, _, err = multiStore.SaveVersion()
	if err != nil {
		return nil, fmt.Errorf("failed to save multistore version for project %s. %w", projectName, err)
	}

	// Return success response
	return &system.SystemOperationOutput{}, nil
}

func (bo *CreateConfigurationOp) CreateProjectMetadataEntry(
	projectName string, configuration novaConfigApi.NovaConfig, multiStore storeApi.MultiStore) (*store.MetadataEntry, error) {

	// Creates a store or retrieves it already exists
	metadataStore, err := bo.CreateMetadataStore(configuration.MetadataDbName, multiStore)
	if err != nil {
		return nil, err
	}

	// Create the Metadata entry
	metadata := bo.createProjectMetadata(projectName, configuration.DatabasesDir)

	// Store the record
	err = metadataStore.InsertMetadata(metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to store metadata for project %s. %w", projectName, err)
	}

	// Save the latest version of the Metadata database to disk
	_, _, err = metadataStore.SaveVersion()
	if err != nil {
		return nil, fmt.Errorf("failed to save metadata store version for project %s. %w", projectName, err)
	}

	return metadata, nil
}

func (bo *CreateConfigurationOp) CreateMetadataStore(storeName string, multiStore storeApi.MultiStore) (*store.MetadataStoreImpl, error) {
	// Creates or retrieve an existing store
	metadataStore, _, err := multiStore.CreateStore(storeName)
	if err != nil {
		return nil, err
	}

	// Load the latest version of the metadata database from disk
	_, err = metadataStore.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load metadata store: %s. %w", storeName, err)
	}

	return store.NewMetadataStore(metadataStore), nil

}

func (bo *CreateConfigurationOp) createProjectMetadata(projectName, databasesDir string) *store.MetadataEntry {

	// Generate storage id
	projectID, projectDbPath := storeApi.GenererateStorageInfo(
		projectName, databasesDir,
	)

	return &store.MetadataEntry{
		ProjectID:    projectID,
		ProjectName:  projectName,
		DatabaseName: projectID,
		DatabasePath: projectDbPath,
	}

}
