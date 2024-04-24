package mocks

import (
	"github.com/stretchr/testify/mock"

	typesApi "github.com/edward1christian/block-forge/nova/pkg/types"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
)

// MockStage is a mock implementation of StageInterface.
type MockStage struct {
	mock.Mock
}

// GetTasks mocks the GetTasks method of StageInterface.
func (m *MockStage) GetTasks() []typesApi.TaskInterface {
	args := m.Called()
	return args.Get(0).([]typesApi.TaskInterface)
}

// ExecuteTasks mocks the ExecuteTasks method of StageInterface.
func (m *MockStage) ExecuteTasks(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// GetTaskByID mocks the GetTaskByID method of StageInterface.
func (m *MockStage) GetTaskByID(name string) (typesApi.TaskInterface, error) {
	args := m.Called(name)
	return args.Get(0).(typesApi.TaskInterface), args.Error(1)
}

// AddTask mocks the AddTask method of StageInterface.
func (m *MockStage) AddTask(task typesApi.TaskInterface) error {
	args := m.Called(task)
	return args.Error(0)
}
