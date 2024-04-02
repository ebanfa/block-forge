package mocks

import (
	"github.com/edward1christian/block-forge/pkg/application/event"
	"github.com/stretchr/testify/mock"
)

// Mock objects/interfaces for testing purposes

// MockEventBus is a mock for the EventBusInterface.
type MockEventBus struct {
	mock.Mock
}

// Subscribe mocks the Subscribe method of the EventBusInterface.
func (m *MockEventBus) Subscribe(params event.BusSubscriptionParams) error {
	args := m.Called(params)
	return args.Error(0)
}

// SubscribeAsync mocks the SubscribeAsync method of the EventBusInterface.
func (m *MockEventBus) SubscribeAsync(params event.BusSubscriptionParams, transactional bool) error {
	args := m.Called(params, transactional)
	return args.Error(0)
}

// SubscribeOnce mocks the SubscribeOnce method of the EventBusInterface.
func (m *MockEventBus) SubscribeOnce(params event.BusSubscriptionParams) error {
	args := m.Called(params)
	return args.Error(0)
}

// SubscribeOnceAsync mocks the SubscribeOnceAsync method of the EventBusInterface.
func (m *MockEventBus) SubscribeOnceAsync(params event.BusSubscriptionParams) error {
	args := m.Called(params)
	return args.Error(0)
}

// Unsubscribe mocks the Unsubscribe method of the EventBusInterface.
func (m *MockEventBus) Unsubscribe(params event.BusSubscriptionParams) error {
	args := m.Called(params)
	return args.Error(0)
}

// Publish mocks the Publish method of the EventBusInterface.
func (m *MockEventBus) Publish(event event.Event) {
	m.Called(event)
}

// HasCallback mocks the HasCallback method of the EventBusInterface.
func (m *MockEventBus) HasCallback(topic string) bool {
	args := m.Called(topic)
	return args.Bool(0)
}

// WaitAsync mocks the WaitAsync method of the EventBusInterface.
func (m *MockEventBus) WaitAsync() {
	m.Called()
}
