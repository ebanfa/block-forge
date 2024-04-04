package components

// ComponentRegistrar defines the registry functionality for
type ComponentRegistrar interface {
	// RegisterComponent registers a component factory with the given ID.
	// Returns an error if the registration fails.
	RegisterComponent(config *ComponentConfig) error

	// GetComponent retrieves the component with the specified ID.
	// Returns the component and an error if the component ID is not found or other error.
	GetComponent(id string) (ComponentInterface, error)

	// GetComponentByType retrieves components of the specified type.
	// Returns a list of components and an error if the type is not found or other error.
	GetComponentByType(componentType ComponentType) ([]ComponentInterface, error)

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
