package mocks

import (
	"github.com/edward1christian/block-forge/nova/pkg/build"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/stretchr/testify/mock"
)

// MockBuildStage is a mock implementation of the BuildStageInterface.
type MockBuildStage struct {
	mock.Mock
}

// GetName implements the GetName method of the BuildStageInterface.
func (m *MockBuildStage) GetName() string {
	args := m.Called()
	return args.String(0)
}

// GetTasks implements the GetTasks method of the BuildStageInterface.
func (m *MockBuildStage) GetTasks() []build.BuildTaskInterface {
	args := m.Called()
	return args.Get(0).([]build.BuildTaskInterface)
}

// ExecuteTasks implements the ExecuteTasks method of the BuildStageInterface.
func (m *MockBuildStage) ExecuteTasks(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// GetTaskByName implements the GetTaskByName method of the BuildStageInterface.
func (m *MockBuildStage) GetTaskByName(name string) (build.BuildTaskInterface, error) {
	args := m.Called(name)
	return args.Get(0).(build.BuildTaskInterface), args.Error(1)
}

// AddTask implements the AddTask method of the BuildStageInterface.
func (m *MockBuildStage) AddTask(task build.BuildTaskInterface) error {
	args := m.Called(task)
	return args.Error(0)
}
