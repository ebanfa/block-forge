package process

import (
	"github.com/edward1christian/block-forge/pkg/etl_beta/components"
)

// ETLProcessStatus represents the status of an ETL process.
type ETLProcessStatus string

const (
	ETLProcessStatusInitialized ETLProcessStatus = "initialized"
	ETLProcessStatusRunning     ETLProcessStatus = "running"
	ETLProcessStatusPaused      ETLProcessStatus = "paused"
	ETLProcessStatusCompleted   ETLProcessStatus = "completed"
	ETLProcessStatusFailed      ETLProcessStatus = "failed"
	ETLProcessStatusStopped     ETLProcessStatus = "stopped"
)

// ETLProcess represents an individual ETL process.
type ETLProcess struct {
	ID         string                                    // Unique identifier of the ETL process
	Config     *components.ETLProcessConfig              // Configuration of the ETL process
	Status     ETLProcessStatus                          // Status of the ETL process
	Components map[string]components.ETLProcessComponent // Map to track instantiated components by name
}
