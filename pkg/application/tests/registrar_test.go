package tests

import (
	"errors"
	"testing"

	"github.com/edward1christian/block-forge/pkg/application/components"
	componentsApi "github.com/edward1christian/block-forge/pkg/application/components"
	configApi "github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestRegisterComponent_Success tests the successful registration of a component.
func TestRegisterComponent_Success(t *testing.T) {
	// Initialize the registrar and mocks
	registrar := componentsApi.NewComponentRegistrar()
	mockFactory := &mocks.MockComponentFactory{}
	mockComponent := &mocks.MockComponent{}

	// Mock the behavior of the component and factory
	mockComponent.On("Type").Return(components.OperationType)
	mockFactory.On("CreateComponent", mock.Anything).Return(mockComponent, nil)

	config := &configApi.ComponentConfig{
		ID:          "component_id",
		FactoryName: "factory_name",
	}

	// Register the factory and component
	err := registrar.RegisterFactory("factory_name", mockFactory)
	assert.NoError(t, err)

	err = registrar.RegisterComponent(config)

	assert.NoError(t, err)
	assert.Len(t, registrar.GetAllComponents(), 1)
}

// TestRegisterComponent_Failure_FactoryNotFound tests failure when the factory is not found.
func TestRegisterComponent_Failure_FactoryNotFound(t *testing.T) {
	registrar := componentsApi.NewComponentRegistrar()

	config := &configApi.ComponentConfig{
		ID:          "component_id",
		FactoryName: "nonexistent_factory",
	}

	// Attempt to register a component with a nonexistent factory
	err := registrar.RegisterComponent(config)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, componentsApi.ErrFactoryNotFound))
	assert.Len(t, registrar.GetAllComponents(), 0)
}

// TestGetComponent_Success tests successful retrieval of a component.
func TestGetComponent_Success(t *testing.T) {
	registrar := componentsApi.NewComponentRegistrar()
	mockComponent := &mocks.MockComponent{}
	mockFactory := &mocks.MockComponentFactory{}

	// Mock the behavior of the component and factory
	mockComponent.On("Type").Return(components.OperationType)
	mockFactory.On("CreateComponent", mock.Anything).Return(mockComponent, nil)

	// Register the factory and component
	err := registrar.RegisterFactory("factory_name", mockFactory)
	assert.NoError(t, err)

	err = registrar.RegisterComponent(&configApi.ComponentConfig{
		ID:          "component_id",
		FactoryName: "factory_name",
	})
	assert.NoError(t, err)

	// Retrieve the registered component
	component, err := registrar.GetComponent("component_id")

	assert.NoError(t, err)
	assert.Equal(t, mockComponent, component)
}

// TestGetComponent_Failure_NotFound tests failure when the component is not found.
func TestGetComponent_Failure_NotFound(t *testing.T) {
	registrar := componentsApi.NewComponentRegistrar()

	// Attempt to retrieve a nonexistent component
	component, err := registrar.GetComponent("nonexistent_component")

	assert.Error(t, err)
	assert.True(t, errors.Is(err, componentsApi.ErrComponentNotFound))
	assert.Nil(t, component)
}

// TestGetComponentByType_Success tests successful retrieval of components by type.
func TestGetComponentByType_Success(t *testing.T) {
	registrar := componentsApi.NewComponentRegistrar()
	mockComponent := &mocks.MockComponent{}
	mockFactory := &mocks.MockComponentFactory{}

	// Mock the behavior of the component and factory
	mockComponent.On("Type").Return(components.OperationType)
	mockFactory.On("CreateComponent", mock.Anything).Return(mockComponent, nil)

	// Register the factory and component
	err := registrar.RegisterFactory("factory_name", mockFactory)
	assert.NoError(t, err)

	err = registrar.RegisterComponent(&configApi.ComponentConfig{
		ID:          "component_id",
		FactoryName: "factory_name",
	})
	assert.NoError(t, err)

	// Retrieve components by type
	components := registrar.GetComponentByType(componentsApi.OperationType)

	assert.Len(t, components, 1)
	assert.Equal(t, mockComponent, components[0])
}

// TestRegisterFactory_Success tests successful registration of a factory.
func TestRegisterFactory_Success(t *testing.T) {
	registrar := componentsApi.NewComponentRegistrar()
	mockFactory := &mocks.MockComponentFactory{}

	// Register a factory
	err := registrar.RegisterFactory("factory_name", mockFactory)

	assert.NoError(t, err)
	assert.Len(t, registrar.GetAllFactories(), 1)
}

// TestRegisterFactory_Failure_AlreadyExists tests failure when the factory already exists.
func TestRegisterFactory_Failure_AlreadyExists(t *testing.T) {
	registrar := componentsApi.NewComponentRegistrar()
	mockFactory := &mocks.MockComponentFactory{}

	// Attempt to register the same factory twice
	err := registrar.RegisterFactory("factory_name", mockFactory)
	assert.NoError(t, err)

	err = registrar.RegisterFactory("factory_name", mockFactory)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, componentsApi.ErrComponentFactoryAlreadyExists))
}

