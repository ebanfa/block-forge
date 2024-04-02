package mocks

import (
	"github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/edward1christian/block-forge/pkg/application/event"
	"github.com/edward1christian/block-forge/pkg/application/logger"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/stretchr/testify/mock"
)

// / MockSystem is a mock implementation of the System interface.
// MockSystem is a mock implementation of the System interface.
type MockSystem struct {
	mock.Mock
	system.System // Embedding the interfaces.System interface
}

// EventBus returns the mock implementation for the EventBusInterface.
func (m *MockSystem) EventBus() event.EventBusInterface {
	args := m.Called()
	return args.Get(0).(event.EventBusInterface)
}

// Operations returns the mock implementation for the Operations interface.
func (m *MockSystem) Operations() system.Operations {
	args := m.Called()
	return args.Get(0).(system.Operations)
}

// Logger returns the mock implementation for the LoggerInterface.
func (m *MockSystem) Logger() logger.LoggerInterface {
	args := m.Called()
	return args.Get(0).(logger.LoggerInterface)
}

// Configuration returns the mock implementation for the Configuration.
func (m *MockSystem) Configuration() config.Configuration {
	args := m.Called()
	return args.Get(0).(config.Configuration)
}
