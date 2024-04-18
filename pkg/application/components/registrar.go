package components

import (
	"fmt"
	"sync"

	configApi "github.com/edward1christian/block-forge/pkg/application/config"
)

// ComponentRegistrar defines the registry functionality for
type ComponentRegistrar interface {
	// RegisterComponent registers a component factory with the given ID.
	// Returns an error if the registration fails.
	RegisterComponent(config *configApi.ComponentConfig) error

	// GetComponent retrieves the component with the specified ID.
	// Returns the component and an error if the component ID is not found or other error.
	GetComponent(id string) (ComponentInterface, error)

	// GetComponentByType retrieves components of the specified type.
	// Returns a list of components and an error if the type is not found or other error.
	GetComponentByType(componentType ComponentType) []ComponentInterface

	// RegisterFactory registers a factory with the given ID.
	// Returns an error if the registration fails.
	RegisterFactory(id string, factory ComponentFactoryInterface) error

	// UnregisterComponent unregisters the component with the specified ID.
	// Returns an error if the component ID is not found or other error.
	UnregisterComponent(id string) error

	// UnregisterFactory unregisters the factory with the specified ID.
	// Returns an error if the factory ID is not found or other error.
	UnregisterFactory(id string) error

	// GetAllComponents returns a list of all registered
	GetAllComponents() []ComponentInterface

	// GetAllFactories returns a list of all registered component factories.
	GetAllFactories() []ComponentFactoryInterface

	// GetComponentFactory retrieves the factory for the specified component ID.
	// Returns the component factory and an error if the component factory ID is not found or other error.
	GetComponentFactory(id string) (ComponentFactoryInterface, error)
}

// ComponentRegistrarImpl implements the ComponentRegistrar interface.
type ComponentRegistrarImpl struct {
	ComponentRegistrar
	mutex              sync.RWMutex
	components         map[string]ComponentInterface
	factories          map[string]ComponentFactoryInterface
	componentTypeIndex map[ComponentType][]ComponentInterface
}

// NewComponentRegistrar creates a new instance of ComponentRegistrarImpl.
func NewComponentRegistrar() *ComponentRegistrarImpl {
	return &ComponentRegistrarImpl{
		components:         make(map[string]ComponentInterface),
		factories:          make(map[string]ComponentFactoryInterface),
		componentTypeIndex: make(map[ComponentType][]ComponentInterface),
	}
}

// RegisterComponent registers a component factory with the given ID.
// It creates a component using the specified factory and registers it.
func (r *ComponentRegistrarImpl) RegisterComponent(config *configApi.ComponentConfig) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	id := config.ID
	factoryName := config.FactoryName

	// Check if component with the same ID already exists
	if _, exists := r.components[id]; exists {
		return ErrComponentAlreadyExist
	}

	// Retrieve the factory by name
	factory, exists := r.factories[factoryName]
	if !exists {
		return ErrFactoryNotFound
	}

	// Create the component using the factory
	component, err := factory.CreateComponent(config)
	if err != nil {
		return fmt.Errorf("failed to create component: %w", err)
	}

	// Register the component
	r.components[id] = component

	// Update the componentTypeIndex
	componentType := component.Type()
	r.componentTypeIndex[componentType] = append(r.componentTypeIndex[componentType], component)

	return nil
}

// GetComponent retrieves the component with the specified ID.
// Returns an error if the component is not found.
func (r *ComponentRegistrarImpl) GetComponent(id string) (ComponentInterface, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	component, exists := r.components[id]
	if !exists {
		return nil, ErrComponentNotFound
	}

	return component, nil
}

// GetComponentByType retrieves components of the specified type.
// Returns an error if the component type is not found.
func (r *ComponentRegistrarImpl) GetComponentByType(componentType ComponentType) []ComponentInterface {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	components, exists := r.componentTypeIndex[componentType]
	if !exists {
		return []ComponentInterface{}
	}

	return components
}

// RegisterFactory registers a factory with the given ID.
// Returns an error if a factory with the same ID already exists.
func (r *ComponentRegistrarImpl) RegisterFactory(id string, factory ComponentFactoryInterface) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.factories[id]; exists {
		return ErrComponentFactoryAlreadyExists
	}

	r.factories[id] = factory
	return nil
}

// UnregisterComponent unregisters the component with the specified ID.
// Returns an error if the component is not found.
func (r *ComponentRegistrarImpl) UnregisterComponent(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.components[id]; !exists {
		return ErrComponentNotFound
	}

	delete(r.components, id)
	return nil
}

// UnregisterFactory unregisters the factory with the specified ID.
// Returns an error if the factory is not found.
func (r *ComponentRegistrarImpl) UnregisterFactory(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.factories[id]; !exists {
		return ErrFactoryNotFound
	}

	delete(r.factories, id)
	return nil
}

// GetAllComponents returns a list of all registered
func (r *ComponentRegistrarImpl) GetAllComponents() []ComponentInterface {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	components := make([]ComponentInterface, 0, len(r.components))
	for _, component := range r.components {
		components = append(components, component)
	}

	return components
}

// GetAllFactories returns a list of all registered component factories.
func (r *ComponentRegistrarImpl) GetAllFactories() []ComponentFactoryInterface {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	factories := make([]ComponentFactoryInterface, 0, len(r.factories))
	for _, factory := range r.factories {
		factories = append(factories, factory)
	}

	return factories
}

// GetComponentFactory retrieves the factory for the specified component ID.
// Returns the component factory and an error if the component factory ID is not found or other error.
func (r *ComponentRegistrarImpl) GetComponentFactory(id string) (ComponentFactoryInterface, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	factory, exists := r.factories[id]
	if !exists {
		return nil, ErrFactoryNotFound
	}

	return factory, nil
}
