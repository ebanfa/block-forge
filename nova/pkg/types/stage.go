package types

import (
	"errors"
)

// TaskInterface represents a build task.
type TaskInterface interface {
	// ID returns the unique identifier of the component.
	ID() string
}

// Task represents a build task.
type Task struct {
	id string
}

// NewTask creates a new Task instance with the given ID, name, and description.
func NewTask(id string) *Task {
	return &Task{id: id}
}

// ID returns the unique identifier of the component.
func (t *Task) ID() string {
	return t.id
}

// StageInterface represents a stage within a build pipeline.
type StageInterface interface {

	// GetTasks returns the tasks within the build stage.
	GetTasks() []TaskInterface

	// GetTaskByID returns the task with the given name from the stage.
	GetTaskByID(name string) (TaskInterface, error)

	// AddTask adds a task to the stage.
	AddTask(task TaskInterface) error
}

// Stage represents a stage within a build pipeline.
type Stage struct {
	Name  string
	Tasks map[string]TaskInterface
}

// NewStage creates a new instance of Stage.
func NewStage(name string) *Stage {
	return &Stage{
		Name:  name,
		Tasks: make(map[string]TaskInterface),
	}
}

// GetTasks returns the tasks within the build stage.
func (bs *Stage) GetTasks() []TaskInterface {
	// Initialize an empty slice to hold the tasks.
	tasks := make([]TaskInterface, 0, len(bs.Tasks))

	// Iterate over the tasks map and append each task to the slice.
	for _, task := range bs.Tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// GetTaskByID returns the task with the given id from the stage.
func (bs *Stage) GetTaskByID(id string) (TaskInterface, error) {
	// Check if the task with the given id exists in the tasks map.
	task, exists := bs.Tasks[id]
	if !exists {
		// If the task does not exist, return an error.
		return nil, errors.New("task not found")
	}
	return task, nil
}

// AddTask adds a task to the stage.
func (bs *Stage) AddTask(task TaskInterface) error {
	// Check if a task with the same id already exists in the tasks map.
	if _, exists := bs.Tasks[task.ID()]; exists {
		// If a task with the same id exists, return an error.
		return errors.New("task with the same id already exists")
	}
	// Add the task to the tasks map.
	bs.Tasks[task.ID()] = task
	return nil
}
