package etl

import (
	"time"

	"github.com/edward1christian/block-forge/pkg/application/component"
	blockchain "github.com/edward1christian/block-forge/pkg/blockchain/interfaces"
)

// BlockchainAdapter is responsible for extracting data from a source blockchain.
type BlockchainAdapter interface {
	component.StartableComponent

	// Blockchain returns the associated blockchain instance.
	Blockchain() blockchain.Blockchain
}

// TransformationPipeline is responsible for transforming extracted data.
type TransformationPipeline interface {
	component.StartableComponent

	// AddStage adds a new transformation stage to the pipeline.
	AddStage(stage TransformationStage) error

	// RemoveStage removes a transformation stage from the pipeline.
	RemoveStage(stageID string) error
}

// TransformationStage represents a single stage in the transformation pipeline.
type TransformationStage interface {
	component.Component

	// Transform performs the transformation logic on the input data.
	Transform(data interface{}) (interface{}, error)

	// Stop stops the stage, terminating any ongoing processing.
	Stop() error
}

// BlockchainRelay is responsible for loading data into a target blockchain.
type BlockchainRelay interface {
	component.StartableComponent

	// Blockchain returns the associated blockchain instance.
	Blockchain() blockchain.Blockchain
}

// EventPublisher provides a way to publish events to the event bus.
type EventPublisher interface {
	// PublishEvent publishes an event to the event bus.
	PublishEvent(topic string, event interface{}) error
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

// OrchestrationFramework is responsible for coordinating and executing ETL processes.
type OrchestrationFramework interface {
	// Configure configures the orchestration framework with the provided ETL configuration.
	Configure(config ETLConfig) error

	// ExecuteETL executes an ETL process based on the provided configuration.
	ExecuteETL(config ETLConfig) error

	// MonitorETL monitors the progress and status of an ETL process.
	MonitorETL(etlID string) (ETLStatus, error)
}

// ETLStatus represents the status of an ETL process.
type ETLStatus struct {
	// ... (Define the necessary fields for the ETL status)
}

// TaskScheduler is responsible for scheduling and executing ETL processes.
type TaskScheduler interface {
	// ScheduleTask schedules a new task for execution.
	ScheduleTask(task Task) error

	// CancelTask cancels a scheduled task.
	CancelTask(taskID string) error

	// ListTasks returns a list of scheduled tasks.
	ListTasks() []Task
}

// Task represents a scheduled task.
type Task interface {
	// ID returns the unique identifier of the task.
	ID() string

	// Schedule returns the schedule for the task.
	Schedule() Schedule

	// Execute executes the task.
	Execute() error
}

// Schedule represents a schedule for task execution.
type Schedule interface {
	// Next returns the next time the task should be executed.
	Next() time.Time
}
