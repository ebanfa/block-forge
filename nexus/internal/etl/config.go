package etl

// ETLConfig represents the configuration for an ETL process.
type ETLConfig struct {
	// Adapters specifies the configuration for blockchain adapters.
	Adapters []AdapterConfig

	// Pipelines specifies the configuration for transformation pipelines.
	Pipelines []PipelineConfig

	// Relays specifies the configuration for blockchain relays.
	Relays []RelayConfig
}

// AdapterConfig represents the configuration for a blockchain adapter.
type AdapterConfig struct {
	// Name specifies the name of the adapter.
	Name string

	// Type specifies the type of the adapter.
	Type string

	// Config specifies the configuration parameters for the adapter.
	Config map[string]interface{}
}

// PipelineConfig represents the configuration for a transformation pipeline.
type PipelineConfig struct {
	// Name specifies the name of the pipeline.
	Name string

	// Stages specifies the configuration for pipeline stages.
	Stages []StageConfig
}

// StageConfig represents the configuration for a pipeline stage.
type StageConfig struct {
	// Name specifies the name of the stage.
	Name string

	// Type specifies the type of the stage.
	Type string

	// Config specifies the configuration parameters for the stage.
	Config map[string]interface{}
}

// RelayConfig represents the configuration for a blockchain relay.
type RelayConfig struct {
	// Name specifies the name of the relay.
	Name string

	// Type specifies the type of the relay.
	Type string

	// Config specifies the configuration parameters for the relay.
	Config map[string]interface{}
}
