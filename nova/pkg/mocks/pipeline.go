package mocks

import (
	"github.com/edward1christian/block-forge/nova/pkg/build"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/stretchr/testify/mock"
)

// MockBuildPipeline is a mock implementation of the PipelineInterface.
type MockBuildPipeline struct {
	mock.Mock
}

// GetName implements the GetName method of the PipelineInterface.
func (m *MockBuildPipeline) GetName() string {
	args := m.Called()
	return args.String(0)
}

// AddStage implements the AddStage method of the PipelineInterface.
func (m *MockBuildPipeline) AddStage(name string, stage build.StageInterface) error {
	args := m.Called(name, stage)
	return args.Error(0)
}

// GetStage implements the GetStage method of the PipelineInterface.
func (m *MockBuildPipeline) GetStage(name string) (build.StageInterface, error) {
	args := m.Called(name)
	return args.Get(0).(build.StageInterface), args.Error(1)
}

// GetStages implements the GetStages method of the PipelineInterface.
func (m *MockBuildPipeline) GetStages() []build.StageInterface {
	args := m.Called()
	return args.Get(0).([]build.StageInterface)
}

// Execute implements the Execute method of the PipelineInterface.
func (m *MockBuildPipeline) Execute(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}
