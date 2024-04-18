package internal

// BuildStageInterface represents a stage within a build pipeline.
type BuildStageInterface interface {
	// GetName returns the name of the build stage.
	GetName() string

	// GetTasks returns the tasks within the build stage.
	GetTasks() []BuildTaskInterface
}
