package system

import (
	"errors"
	"fmt"

	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/application/event"
	"github.com/edward1christian/block-forge/pkg/application/logger"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/edward1christian/block-forge/pkg/etl"
	"github.com/edward1christian/block-forge/pkg/etl/utils"
)

// ETLSystemImpl represents a concrete implementation of the etl.ETLSystem interface.
type ETLSystemImpl struct {
	etl.ETLSystem
	eventBus      event.EventBusInterface
	logger        logger.LoggerInterface
	configuration system.Configuration
	etlProcesses  map[string]*etl.ETLProcess
	scheduledJobs map[string]*etl.ScheduledETLProcess
	factories     map[string]etl.ETLComponentFactory
}

// NewETLSystem creates a new instance of ETLSystemImpl.
func NewETLSystem(eventBus event.EventBusInterface, logger logger.LoggerInterface, configuration system.Configuration) *ETLSystemImpl {
	return &ETLSystemImpl{
		eventBus:      eventBus,
		logger:        logger,
		configuration: configuration,
		etlProcesses:  make(map[string]*etl.ETLProcess),
		scheduledJobs: make(map[string]*etl.ScheduledETLProcess),
		factories:     make(map[string]etl.ETLComponentFactory),
	}
}

// RegisterETLComponentFactory registers a factory for creating components.
func (s *ETLSystemImpl) RegisterETLComponentFactory(name string, factory etl.ETLComponentFactory) error {
	if _, exists := s.factories[name]; exists {
		return fmt.Errorf("ETL component factory for %s already exists", name)
	}
	s.factories[name] = factory
	return nil
}

// GetAdapterFactory retrieves the factory for creating adapters by name.
func (s *ETLSystemImpl) GetETLComponentFactory(name string) (etl.ETLComponentFactory, bool) {
	factory, ok := s.factories[name]
	return factory, ok
}

// InitializeETLProcess initializes an ETL process with the provided configuration.
func (s *ETLSystemImpl) InitializeETLProcess(ctx *context.Context, config *etl.ETLConfig) (*etl.ETLProcess, error) {
	// Execute the system operation
	opOutput, err := utils.ExecuteSystemOp(ctx, s, ETLCreateProcessOp, config)
	if err != nil {
		return nil, err
	}

	process, ok := opOutput.Data.(*etl.ETLProcess)
	if !ok {
		return nil, etl.ErrInvalidProcess
	}

	processConfig := process.Config
	// Initialize pipelines
	for _, componentConfig := range processConfig.Components {
		pipeline, err := s.initializeETLComponent(ctx, componentConfig)
		if err != nil {
			// Rollback initialized adapters
			_, opErr := utils.ExecuteSystemOp(ctx, s, ETLCreateProcessOp, config)
			if opErr != nil {
				return nil, opErr
			}
			return nil, err
		}
		process.Components[componentConfig.Name] = pipeline
	}
	// Save process
	s.etlProcesses[process.ID] = process

	return process, nil
}

// InitializeETLComponent initializes an ETL component.
func (s *ETLSystemImpl) initializeETLComponent(
	ctx *context.Context,
	componentConfig *etl.ETLComponentConfig) (etl.ETLComponent, error) {
	// Get the component factory
	factory, ok := s.GetETLComponentFactory(componentConfig.FactoryNm)
	if !ok {
		return nil, errors.New("component factory not found for: " + componentConfig.Name)
	}

	// Instantiate the component
	component, err := factory(ctx, componentConfig)
	if err != nil {
		return nil, err
	}

	// Initialize the component
	if err := component.Initialize(ctx, s); err != nil {
		return nil, err
	}

	return component, nil
}

// StartETLProcess starts an ETL process with the given ID.
func (s *ETLSystemImpl) StartETLProcess(ctx *context.Context, processID string) error {
	// Retrieve the process
	process, err := s.GetETLProcess(processID)
	if err != nil {
		return etl.ErrProcessNotFound
	}

	// Execute the operation to start the process
	_, err = utils.ExecuteSystemOp(ctx, s, ETLStartProcessOp, process)
	if err != nil {
		return err
	}

	return nil
}

// StopETLProcess stops an ETL process with the given ID.
func (s *ETLSystemImpl) StopETLProcess(ctx *context.Context, processID string) error {
	// Retrieve the process
	process, err := s.GetETLProcess(processID)
	if err != nil {
		return etl.ErrProcessNotFound
	}

	// Execute the operation to start the process
	_, err = utils.ExecuteSystemOp(ctx, s, ETLStopProcessOp, process)
	if err != nil {
		return err
	}

	return nil
}

// GetETLProcess retrieves an ETL process by its ID.
func (s *ETLSystemImpl) GetETLProcess(processID string) (*etl.ETLProcess, error) {
	process, ok := s.etlProcesses[processID]
	if !ok {
		return nil, etl.ErrProcessNotFound
	}
	return process, nil
}

// GetAllETLProcesses retrieves all ETL processes.
func (s *ETLSystemImpl) GetAllETLProcesses() []*etl.ETLProcess {
	processes := make([]*etl.ETLProcess, 0, len(s.etlProcesses))
	for _, process := range s.etlProcesses {
		processes = append(processes, process)
	}
	return processes
}

// ScheduleETLProcess schedules an ETL process for execution.
func (s *ETLSystemImpl) ScheduleETLProcess(process *etl.ETLProcess, schedule etl.Schedule) error {
	// Implement scheduling logic here
	job := &etl.ScheduledETLProcess{
		Process:  process,
		Schedule: schedule,
	}
	s.scheduledJobs[process.ID] = job
	return nil
}

// GetScheduledETLProcesses retrieves all scheduled ETL processes.
func (s *ETLSystemImpl) GetScheduledETLProcesses() []*etl.ScheduledETLProcess {
	jobs := make([]*etl.ScheduledETLProcess, 0, len(s.scheduledJobs))
	for _, job := range s.scheduledJobs {
		jobs = append(jobs, job)
	}
	return jobs
}

// RemoveScheduledETLProcess removes a scheduled ETL process by its ID.
func (s *ETLSystemImpl) RemoveScheduledETLProcess(processID string) error {
	if _, ok := s.scheduledJobs[processID]; !ok {
		return etl.ErrScheduledProcessNotFound
	}
	delete(s.scheduledJobs, processID)
	return nil
}

// EventBus returns the event bus associated with the system.
func (s *ETLSystemImpl) EventBus() event.EventBusInterface {
	return s.eventBus
}

// Logger returns the logger associated with the system.
func (s *ETLSystemImpl) Logger() logger.LoggerInterface {
	return s.logger
}

// Configuration returns the system configuration.
func (s *ETLSystemImpl) Configuration() system.Configuration {
	return s.configuration
}
