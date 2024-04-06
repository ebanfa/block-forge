package mocks

import (
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/components"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/stretchr/testify/mock"
)

// MockETLProcessComponent is a mock implementation of the ETLProcessComponent interface.
type MockETLProcessComponent struct {
	mock.Mock
}

// ID mocks the ID method of ComponentInterface interface.
func (m *MockETLProcessComponent) ID() string {
	args := m.Called()
	return args.String(0)
}

// Name mocks the Name method of ComponentInterface interface.
func (m *MockETLProcessComponent) Name() string {
	args := m.Called()
	return args.String(0)
}

// Type mocks the Type method of ComponentInterface interface.
func (m *MockETLProcessComponent) Type() components.ComponentType {
	args := m.Called()
	return args.Get(0).(components.ComponentType)
}

// Description mocks the Description method of ComponentInterface interface.
func (m *MockETLProcessComponent) Description() string {
	args := m.Called()
	return args.String(0)
}

// GetProcessID mocks the GetProcessID method of ETLProcessComponent interface.
func (m *MockETLProcessComponent) GetProcessID() string {
	args := m.Called()
	return args.String(0)
}

// Initialize mocks the Initialize method of SystemComponentInterface interface.
func (m *MockETLProcessComponent) Initialize(ctx *context.Context, system system.SystemInterface) error {
	args := m.Called(ctx, system)
	return args.Error(0)
}

// Start mocks the Start method of ETLProcessComponent interface.
func (m *MockETLProcessComponent) Start(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// Stop mocks the Stop method of ETLProcessComponent interface.
func (m *MockETLProcessComponent) Stop(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}
