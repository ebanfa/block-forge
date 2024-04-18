package tests

import (
	"time"

	"github.com/edward1christian/block-forge/pkg/application/components"
	"github.com/edward1christian/block-forge/pkg/etl/process"
)

func DummyProcessConfiguration() []*process.ETLProcessConfig {
	component1 := &components.ComponentConfig{
		ID:           "AdapterID",
		Name:         "Adapter",
		Description:  "Extracts data from source",
		FactoryName:  "AdaptorFactory",
		CustomConfig: map[string]interface{}{"param1": "value1", "param2": 123},
	}

	component2 := &components.ComponentConfig{
		ID:           "TransformerID",
		Name:         "Transformer",
		Description:  "Transforms extracted data",
		FactoryName:  "TransformerFactory",
		CustomConfig: map[string]interface{}{"param3": "value3", "param4": 456},
	}

	config := &process.ETLProcessConfig{
		Components: []*components.ComponentConfig{component1, component2},
	}
	return []*process.ETLProcessConfig{config}
}

func DummySystemConfiguration() *components.Configuration {
	// Dummy service configuration
	serviceConfig := &components.ServiceConfiguration{
		ComponentConfig: components.ComponentConfig{
			ID:           "service_id",
			Name:         "service_name",
			Description:  "service_description",
			FactoryName:  "service_factory",
			CustomConfig: nil, // Add custom service configuration if needed
		},
		RetryInterval: 5 * time.Second, // Example retry interval
		CustomConfig:  nil,             // Add custom service configuration if needed
	}

	// Dummy operation configuration
	operationConfig := &components.OperationConfiguration{
		ComponentConfig: components.ComponentConfig{
			ID:           "operation_id",
			Name:         "operation_name",
			Description:  "operation_description",
			FactoryName:  "operation_factory",
			CustomConfig: nil, // Add custom operation configuration if needed
		},
		// Add other operation-specific configuration options if needed
	}

	// Create and return the dummy configuration
	return &components.Configuration{
		Services:     []*components.ServiceConfiguration{serviceConfig},
		Operations:   []*components.OperationConfiguration{operationConfig},
		CustomConfig: DummyProcessConfiguration(), // Add custom configuration if needed
	}
}
