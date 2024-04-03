package components

import (
	"github.com/edward1christian/block-forge/pkg/application/system"
	blockchain "github.com/edward1christian/block-forge/pkg/blockchain/interfaces"
)

// ETLProcessComponent is an component that belongs to an ETL process
type ETLProcessComponent interface {
	system.Startable       // Interface for a startable component
	system.SystemComponent // Interface for a system component
	GetProcessID() string  // Gets the ID of the ETL process the component belongs to
}

// ETLProcessConfig represents the configuration for an ETL process.
type ETLProcessConfig struct {
	Components []*system.ComponentConfig // Components configuration for the ETL process
}

// PipelineConfig represents the configuration for a transformation pipeline.
type PipelineConfig struct {
	system.ComponentConfig
	Stages []system.ComponentConfig // Stages configuration for the pipeline
}

// BlockchainAdapter is responsible for extracting data from a source blockchain.
type BlockchainAdapter interface {
	ETLProcessComponent                 // Is an ETL process component
	Blockchain() *blockchain.Blockchain // Returns the associated blockchain instance
}

// TransformationPipeline is responsible for transforming extracted data.
type TransformationPipeline interface {
	ETLProcessComponent                        // Is an ETL process component
	AddStage(stage *TransformationStage) error // Adds a transformation stage to the pipeline
	RemoveStage(stageID string) error          // Removes a transformation stage from the pipeline
}

// TransformationStage represents a single stage in the transformation pipeline.
type TransformationStage interface {
	ETLProcessComponent                              // Is an ETL process component
	Transform(data interface{}) (interface{}, error) // Transforms the input data
}

// BlockchainRelay is responsible for loading data into a target blockchain.
type BlockchainRelay interface {
	ETLProcessComponent                 // Is an ETL process component
	Blockchain() *blockchain.Blockchain // Returns the associated blockchain instance
}
