package process

import (
	"time"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// ETLManager represents an interface for managing and executing ETL processes.
type ProcessManagerInterface interface {
	// InitializeProcess initializes an ETL process with the provided configuration.
	InitializeProcess(ctx *context.Context, config *ETLProcessConfig) (*ETLProcess, error)

	// StartProcess starts the ETL process with the given ID.
	StartProcess(ctx *context.Context, processID string) error

	// StopProcess stops the ETL process with the given ID.
	StopProcess(ctx *context.Context, processID string) error

	// RestartProcess restarts the ETL process with the given ID.
	RestartProcess(ctx *context.Context, processID string) error

	// GetProcess retrieves an ETL process by its ID.
	GetProcess(processID string) (*ETLProcess, error)

	// GetAllProcesses retrieves all ETL processes.
	GetAllProcesses() []*ETLProcess

	// RemoveScheduledProcess removes a scheduled ETL process by its ID.
	RemoveProcess(processID string) error
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
