package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// MockComponentFactory represents a mock for the ComponentFactory type.
type MockComponentFactory struct {
	mock.Mock
}

// CreateComponent mocks the CreateComponent method of the ComponentFactory type.
func (m *MockComponentFactory) CreateComponent(ctx *context.Context, config *system.ComponentConfig) (system.Component, error) {
	args := m.Called(ctx, config)
	if component, ok := args.Get(0).(system.Component); ok {
		return component, args.Error(1)
	}
	return nil, args.Error(1)
}
