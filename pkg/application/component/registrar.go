package component

import (
	"fmt"
	"sync"

	configApi "github.com/edward1christian/block-forge/pkg/application/config"
)

// ComponentRegistrarInterface defines the registry functionality for components and factories.
type ComponentRegistrarInterface interface {

	// GetComponentsByType retrieves components of the specified type.
	// It returns a list of components and an error if the type is not found or other error.
	GetComponentsByType(componentType ComponentType) []ComponentInterface

	// GetComponent retrieves the component with the specified ID.
	// It returns the component and an error if the component ID is not found or other error.
	GetComponent(id string) (ComponentInterface, error)

	// GetComponentFactory retrieves the factory for the specified component ID.
	// It returns the component factory and an error if the component factory ID is not found or other error.
	GetComponentFactory(id string) (ComponentFactoryInterface, error)

	// GetAllComponents returns a list of all registered components.
	GetAllComponents() []ComponentInterface

	// GetAllFactories returns a list of all registered component factories.
	GetAllFactories() []ComponentFactoryInterface

	// CreateComponent creates and registers a new instance of the component.
	// Returns the created component and an error if the creation fails.
	CreateComponent(config *configApi.ComponentConfig) (ComponentInterface, error)

	// RegisterFactory registers a factory with the given ID.
	// It returns an error if the registration fails.
	RegisterFactory(id string, factory ComponentFactoryInterface) error

	// UnregisterFactory unregisters the factory with the specified ID.
	// It returns an error if the factory ID is not found or other error.
	UnregisterFactory(id string) error

	// UnregisterComponent unregisters the component with the specified ID.
	// It returns an error if the component ID is not found or other error.
	UnregisterComponent(id string) error
}

// ComponentRegistrar defines the registry functionality for components and factories.
type ComponentRegistrar struct {
	factoriesMutex  sync.RWMutex
	componentsMutex sync.RWMutex
	factories       map[string]ComponentFactoryInterface
	components      map[string]ComponentInterface
}

// NewComponentRegistrar creates a new instance of ComponentRegistrar.
func NewComponentRegistrar() *ComponentRegistrar {
	return &ComponentRegistrar{
		factories:  make(map[string]ComponentFactoryInterface),
		components: make(map[string]ComponentInterface),
	}
}

// GetComponent retrieves the component with the specified ID.
func (cr *ComponentRegistrar) GetComponent(id string) (ComponentInterface, error) {
	cr.componentsMutex.RLock()
	defer cr.componentsMutex.RUnlock()

	// Check if the component exists
	component, exists := cr.components[id]
	if !exists {
		return nil, fmt.Errorf("component with ID %s not found", id)
	}
	return component, nil
}

// GetComponentsByType retrieves components of the specified type.
func (cr *ComponentRegistrar) GetComponentsByType(componentType ComponentType) []ComponentInterface {
	cr.componentsMutex.RLock()
	defer cr.componentsMutex.RUnlock()

	components := []ComponentInterface{}
	for _, component := range cr.components {
		if component.Type() == componentType {
			components = append(components, component)
		}
	}

	return components
}

// GetAllComponents returns a list of all registered components.
func (cr *ComponentRegistrar) GetAllComponents() []ComponentInterface {
	cr.componentsMutex.RLock()
	defer cr.componentsMutex.RUnlock()

	allComponents := []ComponentInterface{}
	for _, component := range cr.components {
		allComponents = append(allComponents, component)
	}
	return allComponents
}

// GetAllFactories returns a list of all registered component factories.
func (cr *ComponentRegistrar) GetAllFactories() []ComponentFactoryInterface {
	cr.factoriesMutex.RLock()
	defer cr.factoriesMutex.RUnlock()

	allFactories := []ComponentFactoryInterface{}
	for _, factory := range cr.factories {
		allFactories = append(allFactories, factory)
	}
	return allFactories
}

// GetComponentFactory retrieves the factory for the specified component ID.
func (cr *ComponentRegistrar) GetComponentFactory(id string) (ComponentFactoryInterface, error) {
	cr.factoriesMutex.RLock()
	defer cr.factoriesMutex.RUnlock()

	// Check if the factory exists
	factory, exists := cr.factories[id]
	if !exists {
		return nil, fmt.Errorf("factory with ID %s not found", id)
	}
	return factory, nil
}

// CreateComponent creates and registers a new instance of the component.
func (cr *ComponentRegistrar) CreateComponent(config *configApi.ComponentConfig) (ComponentInterface, error) {
	cr.factoriesMutex.RLock()
	defer cr.factoriesMutex.RUnlock()

	// Check if the factory exists
	factory, exists := cr.factories[config.FactoryID]
	if !exists {
		return nil, fmt.Errorf("factory with ID %s not found", config.FactoryID)
	}

	// Use the factory to create the component
	component, err := factory.CreateComponent(config)
	if err != nil {
		return nil, err
	}

	// Register the component
	cr.componentsMutex.Lock()
	defer cr.componentsMutex.Unlock()
	cr.components[config.ID] = component

	return component, nil
}

// RegisterFactory registers a factory with the given ID.
func (cr *ComponentRegistrar) RegisterFactory(id string, factory ComponentFactoryInterface) error {
	cr.factoriesMutex.Lock()
	defer cr.factoriesMutex.Unlock()

	// Check if the factory already exists
	if _, exists := cr.factories[id]; exists {
		return fmt.Errorf("factory with ID %s already exists", id)
	}

	// Register the factory
	cr.factories[id] = factory
	return nil
}

// UnregisterComponent unregisters the component with the specified ID.
func (cr *ComponentRegistrar) UnregisterComponent(id string) error {
	cr.componentsMutex.Lock()
	defer cr.componentsMutex.Unlock()

	// Check if the component exists
	if _, exists := cr.components[id]; !exists {
		return fmt.Errorf("component with ID %s not found", id)
	}

	// Unregister the component
	delete(cr.components, id)
	return nil
}

// UnregisterFactory unregisters the factory with the specified ID.
func (cr *ComponentRegistrar) UnregisterFactory(id string) error {
	cr.factoriesMutex.Lock()
	defer cr.factoriesMutex.Unlock()

	// Check if the factory exists
	if _, exists := cr.factories[id]; !exists {
		return fmt.Errorf("factory with ID %s not found", id)
	}

	// Unregister the factory
	delete(cr.factories, id)

	// Remove components created from this factory
	cr.componentsMutex.Lock()
	defer cr.componentsMutex.Unlock()
	for key := range cr.components {
		if key == id {
			delete(cr.components, key)
		}
	}
	return nil
}
