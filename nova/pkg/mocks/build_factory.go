package mocks

import (
	"github.com/edward1christian/block-forge/nova/pkg/build"
	"github.com/stretchr/testify/mock"
)

// MockBuilderFactory is a mock implementation of BuilderFactoryInterface for testing purposes.
type MockBuilderFactory struct {
	mock.Mock
}

// CreatePipelineBuilder is a mock implementation for creating a pipeline builder instance.
func (m *MockBuilderFactory) CreatePipelineBuilder(name, builderType string) (build.PipelineBuilderInterface, error) {
	args := m.Called(name, builderType)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(build.PipelineBuilderInterface), args.Error(1)
}
