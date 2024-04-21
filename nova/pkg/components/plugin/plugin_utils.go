package plugin

import (
	"fmt"

	"github.com/edward1christian/block-forge/nova/pkg/components/factories"
	"github.com/edward1christian/block-forge/pkg/application/component"
	configApi "github.com/edward1christian/block-forge/pkg/application/config"
)

// RegisterAndCreateComponent registers and creates a component using the provided registrar.
func RegisterAndCreateComponent(registrar component.ComponentRegistrarInterface, config *configApi.ComponentConfig) error {
	// Register the factory for the component
	err := registrar.RegisterFactory(config.FactoryID, &factories.BuilderServiceFactory{})
	if err != nil {
		return fmt.Errorf("failed to register factory %s: %w", config.FactoryID, err)
	}

	// Create and register the component
	_, err = registrar.CreateComponent(&configApi.ComponentConfig{
		ID:        config.ID,
		FactoryID: config.FactoryID,
	})

	if err != nil {
		return fmt.Errorf("failed to create and register component %s: %w", config.ID, err)
	}

	return nil
}
