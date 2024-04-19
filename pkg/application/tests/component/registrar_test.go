package component

import (
	"testing"

	"github.com/edward1christian/block-forge/pkg/application/component"
	"github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	// Import the package containing the interface and its dependencies
)

// TestComponentRegistrar_GetComponentsByType_Success tests the GetComponentsByType function of ComponentRegistrarInterface for success.
func TestComponentRegistrar_GetComponentsByType_Success(t *testing.T) {
	// Create a new instance of ComponentRegistrar
	registrar := component.NewComponentRegistrar()

	// Register a mock component
	mockComponent := new(mocks.MockComponent)
	mockFactory := new(mocks.MockComponentFactory)

	mockComponent.On("Type").Return(component.SystemComponentType)
	mockFactory.On("CreateComponent", mock.Anything).Return(mockComponent, nil)

	registrar.RegisterFactory("mockFactoryID", mockFactory)
	registrar.CreateComponent(&config.ComponentConfig{
		FactoryID: "mockFactoryID",
	})

	// Call GetComponentsByType
	components := registrar.GetComponentsByType(component.SystemComponentType)
	assert.NotNil(t, components)
	assert.Len(t, components, 1)
}

// TestComponentRegistrar_GetComponentsByType_Error tests the GetComponentsByType function of ComponentRegistrarInterface for error.
func TestComponentRegistrar_GetComponentsByType_Error(t *testing.T) {
	// Create a new instance of ComponentRegistrar
	registrar := component.NewComponentRegistrar()

	// Call GetComponentsByType for non-existent type
	components := registrar.GetComponentsByType(20)
	assert.Empty(t, components)
}

// TestComponentRegistrar_GetComponentFactory_Success tests the GetComponentFactory function of ComponentRegistrarInterface for success.
func TestComponentRegistrar_GetComponentFactory_Success(t *testing.T) {
	// Create a new instance of ComponentRegistrar
	registrar := component.NewComponentRegistrar()

	// Register a mock factory
	mockFactory := new(mocks.MockComponentFactory)
	registrar.RegisterFactory("mockFactoryID", mockFactory)

	// Call GetComponentFactory
	factory, err := registrar.GetComponentFactory("mockFactoryID")
	assert.NoError(t, err)
	assert.NotNil(t, factory)
}

// TestComponentRegistrar_GetComponentFactory_Error tests the GetComponentFactory function of ComponentRegistrarInterface for error.
func TestComponentRegistrar_GetComponentFactory_Error(t *testing.T) {
	// Create a new instance of ComponentRegistrar
	registrar := component.NewComponentRegistrar()

	// Call GetComponentFactory for non-existent factory ID
	factory, err := registrar.GetComponentFactory("nonExistentFactory")
	assert.Error(t, err)
	assert.Nil(t, factory)
}

// TestComponentRegistrar_GetAllComponents tests the GetAllComponents function of ComponentRegistrarInterface.
func TestComponentRegistrar_GetAllComponents(t *testing.T) {
	// Create a new instance of ComponentRegistrar
	registrar := component.NewComponentRegistrar()

	// Call GetAllComponents
	allComponents := registrar.GetAllComponents()
	assert.NotNil(t, allComponents)
	assert.Empty(t, allComponents)
}

// TestComponentRegistrar_GetAllFactories tests the GetAllFactories function of ComponentRegistrarInterface.
func TestComponentRegistrar_GetAllFactories(t *testing.T) {
	// Create a new instance of ComponentRegistrar
	registrar := component.NewComponentRegistrar()

	// Call GetAllFactories
	allFactories := registrar.GetAllFactories()
	assert.NotNil(t, allFactories)
	assert.Empty(t, allFactories)
}

// TestComponentRegistrar_RegisterFactory_Success tests the RegisterFactory function of ComponentRegistrarInterface for success.
func TestComponentRegistrar_RegisterFactory_Success(t *testing.T) {
	// Create a new instance of ComponentRegistrar
	registrar := component.NewComponentRegistrar()

	// Mock ComponentFactory
	mockFactory := new(mocks.MockComponentFactory)

	// Call RegisterFactory
	err := registrar.RegisterFactory("mockFactoryID", mockFactory)
	assert.NoError(t, err)
}

// TestComponentRegistrar_RegisterFactory_Error tests the RegisterFactory function of ComponentRegistrarInterface for error.
func TestComponentRegistrar_RegisterFactory_Error(t *testing.T) {
	// Create a new instance of ComponentRegistrar
	registrar := component.NewComponentRegistrar()

	// Attempt to register the same factory ID twice
	mockFactory := new(mocks.MockComponentFactory)
	err := registrar.RegisterFactory("mockFactoryID", mockFactory)
	assert.NoError(t, err)

	err = registrar.RegisterFactory("mockFactoryID", mockFactory)
	assert.Error(t, err)
}

// TestComponentRegistrar_UnregisterFactory_Success tests the UnregisterFactory function of ComponentRegistrarInterface for success.
func TestComponentRegistrar_UnregisterFactory_Success(t *testing.T) {
	// Create a new instance of ComponentRegistrar
	registrar := component.NewComponentRegistrar()

	// Mock ComponentFactory
	mockFactory := new(mocks.MockComponentFactory)

	// Register the factory
	err := registrar.RegisterFactory("mockFactoryID", mockFactory)
	assert.NoError(t, err)

	// Call UnregisterFactory
	err = registrar.UnregisterFactory("mockFactoryID")
	assert.NoError(t, err)
}

// TestComponentRegistrar_UnregisterFactory_Error tests the UnregisterFactory function of ComponentRegistrarInterface for error.
func TestComponentRegistrar_UnregisterFactory_Error(t *testing.T) {
	// Create a new instance of ComponentRegistrar
	registrar := component.NewComponentRegistrar()

	// Attempt to unregister non-existent factory
	err := registrar.UnregisterFactory("nonExistentFactory")
	assert.Error(t, err)
}
