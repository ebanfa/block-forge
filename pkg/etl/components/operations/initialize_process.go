package operations

import (
	"fmt"

	"github.com/edward1christian/block-forge/pkg/application/common"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/components"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/edward1christian/block-forge/pkg/etl"
	etlComponentsApi "github.com/edward1christian/block-forge/pkg/etl/components"
	"github.com/edward1christian/block-forge/pkg/etl/process"
)

// InitializeETLProcessOperation represents a concrete implementation of the OperationInterface.
type InitializeETLProcessOperation struct {
	systemApi.BaseSystemOperation
	idGenerator common.IDGeneratorInterface
}

// NewOperationComponent creates a new instance of InitializeETLProcessOperation.
func NewInitializeETLProcessOperation(id, name, description string, idGenerator common.IDGeneratorInterface) *InitializeETLProcessOperation {
	return &InitializeETLProcessOperation{
		idGenerator: idGenerator,
		BaseSystemOperation: systemApi.BaseSystemOperation{
			BaseSystemComponent: systemApi.BaseSystemComponent{
				BaseComponent: components.BaseComponent{
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
func (ie *InitializeETLProcessOperation) Execute(ctx *context.Context, input *systemApi.OperationInput) (*systemApi.OperationOutput, error) {
	// Check if the processes configuration is of type []*process.ETLProcessConfig
	processesConfig, ok := input.Data.([]*process.ETLProcessConfig)
	if !ok {
		return nil, etl.ErrInvalidProcessesConfig
	}

	// Iterate over each process configuration and initialize the corresponding ETL process
	for _, processConfig := range processesConfig {
		// Initialize the ETL process
		_, err := ie.InitializeProcess(ctx, processConfig)
		if err != nil {
			return nil, fmt.Errorf("error initializing ETL process: %v", err)
		}
	}

	return nil, nil
}

// InitializeETLProcess initializes an ETL process with the provided configuration.
func (ie *InitializeETLProcessOperation) InitializeProcess(ctx *context.Context, config *process.ETLProcessConfig) (*process.ETLProcess, error) {
	// Cant have a process with no components for now
	if len(config.Components) == 0 {
		return nil, etl.ErrEmptyProcessConfig
	}

	// Generate a unique ID for the ETL process
	processID, err := ie.idGenerator.GenerateID()
	if err != nil {
		return nil, err
	}

	// Create a new ETL process instance
	process := &process.ETLProcess{
		ID:         processID,
		Config:     config,
		Status:     process.ETLProcessStatusInitialized,
		Components: make(map[string]etlComponentsApi.ETLProcessComponent),
	}

	// Initialize each component defined in the process config
	sys := ie.BaseSystemOperation.System
	for _, compConfig := range config.Components {
		registrar := sys.ComponentRegistry()

		// Retrieve the factory for the component type from the system
		factory, err := registrar.GetComponentFactory(compConfig.FactoryName)
		if err != nil {
			return nil, err
		}

		// Create an instance of the component using the factory
		component, err := factory.CreateComponent(compConfig)
		if err != nil {
			return nil, err
		}

		// Check if the component implements the ETLProcessComponent interface
		etlComponent, ok := component.(etlComponentsApi.ETLProcessComponent)
		if !ok {
			return nil, etl.ErrNotProcessComponent
		}

		// Add the component to the ETL process
		process.Components[component.ID()] = etlComponent
	}

	return process, nil
}
