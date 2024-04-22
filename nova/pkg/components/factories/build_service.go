package factories

import (
	"github.com/edward1christian/block-forge/nova/pkg/build"
	"github.com/edward1christian/block-forge/nova/pkg/components/services"
	"github.com/edward1christian/block-forge/pkg/application/component"
	configApi "github.com/edward1christian/block-forge/pkg/application/config"
)

// BuildServiceFactory is responsible for creating instances of BuildService.
type BuildServiceFactory struct {
}

// CreateComponent creates a new instance of the BuildService.
func (bf *BuildServiceFactory) CreateComponent(config *configApi.ComponentConfig) (component.ComponentInterface, error) {
	// Construct the service
	return services.NewBuildService(config.ID,
		config.Name, config.Description, build.NewPipelineBuilderFactory()), nil
}
