package etl

import (
	"time"

	"github.com/edward1christian/block-forge/pkg/application/event"
	"github.com/edward1christian/block-forge/pkg/application/system"
	blockchain "github.com/edward1christian/block-forge/pkg/blockchain/interfaces"
)

// ETLConfig represents the configuration for an ETL process.
type ETLConfig struct {
	Adapters  []AdapterConfig  // Adapters configuration for the ETL process
	Pipelines []PipelineConfig // Pipelines configuration for the ETL process
	Relays    []RelayConfig    // Relays configuration for the ETL process
}

// AdapterConfig represents the configuration for a blockchain adapter.
type AdapterConfig struct {
	Name   string                 // Name of the adapter
	Type   string                 // Type of the adapter
	Config map[string]interface{} // Configuration parameters for the adapter
}

// PipelineConfig represents the configuration for a transformation pipeline.
type PipelineConfig struct {
	Name   string        // Name of the pipeline
	Stages []StageConfig // Stages configuration for the pipeline
}

// StageConfig represents the configuration for a pipeline stage.
type StageConfig struct {
	Name   string                 // Name of the stage
	Type   string                 // Type of the stage
	Config map[string]interface{} // Configuration parameters for the stage
}

// RelayConfig represents the configuration for a blockchain relay.
type RelayConfig struct {
	Name   string                 // Name of the relay
	Type   string                 // Type of the relay
	Config map[string]interface{} // Configuration parameters for the relay
}

// ETLManager represents an interface for managing and executing ETL processes.
type ETLManager interface {
	// InitializeETLProcess initializes an ETL process with the provided configuration.
	InitializeETLProcess(config ETLConfig) (*ETLProcess, error)

	// StartETLProcess starts an ETL process with the given ID.
	StartETLProcess(processID string) error

	// StopETLProcess stops an ETL process with the given ID.
	StopETLProcess(processID string) error

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

type ETLSystem interface {
	ETLManager
	system.System
}

// ETLProcess represents an individual ETL process.
type ETLProcess struct {
	ID     string    // Unique identifier of the ETL process
	Config ETLConfig // Configuration of the ETL process
	Status ETLStatus // Status of the ETL process
}

// ETLStatus represents the status of an ETL process.
type ETLStatus struct {
	// Define the necessary fields for the ETL status
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

// ETLComponent represents a generic component in the ETL system.
type ETLComponent interface {
	system.StartableComponent                           // Interface for a startable component
	setProcessID(processID string)                      // Sets the ID of the ETL process the component belongs to
	getProcessID() string                               // Gets the ID of the ETL process the component belongs to
	SubscribeToEvents(eventBus event.EventBusInterface) // Subscribes to relevant events from the event bus
}

// BlockchainAdapter is responsible for extracting data from a source blockchain.
type BlockchainAdapter interface {
	ETLComponent                       // Interface for a blockchain adapter
	Blockchain() blockchain.Blockchain // Returns the associated blockchain instance
}

// TransformationPipeline is responsible for transforming extracted data.
type TransformationPipeline interface {
	ETLComponent                              // Interface for a transformation pipeline
	AddStage(stage TransformationStage) error // Adds a transformation stage to the pipeline
	RemoveStage(stageID string) error         // Removes a transformation stage from the pipeline
}

// TransformationStage represents a single stage in the transformation pipeline.
type TransformationStage interface {
	ETLComponent                                     // Interface for a transformation stage
	Transform(data interface{}) (interface{}, error) // Transforms the input data
}

// BlockchainRelay is responsible for loading data into a target blockchain.
type BlockchainRelay interface {
	ETLComponent                       // Interface for a blockchain relay
	Blockchain() blockchain.Blockchain // Returns the associated blockchain instance
}

// EventType represents the type of an event.
type EventType string

const (
	// EventTypeDataExtracted represents an event emitted when data is extracted from a source blockchain.
	EventTypeDataExtracted EventType = "data_extracted"

	// EventTypeDataTransformed represents an event emitted when data is transformed by a pipeline stage.
	EventTypeDataTransformed EventType = "data_transformed"

	// EventTypeDataLoaded represents an event emitted when data is loaded into a target blockchain.
	EventTypeDataLoaded EventType = "data_loaded"
)
