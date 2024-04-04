// Package mocks provides mock implementations for testing purposes.
package mocks

import (
	"github.com/edward1christian/block-forge/pkg/application/appl"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/stretchr/testify/mock"
)

// MockApplication represents a mock for the Application interface.
type MockApplication struct {
	appl.ApplicationInterface
	mock.Mock
}

// Initialize mocks the Initialize method of the Application interface.
func (m *MockApplication) Initialize(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// Start mocks the Start method of the Application interface.
func (m *MockApplication) Start(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// Stop mocks the Stop method of the Application interface.
func (m *MockApplication) Stop(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// System mocks the System method of the Application interface.
func (m *MockApplication) System() system.SystemInterface {
	args := m.Called()
	return args.Get(0).(system.SystemInterface)
}

// ModuleManager mocks the ModuleManager method of the Application interface.
func (m *MockApplication) ModuleManager() appl.ModuleManager {
	args := m.Called()
	return args.Get(0).(appl.ModuleManager)
}
