package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/application/event"
	"github.com/edward1christian/block-forge/pkg/application/logger"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// MockSystem represents a mock for the SystemImpl type.
type MockSystem struct {
	mock.Mock
}

// ID mocks the ID method of the SystemImpl type.
func (m *MockSystem) ID() string {
	args := m.Called()
	return args.String(0)
}

// Name mocks the Name method of the SystemImpl type.
func (m *MockSystem) Name() string {
	args := m.Called()
	return args.String(0)
}

// Description mocks the Description method of the SystemImpl type.
func (m *MockSystem) Description() string {
	args := m.Called()
	return args.String(0)
}

// Initialize mocks the Initialize method of the SystemImpl type.
func (m *MockSystem) Initialize(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// Start mocks the Start method of the SystemImpl type.
func (m *MockSystem) Start(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// Stop mocks the Stop method of the SystemImpl type.
func (m *MockSystem) Stop(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// RegisterOperation mocks the RegisterOperation method of the SystemImpl type.
func (m *MockSystem) RegisterOperation(operationID string, operation system.Operation) error {
	args := m.Called(operationID, operation)
	return args.Error(0)
}

// UnregisterOperation mocks the UnregisterOperation method of the SystemImpl type.
func (m *MockSystem) UnregisterOperation(operationID string) error {
	args := m.Called(operationID)
	return args.Error(0)
}

// ExecuteOperation mocks the ExecuteOperation method of the SystemImpl type.
func (m *MockSystem) ExecuteOperation(ctx *context.Context, operationID string, data *system.OperationInput) (*system.OperationOutput, error) {
	args := m.Called(ctx, operationID, data)
	if output, ok := args.Get(0).(*system.OperationOutput); ok {
		return output, args.Error(1)
	}
	return nil, args.Error(1)
}

// RegisterService mocks the RegisterService method of the SystemImpl type.
func (m *MockSystem) RegisterService(serviceID string, service system.SystemService) error {
	args := m.Called(serviceID, service)
	return args.Error(0)
}

// UnregisterService mocks the UnregisterService method of the SystemImpl type.
func (m *MockSystem) UnregisterService(serviceID string) error {
	args := m.Called(serviceID)
	return args.Error(0)
}

// StartService mocks the StartService method of the SystemImpl type.
func (m *MockSystem) StartService(serviceID string, ctx *context.Context) error {
	args := m.Called(serviceID, ctx)
	return args.Error(0)
}

// StopService mocks the StopService method of the SystemImpl type.
func (m *MockSystem) StopService(serviceID string, ctx *context.Context) error {
	args := m.Called(serviceID, ctx)
	return args.Error(0)
}

// RegisterComponentFactory mocks the RegisterComponentFactory method of the SystemImpl type.
func (m *MockSystem) RegisterComponentFactory(name string, factory system.ComponentFactory) error {
	args := m.Called(name, factory)
	return args.Error(0)
}

// GetComponentFactory mocks the GetComponentFactory method of the SystemImpl type.
func (m *MockSystem) GetComponentFactory(name string) (system.ComponentFactory, error) {
	args := m.Called(name)
	if factory, ok := args.Get(0).(system.ComponentFactory); ok {
		return factory, args.Error(1)
	}
	return nil, args.Error(1)
}

// Logger mocks the Logger method of the SystemImpl type.
func (m *MockSystem) Logger() logger.LoggerInterface {
	args := m.Called()
	if logger, ok := args.Get(0).(logger.LoggerInterface); ok {
		return logger
	}
	return nil
}

// EventBus mocks the EventBus method of the SystemImpl type.
func (m *MockSystem) EventBus() event.EventBusInterface {
	args := m.Called()
	if eventBus, ok := args.Get(0).(event.EventBusInterface); ok {
		return eventBus
	}
	return nil
}

// Configuration mocks the Configuration method of the SystemImpl type.
func (m *MockSystem) Configuration() system.Configuration {
	args := m.Called()
	if configuration, ok := args.Get(0).(system.Configuration); ok {
		return configuration
	}
	return system.Configuration{}
}
