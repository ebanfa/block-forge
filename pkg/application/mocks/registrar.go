package mocks

import (
	"github.com/edward1christian/block-forge/pkg/application/components"
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
func (m *MockComponentRegistrar) GetComponent(id string) (components.ComponentInterface, error) {
	args := m.Called(id)
	return args.Get(0).(components.ComponentInterface), args.Error(1)
}

// GetComponentByType mocks the GetComponentByType method of ComponentRegistrarImpl.
func (m *MockComponentRegistrar) GetComponentByType(componentType components.ComponentType) []components.ComponentInterface {
	args := m.Called(componentType)
	return args.Get(0).([]components.ComponentInterface)
}

// RegisterFactory mocks the RegisterFactory method of ComponentRegistrarImpl.
func (m *MockComponentRegistrar) RegisterFactory(id string, factory components.ComponentFactoryInterface) error {
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
func (m *MockComponentRegistrar) GetAllComponents() []components.ComponentInterface {
	args := m.Called()
	return args.Get(0).([]components.ComponentInterface)
}

// GetAllFactories mocks the GetAllFactories method of ComponentRegistrarImpl.
func (m *MockComponentRegistrar) GetAllFactories() []components.ComponentFactoryInterface {
	args := m.Called()
	return args.Get(0).([]components.ComponentFactoryInterface)
}

// GetComponentFactory mocks the GetComponentFactory method of ComponentRegistrarImpl.
func (m *MockComponentRegistrar) GetComponentFactory(id string) (components.ComponentFactoryInterface, error) {
	args := m.Called(id)
	return args.Get(0).(components.ComponentFactoryInterface), args.Error(1)
}
