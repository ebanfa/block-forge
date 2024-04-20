package factories

import (
	"github.com/edward1christian/block-forge/nova/pkg/build"
	"github.com/edward1christian/block-forge/nova/pkg/components/services"
	"github.com/edward1christian/block-forge/pkg/application/component"
	configApi "github.com/edward1christian/block-forge/pkg/application/config"
)

// BuilderServiceFactory is responsible for creating instances of BuilderService.
type BuilderServiceFactory struct {
}

// CreateComponent creates a new instance of the BuilderService.
func (bf *BuilderServiceFactory) CreateComponent(config *configApi.ComponentConfig) (component.ComponentInterface, error) {
	// Create a new instance of BuilderService
	builderFactory := build.NewPipelineBuilderFactory()

	// Create the service
	builderService := services.NewBuilderService(
		config.ID, config.Name, config.Description, builderFactory)

	return builderService, nil
}
