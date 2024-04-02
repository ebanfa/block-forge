package etl

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
