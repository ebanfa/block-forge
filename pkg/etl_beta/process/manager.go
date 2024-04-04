package process

import (
	"time"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
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

// ETLManager represents an interface for managing and executing ETL processes.
type ETLManager interface {
	// InitializeETLProcess initializes an ETL process with the provided configuration.
	InitializeETLProcess(ctx *context.Context, config *components.ETLProcessConfig) (*ETLProcess, error)

	// StartETLProcess starts an ETL process with the given ID.
	StartETLProcess(ctx *context.Context, processID string) error

	// StopETLProcess stops an ETL process with the given ID.
	StopETLProcess(ctx *context.Context, processID string) error

	// GetETLProcess retrieves an ETL process by its ID.
	GetETLProcess(processID string) (*ETLProcess, error)

	// GetAllETLProcesses retrieves all ETL processes.
	GetAllETLProcesses() []*ETLProcess

	// ScheduleETLProcess schedules an ETL process for execution.
	ScheduleETLProcess(process *ETLProcess, schedule Schedule) error

	// GetScheduledETLProcesses retrieves all scheduled ETL processes.
	GetScheduledETLProcesses() []*ScheduledETLProcess

	// RemoveScheduledETLProcess removes a scheduled ETL process by its ID.
	RemoveScheduledETLProcess(processID string) error
}

type ETLManagerService interface {
	ETLManager
	system.SystemComponentInterface
}

// ScheduledETLProcess represents an ETL process scheduled for execution at specific intervals.
type ScheduledETLProcess struct {
	Process  *ETLProcess // Pointer to the scheduled ETL process
	Schedule Schedule    // Schedule for the process execution
}

// Schedule represents a schedule for task execution.
type Schedule interface {
	Next() time.Time // Returns the next time the task should be executed
}
