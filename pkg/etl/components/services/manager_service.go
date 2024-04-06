package services

import (
	"errors"
	"fmt"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/components"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/edward1christian/block-forge/pkg/etl/process"
)

// BaseComponent represents a concrete implementation of the OperationInterface.
type ProcessManagerService struct {
	system.BaseSystemService // Embedding BaseComponent
	manager                  process.ProcessManagerInterface
}

// Type returns the type of the component.
func (pms *ProcessManagerService) Type() components.ComponentType {
	return components.ServiceType
}

func NewProcessManagerService(id, name, description string, manager process.ProcessManagerInterface) *ProcessManagerService {
	return &ProcessManagerService{
		manager: manager,

		BaseSystemService: system.BaseSystemService{
			BaseSystemComponent: system.BaseSystemComponent{
				BaseComponent: components.BaseComponent{
					Id:   id,
					Nm:   name,
					Desc: description,
				},
			},
		},
	}
}

// Initialize initializes the module.
// Returns an error if the initialization fails.
func (pms *ProcessManagerService) Initialize(ctx *context.Context, system system.SystemInterface) error {
	pms.System = system
	// Get the system configuration
	config := system.Configuration()

	// Retrieve the process definitions from the configuration
	processesConfig := config.CustomConfig

	// Check if the processes configuration is of type []interface{}
	etlProcessesConfig, ok := processesConfig.([]*process.ETLProcessConfig)
	if !ok {
		return errors.New("processes configuration is not an array")
	}

	// Iterate over each process configuration and initialize the corresponding ETL process
	for _, etlProcessConfig := range etlProcessesConfig {

		// Initialize the ETL process using the ETLManagerService method
		_, err := pms.manager.InitializeProcess(ctx, etlProcessConfig)
		if err != nil {
			return fmt.Errorf("error initializing ETL process: %v", err)
		}
	}

	return nil
}

// Start starts the component.
// Returns an error if the start operation fails.
func (pms *ProcessManagerService) Start(ctx *context.Context) error {
	// Iterate over each process and start it
	for _, etlProcess := range pms.manager.GetAllProcesses() {

		// Start the process
		err := pms.manager.StartProcess(ctx, etlProcess.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

// Stop stops the component.
// Returns an error if the stop operation fails.
func (pms *ProcessManagerService) Stop(ctx *context.Context) error {
	// Iterate over each process and start it
	for _, etlProcess := range pms.manager.GetAllProcesses() {

		// Start the process
		err := pms.manager.StopProcess(ctx, etlProcess.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
