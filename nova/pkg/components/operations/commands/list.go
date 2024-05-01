package commands

import (
	"github.com/edward1christian/block-forge/nova/pkg/database"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/component"
	configApi "github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// BuildServiceFactory is responsible for creating instances of BuildService.
type ListConfigurationsOpFactory struct {
}

// CreateComponent creates a new instance of ListConfigurationsOp.
func (bf *ListConfigurationsOpFactory) CreateComponent(config *configApi.ComponentConfig) (component.ComponentInterface, error) {
	// Get the MetadataDatabase DB path
	metaDBPath, err := database.GetDefaultDatabasePath(database.MetadataDatabaseID)
	if err != nil {
		return nil, err
	}

	// Get the MetadataDatabase instance
	metaDB, err := database.GetMetadataDBInstance(config.Name, metaDBPath)
	if err != nil {
		return nil, err
	}

	// Construct the service with the injected database
	return NewListConfigurationsOp(config.ID, config.Name, config.Description, metaDB), nil
}

// ListConfigurationsOp represents a concrete implementation of the SystemOperationInterface.
type ListConfigurationsOp struct {
	system.BaseSystemOperation
	metadataDB database.MetadataDatabaseInterface
}

// Type returns the type of the component.
func (bo *ListConfigurationsOp) Type() component.ComponentType {
	return component.OperationType
}

func NewListConfigurationsOp(id, name, description string, metadataDB database.MetadataDatabaseInterface) *ListConfigurationsOp {
	return &ListConfigurationsOp{
		BaseSystemOperation: system.BaseSystemOperation{
			BaseSystemComponent: system.BaseSystemComponent{
				BaseComponent: component.BaseComponent{
					Id:   id,
					Nm:   name,
					Desc: description,
				},
			},
		},

		metadataDB: metadataDB,
	}
}

// Execute performs the operation with the given context and input parameters,
// and returns any output or error encountered.
func (bo *ListConfigurationsOp) Execute(ctx *context.Context,
	input *system.SystemOperationInput) (*system.SystemOperationOutput, error) {
	// Retrieve all metadata entries from the database
	entries, err := bo.metadataDB.GetAll()
	if err != nil {
		return nil, err
	}

	// Extract ProjectIDs from the retrieved entries
	var projectIDs []string
	for _, entry := range entries {
		projectIDs = append(projectIDs, entry.ProjectID)
	}

	// Create SystemOperationOutput with the list of ProjectIDs
	output := &system.SystemOperationOutput{
		Data: projectIDs,
	}

	return output, nil
}
