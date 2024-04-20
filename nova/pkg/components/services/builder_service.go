package services

import (
	"errors"

	"github.com/edward1christian/block-forge/nova/pkg/build"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/component"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
)

// BuilderService represents a service for managing build pipelines.
type BuilderService struct {
	systemApi.BaseSystemService // Embedding BaseComponent
	factory                     build.BuilderFactoryInterface
}

// NewBuilderService creates a new instance of BuilderService.
func NewBuilderService(id, name, description string, factory build.BuilderFactoryInterface) *BuilderService {
	return &BuilderService{
		BaseSystemService: systemApi.BaseSystemService{
			BaseSystemComponent: systemApi.BaseSystemComponent{
				BaseComponent: component.BaseComponent{
					Id:   id,
					Nm:   name,
					Desc: description,
				},
			},
		},
		factory: factory,
	}
}

// Initialize initializes the BuilderService.
func (bs *BuilderService) Initialize(ctx *context.Context, system systemApi.SystemInterface) error {
	bs.System = system
	bs.factory.RegisterBuilderType("Pipeline1", createPipelineBuilder)
	// Additional initialization logic can be added here
	return nil
}

// Start starts the BuilderService.
func (bs *BuilderService) Start(ctx *context.Context) error {
	// Create a new instance of the pipeline builder
	builder, err := bs.factory.CreatePipelineBuilder("Pipeline1", "type1")
	if err != nil {
		return err
	}

	// Build and execute the pipeline
	buildPipeline, err := builder.Build()
	if err != nil {
		return errors.New("failed to build pipeline")
	}

	if err := buildPipeline.Execute(ctx); err != nil {
		return err
	}

	return nil
}

// Stop stops the BuilderService.
func (bs *BuilderService) Stop(ctx *context.Context) error {
	// Additional cleanup logic can be added here
	return nil
}
func createPipelineBuilder(name string) build.PipelineBuilderInterface {
	builder := build.NewPipelineBuilder(name)
	builder.AddStage("Test")
	return builder
}
