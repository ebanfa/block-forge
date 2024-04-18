package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/edward1christian/block-forge/pkg/application/components"
	configApi "github.com/edward1christian/block-forge/pkg/application/config"
)

// MockComponentFactory represents a mock for the ComponentFactory type.
type MockComponentFactory struct {
	components.ComponentFactoryInterface
	mock.Mock
}

func NewMockComponentFactory() *MockComponentFactory {
	return &MockComponentFactory{}
}

// CreateComponent mocks the CreateComponent method of the ComponentFactory type.
func (m *MockComponentFactory) CreateComponent(config *configApi.ComponentConfig) (components.ComponentInterface, error) {
	args := m.Called(config)
	if component, ok := args.Get(0).(components.ComponentInterface); ok {
		return component, args.Error(1)
	}
	return nil, args.Error(1)
}
