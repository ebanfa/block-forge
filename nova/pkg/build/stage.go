package internal

import "github.com/edward1christian/block-forge/pkg/application/common/context"

// BuildStageInterface represents a stage within a build pipeline.
type BuildStageInterface interface {
	// GetName returns the name of the build stage.
	GetName() string

	// AddTask adds a task to the stage.
	AddTask(task BuildTaskInterface) error

	// GetTasks returns the tasks within the build stage.
	GetTasks() []BuildTaskInterface

	// GetTaskByName returns the task with the given name from the stage.
	GetTaskByName(name string) (BuildTaskInterface, error)

	// ExecuteTasks executes all tasks within the stage.
	ExecuteTasks(ctx *context.Context) error
}
