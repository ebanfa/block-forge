package component

import "github.com/edward1christian/block-forge/pkg/application/common/context"

// RegisterFactoryAndCreateComponent is a utility function to register a factory and create a component in a single call.
func RegisterFactoryAndCreateComponent(
	ctx *context.Context,
	registrar ComponentRegistrarInterface,
	factoryInfo *FactoryRegistrationInfo,
	componentInfo *ComponentCreationInfo) (ComponentInterface, error) {

	// Register the factory
	if err := registrar.RegisterFactory(ctx, factoryInfo); err != nil {
		return nil, err
	}

	// Create the component using the registered factory
	component, err := registrar.CreateComponent(ctx, componentInfo)
	if err != nil {
		// If component creation fails, unregister the factory to maintain consistency
		err := registrar.UnregisterFactory(ctx, factoryInfo.ID)
		return nil, err
	}

	return component, nil
}
