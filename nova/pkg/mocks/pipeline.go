package mocks

import (
	"github.com/edward1christian/block-forge/nova/pkg/build"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/stretchr/testify/mock"
)

// MockBuildPipeline is a mock implementation of the BuildPipelineInterface.
type MockBuildPipeline struct {
	mock.Mock
}

// GetName implements the GetName method of the BuildPipelineInterface.
func (m *MockBuildPipeline) GetName() string {
	args := m.Called()
	return args.String(0)
}

// AddStage implements the AddStage method of the BuildPipelineInterface.
func (m *MockBuildPipeline) AddStage(name string, stage build.BuildStageInterface) error {
	args := m.Called(name, stage)
	return args.Error(0)
}

// GetStage implements the GetStage method of the BuildPipelineInterface.
func (m *MockBuildPipeline) GetStage(name string) (build.BuildStageInterface, error) {
	args := m.Called(name)
	return args.Get(0).(build.BuildStageInterface), args.Error(1)
}

// GetStages implements the GetStages method of the BuildPipelineInterface.
func (m *MockBuildPipeline) GetStages() []build.BuildStageInterface {
	args := m.Called()
	return args.Get(0).([]build.BuildStageInterface)
}

// Execute implements the Execute method of the BuildPipelineInterface.
func (m *MockBuildPipeline) Execute(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}
