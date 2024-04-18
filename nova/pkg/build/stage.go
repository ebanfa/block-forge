package build

import (
	"errors"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// BuildStageInterface represents a stage within a build pipeline.
type BuildStageInterface interface {
	// GetName returns the name of the build stage.
	GetName() string

	// GetTasks returns the tasks within the build stage.
	GetTasks() []BuildTaskInterface

	// ExecuteTasks executes all tasks within the stage.
	ExecuteTasks(ctx *context.Context) error

	// GetTaskByName returns the task with the given name from the stage.
	GetTaskByName(name string) (BuildTaskInterface, error)

	// AddTask adds a task to the stage.
	AddTask(task BuildTaskInterface) error
}

// BuildStage represents a stage within a build pipeline.
type BuildStage struct {
	Name  string
	Tasks map[string]BuildTaskInterface
}

// NewBuildStage creates a new instance of BuildStage.
func NewBuildStage(name string) *BuildStage {
	return &BuildStage{
		Name:  name,
		Tasks: make(map[string]BuildTaskInterface),
	}
}

// GetName returns the name of the build stage.
func (bs *BuildStage) GetName() string {
	return bs.Name
}

// GetTasks returns the tasks within the build stage.
func (bs *BuildStage) GetTasks() []BuildTaskInterface {
	tasks := make([]BuildTaskInterface, 0, len(bs.Tasks))
	for _, task := range bs.Tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// ExecuteTasks executes all tasks within the stage.
func (bs *BuildStage) ExecuteTasks(ctx *context.Context) error {
	for _, task := range bs.Tasks {
		_, err := task.Execute(ctx, &system.OperationInput{})
		if err != nil {
			return err
		}
	}
	return nil
}

// GetTaskByName returns the task with the given name from the stage.
func (bs *BuildStage) GetTaskByName(name string) (BuildTaskInterface, error) {
	task, exists := bs.Tasks[name]
	if !exists {
		return nil, errors.New("task not found")
	}
	return task, nil
}

// AddTask adds a task to the stage.
func (bs *BuildStage) AddTask(task BuildTaskInterface) error {
	if _, exists := bs.Tasks[task.GetName()]; exists {
		return errors.New("task with the same name already exists")
	}
	bs.Tasks[task.GetName()] = task
	return nil
}
