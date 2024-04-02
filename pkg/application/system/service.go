package system

import (
	"fmt"
	"sync"

	"github.com/edward1christian/block-forge/pkg/application/context"
)

// ServiceManagerImpl implements the ServiceManager interface.
type ServiceManagerImpl struct {
	system   System
	services map[string]SystemService // Map to store registered services
	mutex    sync.RWMutex             // Mutex for thread-safe access to the services map
}

// NewServiceManager creates a new instance of ServiceManagerImpl.
func NewServiceManager() *ServiceManagerImpl {
	return &ServiceManagerImpl{
		services: make(map[string]SystemService),
	}
}

// RegisterService registers a SystemService with the given ID.
func (m *ServiceManagerImpl) RegisterService(serviceID string, service SystemService) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// Check if the service ID is already registered
	if _, exists := m.services[serviceID]; exists {
		return fmt.Errorf("service with ID '%s' already registered", serviceID)
	}

	// Register the service
	m.services[serviceID] = service
	return nil
}

// UnregisterService unregisters a SystemService with the given ID.
func (m *ServiceManagerImpl) UnregisterService(serviceID string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// Check if the service ID exists
	if _, exists := m.services[serviceID]; !exists {
		return fmt.Errorf("service with ID '%s' not found", serviceID)
	}

	// Unregister the service
	delete(m.services, serviceID)
	return nil
}

// Start starts the service manager and all registered services.
func (m *ServiceManagerImpl) Start(ctx *context.Context) error {
	// Start all registered services
	for _, service := range m.services {
		if err := service.Start(ctx); err != nil {
			return err
		}
	}
	return nil
}

// Stop stops the service manager and all registered services.
func (m *ServiceManagerImpl) Stop(ctx *context.Context) error {
	// Stop all registered services
	for _, service := range m.services {
		if err := service.Stop(ctx); err != nil {
			return err
		}
	}
	return nil
}

// Initialize initializes the service manager with a reference to the system instance.
func (m *ServiceManagerImpl) Initialize(ctx *context.Context, system System) error {
	m.system = system
	// Additional initialization logic for the ServiceManager can be added here if needed.
	return nil
}