// TestUnregisterComponent_Success tests successful unregistering of a component.
func TestUnregisterComponent_Success(t *testing.T) {
	registrar := componentsApi.NewComponentRegistrar()
	mockComponent := &mocks.MockComponent{}
	mockFactory := &mocks.MockComponentFactory{}

	// Mock the behavior of the component and factory
	mockComponent.On("Type").Return(components.OperationType)
	mockFactory.On("CreateComponent", mock.Anything).Return(mockComponent, nil)

	// Register the factory and component
	err := registrar.RegisterFactory("factory_name", mockFactory)
	assert.NoError(t, err)

	err = registrar.RegisterComponent(&configApi.ComponentConfig{
		ID:          "component_id",
		FactoryName: "factory_name",
	})
	assert.NoError(t, err)

	// Unregister the component
	err = registrar.UnregisterComponent("component_id")

	assert.NoError(t, err)
	assert.Len(t, registrar.GetAllComponents(), 0)
}

// TestUnregisterComponent_Failure_NotFound tests failure when the component to unregister is not found.
func TestUnregisterComponent_Failure_NotFound(t *testing.T) {
	registrar := componentsApi.NewComponentRegistrar()

	// Attempt to unregister a nonexistent component
	err := registrar.UnregisterComponent("nonexistent_component")

	assert.Error(t, err)
	assert.True(t, errors.Is(err, componentsApi.ErrComponentNotFound))
}

// TestUnregisterFactory_Success tests successful unregistering of a factory.
func TestUnregisterFactory_Success(t *testing.T) {
	registrar := componentsApi.NewComponentRegistrar()
	mockFactory := &mocks.MockComponentFactory{}

	// Register a factory
	registrar.RegisterFactory("factory_name", mockFactory)

	// Unregister the factory
	err := registrar.UnregisterFactory("factory_name")

	assert.NoError(t, err)
	assert.Len(t, registrar.GetAllFactories(), 0)
}

// TestUnregisterFactory_Failure_NotFound tests failure when the factory to unregister is not found.
func TestUnregisterFactory_Failure_NotFound(t *testing.T) {
	registrar := componentsApi.NewComponentRegistrar()

	// Attempt to unregister a nonexistent factory
	err := registrar.UnregisterFactory("nonexistent_factory")

	assert.Error(t, err)
	assert.True(t, errors.Is(err, componentsApi.ErrFactoryNotFound))
}

// TestGetAllComponents tests the retrieval of all components.
func TestGetAllComponents(t *testing.T) {
	registrar := componentsApi.NewComponentRegistrar()
	mockComponent := &mocks.MockComponent{}
	mockFactory := &mocks.MockComponentFactory{}

	// Mock the behavior of the component and factory
	mockComponent.On("Type").Return(components.OperationType)
	mockFactory.On("CreateComponent", mock.Anything).Return(mockComponent, nil)

	// Register the factory and component
	err := registrar.RegisterFactory("factory_name", mockFactory)
	assert.NoError(t, err)

	err = registrar.RegisterComponent(&configApi.ComponentConfig{
		ID:          "component_id",
		FactoryName: "factory_name",
	})
	assert.NoError(t, err)

	// Retrieve all registered components
	components := registrar.GetAllComponents()

	assert.Len(t, components, 1)
	assert.Equal(t, mockComponent, components[0])
}

// TestGetAllFactories tests the retrieval of all factories.
func TestGetAllFactories(t *testing.T) {
	registrar := componentsApi.NewComponentRegistrar()
	mockFactory := &mocks.MockComponentFactory{}

	// Register a factory
	registrar.RegisterFactory("factory_name", mockFactory)

	// Retrieve all registered factories
	factories := registrar.GetAllFactories()

	assert.Len(t, factories, 1)
	assert.Equal(t, mockFactory, factories[0])
}

// TestGetComponentFactory_Success tests successful retrieval of a component factory.
func TestGetComponentFactory_Success(t *testing.T) {
	registrar := componentsApi.NewComponentRegistrar()
	mockFactory := &mocks.MockComponentFactory{}

	// Register a factory
	registrar.RegisterFactory("factory_name", mockFactory)

	// Retrieve a registered factory by ID
	factory, err := registrar.GetComponentFactory("factory_name")

	assert.NoError(t, err)
	assert.Equal(t, mockFactory, factory)
}

// TestGetComponentFactory_Failure_NotFound tests failure when the component factory is not found.
func TestGetComponentFactory_Failure_NotFound(t *testing.T) {
	registrar := componentsApi.NewComponentRegistrar()

	// Attempt to retrieve a nonexistent factory
	factory, err := registrar.GetComponentFactory("nonexistent_factory")

	assert.Error(t, err)
	assert.True(t, errors.Is(err, componentsApi.ErrFactoryNotFound))
	assert.Nil(t, factory)
}
