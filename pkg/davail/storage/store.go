package storage

// Storage is an interface that defines the operations for storing and retrieving
// encoded data segments, as well as repairing corrupted or missing segments.
type Storage interface {
	// Store takes the encoded data segments and parity segments and persistently
	// stores them, returning a unique identifier for the stored data.
	// The `options` parameter allows passing additional configuration options to the storage.
	Store(dataSegments, paritySegments [][]byte, options ...StorageOption) (storageID string, err error)

	// Retrieve takes the unique identifier of the stored data and retrieves the
	// encoded data segments and parity segments.
	// The `options` parameter allows passing additional configuration options to the retrieval.
	Retrieve(storageID string, options ...RetrieveOption) (dataSegments, paritySegments [][]byte, err error)

	// Repair takes the unique identifier of the stored data and repairs any
	// corrupted or missing segments, returning the fully recovered encoded data.
	// The `options` parameter allows passing additional configuration options to the repair process.
	Repair(storageID string, options ...RepairOption) (dataSegments, paritySegments [][]byte, err error)

	// Delete removes the stored data with the given unique identifier from the
	// storage system.
	// The `options` parameter allows passing additional configuration options to the deletion.
	Delete(storageID string, options ...DeleteOption) error

	// SetOptions allows setting global options for the storage, which will be used
	// in all subsequent calls to Store, Retrieve, Repair, and Delete.
	SetOptions(options ...StorageOption)
}

// StorageOption is a function type that can be used to set options for the Storage.
type StorageOption func(*StorageConfig)

// RetrieveOption is a function type that can be used to set options for the Retrieve operation.
type RetrieveOption func(*RetrieveConfig)

// RepairOption is a function type that can be used to set options for the Repair operation.
type RepairOption func(*RepairConfig)

// DeleteOption is a function type that can be used to set options for the Delete operation.
type DeleteOption func(*DeleteConfig)
