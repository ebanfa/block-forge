package build

import (
	"errors"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// StageInterface represents a stage within a build pipeline.
type StageInterface interface {

	// GetTasks returns the tasks within the build stage.
	GetTasks() []TaskInterface

	// ExecuteTasks executes all tasks within the stage.
	ExecuteTasks(ctx *context.Context) error

	// GetTaskByID returns the task with the given name from the stage.
	GetTaskByID(name string) (TaskInterface, error)

	// AddTask adds a task to the stage.
	AddTask(task TaskInterface) error
}

// BuildStage represents a stage within a build pipeline.
type BuildStage struct {
	Name  string
	Tasks map[string]TaskInterface
}

// NewBuildStage creates a new instance of BuildStage.
func NewBuildStage(name string) *BuildStage {
	return &BuildStage{
		Name:  name,
		Tasks: make(map[string]TaskInterface),
	}
}

// GetTasks returns the tasks within the build stage.
func (bs *BuildStage) GetTasks() []TaskInterface {
	tasks := make([]TaskInterface, 0, len(bs.Tasks))
	for _, task := range bs.Tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// ExecuteTasks executes all tasks within the stage.
func (bs *BuildStage) ExecuteTasks(ctx *context.Context) error {
	for _, task := range bs.Tasks {
		_, err := task.Execute(ctx, &system.SystemOperationInput{})
		if err != nil {
			return err
		}
	}
	return nil
}

// GetTaskByID returns the task with the given id from the stage.
func (bs *BuildStage) GetTaskByID(id string) (TaskInterface, error) {
	task, exists := bs.Tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}
	return task, nil
}

// AddTask adds a task to the stage.
func (bs *BuildStage) AddTask(task TaskInterface) error {
	if _, exists := bs.Tasks[task.ID()]; exists {
		return errors.New("task with the same id already exists")
	}
	bs.Tasks[task.ID()] = task
	return nil
}
