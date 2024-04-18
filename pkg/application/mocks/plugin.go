package mocks

import (
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/components"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/stretchr/testify/mock"
)

// MockPlugin is a mock implementation of the Plugin interface.
type MockPlugin struct {
	mock.Mock
}

// ID returns the unique identifier of the plugin.
func (m *MockPlugin) ID() string {
	args := m.Called()
	return args.String(0)
}

// Name returns the name of the plugin.
func (m *MockPlugin) Name() string {
	args := m.Called()
	return args.String(0)
}

// Description returns the description of the plugin.
func (m *MockPlugin) Description() string {
	args := m.Called()
	return args.String(0)
}

// Description returns the description of the component.
func (m *MockPlugin) Type() components.ComponentType {
	args := m.Called()
	return args.Get(0).(components.ComponentType)
}

// Initialize initializes the plugin.
func (m *MockPlugin) Initialize(ctx *context.Context, system system.SystemInterface) error {
	args := m.Called(ctx, system)
	return args.Error(0)
}

// RegisterResources registers resources into the system.
func (m *MockPlugin) RegisterResources(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// Start starts the plugin.
func (m *MockPlugin) Start(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// Stop stops the plugin.
func (m *MockPlugin) Stop(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}
