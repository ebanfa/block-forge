package common

import (
	"errors"

	typesApi "github.com/edward1christian/block-forge/nova/pkg/types"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/component"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
)

// PipelineInterface represents a build pipeline.
type PipelineInterface interface {
	systemApi.SystemComponentInterface

	// AddStage adds a stage to the build pipeline.
	AddStage(name string, stage typesApi.StageInterface) error

	// GetStage returns the stage with the given name from the build pipeline.
	GetStage(name string) (typesApi.StageInterface, error)

	// GetStages returns all stages within the build pipeline.
	GetStages() []typesApi.StageInterface

	// Execute executes all stages within the build pipeline.
	Execute(ctx *context.Context, data *systemApi.SystemOperationInput) error
}

// Pipeline represents a build pipeline.
type Pipeline struct {
	PipelineInterface
	systemApi.BaseSystemComponent
	Stages map[string]typesApi.StageInterface
}

// NewPipeline creates a new instance of Pipeline.
func NewPipeline(id, name, description string) PipelineInterface {
	return &Pipeline{
		Stages: make(map[string]typesApi.StageInterface),
		BaseSystemComponent: systemApi.BaseSystemComponent{
			BaseComponent: component.BaseComponent{
				Id:   id,
				Nm:   name,
				Desc: description,
			},
		},
	}
}

// Type returns the type of the component.
func (p *Pipeline) Type() component.ComponentType {
	return component.SystemComponentType
}

// ID returns the unique identifier of the component.
func (p *Pipeline) ID() string {
	return p.Id
}

// Name returns the Nm of the component.
func (p *Pipeline) Name() string {
	return p.Nm
}

// Description returns the Desc of the component.
func (p *Pipeline) Description() string {
	return p.Desc
}

// Initialize initializes the module.
// Returns an error if the initialization fails.
func (p *Pipeline) Initialize(ctx *context.Context, system systemApi.SystemInterface) error {
	p.System = system
	return nil
}

// AddStage adds a stage to the build pipeline.
func (bp *Pipeline) AddStage(name string, stage typesApi.StageInterface) error {
	if _, exists := bp.Stages[name]; exists {
		return errors.New("stage already exists")
	}
	bp.Stages[name] = stage
	return nil
}

// GetStage returns the stage with the given name from the build pipeline.
func (bp *Pipeline) GetStage(name string) (typesApi.StageInterface, error) {
	stage, exists := bp.Stages[name]
	if !exists {
		return nil, errors.New("stage not found")
	}
	return stage, nil
}

// GetStages returns all stages within the build pipeline.
func (bp *Pipeline) GetStages() []typesApi.StageInterface {
	stages := make([]typesApi.StageInterface, 0, len(bp.Stages))
	for _, stage := range bp.Stages {
		stages = append(stages, stage)
	}
	return stages
}

// Execute executes all stages within the build pipeline.
func (bp *Pipeline) Execute(ctx *context.Context, data *systemApi.SystemOperationInput) error {
	for _, stage := range bp.Stages {
		for _, task := range stage.GetTasks() {
			_, err := bp.System.ExecuteOperation(ctx, task.ID(), data)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
