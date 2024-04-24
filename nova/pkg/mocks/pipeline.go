package mocks

import (
	"github.com/stretchr/testify/mock"

	typesApi "github.com/edward1christian/block-forge/nova/pkg/types"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/component"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// MockPipeline is a testify mock for the PipelineInterface.
type MockPipeline struct {
	mock.Mock
}

// AddStage adds a stage to the build pipeline.
func (m *MockPipeline) AddStage(name string, stage typesApi.StageInterface) error {
	args := m.Called(name, stage)
	return args.Error(0)
}

// GetStage returns the stage with the given name from the build pipeline.
func (m *MockPipeline) GetStage(name string) (typesApi.StageInterface, error) {
	args := m.Called(name)
	stage, _ := args.Get(0).(typesApi.StageInterface)
	return stage, args.Error(1)
}

// GetStages returns all stages within the build pipeline.
func (m *MockPipeline) GetStages() []typesApi.StageInterface {
	args := m.Called()
	stages, _ := args.Get(0).([]typesApi.StageInterface)
	return stages
}

// Execute executes all stages within the build pipeline.
func (m *MockPipeline) Execute(ctx *context.Context, data *system.SystemOperationInput) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}

// Type returns the type of the component.
func (m *MockPipeline) Type() component.ComponentType {
	args := m.Called()
	return args.Get(0).(component.ComponentType)
}

// ID returns the unique identifier of the component.
func (m *MockPipeline) ID() string {
	args := m.Called()
	return args.String(0)
}

// Name returns the name of the component.
func (m *MockPipeline) Name() string {
	args := m.Called()
	return args.String(0)
}

// Description returns the description of the component.
func (m *MockPipeline) Description() string {
	args := m.Called()
	return args.String(0)
}

// Initialize initializes the module.
func (m *MockPipeline) Initialize(ctx *context.Context, system system.SystemInterface) error {
	args := m.Called(ctx, system)
	return args.Error(0)
}
