package system

import (
	"errors"
	"fmt"
	"sync"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/common/event"
	"github.com/edward1christian/block-forge/pkg/application/common/logger"
	"github.com/edward1christian/block-forge/pkg/application/components"
)

// SystemImpl represents the core system in the application.
type SystemImpl struct {
	SystemInterface
	mutex         sync.RWMutex
	configuration *components.Configuration
	componentReg  components.ComponentRegistrar
	logger        logger.LoggerInterface
	eventBus      event.EventBusInterface
	status        SystemStatusType
}

// NewSystem creates a new instance of the SystemImpl.
func NewSystem(
	logger logger.LoggerInterface,
	eventBus event.EventBusInterface,
	configuration *components.Configuration,
	componentReg components.ComponentRegistrar) *SystemImpl {
	return &SystemImpl{
		logger:        logger,
		eventBus:      eventBus,
		componentReg:  componentReg,
		configuration: configuration,
		status:        SystemStoppedType,
	}
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
func (s *SystemImpl) Configuration() *components.Configuration {
	return s.configuration
}

// ComponentRegistry returns the component registry.
func (s *SystemImpl) ComponentRegistry() components.ComponentRegistrar {
	return s.componentReg
}

// Initialize initializes the system component by executing the initialize operation.
func (s *SystemImpl) Initialize(ctx *context.Context) error {
	// Override this function to customize system initialization

	s.status = SystemInitializedType

	return nil
}

// InitializeService initializes a single service based on the provided configuration.
func (s *SystemImpl) InitializeService(ctx *context.Context, serviceConfig *components.ServiceConfiguration) error {
	// Retrieve the factory for the service component
	factory, err := s.ComponentRegistry().GetComponentFactory(serviceConfig.FactoryName)
	if err != nil {
		return fmt.Errorf("failed to get component factory for service %s: %w", serviceConfig.Name, err)
	}

	// Create an instance of the service component using the factory
	service, err := factory.CreateComponent(&serviceConfig.ComponentConfig)
	if err != nil {
		return fmt.Errorf("failed to create service %s: %w", serviceConfig.Name, err)
	}
	// Check if the service implements the SystemServiceInterface interface
	systemService, ok := service.(SystemServiceInterface)
	if !ok {
		return errors.New("service does not implement SystemServiceInterface")
	}

	// Initialize the service
	return systemService.Initialize(ctx, s)
}

// InitializeOperation initializes a single operation based on the provided configuration.
func (s *SystemImpl) InitializeOperation(ctx *context.Context, operationConfig *components.OperationConfiguration) error {
	// Retrieve the factory for the operation component
	factory, err := s.ComponentRegistry().GetComponentFactory(operationConfig.FactoryName)
	if err != nil {
		return fmt.Errorf("failed to get component factory for operation %s: %w", operationConfig.Name, err)
	}

	// Create an instance of the operation component using the factory
	operation, err := factory.CreateComponent(&operationConfig.ComponentConfig)
	if err != nil {
		return fmt.Errorf("failed to create operation %s: %w", operationConfig.Name, err)
	}

	// Check if the operation implements the OperationInterface interface
	systemOperation, ok := operation.(OperationInterface)
	if !ok {
		return errors.New("operation does not implement OperationInterface")
	}

	// Initialize the operation
	return systemOperation.Initialize(ctx, s)
}

// Start starts the system component along with all registered services.
func (s *SystemImpl) Start(ctx *context.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if s.status != SystemInitializedType {
		return ErrSystemNotInitialized
	}

	// Retrieve all components of type ServiceType
	components, err := s.ComponentRegistry().GetComponentByType(components.ServiceType)
	if err != nil {
		return fmt.Errorf("failed to retrieve components: %w", err)
	}

	// Iterate over each service component and start it
	for _, service := range components {
		// Check if the component implements SystemServiceInterface
		systemService, ok := service.(SystemServiceInterface)
		if !ok {
			return fmt.Errorf("failed to start service: component %v is not a service", service)
		}

		// Start the service
		if err := systemService.Start(ctx); err != nil {
			return fmt.Errorf("failed to start service: %w", err)
		}
	}

	s.status = SystemStartedType
	return nil
}

// Stop stops the system component along with all registered services.
func (s *SystemImpl) Stop(ctx *context.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if s.status != SystemStartedType {
		return ErrSystemNotStarted
	}
	// Retrieve all components of type ServiceType
	components, err := s.ComponentRegistry().GetComponentByType(components.ServiceType)
	if err != nil {
		return fmt.Errorf("failed to retrieve components: %w", err)
	}

	// Iterate over each service component and start it
	for _, service := range components {
		// Check if the component implements SystemServiceInterface
		systemService, ok := service.(SystemServiceInterface)
		if !ok {
			return fmt.Errorf("failed to start service: component %v is not a service", service)
		}

		// Stop the service
		if err := systemService.Stop(ctx); err != nil {
			// Log the error, but continue stopping other services
			s.logger.Log(logger.LevelError, "Error stopping service:", err)
		}
	}

	s.status = SystemStoppedType
	return nil
}

// ExecuteOperation executes the operation with the given ID and input data.
// Returns the output of the operation and an error if the operation is not found or if execution fails.
func (s *SystemImpl) ExecuteOperation(ctx *context.Context, operationID string, data *OperationInput) (*OperationOutput, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// Retrieve the operation by its ID
	component, err := s.ComponentRegistry().GetComponent(operationID)
	if err != nil {
		return nil, err
	}
	// Check if the component implements Operation interface
	operation, ok := component.(OperationInterface)
	if !ok {
		return nil, fmt.Errorf("failed to start service: component %v is not an operation", operation)
	}
	// Execute the operation
	return operation.Execute(ctx, data)
}

// StartService starts the service with the given ID.
// Returns an error if the service ID is not found or other error
func (s *SystemImpl) StartService(ctx *context.Context, serviceID string) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// Retrieve the service by its ID
	component, err := s.ComponentRegistry().GetComponent(serviceID)
	if err != nil {
		return err
	}
	// Check if the component implements SystemServiceInterface interface
	service, ok := component.(SystemServiceInterface)
	if !ok {
		return fmt.Errorf("failed to start service: component %v is not a service", service)
	}

	// Start the service
	return service.Start(ctx)
}

// StopService stops the service with the given ID.
// Returns an error if the service ID is not found or other error.
func (s *SystemImpl) StopService(ctx *context.Context, serviceID string) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// Retrieve the service by its ID
	component, err := s.ComponentRegistry().GetComponent(serviceID)
	if err != nil {
		return err
	}
	// Check if the component implements SystemServiceInterface interface
	service, ok := component.(SystemServiceInterface)
	if !ok {
		return fmt.Errorf("failed to start service: component %v is not a service", service)
	}

	// Start the service
	return service.Stop(ctx)
}

// RestartService restarts the service with the given ID.
// Returns an error if the service ID is not found or other error.
func (s *SystemImpl) RestartService(ctx *context.Context, serviceID string) error {
	// Stop the service first
	if err := s.StopService(ctx, serviceID); err != nil {
		return err
	}

	// Start the service
	return s.StartService(ctx, serviceID)
}
