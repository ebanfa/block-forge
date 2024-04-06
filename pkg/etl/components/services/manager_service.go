package services

import (
	"fmt"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/components"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/edward1christian/block-forge/pkg/etl"
	"github.com/edward1christian/block-forge/pkg/etl/process"
)

// ProcessManagerService represents a service for managing ETL processes.
type ProcessManagerService struct {
	systemApi.BaseSystemService                                 // Embedding BaseComponent for component properties
	manager                     process.ProcessManagerInterface // Process manager interface
	system                      systemApi.SystemInterface       // System interface
}

// NewProcessManagerService creates a new instance of ProcessManagerService.
func NewProcessManagerService(id, name, description string, manager process.ProcessManagerInterface) *ProcessManagerService {
	return &ProcessManagerService{
		BaseSystemService: systemApi.BaseSystemService{
			BaseSystemComponent: systemApi.BaseSystemComponent{
				BaseComponent: components.BaseComponent{
					Id:   id,
					Nm:   name,
					Desc: description,
				},
			},
		},
		manager: manager,
	}
}

// Initialize initializes the ProcessManagerService.
// It initializes ETL processes based on the configuration provided by the systemApi.
func (pms *ProcessManagerService) Initialize(ctx *context.Context, system systemApi.SystemInterface) error {
	pms.system = system

	// Retrieve processes configuration from system configuration
	config := system.Configuration()
	processesConfig, ok := config.CustomConfig.([]*process.ETLProcessConfig)
	if !ok {
		return etl.ErrInvalidProcessesConfig
	}

	// Iterate over each process configuration and initialize the corresponding ETL process
	for _, etlProcessConfig := range processesConfig {
		_, err := pms.manager.InitializeProcess(ctx, etlProcessConfig)
		if err != nil {
			return fmt.Errorf("failed to initialize ETL process: %w", err)
		}
	}

	return nil
}

// Start starts the ProcessManagerService.
// It starts all initialized ETL processes.
func (pms *ProcessManagerService) Start(ctx *context.Context) error {
	for _, etlProcess := range pms.manager.GetAllProcesses() {
		err := pms.manager.StartProcess(ctx, etlProcess.ID)
		if err != nil {
			return fmt.Errorf("failed to start ETL process: %w", err)
		}
	}
	return nil
}

// Stop stops the ProcessManagerService.
// It stops all running ETL processes.
func (pms *ProcessManagerService) Stop(ctx *context.Context) error {
	for _, etlProcess := range pms.manager.GetAllProcesses() {
		err := pms.manager.StopProcess(ctx, etlProcess.ID)
		if err != nil {
			return fmt.Errorf("failed to stop ETL process: %w", err)
		}
	}
	return nil
}

// Type returns the type of the component.
func (pms *ProcessManagerService) Type() components.ComponentType {
	return components.ServiceType
}
