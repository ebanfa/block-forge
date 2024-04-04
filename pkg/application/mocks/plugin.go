package mocks

import (
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/stretchr/testify/mock"
)

// MockPlugin is a mock implementation of the Plugin interface.
type MockPlugin struct {
	mock.Mock
}

// ID implements application.Operations.
func (m *MockPlugin) ID() string {
	args := m.Called()
	return args.String(0)
}

// Name implements application.Operations.
func (m *MockPlugin) Name() string {
	args := m.Called()
	return args.String(0)
}

// Description implements application.Operations.
func (m *MockPlugin) Description() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockPlugin) Start(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockPlugin) Stop(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// Initialize initializes the plugin.
func (m *MockPlugin) Initialize(ctx *context.Context, system system.SystemInterface) error {
	args := m.Called(ctx, system)
	return args.Error(0)
}

// RegisterEventHandlers registers event handlers provided by the plugin.
func (m *MockPlugin) RegisterEventHandlers() error {
	args := m.Called()
	return args.Error(0)
}

// RegisterOperations registers additional operations provided by the plugin.
func (m *MockPlugin) RegisterOperations() error {
	args := m.Called()
	return args.Error(0)
}

// Configure configures the plugin with the provided configuration.
func (m *MockPlugin) Configure(config interface{}) error {
	args := m.Called(config)
	return args.Error(0)
}

// Dependencies returns a list of dependencies required by the plugin.
func (m *MockPlugin) Dependencies() []string {
	args := m.Called()
	return args.Get(0).([]string)
}

// OnInitialize is called during plugin initialization.
func (m *MockPlugin) OnInitialize() error {
	args := m.Called()
	return args.Error(0)
}

// OnStart is called when the plugin starts.
func (m *MockPlugin) OnStart() error {
	args := m.Called()
	return args.Error(0)
}

// OnStop is called when the plugin stops.
func (m *MockPlugin) OnStop() error {
	args := m.Called()
	return args.Error(0)
}

// LogError logs an error that occurred during plugin execution.
func (m *MockPlugin) LogError(err error) {
	m.Called(err)
}
