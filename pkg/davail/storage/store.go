package storage

// StorageConfig is a struct to hold configuration options for the Storage.
type StorageConfig struct {
	// Add storage configuration options here as needed
}

// StorageOption is a function type that can be used to set options for the Storage.
type StorageOption func(*StorageConfig)

// Storage is an interface that defines the operations for storing and retrieving
// encoded data segments, as well as repairing corrupted or missing segments.
type StorageInterface interface {
	// Store takes the encoded data segments and parity segments and persistently
	// stores them, returning a unique identifier for the stored data.
	Store(dataSegments, paritySegments [][]byte) (storageID string, err error)

	// Retrieve takes the unique identifier of the stored data and retrieves the
	// encoded data segments and parity segments.
	Retrieve(storageID string) (dataSegments, paritySegments [][]byte, err error)

	// Repair takes the unique identifier of the stored data and repairs any
	// corrupted or missing segments, returning the fully recovered encoded data.
	Repair(storageID string) (dataSegments, paritySegments [][]byte, err error)

	// Delete removes the stored data with the given unique identifier from the
	// storage system.
	Delete(storageID string) error

	// SetOptions allows setting global options for the storage, which will be used
	// in all subsequent calls to Store, Retrieve, Repair, and Delete.
	SetOptions(options ...StorageOption)
}
