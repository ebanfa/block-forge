package internal

// BuildPipelineInterface represents a build pipeline.
type BuildPipelineInterface interface {
	// GetName returns the name of the build pipeline.
	GetName() string

	// AddStage adds a stage to the build pipeline.
	AddStage(name string, stage BuildStageInterface) error

	// GetStage returns the stage with the given name from the build pipeline.
	GetStage(name string) (BuildStageInterface, error)

	// GetStages returns all stages within the build pipeline.
	GetStages() []BuildStageInterface
}
