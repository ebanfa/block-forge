package build

import (
	"errors"
)

// PipelineBuilderInterface defines the interface for the pipeline builder.
type PipelineBuilderInterface interface {
	// AddStage adds a stage to the pipeline.
	AddStage(name string) (PipelineBuilderInterface, error)

	// AddTask adds a task to the current stage.
	AddTask(task BuildTaskInterface) (PipelineBuilderInterface, error)

	// Build constructs the pipeline.
	Build() (BuildPipelineInterface, error)
}

// PipelineBuilder is a builder for creating a pipeline.
type PipelineBuilder struct {
	pipeline     *BuildPipeline
	currentStage *BuildStage
}

// NewPipelineBuilder creates a new instance of PipelineBuilder.
func NewPipelineBuilder(name string) PipelineBuilderInterface {
	return &PipelineBuilder{
		pipeline: NewBuildPipeline(name),
	}
}

// AddStage adds a stage to the pipeline.
func (b *PipelineBuilder) AddStage(name string) (PipelineBuilderInterface, error) {
	b.currentStage = NewBuildStage(name)

	if err := b.pipeline.AddStage(name, b.currentStage); err != nil {
		return nil, errors.New("failed to add stage: " + err.Error())
	}
	return b, nil
}

// AddTask adds a task to the current stage.
func (b *PipelineBuilder) AddTask(task BuildTaskInterface) (PipelineBuilderInterface, error) {
	if b.currentStage == nil {
		return nil, errors.New("no stage added. Please add a stage first")
	}
	if err := b.currentStage.AddTask(task); err != nil {
		return nil, errors.New("failed to add task to stage: " + err.Error())
	}
	return b, nil
}

// Build constructs the pipeline.
func (b *PipelineBuilder) Build() (BuildPipelineInterface, error) {
	if len(b.pipeline.Stages) == 0 {
		return nil, errors.New("no stages added. Please add stages before building the pipeline")
	}
	return b.pipeline, nil
}
