package process

import (
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
	etlComponentsApi "github.com/edward1christian/block-forge/pkg/etl/components"
)

// ETLProcessStatus represents the status of an ETL process.
type ETLProcessStatus string

const (
	ETLProcessStatusUnInitialized ETLProcessStatus = "Uninitialized"
	ETLProcessStatusInitialized   ETLProcessStatus = "initialized"
	ETLProcessStatusRunning       ETLProcessStatus = "running"
	ETLProcessStatusPaused        ETLProcessStatus = "paused"
	ETLProcessStatusCompleted     ETLProcessStatus = "completed"
	ETLProcessStatusFailed        ETLProcessStatus = "failed"
	ETLProcessStatusStopped       ETLProcessStatus = "stopped"
)

// SourceInterface represents a data source from which data can be extracted.
type SourceInterface interface {
	// GetID returns the unique identifier of the source.
	GetID() string

	// GetType returns the type of the source (e.g., blockchain, database, file).
	GetType() string

	// GetConfig returns the configuration parameters for the source.
	GetConfig() map[string]interface{}
}

// DestinationInterface represents a destination where data can be loaded.
type DestinationInterface interface {
	// GetID returns the unique identifier of the destination.
	GetID() string

	// GetType returns the type of the destination (e.g., data warehouse, blockchain).
	GetType() string

	// GetConfig returns the configuration parameters for the destination.
	GetConfig() map[string]interface{}
}

// RecordInterface represents a single data record in the ETL pipeline.
type RecordInterface struct {
	// Data holds the actual data of the record.
	Data interface{}

	// Metadata contains additional metadata associated with the record.
	Metadata map[string]interface{}
}

// ETLProcessComponentInterface is a component that belongs to an ETL process.
type ETLProcessComponentInterface interface {
	system.BaseSystemService
}

// ExtractorInterface defines the interface for extracting data from a source.
type ExtractorInterface interface {
	// Extract retrieves data from the specified source and returns it as a slice of Records.
	Extract(ctx *context.Context, source SourceInterface) ([]RecordInterface, error)
}

// StreamingExtractor defines the interface for extracting data from a source in a streaming manner.
type StreamingExtractorInterface interface {
	ExtractorInterface

	// OpenStream establishes a streaming connection with the data source and returns a channel
	// for receiving records as they become available.
	OpenStream(ctx context.Context, source SourceInterface) (<-chan RecordInterface, error)
}

// ScheduledExtractor defines the interface for extracting data from a source on a scheduled basis.
type ScheduledExtractor interface {
	ExtractorInterface

	// ScheduleExtraction schedules a periodic extraction from the data source based on the provided schedule.
	ScheduleExtraction(ctx context.Context, source SourceInterface, schedule Schedule) error

	// ExtractionsChannel returns the channel from which scheduled extractions can be read.
	ExtractionsChannel() <-chan []RecordInterface
}

// TransformerInterface defines the interface for transforming data.
type TransformerInterface interface {
	// Transform applies transformations to the input data and returns the transformed Records.
	Transform(ctx *context.Context, records []RecordInterface) ([]RecordInterface, error)
}

// LoaderInterface defines the interface for loading data into a destination.
type LoaderInterface interface {
	// Load writes the provided Records to the specified destination.
	Load(ctx *context.Context, destination DestinationInterface, records []RecordInterface) error
}

// ETLProcess represents an individual ETL process.
type ETLProcess struct {
	ID         string                                          // Unique identifier of the ETL process
	Config     *ETLProcessConfig                               // Configuration of the ETL process
	Status     ETLProcessStatus                                // Status of the ETL process
	Components map[string]etlComponentsApi.ETLProcessComponent // Map to track instantiated components by name
}
