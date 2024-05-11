package system

import (
	"fmt"
	"sync"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/common/event"
	"github.com/edward1christian/block-forge/pkg/application/common/logger"
	"github.com/edward1christian/block-forge/pkg/application/component"
	"github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/edward1christian/block-forge/pkg/application/store"
)

// SystemComponentInterface represents a component in the system.
type SystemComponentInterface interface {
	component.ComponentInterface

	// Initialize initializes the module.
	// Returns an error if the initialization fails.
	Initialize(ctx *context.Context, system SystemInterface) error
}

// SystemServiceInterface represents a service within the system.
type SystemServiceInterface interface {
	component.StartableInterface
	SystemComponentInterface
}

// SystemOperationInput represents the input data for an operation.
type SystemOperationInput struct {
	// Data is the input data for the operation.
	Data interface{}
}

// SystemOperationOutput represents the response data from an operation.
type SystemOperationOutput struct {
	// Data is the response data from the operation.
	Data interface{}
}

// SystemOperation represents a unit of work that can be executed.
type SystemOperationInterface interface {
	SystemComponentInterface

	// Execute performs the operation with the given context and input parameters,
	// and returns any output or error encountered.
	Execute(ctx *context.Context, input *SystemOperationInput) (*SystemOperationOutput, error)
}

// SystemInterface represents the core system in the application.
type SystemInterface interface {
	component.BootableInterface
	component.StartableInterface

	// Logger returns the system logger.
	Logger() logger.LoggerInterface

	// EventBus returns the system event bus.
	EventBus() event.EventBusInterface

	// Configuration returns the system configuration.
	Configuration() *config.Configuration

	// ComponentRegistry returns the component registry
	ComponentRegistry() component.ComponentRegistrarInterface

	// MultiStore returns the multi-level store used by the system, which is a store of stores.
	MultiStore() store.MultiStore

	// PluginManager returns the plugin manager
	PluginManager() PluginManagerInterface

	// ExecuteOperation executes the operation with the given ID and input data.
	// Returns the output of the operation and an error if the operation is not found or if execution fails.
	ExecuteOperation(ctx *context.Context, operationID string, data *SystemOperationInput) (*SystemOperationOutput, error)

	// StartService starts the service with the given ID.
	// Returns an error if the service ID is not found or other error.
	StartService(ctx *context.Context, serviceID string) error

	// StopService stops the service with the given ID.
	// Returns an error if the service ID is not found or other error.
	StopService(ctx *context.Context, serviceID string) error

	// RestartService restarts the service with the given ID.
	// Returns an error if the service ID is not found or other error.
	RestartService(ctx *context.Context, serviceID string) error
}

// System status.
type SystemStatusType int

const (
	SystemInitializedType SystemStatusType = iota
	SystemStartedType
	SystemStoppedType
)

// SystemImpl represents the core system in the application.
type SystemImpl struct {
	SystemInterface
	mutex         sync.RWMutex
	configuration *config.Configuration
	componentReg  component.ComponentRegistrarInterface
	logger        logger.LoggerInterface
	eventBus      event.EventBusInterface
	pluginManager PluginManagerInterface
	status        SystemStatusType
	store         store.MultiStore
}

// NewSystem creates a new instance of the SystemImpl.
func NewSystem(
	logger logger.LoggerInterface,
	eventBus event.EventBusInterface,
	configuration *config.Configuration,
	pluginManager PluginManagerInterface,
	componentReg component.ComponentRegistrarInterface,
	store store.MultiStore) *SystemImpl {
	return &SystemImpl{
		logger:        logger,
		eventBus:      eventBus,
		componentReg:  componentReg,
		configuration: configuration,
		pluginManager: pluginManager,
		status:        SystemStoppedType,
		store:         store,
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
func (s *SystemImpl) Configuration() *config.Configuration {
	return s.configuration
}

// ComponentRegistry returns the component registry.
func (s *SystemImpl) ComponentRegistry() component.ComponentRegistrarInterface {
	return s.componentReg
}

// MultiStore returns the multistore
func (s *SystemImpl) MultiStore() store.MultiStore {
	return s.store
}

// ComponentRegistry returns the component registry.
func (s *SystemImpl) PluginManager() PluginManagerInterface {
	return s.pluginManager
}

// Initialize initializes the system component by executing the initialize operation.
func (s *SystemImpl) Initialize(ctx *context.Context) error {
	// Override this function to customize system initialization

	s.status = SystemInitializedType
	s.pluginManager.Initialize(ctx, s)
	return nil
}

// Start starts the system component along with all registered services.
func (s *SystemImpl) Start(ctx *context.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if s.status != SystemInitializedType {
		return ErrSystemNotInitialized
	}

	if err := s.pluginManager.StartPlugins(ctx); err != nil {
		// Log the error, but continue stopping other services
		s.logger.Log(logger.LevelError, "Error starting plugin:", err)
		return err
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
	components := s.ComponentRegistry().GetComponentsByType(component.ServiceType)

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
func (s *SystemImpl) ExecuteOperation(ctx *context.Context, operationID string, data *SystemOperationInput) (*SystemOperationOutput, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// Retrieve the operation by its ID
	component, err := s.ComponentRegistry().GetComponent(operationID)
	if err != nil {
		return nil, err
	}

	// Check if the component implements Operation interface
	operation, ok := component.(SystemOperationInterface)
	if !ok {
		return nil, fmt.Errorf("failed to execute operation: component %v is not an operation", operation)
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
