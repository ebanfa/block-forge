package mocks

import (
	"github.com/edward1christian/block-forge/nova/pkg/build"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/stretchr/testify/mock"
)

// MockBuildStage is a mock implementation of the StageInterface.
type MockBuildStage struct {
	mock.Mock
}

// GetTasks implements the GetTasks method of the StageInterface.
func (m *MockBuildStage) GetTasks() []build.TaskInterface {
	args := m.Called()
	return args.Get(0).([]build.TaskInterface)
}

// ExecuteTasks implements the ExecuteTasks method of the StageInterface.
func (m *MockBuildStage) ExecuteTasks(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// GetTaskByID implements the GetTaskByID method of the StageInterface.
func (m *MockBuildStage) GetTaskByID(id string) (build.TaskInterface, error) {
	args := m.Called(id)
	return args.Get(0).(build.TaskInterface), args.Error(1)
}

// AddTask implements the AddTask method of the StageInterface.
func (m *MockBuildStage) AddTask(task build.TaskInterface) error {
	args := m.Called(task)
	return args.Error(0)
}
