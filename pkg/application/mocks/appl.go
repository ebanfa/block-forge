package mocks

import (
	"github.com/edward1christian/block-forge/pkg/application/appl"
	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/stretchr/testify/mock"
)

type MockApplication struct {
	mock.Mock
}

func (m *MockApplication) Initialize(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockApplication) Start(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockApplication) Stop(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockApplication) System() system.System {
	args := m.Called()
	return args.Get(0).(system.System)
}

func (m *MockApplication) ModuleManager() appl.ModuleManager {
	args := m.Called()
	return args.Get(0).(appl.ModuleManager)
}
