package etl

import (
	"time"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
	blockchain "github.com/edward1christian/block-forge/pkg/blockchain/interfaces"
)

// ETLComponent is an component that belongs to an ETL process
type ETLComponent interface {
	system.Startable               // Interface for a startable component
	system.SystemComponent         // Interface for a system component
	setProcessID(processID string) // Sets the ID of the ETL process the component belongs to
	getProcessID() string          // Gets the ID of the ETL process the component belongs to
}

// ETLConfig represents the configuration for an ETL process.
type ETLConfig struct {
	Components []*ETLComponentConfig // Components configuration for the ETL process
}

// ETLComponentConfig represents the configuration for a blockchain adapter.
type ETLComponentConfig struct {
	Name      string                 // Name of the component
	Type      string                 // Type of the component (Adapter, Pipeline, Relay)
	FactoryNm string                 // The name of the factory that creates instances of this component
	Config    map[string]interface{} // Configuration parameters for the adapter
}

// PipelineConfig represents the configuration for a transformation pipeline.
type PipelineConfig struct {
	ETLComponentConfig
	Stages []ETLComponentConfig // Stages configuration for the pipeline
}

// ETLComponentFactory is a function type for creating an ETLComponent.
type ETLComponentFactory func(ctx *context.Context, config *ETLComponentConfig) (ETLComponent, error)

// ETLProcess represents an individual ETL process.
type ETLProcess struct {
	ID         string                  // Unique identifier of the ETL process
	Config     *ETLConfig              // Configuration of the ETL process
	Status     ETLStatus               // Status of the ETL process
	Components map[string]ETLComponent // Map to track instantiated components by name
}

// ETLManager represents an interface for managing and executing ETL processes.
type ETLManager interface {
	// InitializeETLProcess initializes an ETL process with the provided configuration.
	InitializeETLProcess(ctx *context.Context, config *ETLConfig) (*ETLProcess, error)

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

	// RegisterAdapterFactory registers an adapter factory with the ETL system.
	RegisterETLComponentFactory(name string, factory ETLComponentFactory) error

	// GetAdapterFactory retrieves the factory for creating adapters by name.
	GetETLComponentFactory(name string) (ETLComponentFactory, bool)
}

type ETLSystem interface {
	ETLManager
	system.System
}

// ETLStatus represents the status of an ETL process.
type ETLStatus string

const (
	ETLStatusInitialized ETLStatus = "initialized"
	ETLStatusRunning     ETLStatus = "running"
	ETLStatusPaused      ETLStatus = "paused"
	ETLStatusCompleted   ETLStatus = "completed"
	ETLStatusFailed      ETLStatus = "failed"
	ETLStatusStopped     ETLStatus = "stopped"
)

// ScheduledETLProcess represents an ETL process scheduled for execution at specific intervals.
type ScheduledETLProcess struct {
	Process  *ETLProcess // Pointer to the scheduled ETL process
	Schedule Schedule    // Schedule for the process execution
}

// Schedule represents a schedule for task execution.
type Schedule interface {
	Next() time.Time // Returns the next time the task should be executed
}

// BlockchainAdapter is responsible for extracting data from a source blockchain.
type BlockchainAdapter interface {
	ETLComponent                        // Is an ETL process component
	Blockchain() *blockchain.Blockchain // Returns the associated blockchain instance
}

// TransformationPipeline is responsible for transforming extracted data.
type TransformationPipeline interface {
	ETLComponent                               // Is an ETL process component
	AddStage(stage *TransformationStage) error // Adds a transformation stage to the pipeline
	RemoveStage(stageID string) error          // Removes a transformation stage from the pipeline
}

// TransformationStage represents a single stage in the transformation pipeline.
type TransformationStage interface {
	ETLComponent                                     // Is an ETL process component
	Transform(data interface{}) (interface{}, error) // Transforms the input data
}

// BlockchainRelay is responsible for loading data into a target blockchain.
type BlockchainRelay interface {
	ETLComponent                        // Is an ETL process component
	Blockchain() *blockchain.Blockchain // Returns the associated blockchain instance
}
