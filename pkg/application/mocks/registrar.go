package mocks

import (
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/component"
	"github.com/stretchr/testify/mock"
)

// MockComponentRegistrar is a mock implementation of ComponentRegistrarInterface.
type MockComponentRegistrar struct {
	mock.Mock
}

// GetComponentsByType mocks the GetComponentsByType method of ComponentRegistrarInterface.
func (m *MockComponentRegistrar) GetComponentsByType(componentType component.ComponentType) []component.ComponentInterface {
	args := m.Called(componentType)
	return args.Get(0).([]component.ComponentInterface)
}

// GetComponent mocks the GetComponent method of ComponentRegistrarInterface.
func (m *MockComponentRegistrar) GetComponent(id string) (component.ComponentInterface, error) {
	args := m.Called(id)
	return args.Get(0).(component.ComponentInterface), args.Error(1)
}

// GetAllComponents mocks the GetAllComponents method of ComponentRegistrarInterface.
func (m *MockComponentRegistrar) GetAllComponents() []component.ComponentInterface {
	args := m.Called()
	return args.Get(0).([]component.ComponentInterface)
}

// GetFactory mocks the GetFactory method of ComponentRegistrarInterface.
func (m *MockComponentRegistrar) GetFactory(id string) (component.ComponentFactoryInterface, error) {
	args := m.Called(id)
	return args.Get(0).(component.ComponentFactoryInterface), args.Error(1)
}

// RegisterFactory mocks the RegisterFactory method of ComponentRegistrarInterface.
func (m *MockComponentRegistrar) RegisterFactory(ctx *context.Context, info *component.FactoryRegistrationInfo) error {
	args := m.Called(ctx, info)
	return args.Error(0)
}

// UnregisterFactory mocks the UnregisterFactory method of ComponentRegistrarInterface.
func (m *MockComponentRegistrar) UnregisterFactory(ctx *context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// CreateComponent mocks the CreateComponent method of ComponentRegistrarInterface.
func (m *MockComponentRegistrar) CreateComponent(ctx *context.Context, info *component.ComponentCreationInfo) (component.ComponentInterface, error) {
	args := m.Called(ctx, info)
	return args.Get(0).(component.ComponentInterface), args.Error(1)
}

// RemoveComponent mocks the RemoveComponent method of ComponentRegistrarInterface.
func (m *MockComponentRegistrar) RemoveComponent(ctx *context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
