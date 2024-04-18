package mocks

import (
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/common/event"
	"github.com/edward1christian/block-forge/pkg/application/common/logger"
	"github.com/edward1christian/block-forge/pkg/application/components"
	configApi "github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/stretchr/testify/mock"
)

// MockSystem is a mock implementation of the SystemImpl struct.
type MockSystem struct {
	mock.Mock
}

// NewMockSystemImpl creates a new instance of MockSystem.
func NewMockSystemImpl() *MockSystem {
	return &MockSystem{}
}

// Logger provides a mock implementation of the Logger method.
func (m *MockSystem) Logger() logger.LoggerInterface {
	args := m.Called()
	return args.Get(0).(logger.LoggerInterface)
}

// EventBus provides a mock implementation of the EventBus method.
func (m *MockSystem) EventBus() event.EventBusInterface {
	args := m.Called()
	return args.Get(0).(event.EventBusInterface)
}

// Configuration provides a mock implementation of the Configuration method.
func (m *MockSystem) Configuration() *configApi.Configuration {
	args := m.Called()
	return args.Get(0).(*configApi.Configuration)
}

// ComponentRegistry provides a mock implementation of the ComponentRegistry method.
func (m *MockSystem) ComponentRegistry() components.ComponentRegistrar {
	args := m.Called()
	return args.Get(0).(components.ComponentRegistrar)
}

// PluginManager provides a mock implementation of the PluginManager method.
func (m *MockSystem) PluginManager() system.PluginManagerInterface {
	args := m.Called()
	return args.Get(0).(system.PluginManagerInterface)
}

// Initialize provides a mock implementation of the Initialize method.
func (m *MockSystem) Initialize(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// Start provides a mock implementation of the Start method.
func (m *MockSystem) Start(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// Stop provides a mock implementation of the Stop method.
func (m *MockSystem) Stop(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// ExecuteOperation provides a mock implementation of the ExecuteOperation method.
func (m *MockSystem) ExecuteOperation(ctx *context.Context, operationID string, data *system.OperationInput) (*system.OperationOutput, error) {
	args := m.Called(ctx, operationID, data)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*system.OperationOutput), args.Error(1)
}

// StartService provides a mock implementation of the StartService method.
func (m *MockSystem) StartService(ctx *context.Context, serviceID string) error {
	args := m.Called(ctx, serviceID)
	return args.Error(0)
}

// StopService provides a mock implementation of the StopService method.
func (m *MockSystem) StopService(ctx *context.Context, serviceID string) error {
	args := m.Called(ctx, serviceID)
	return args.Error(0)
}

// RestartService provides a mock implementation of the RestartService method.
func (m *MockSystem) RestartService(ctx *context.Context, serviceID string) error {
	args := m.Called(ctx, serviceID)
	return args.Error(0)
}
