package build

import (
	"errors"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
)

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

	// Execute executes all stages within the build pipeline.
	Execute(ctx *context.Context) error
}

// BuildPipeline represents a build pipeline.
type BuildPipeline struct {
	Name   string
	Stages map[string]BuildStageInterface
}

// NewBuildPipeline creates a new instance of BuildPipeline.
func NewBuildPipeline(name string) *BuildPipeline {
	return &BuildPipeline{
		Name:   name,
		Stages: make(map[string]BuildStageInterface),
	}
}

// GetName returns the name of the build pipeline.
func (bp *BuildPipeline) GetName() string {
	return bp.Name
}

// AddStage adds a stage to the build pipeline.
func (bp *BuildPipeline) AddStage(name string, stage BuildStageInterface) error {
	if _, exists := bp.Stages[name]; exists {
		return errors.New("stage already exists")
	}
	bp.Stages[name] = stage
	return nil
}

// GetStage returns the stage with the given name from the build pipeline.
func (bp *BuildPipeline) GetStage(name string) (BuildStageInterface, error) {
	stage, exists := bp.Stages[name]
	if !exists {
		return nil, errors.New("stage not found")
	}
	return stage, nil
}

// GetStages returns all stages within the build pipeline.
func (bp *BuildPipeline) GetStages() []BuildStageInterface {
	stages := make([]BuildStageInterface, 0, len(bp.Stages))
	for _, stage := range bp.Stages {
		stages = append(stages, stage)
	}
	return stages
}

// Execute executes all stages within the build pipeline.
func (bp *BuildPipeline) Execute(ctx *context.Context) error {
	for _, stage := range bp.Stages {
		err := stage.ExecuteTasks(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
