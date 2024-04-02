package etl

import (
	"errors"
	"fmt"

	"github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/application/event"
	"github.com/edward1christian/block-forge/pkg/application/logger"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// ETLSystemImpl represents a concrete implementation of the ETLSystem interface.
type ETLSystemImpl struct {
	ETLSystem
	eventBus      event.EventBusInterface
	operations    system.Operations
	logger        logger.LoggerInterface
	configuration config.Configuration
	etlProcesses  map[string]*ETLProcess
	scheduledJobs map[string]*ScheduledETLProcess
	factories     map[string]ETLComponentFactory
}

// NewETLSystem creates a new instance of ETLSystemImpl.
func NewETLSystem(eventBus event.EventBusInterface, operations system.Operations, logger logger.LoggerInterface, configuration config.Configuration) *ETLSystemImpl {
	return &ETLSystemImpl{
		eventBus:      eventBus,
		operations:    operations,
		logger:        logger,
		configuration: configuration,
		etlProcesses:  make(map[string]*ETLProcess),
		scheduledJobs: make(map[string]*ScheduledETLProcess),
		factories:     make(map[string]ETLComponentFactory),
	}
}

// RegisterETLComponentFactory registers a factory for creating components.
func (s *ETLSystemImpl) RegisterETLComponentFactory(name string, factory ETLComponentFactory) error {
	if _, exists := s.factories[name]; exists {
		return fmt.Errorf("ETL component factory for %s already exists", name)
	}
	s.factories[name] = factory
	return nil
}

// GetAdapterFactory retrieves the factory for creating adapters by name.
func (s *ETLSystemImpl) GetETLComponentFactory(name string) (ETLComponentFactory, bool) {
	factory, ok := s.factories[name]
	return factory, ok
}

// InitializeETLProcess initializes an ETL process with the provided configuration.
func (s *ETLSystemImpl) InitializeETLProcess(ctx *context.Context, config *ETLConfig) (*ETLProcess, error) {
	generator := NewProcessIDGenerator("etl")
	helper := NewETLProcessInitHelper(generator)

	process, err := helper.CreateETLProcess(ctx, config)
	if err != nil {
		return nil, err
	}

	processConfig := process.Config
	// Initialize pipelines
	for _, componentConfig := range processConfig.Components {
		pipeline, err := s.initializeETLComponent(ctx, componentConfig)
		if err != nil {
			// Rollback initialized adapters
			helper.RollbackComponents(ctx, process)
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
	componentConfig *ETLComponentConfig) (ETLComponent, error) {
	factory, ok := s.GetETLComponentFactory(componentConfig.FactoryNm)
	if !ok {
		return nil, errors.New("component factory not found for: " + componentConfig.Name)
	}

	component, err := factory(ctx, componentConfig)
	if err != nil {
		return nil, err
	}

	if err := component.Initialize(ctx, s); err != nil {
		return nil, err
	}

	return component, nil
}

// StartETLProcess starts an ETL process with the given ID.
func (s *ETLSystemImpl) StartETLProcess(ctx *context.Context, processID string) error {
	helper := NewETLProcessStartHelper(s)
	process, err := s.GetETLProcess(processID)
	if err != nil {
		return ErrProcessNotFound
	}

	return helper.StartETLProcess(ctx, process)
}

// StopETLProcess stops an ETL process with the given ID.
func (s *ETLSystemImpl) StopETLProcess(ctx *context.Context, processID string) error {
	helper := NewETLProcessStopHelper()
	process, err := s.GetETLProcess(processID)
	if err != nil {
		return ErrProcessNotFound
	}
	return helper.StopETLProcess(ctx, process)
}

// GetETLProcess retrieves an ETL process by its ID.
func (s *ETLSystemImpl) GetETLProcess(processID string) (*ETLProcess, error) {
	process, ok := s.etlProcesses[processID]
	if !ok {
		return nil, ErrProcessNotFound
	}
	return process, nil
}

// GetAllETLProcesses retrieves all ETL processes.
func (s *ETLSystemImpl) GetAllETLProcesses() []*ETLProcess {
	processes := make([]*ETLProcess, 0, len(s.etlProcesses))
	for _, process := range s.etlProcesses {
		processes = append(processes, process)
	}
	return processes
}

// ScheduleETLProcess schedules an ETL process for execution.
func (s *ETLSystemImpl) ScheduleETLProcess(process *ETLProcess, schedule Schedule) error {
	// Implement scheduling logic here
	job := &ScheduledETLProcess{
		Process:  process,
		Schedule: schedule,
	}
	s.scheduledJobs[process.ID] = job
	return nil
}

// GetScheduledETLProcesses retrieves all scheduled ETL processes.
func (s *ETLSystemImpl) GetScheduledETLProcesses() []*ScheduledETLProcess {
	jobs := make([]*ScheduledETLProcess, 0, len(s.scheduledJobs))
	for _, job := range s.scheduledJobs {
		jobs = append(jobs, job)
	}
	return jobs
}

// RemoveScheduledETLProcess removes a scheduled ETL process by its ID.
func (s *ETLSystemImpl) RemoveScheduledETLProcess(processID string) error {
	if _, ok := s.scheduledJobs[processID]; !ok {
		return ErrScheduledProcessNotFound
	}
	delete(s.scheduledJobs, processID)
	return nil
}

// EventBus returns the event bus associated with the system.
func (s *ETLSystemImpl) EventBus() event.EventBusInterface {
	return s.eventBus
}

// Operations returns the operations manager associated with the system.
func (s *ETLSystemImpl) Operations() system.Operations {
	return s.operations
}

// Logger returns the logger associated with the system.
func (s *ETLSystemImpl) Logger() logger.LoggerInterface {
	return s.logger
}

// Configuration returns the system configuration.
func (s *ETLSystemImpl) Configuration() config.Configuration {
	return s.configuration
}
