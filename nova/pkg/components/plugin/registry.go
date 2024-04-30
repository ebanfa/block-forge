package plugin

import (
	"github.com/edward1christian/block-forge/nova/pkg/components/operations"
	"github.com/edward1christian/block-forge/nova/pkg/components/operations/commands"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/component"
	configApi "github.com/edward1christian/block-forge/pkg/application/config"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
)

// Define a struct to represent a component to register
type ComponentRegistration struct {
	ID        string
	FactoryID string
	Factory   component.ComponentFactoryInterface
}

// RegisterComponents registers all components returned by GetComponentsToRegister
func RegisterComponents(ctx *context.Context, system systemApi.SystemInterface) error {
	// Concatenate all component arrays
	allComponentsToRegister := append(
		GetServicesToRegister(), append(
			GetOperationsToRegister(), GetNormalComponentsToRegister()...)...)

	// Register all components
	if err := getListOfComponentsToRegister(ctx, system, allComponentsToRegister); err != nil {
		return err
	}

	return nil
}

// getListOfComponentsToRegister registers a list of components
func getListOfComponentsToRegister(ctx *context.Context, system systemApi.SystemInterface, components []ComponentRegistration) error {
	for _, comp := range components {
		config := &configApi.ComponentConfig{
			ID:        comp.ID,
			FactoryID: comp.FactoryID,
		}
		if err := systemApi.RegisterComponent(ctx, system, config, comp.Factory); err != nil {
			return err
		}
	}
	return nil
}

// GetServicesToRegister returns the list of services to register
func GetServicesToRegister() []ComponentRegistration {
	serviceFactories := map[string]component.ComponentFactoryInterface{
		//"BuildProjectOp":      &commands.BuildProjectOpFactory{},
	}

	return populateComponentRegistrations(serviceFactories)
}

// GetOperationsToRegister returns the list of operations to register
func GetOperationsToRegister() []ComponentRegistration {
	operationFactories := map[string]component.ComponentFactoryInterface{
		"BuildProjectOp":           &commands.BuildProjectOpFactory{},
		"CreateConfigurationOp":    &commands.CreateConfigurationOpFactory{},
		"GenerateArtifactsOp":      &commands.GenerateArtifactsOpFactory{},
		"ListConfigurationsOp":     &commands.ListConfigurationsOpFactory{},
		"AddEntityOp":              &commands.AddEntityOpFactory{},
		"AddMessageOp":             &commands.AddMessageOpFactory{},
		"AddModuleOp":              &commands.AddModuleOpFactory{},
		"AddQueryOp":               &commands.AddQueryOpFactory{},
		"RemoveProjectConfigOp":    &commands.RemoveProjectConfigOpFactory{},
		"RemoveEntityOp":           &commands.RemoveEntityOpFactory{},
		"RemoveMessageOp":          &commands.RemoveMessageOpFactory{},
		"RemoveModuleOp":           &commands.RemoveModuleOpFactory{},
		"RemoveQueryOp":            &commands.RemoveQueryOpFactory{},
		"RunProjectOp":             &commands.RunProjectOpFactory{},
		"ValidateConfigOp":         &commands.ValidateConfigOpFactory{},
		"VisualizeConfigOp":        &commands.VisualizeConfigOpFactory{},
		"InitDirectoriesOperation": &operations.InitDirectoriesOperationFactory{},
	}

	return populateComponentRegistrations(operationFactories)
}

// GetNormalComponentsToRegister returns the list of normal components to register
func GetNormalComponentsToRegister() []ComponentRegistration {
	normalComponentFactories := map[string]component.ComponentFactoryInterface{
		// Add normal components here
	}

	return populateComponentRegistrations(normalComponentFactories)
}

// PopulateComponentRegistrations iterates over the given map of components and populates the component registrations
func populateComponentRegistrations(componentMap map[string]component.ComponentFactoryInterface) []ComponentRegistration {
	var registrations []ComponentRegistration

	for componentName, factory := range componentMap {
		// Get the factory ID
		factoryID := componentName + "Factory"

		// Add the component registration
		registrations = append(registrations, ComponentRegistration{
			ID:        componentName,
			FactoryID: factoryID,
			Factory:   factory,
		})
	}

	return registrations
}
