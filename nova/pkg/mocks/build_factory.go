package mocks

import (
	"github.com/edward1christian/block-forge/nova/pkg/build"
	"github.com/stretchr/testify/mock"
)

// MockBuilderFactory is a mock implementation of BuilderFactoryInterface.
type MockBuilderFactory struct {
	mock.Mock
}

// CreatePipelineBuilder mocks the CreatePipelineBuilder method of BuilderFactoryInterface.
func (m *MockBuilderFactory) CreatePipelineBuilder(name, builderType string) (build.PipelineBuilderInterface, error) {
	args := m.Called(name, builderType)
	return args.Get(0).(build.PipelineBuilderInterface), args.Error(1)
}

// RegisterBuilderType mocks the RegisterBuilderType method of BuilderFactoryInterface.
func (m *MockBuilderFactory) RegisterPipelineBuilderFactory(builderType string, creator func(name string) build.PipelineBuilderInterface) {
	m.Called(builderType, creator)
}
