package internal

// BuildPipelineInterface represents a build pipeline.
type BuildPipelineInterface interface {
	// GetName returns the name of the build pipeline.
	GetName() string

	// GetStages returns the stages within the build pipeline.
	GetStages() []BuildStageInterface
}
