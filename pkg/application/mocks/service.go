package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/component"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
)

// MockSystemService represents a mock for the SystemService interface.
type MockSystemService struct {
	mock.Mock
}

// ID returns the unique identifier of the component.
func (m *MockSystemService) ID() string {
	args := m.Called()
	return args.String(0)
}

// Name returns the name of the component.
func (m *MockSystemService) Name() string {
	args := m.Called()
	return args.String(0)
}

// Description returns the description of the component.
func (m *MockSystemService) Description() string {
	args := m.Called()
	return args.String(0)
}

// Description returns the description of the component.
func (m *MockSystemService) Type() component.ComponentType {
	args := m.Called()
	return args.Get(0).(component.ComponentType)
}

// Start mocks the Start method of the SystemService interface.
func (m *MockSystemService) Initialize(ctx *context.Context, system systemApi.SystemInterface) error {
	args := m.Called(ctx, system)
	return args.Error(0)
}

// Start mocks the Start method of the SystemService interface.
func (m *MockSystemService) Start(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// Stop mocks the Stop method of the SystemService interface.
func (m *MockSystemService) Stop(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}
