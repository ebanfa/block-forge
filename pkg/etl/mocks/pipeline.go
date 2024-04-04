package mocks

import (
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/edward1christian/block-forge/pkg/etl"
	"github.com/stretchr/testify/mock"
)

type MockPipeline struct {
	mock.Mock
	etl.TransformationPipeline
	processID string
}

func (m *MockPipeline) ID() string {
	return "mock_pipeline_id"
}

func (m *MockPipeline) Name() string {
	return "Mock Pipeline"
}

func (m *MockPipeline) Description() string {
	return "Mock Pipeline Description"
}

func (m *MockPipeline) setProcessID(processID string) {
	m.processID = processID
}

func (m *MockPipeline) getProcessID() string {
	return m.processID
}

func (m *MockPipeline) AddStage(stage etl.TransformationStage) error {
	args := m.Called(stage)
	return args.Error(0)
}

func (m *MockPipeline) RemoveStage(stageID string) error {
	args := m.Called(stageID)
	return args.Error(0)
}

func (m *MockPipeline) Start(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockPipeline) Stop(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockPipeline) Initialize(ctx *context.Context, system system.System) error {
	args := m.Called(ctx, system)
	return args.Error(0)
}

type MockTransformationStage struct {
	mock.Mock
	processID string
}

func (m *MockTransformationStage) ID() string {
	return "mock_stage_id"
}

func (m *MockTransformationStage) Name() string {
	return "Mock Transformation Stage"
}

func (m *MockTransformationStage) Description() string {
	return "Mock Transformation Stage Description"
}

func (m *MockTransformationStage) setProcessID(processID string) {
	m.processID = processID
}

func (m *MockTransformationStage) getProcessID() string {
	return m.processID
}

func (m *MockTransformationStage) Transform(data interface{}) (interface{}, error) {
	args := m.Called(data)
	return args.Get(0), args.Error(1)
}

func (m *MockTransformationStage) Start(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockTransformationStage) Stop(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockTransformationStage) Initialize(ctx *context.Context, system system.System) error {
	args := m.Called(ctx, system)
	return args.Error(0)
}
