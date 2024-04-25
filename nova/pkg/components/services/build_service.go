package services

import (
	"fmt"

	ncApi "github.com/edward1christian/block-forge/nova/pkg/common"
	"github.com/edward1christian/block-forge/nova/pkg/components/common"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/component"
	configApi "github.com/edward1christian/block-forge/pkg/application/config"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
)

// BuildServiceFactory is responsible for creating instances of BuildService.
type BuildServiceFactory struct {
}

// CreateComponent creates a new instance of the BuildService.
func (bf *BuildServiceFactory) CreateComponent(config *configApi.ComponentConfig) (component.ComponentInterface, error) {
	// Construct the service
	return NewBuildService(config.ID, config.Name, config.Description), nil
}

// BuildService represents a service for managing build pipelines.
type BuildService struct {
	systemApi.BaseSystemService // Embedding BaseComponent
}

// NewBuildService creates a new instance of BuildService.
func NewBuildService(id, name, description string) systemApi.SystemServiceInterface {
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
	}
}

// Initialize initializes the BuildService.
// It sets the system instance and registers a pipeline factory.
func (bs *BuildService) Initialize(ctx *context.Context, system systemApi.SystemInterface) error {
	bs.System = system

	// Register pipeline factory
	registrar := system.ComponentRegistry()
	if err := registrar.RegisterFactory(ctx, ncApi.BuildPipelineFactory, &common.PipelineFactory{}); err != nil {
		return fmt.Errorf("failed to register pipeline factory: %w", err)
	}

	return nil
}

// Start starts the BuildService.
// It creates a new instance of the pipeline builder and starts it.
func (bs *BuildService) Start(ctx *context.Context) error {
	// Create a new instance of the pipeline builder
	pipeline, err := bs.createPipeline(ctx, &configApi.ComponentConfig{
		ID:        ncApi.BuildPipeline,
		FactoryID: ncApi.BuildPipelineFactory,
	})
	if err != nil {
		return fmt.Errorf("failed to start pipeline: %w", err)
	}

	// Start the pipeline
	return pipeline.Execute(ctx, &systemApi.SystemOperationInput{})
}

// createPipeline creates a new instance of the pipeline builder.
// It retrieves the pipeline component from the system's component registry.
func (bs *BuildService) createPipeline(ctx *context.Context, config *configApi.ComponentConfig) (common.PipelineInterface, error) {
	// Create a new pipeline component using the component registry
	pipelineComp, err := bs.System.ComponentRegistry().CreateComponent(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create pipeline: %w", err)
	}
	// Check if the created component is a system service interface
	pipeline, ok := pipelineComp.(common.PipelineInterface)
	if !ok {
		// Return an error if the created component is not a system service
		return nil, fmt.Errorf("instantiated pipeline (%s) is not a system service", pipelineComp.ID())
	}
	// Initialize the pipeline
	err = pipeline.Initialize(ctx, bs.System)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize pipeline: %w", err)
	}

	// Return the created pipeline
	return pipeline, nil
}

// Stop stops the BuildService.
// It removes the pipeline component registration and unregisters the associated factory.
func (bs *BuildService) Stop(ctx *context.Context) error {
	// Remove the pipeline component registration
	if err := bs.System.ComponentRegistry().RemoveComponent(ctx, "pipeline"); err != nil {
		return fmt.Errorf("failed to remove pipeline component: %w", err)
	}

	// Unregister the factory used to create the pipeline
	if err := bs.System.ComponentRegistry().UnregisterFactory(ctx, "factory"); err != nil {
		return fmt.Errorf("failed to unregister pipeline factory: %w", err)
	}

	return nil
}
