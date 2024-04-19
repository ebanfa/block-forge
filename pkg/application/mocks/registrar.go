package mocks

import (
	"github.com/edward1christian/block-forge/pkg/application/component"
	configApi "github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/stretchr/testify/mock"
)

// MockComponentRegistrar is a mock implementation of the ComponentRegistrar interface.
type MockComponentRegistrar struct {
	mock.Mock
}

// RegisterComponent mocks the RegisterComponent method of ComponentRegistrarImpl.
func (m *MockComponentRegistrar) RegisterComponent(config *configApi.ComponentConfig) error {
	args := m.Called(config)
	return args.Error(0)
}

// GetComponent mocks the GetComponent method of ComponentRegistrarImpl.
func (m *MockComponentRegistrar) GetComponent(id string) (component.ComponentInterface, error) {
	args := m.Called(id)
	return args.Get(0).(component.ComponentInterface), args.Error(1)
}

// GetComponentByType mocks the GetComponentByType method of ComponentRegistrarImpl.
func (m *MockComponentRegistrar) GetComponentsByType(componentType component.ComponentType) []component.ComponentInterface {
	args := m.Called(componentType)
	return args.Get(0).([]component.ComponentInterface)
}

// CreateComponent creates and registers a new instance of the component.
func (m *MockComponentRegistrar) CreateComponent(config *configApi.ComponentConfig) (component.ComponentInterface, error) {
	args := m.Called(config)
	return args.Get(0).(component.ComponentInterface), args.Error(1)
}

// RegisterFactory mocks the RegisterFactory method of ComponentRegistrarImpl.
func (m *MockComponentRegistrar) RegisterFactory(id string, factory component.ComponentFactoryInterface) error {
	args := m.Called(id, factory)
	return args.Error(0)
}

// UnregisterComponent mocks the UnregisterComponent method of ComponentRegistrarImpl.
func (m *MockComponentRegistrar) UnregisterComponent(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// UnregisterFactory mocks the UnregisterFactory method of ComponentRegistrarImpl.
func (m *MockComponentRegistrar) UnregisterFactory(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// GetAllComponents mocks the GetAllComponents method of ComponentRegistrarImpl.
func (m *MockComponentRegistrar) GetAllComponents() []component.ComponentInterface {
	args := m.Called()
	return args.Get(0).([]component.ComponentInterface)
}

// GetAllFactories mocks the GetAllFactories method of ComponentRegistrarImpl.
func (m *MockComponentRegistrar) GetAllFactories() []component.ComponentFactoryInterface {
	args := m.Called()
	return args.Get(0).([]component.ComponentFactoryInterface)
}

// GetComponentFactory mocks the GetComponentFactory method of ComponentRegistrarImpl.
func (m *MockComponentRegistrar) GetComponentFactory(id string) (component.ComponentFactoryInterface, error) {
	args := m.Called(id)
	return args.Get(0).(component.ComponentFactoryInterface), args.Error(1)
}
