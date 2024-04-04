package process

import (
	"time"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/edward1christian/block-forge/pkg/etl_beta/components"
)

// ETLManager represents an interface for managing and executing ETL processes.
type ProcessManagerInterface interface {
	// InitializeETLProcess initializes an ETL process with the provided configuration.
	InitializeETLProcess(ctx *context.Context, config *components.ETLProcessConfig) (*ETLProcess, error)

	// StartETLProcess starts the ETL process with the given ID.
	StartETLProcess(ctx *context.Context, processID string) error

	// StopETLProcess stops the ETL process with the given ID.
	StopETLProcess(ctx *context.Context, processID string) error

	// RestartETLProcess restarts the ETL process with the given ID.
	RestartETLProcess(ctx *context.Context, processID string) error

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

type ProcessManagerServiceInterface interface {
	ProcessManagerInterface
	system.SystemServiceInterface
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
