// Package process provides a comprehensive API for Extract, Transform, and Load (ETL) operations
// on blockchain data. It allows you to extract data from various blockchain sources,
// transform it into a desired format, and load it into centralized data stores or other
// blockchain destinations.
package process

import (
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/components"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
	etlComponentsApi "github.com/edward1christian/block-forge/pkg/etl/components"
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

// Source represents a data source from which data can be extracted.
type Source interface {
	// GetID returns the unique identifier of the source.
	GetID() string

	// GetType returns the type of the source (e.g., blockchain, database, file).
	GetType() string

	// GetConfig returns the configuration parameters for the source.
	GetConfig() map[string]interface{}
}

// Destination represents a destination where data can be loaded.
type Destination interface {
	// GetID returns the unique identifier of the destination.
	GetID() string

	// GetType returns the type of the destination (e.g., data warehouse, blockchain).
	GetType() string

	// GetConfig returns the configuration parameters for the destination.
	GetConfig() map[string]interface{}
}

// Record represents a single data record in the ETL pipeline.
type Record struct {
	// Data holds the actual data of the record.
	Data interface{}

	// Metadata contains additional metadata associated with the record.
	Metadata map[string]interface{}
}

// ETLProcessComponent is an component that belongs to an ETL process
type ETLProcessComponent interface {
	// ID returns the unique identifier of the component.
	ID() string

	// Name returns the name of the component.
	Name() string

	// Type returns the type of the component.
	Type() components.ComponentType

	// Description returns the description of the component.
	Description() string

	// Initialize initializes the module.
	// Returns an error if the initialization fails.
	Initialize(ctx *context.Context, system systemApi.SystemInterface) error

	GetProcessID() string // Gets the ID of the ETL process the component belongs to
}

// Extractor defines the interface for extracting data from a source.
type Extractor interface {
	ETLProcessComponent
	// Extract retrieves data from the specified source and returns it as a slice of Records.
	Extract(ctx *context.Context, source Source) ([]Record, error)
}

// StreamingExtractor defines the interface for extracting data from a source in a streaming manner.
type StreamingExtractor interface {
	Extractor

	// OpenStream establishes a streaming connection with the data source and returns a channel
	// for receiving records as they become available.
	OpenStream(ctx context.Context, source Source) (<-chan Record, error)
}

// ScheduledExtractor defines the interface for extracting data from a source on a scheduled basis.
type ScheduledExtractor interface {
	Extractor

	// ScheduleExtraction schedules a periodic extraction from the data source based on the provided schedule.
	ScheduleExtraction(ctx context.Context, source Source, schedule Schedule) error

	// ExtractionsChannel returns the channel from which scheduled extractions can be read.
	ExtractionsChannel() <-chan []Record
}

// Transformer defines the interface for transforming data.
type Transformer interface {
	ETLProcessComponent
	// Transform applies transformations to the input data and returns the transformed Records.
	Transform(ctx *context.Context, records []Record) ([]Record, error)
}

// Loader defines the interface for loading data into a destination.
type Loader interface {
	ETLProcessComponent
	// Load writes the provided Records to the specified destination.
	Load(ctx *context.Context, destination Destination, records []Record) error
}

// ETLProcess represents an individual ETL process.
type ETLProcess struct {
	ID         string                                          // Unique identifier of the ETL process
	Config     *ETLProcessConfig                               // Configuration of the ETL process
	Status     ETLProcessStatus                                // Status of the ETL process
	Components map[string]etlComponentsApi.ETLProcessComponent // Map to track instantiated components by name
}
