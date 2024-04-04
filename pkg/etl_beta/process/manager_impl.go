package process

import (
	"errors"
	"fmt"
	"sync"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/components"
	"github.com/edward1christian/block-forge/pkg/application/system"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
	etlComponents "github.com/edward1christian/block-forge/pkg/etl_beta/components"
	"github.com/edward1christian/block-forge/pkg/etl_beta/utils"
)

// etlManagerService implements the ETLManagerService interface.
type etlManagerService struct {
	systemApi.SystemComponentInterface
	system    systemApi.SystemInterface
	processes map[string]*ETLProcess          // Map to store ETL processes by ID
	schedule  map[string]*ScheduledETLProcess // Map to store scheduled processes by ID
	mutex     sync.Mutex                      // Mutex for concurrent access to maps
}

// NewETLManagerService creates a new instance of ETLManagerService.
func NewETLManagerService() ETLManagerService {
	return &etlManagerService{
		processes: make(map[string]*ETLProcess),
		schedule:  make(map[string]*ScheduledETLProcess),
	}
}

// Initialize initializes the ETLManagerService.
func (m *etlManagerService) Initialize(ctx *context.Context, system system.SystemInterface) error {
	// Save reference to the system
	m.system = system

	// Get the system configuration
	config := system.Configuration()

	// Retrieve the process definitions from the configuration
	processesConfig := config.CustomConfig

	// Check if the processes configuration is of type []interface{}
	etlProcessesConfig, ok := processesConfig.([]*etlComponents.ETLProcessConfig)
	if !ok {
		return errors.New("processes configuration is not an array")
	}

	// Iterate over each process configuration and initialize the corresponding ETL process
	for _, etlProcessConfig := range etlProcessesConfig {

		// Initialize the ETL process using the ETLManagerService method
		_, err := m.InitializeETLProcess(ctx, etlProcessConfig)
		if err != nil {
			return fmt.Errorf("error initializing ETL process: %v", err)
		}
	}

	return nil
}

// InitializeETLProcess initializes an ETL process with the provided configuration.
func (m *etlManagerService) InitializeETLProcess(ctx *context.Context, config *etlComponents.ETLProcessConfig) (*ETLProcess, error) {
	// Generate a unique ID for the ETL process
	processID, err := utils.NewProcessIDGenerator("").GenerateID()
	if err != nil {
		return nil, err
	}

	// Create a new ETL process instance
	process := &ETLProcess{
		ID:         processID,
		Config:     config,
		Status:     ETLProcessStatusInitialized,
		Components: make(map[string]etlComponents.ETLProcessComponent),
	}

	// Add the process to the map
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.processes[processID] = process

	// Initialize each component of the ETL process
	for _, compConfig := range config.Components {
		// Retrieve the factory for the component type from the system
		factory, err := m.system.ComponentRegistry().GetComponentFactory(compConfig.FactoryName)
		if err != nil {
			return nil, err
		}

		// Create an instance of the component using the factory
		component, err := factory.CreateComponent(compConfig)
		if err != nil {
			return nil, err
		}

		// Check if the component implements the ETLProcessComponent interface
		etlComponent, ok := component.(etlComponents.ETLProcessComponent)
		if !ok {
			return nil, errors.New("component is not an ETLProcess component")
		}

		// Add the component to the ETL process
		process.Components[component.ID()] = etlComponent
	}

	return process, nil
}

// StartETLProcess starts an ETL process with the given ID.
func (m *etlManagerService) StartETLProcess(ctx *context.Context, processID string) error {
	// Fetch the ETL process
	m.mutex.Lock()
	defer m.mutex.Unlock()
	process, found := m.processes[processID]
	if !found {
		return errors.New("process not found")
	}

	// Start each component of the ETL process
	for _, component := range process.Components {
		// Check if the component is startable
		startable, ok := component.(components.StartableInterface)
		if !ok {
			continue
		}
		// Start the component
		if err := startable.Start(ctx); err != nil {
			return err
		}
	}

	// Update the process status
	process.Status = ETLProcessStatusRunning
	return nil
}

// StopETLProcess stops an ETL process with the given ID.
func (m *etlManagerService) StopETLProcess(ctx *context.Context, processID string) error {
	// Fetch the ETL process
	m.mutex.Lock()
	defer m.mutex.Unlock()
	process, found := m.processes[processID]
	if !found {
		return errors.New("process not found")
	}

	// Stop each component of the ETL process
	for _, component := range process.Components {
		// Check if the component is startable
		startable, ok := component.(components.StartableInterface)
		if !ok {
			continue
		}
		// Stop the component
		if err := startable.Stop(ctx); err != nil {
			return err
		}
	}

	// Update the process status
	process.Status = ETLProcessStatusStopped
	return nil
}

// GetETLProcess retrieves an ETL process by its ID.
func (m *etlManagerService) GetETLProcess(processID string) (*ETLProcess, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	process, found := m.processes[processID]
	if !found {
		return nil, errors.New("process not found")
	}
	return process, nil
}

// GetAllETLProcesses retrieves all ETL processes.
func (m *etlManagerService) GetAllETLProcesses() []*ETLProcess {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	processes := make([]*ETLProcess, 0, len(m.processes))
	for _, process := range m.processes {
		processes = append(processes, process)
	}
	return processes
}

// ScheduleETLProcess schedules an ETL process for execution.
func (m *etlManagerService) ScheduleETLProcess(process *ETLProcess, schedule Schedule) error {
	// Add scheduling logic here

	// Create a new ScheduledETLProcess instance
	scheduledProcess := &ScheduledETLProcess{
		Process:  process,
		Schedule: schedule,
	}

	// Add the scheduled process to the map
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.schedule[process.ID] = scheduledProcess

	return nil
}

// GetScheduledETLProcesses retrieves all scheduled ETL processes.
func (m *etlManagerService) GetScheduledETLProcesses() []*ScheduledETLProcess {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	scheduledProcesses := make([]*ScheduledETLProcess, 0, len(m.schedule))
	for _, scheduledProcess := range m.schedule {
		scheduledProcesses = append(scheduledProcesses, scheduledProcess)
	}
	return scheduledProcesses
}

// RemoveScheduledETLProcess removes a scheduled ETL process by its ID.
func (m *etlManagerService) RemoveScheduledETLProcess(processID string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	_, found := m.schedule[processID]
	if !found {
		return errors.New("scheduled process not found")
	}
	delete(m.schedule, processID)
	return nil
}
