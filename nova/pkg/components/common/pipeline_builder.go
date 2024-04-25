package common

import (
	typesApi "github.com/edward1christian/block-forge/nova/pkg/types"
	configApi "github.com/edward1christian/block-forge/pkg/application/config"
)

// PipelineBuilderInterface defines the interface for the pipeline builder.
type PipelineBuilderInterface interface {
	// AddStage adds a stage to the pipeline.
	AddStage(name string) PipelineBuilderInterface

	// AddTask adds a task to the current stage.
	AddTask(task typesApi.TaskInterface) PipelineBuilderInterface

	// Build constructs the pipeline.
	Build() PipelineInterface
}

// PipelineBuilder is a builder for creating a pipeline.
type PipelineBuilder struct {
	pipeline     PipelineInterface
	currentStage typesApi.StageInterface
}

// NewPipelineBuilder creates a new instance of PipelineBuilder.
func NewPipelineBuilder(config *configApi.ComponentConfig) PipelineBuilderInterface {
	return &PipelineBuilder{
		pipeline: NewPipeline(config.ID, config.Name, config.Description),
	}
}

// AddStage adds a stage to the pipeline.
func (b *PipelineBuilder) AddStage(name string) PipelineBuilderInterface {
	b.currentStage = typesApi.NewStage(name)
	b.pipeline.AddStage(name, b.currentStage)
	return b
}

// AddTask adds a task to the current stage.
func (b *PipelineBuilder) AddTask(task typesApi.TaskInterface) PipelineBuilderInterface {
	if b.currentStage != nil {
		b.currentStage.AddTask(task)
	}
	return b
}

// Build constructs the pipeline.
func (b *PipelineBuilder) Build() PipelineInterface {
	return b.pipeline
}
