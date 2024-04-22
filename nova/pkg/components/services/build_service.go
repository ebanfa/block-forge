package services

import (
	"errors"

	"github.com/edward1christian/block-forge/nova/pkg/build"
	"github.com/edward1christian/block-forge/nova/pkg/build/builders"
	"github.com/edward1christian/block-forge/nova/pkg/common"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/common/logger"
	"github.com/edward1christian/block-forge/pkg/application/component"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
)

// BuildService represents a service for managing build pipelines.
type BuildService struct {
	systemApi.BaseSystemService // Embedding BaseComponent
	factory                     build.PipelineBuilderFactoryInterface
}

// NewBuildService creates a new instance of BuildService.
func NewBuildService(id, name, description string, factory build.PipelineBuilderFactoryInterface) *BuildService {
	return &BuildService{
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

// Initialize initializes the BuildService.
func (bs *BuildService) Initialize(ctx *context.Context, system systemApi.SystemInterface) error {
	bs.System = system

	bs.factory.RegisterPipelineBuilderFactory(
		common.IgnitePipelineBuilder, builders.CosmosSDKBlockchainPipelineBuilder)

	return nil
}

// Start starts the BuildService.
func (bs *BuildService) Start(ctx *context.Context) error {
	// Create a new instance of the pipeline builder
	bs.System.Logger().Log(logger.LevelInfo, "BuildService: Creating pipeline builder:"+common.IgnitePipelineBuilder)
	builder, err := bs.factory.CreatePipelineBuilder("Pipeline1", common.IgnitePipelineBuilder)
	if err != nil {
		return err
	}

	// Build and execute the pipeline
	bs.System.Logger().Log(logger.LevelInfo, "BuildService: Building pipeline")
	buildPipeline, err := builder.Build()
	if err != nil {
		return errors.New("failed to build pipeline")
	}

	bs.System.Logger().Log(logger.LevelInfo, "BuildService: Executing pipeline:"+buildPipeline.GetName())
	if err := buildPipeline.Execute(ctx); err != nil {
		return err
	}

	return nil
}

// Stop stops the BuildService.
func (bs *BuildService) Stop(ctx *context.Context) error {
	// Additional cleanup logic can be added here
	return nil
}
