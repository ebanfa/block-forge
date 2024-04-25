package operations

import (
	"fmt"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/component"
	configApi "github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// BuildServiceFactory is responsible for creating instances of BuildService.
type CreateDirectoryTaskFactory struct {
}

// CreateComponent creates a new instance of the BuildService.
func (bf *CreateDirectoryTaskFactory) CreateComponent(config *configApi.ComponentConfig) (component.ComponentInterface, error) {
	// Construct the service
	return NewCreateDirectoryTask(config.ID, config.Name, config.Description), nil
}

// BaseComponent represents a concrete implementation of the SystemOperationInterface.
type CreateDirectoryTask struct {
	system.BaseSystemComponent // Embedding BaseComponent
}

// Type returns the type of the component.
func (ctk *CreateDirectoryTask) Type() component.ComponentType {
	return component.BasicComponentType
}

func NewCreateDirectoryTask(id, name, description string) *CreateDirectoryTask {
	return &CreateDirectoryTask{
		BaseSystemComponent: system.BaseSystemComponent{
			BaseComponent: component.BaseComponent{
				Id:   id,
				Nm:   name,
				Desc: description,
			},
		},
	}
}

// Execute performs the operation with the given context and input parameters,
// and returns any output or error encountered.
func (ctk *CreateDirectoryTask) Execute(ctx *context.Context, input *system.SystemOperationInput) (*system.SystemOperationOutput, error) {
	// Perform operation logic here
	// For demonstration purposes, just return an error
	fmt.Printf("Executing create workspace task")
	return nil, nil
}
