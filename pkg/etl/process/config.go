package process

import "github.com/edward1christian/block-forge/pkg/application/components"

// ETLProcessConfig represents the configuration for an ETL process.
type ETLProcessConfig struct {
	Components []*components.ComponentConfig // Components configuration for the ETL process
}

// PipelineConfig represents the configuration for a transformation pipeline.
type PipelineConfig struct {
	components.ComponentConfig
	Stages []components.ComponentConfig // Stages configuration for the pipeline
}
