package build

// PipelineBuilder is a builder for creating a pipeline.
type PipelineBuilder struct {
	pipeline     *BuildPipeline
	currentStage *BuildStage
}

// NewPipelineBuilder creates a new instance of PipelineBuilder.
func NewPipelineBuilder(name string) *PipelineBuilder {
	return &PipelineBuilder{
		pipeline: NewBuildPipeline(name),
	}
}

// AddStage adds a stage to the pipeline.
func (b *PipelineBuilder) AddStage(name string) *PipelineBuilder {
	b.currentStage = NewBuildStage(name)
	b.pipeline.AddStage(name, b.currentStage)
	return b
}

// AddTask adds a task to the current stage.
func (b *PipelineBuilder) AddTask(task BuildTaskInterface) *PipelineBuilder {
	if b.currentStage == nil {
		panic("No stage added. Please add a stage first.")
	}
	b.currentStage.AddTask(task)
	return b
}

// Build constructs the pipeline.
func (b *PipelineBuilder) Build() BuildPipelineInterface {
	return b.pipeline
}
