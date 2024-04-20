package mocks

import (
	"github.com/edward1christian/block-forge/nova/pkg/build"
	"github.com/stretchr/testify/mock"
)

// MockPipelineBuilder is a mock implementation of PipelineBuilderInterface for testing purposes.
type MockPipelineBuilder struct {
	mock.Mock
}

// AddStage is a mock implementation for adding a stage to the pipeline.
func (m *MockPipelineBuilder) AddStage(name string) (build.PipelineBuilderInterface, error) {
	args := m.Called(name)
	return args.Get(0).(build.PipelineBuilderInterface), args.Error(1)
}

// AddTask is a mock implementation for adding a task to the current stage.
func (m *MockPipelineBuilder) AddTask(task build.TaskInterface) (build.PipelineBuilderInterface, error) {
	args := m.Called(task)
	return args.Get(0).(build.PipelineBuilderInterface), args.Error(1)
}

// Build is a mock implementation for constructing the pipeline.
func (m *MockPipelineBuilder) Build() (build.PipelineInterface, error) {
	args := m.Called()
	return args.Get(0).(build.PipelineInterface), args.Error(1)
}
