package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/edward1christian/block-forge/pkg/application/component"
	configApi "github.com/edward1christian/block-forge/pkg/application/config"
)

// MockComponentFactory represents a mock for the ComponentFactory type.
type MockComponentFactory struct {
	mock.Mock
}

func NewMockComponentFactory() *MockComponentFactory {
	return &MockComponentFactory{}
}

// CreateComponent mocks the CreateComponent method of the ComponentFactory type.
func (m *MockComponentFactory) CreateComponent(config *configApi.ComponentConfig) (component.ComponentInterface, error) {
	args := m.Called(config)
	if component, ok := args.Get(0).(component.ComponentInterface); ok {
		return component, args.Error(1)
	}
	return nil, args.Error(1)
}
