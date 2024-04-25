package mocks

import (
	"github.com/stretchr/testify/mock"

	typesApi "github.com/edward1christian/block-forge/nova/pkg/types"
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
