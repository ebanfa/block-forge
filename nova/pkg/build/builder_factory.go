package build

import "errors"

// BuilderFactoryInterface defines the interface for the builder factory.
type BuilderFactoryInterface interface {
	CreatePipelineBuilder(name, builderType string) (PipelineBuilderInterface, error)
	RegisterBuilderType(builderType string, creator func(name string) PipelineBuilderInterface)
}

// PipelineBuilderFactory is a factory for creating different types of PipelineBuilders.
type PipelineBuilderFactory struct {
	BuilderFactoryInterface
	builderCreators map[string]func(name string) PipelineBuilderInterface
}

// NewPipelineBuilderFactory creates a new instance of PipelineBuilderFactory.
func NewPipelineBuilderFactory() *PipelineBuilderFactory {
	return &PipelineBuilderFactory{
		builderCreators: make(map[string]func(name string) PipelineBuilderInterface),
	}
}

// RegisterBuilderType registers a builder creation function for the given pipeline type.
func (f *PipelineBuilderFactory) RegisterBuilderType(builderType string, creator func(name string) PipelineBuilderInterface) {
	f.builderCreators[builderType] = creator
}

// CreatePipelineBuilder creates a new instance of PipelineBuilder based on the builder type.
func (f *PipelineBuilderFactory) CreatePipelineBuilder(name, builderType string) (PipelineBuilderInterface, error) {
	creator, exists := f.builderCreators[builderType]
	if !exists {
		return nil, errors.New("unsupported builder type")
	}
	builder := creator(name)
	if builder == nil {
		return nil, errors.New("builder creator returned nil")
	}
	return builder, nil
}
