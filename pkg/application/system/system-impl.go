package system

import (
	"errors"
	"fmt"
	"sync"

	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/application/event"
	"github.com/edward1christian/block-forge/pkg/application/logger"
)

// Custom errors
var (
	ErrComponentNil                  = errors.New("component is nil")
	ErrComponentFactoryNil           = errors.New("component factory is nil")
	ErrComponentFactoryAlreadyExists = errors.New("component factory already exists")
	ErrComponentNotFound             = errors.New("component not found")
	ErrServiceAlreadyExists          = errors.New("service already exists")
	ErrServiceNotRegistered          = errors.New("service not registered")
	ErrOperationNotRegistered        = errors.New("operation not registered")
	ErrOperationAlreadyExists        = errors.New("operation already exists")
)

// SystemImpl represents the core system in the application.
type SystemImpl struct {
	BootableComponent
	mutex              sync.RWMutex
	eventBus           event.EventBusInterface
	logger             logger.LoggerInterface
	configuration      Configuration
	operations         map[string]Operation
	services           map[string]SystemService
	componentFactories map[string]ComponentFactory
}

// NewSystem creates a new instance of the SystemImpl.
func NewSystem(
	eventBus event.EventBusInterface,
	logger logger.LoggerInterface,
	configuration Configuration) *SystemImpl {
	return &SystemImpl{
		eventBus:           eventBus,
		logger:             logger,
		configuration:      configuration,
		operations:         make(map[string]Operation),
		services:           make(map[string]SystemService),
		componentFactories: make(map[string]ComponentFactory),
	}
}

// ID returns the unique identifier of the system.
func (s *SystemImpl) ID() string {
	return "system"
}

// Name returns the name of the system.
func (s *SystemImpl) Name() string {
	return "System"
}

// Description returns the description of the system.
func (s *SystemImpl) Description() string {
	return "Core system in the application"
}

// Initialize initializes the system component based on the provided configuration.
func (s *SystemImpl) Initialize(ctx *context.Context) error {
	for _, serviceConfig := range s.configuration.Services {
		factory, err := s.GetComponentFactory(serviceConfig.FactoryName)
		if err != nil {
			return fmt.Errorf("failed to get component factory: %w", err)
		}

		service, err := factory(ctx, &serviceConfig.ComponentConfig)
		if err != nil {
			return fmt.Errorf("failed to create service: %w", err)
		}

		systemService, ok := service.(SystemService)
		if !ok {
			return errors.New("service is not a SystemService")
		}

		if err := systemService.Initialize(ctx, s); err != nil {
			return fmt.Errorf("failed to initialize service: %w", err)
		}

		if err := s.RegisterService(serviceConfig.ID, systemService); err != nil {
			return fmt.Errorf("failed to register service: %w", err)
		}
	}

	for _, operationConfig := range s.configuration.Operations {
		factory, err := s.GetComponentFactory(operationConfig.FactoryName)
		if err != nil {
			return fmt.Errorf("failed to get component factory: %w", err)
		}

		operation, err := factory(ctx, &operationConfig.ComponentConfig)
		if err != nil {
			return fmt.Errorf("failed to create operation: %w", err)
		}

		systemOperation, ok := operation.(Operation)
		if !ok {
			return errors.New("Operation is not an Operation")
		}

		if err := s.RegisterOperation(operationConfig.ID, systemOperation); err != nil {
			return fmt.Errorf("failed to register operation: %w", err)
		}
	}

	return nil
}

// Start starts the system component along with all registered services.
func (s *SystemImpl) Start(ctx *context.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for _, service := range s.services {
		if err := service.Start(ctx); err != nil {
			return fmt.Errorf("failed to start service: %w", err)
		}
	}

	return nil
}

// Stop stops the system component along with all registered services.
func (s *SystemImpl) Stop(ctx *context.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for _, service := range s.services {
		if err := service.Stop(ctx); err != nil {
			s.logger.Log(logger.LevelError, "Error stopping service:", err)
		}
	}

	return nil
}

// RegisterOperation registers an operation with the given ID.
// Returns an error if the operation ID is already registered or if the operation is nil.
func (s *SystemImpl) RegisterOperation(operationID string, operation Operation) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if operation == nil {
		return ErrComponentNil
	}
	if _, exists := s.operations[operationID]; exists {
		return ErrOperationAlreadyExists
	}
	s.operations[operationID] = operation
	return nil
}

// UnregisterOperation unregisters the operation with the given ID.
// Returns an error if the operation ID is not found.
func (s *SystemImpl) UnregisterOperation(operationID string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.operations[operationID]; !exists {
		return ErrOperationNotRegistered
	}
	delete(s.operations, operationID)
	return nil
}

// ExecuteOperation executes the operation with the given ID and input data.
// Returns the output of the operation and an error if the operation is not found or if execution fails.
func (s *SystemImpl) ExecuteOperation(ctx *context.Context, operationID string, data *OperationInput) (*OperationOutput, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	operation, exists := s.operations[operationID]
	if !exists {
		return nil, ErrOperationNotRegistered
	}
	return operation.Execute(ctx, data)
}

// RegisterService registers a SystemService with the given ID.
// Returns an error if the service ID is already registered or if the service is nil.
func (s *SystemImpl) RegisterService(serviceID string, service SystemService) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if service == nil {
		return ErrComponentNil
	}
	if _, exists := s.services[serviceID]; exists {
		return ErrServiceAlreadyExists
	}
	s.services[serviceID] = service
	return nil
}

// UnregisterService unregisters a SystemService with the given ID.
// Returns an error if the service ID is not found.
func (s *SystemImpl) UnregisterService(serviceID string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.services[serviceID]; !exists {
		return ErrServiceNotRegistered
	}
	delete(s.services, serviceID)
	return nil
}

// StartService starts the service with the given ID.
// Returns an error if the service ID is not found or other error
func (s *SystemImpl) StartService(serviceID string, ctx *context.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	service, exists := s.services[serviceID]
	if !exists {
		return ErrServiceNotRegistered
	}
	return service.Start(ctx)
}

// StopService stops the service with the given ID.
// Returns an error if the service ID is not found or other error.
func (s *SystemImpl) StopService(serviceID string, ctx *context.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	service, exists := s.services[serviceID]
	if !exists {
		return ErrServiceNotRegistered
	}
	return service.Stop(ctx)
}

// RegisterComponentFactory registers a component factory with the system.
// Returns an error if the component name is already registered or if the component is nil.
func (s *SystemImpl) RegisterComponentFactory(name string, factory ComponentFactory) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if factory == nil {
		return ErrComponentFactoryNil
	}
	if _, exists := s.componentFactories[name]; exists {
		return ErrComponentFactoryAlreadyExists
	}
	s.componentFactories[name] = factory
	return nil
}

// GetComponentFactory retrieves the factory for creating components by name.
// Returns an error if the component name is not found or other error.
func (s *SystemImpl) GetComponentFactory(name string) (ComponentFactory, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	factory, exists := s.componentFactories[name]
	if !exists {
		return nil, ErrComponentNotFound
	}
	return factory, nil
}

// Logger returns the system logger.
func (s *SystemImpl) Logger() logger.LoggerInterface {
	return s.logger
}

// EventBus returns the system event bus.
func (s *SystemImpl) EventBus() event.EventBusInterface {
	return s.eventBus
}

// Configuration returns the system configuration.
func (s *SystemImpl) Configuration() Configuration {
	return s.configuration
}
