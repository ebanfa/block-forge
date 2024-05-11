package commands

import (
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	novaConfigApi "github.com/edward1christian/block-forge/nova/pkg/config"
	"github.com/edward1christian/block-forge/nova/pkg/store"
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

	// Construct the service with the injected database
	return NewListConfigurationsOp(config.ID, config.Name, config.Description), nil
}

// ListConfigurationsOp represents a concrete implementation of the SystemOperationInterface.
type ListConfigurationsOp struct {
	system.BaseSystemOperation
}

// Type returns the type of the component.
func (bo *ListConfigurationsOp) Type() component.ComponentType {
	return component.OperationType
}

func NewListConfigurationsOp(id, name, description string) *ListConfigurationsOp {
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
	}
}

// Execute performs the operation with the given context and input parameters,
// and returns any output or error encountered.
func (bo ListConfigurationsOp) Execute(ctx_ *context.Context,
	input *system.SystemOperationInput) (*system.SystemOperationOutput, error) {

	multiStore := bo.System.MultiStore()
	configuration := bo.System.Configuration()

	// Validate the configuration
	novaConfig, ok := configuration.CustomConfig.(novaConfigApi.NovaConfig)
	if !ok {
		return nil, errors.New("failed to create project. Invalid configuration")
	}

	// Get the MetadataDatabase instance
	// Creates or retrieve an existing store
	entryStore, _, err := multiStore.CreateStore(novaConfig.MetadataDbName)
	if err != nil {
		return nil, err
	}

	metadataStore := store.NewMetadataStore(entryStore)

	// Load the current working version
	_, err = metadataStore.Load()
	if err != nil {
		return nil, err
	}

	// Retrieve all metadata entries from the database
	entries, err := metadataStore.GetAllMetadata()
	if err != nil {
		return nil, err
	}

	// Extract ProjectIDs from the retrieved entries
	var projectIDs []string
	// Create a new tabwriter.Writer instance.
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 5, ' ', 0)
	// Print header
	fmt.Fprintf(w, "Name\tProject ID\n")
	// Print underline
	fmt.Fprintln(w, "------------\t----------")

	// Iterate over the collection
	for _, entry := range entries {
		// Print name and project ID in tab-separated columns
		fmt.Fprintf(w, "%s\t%s\n", entry.ProjectName, entry.ProjectID)
	}
	fmt.Fprintf(w, "\n")

	// Flush the buffer to ensure all data is written
	w.Flush()

	// Create SystemOperationOutput with the list of ProjectIDs
	output := &system.SystemOperationOutput{
		Data: projectIDs,
	}
	err = metadataStore.Close()
	if err != nil {
		return nil, err
	}

	return output, nil
}
