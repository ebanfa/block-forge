package store

import (
	"encoding/json"

	dbApi "github.com/edward1christian/block-forge/pkg/application/db"
)

// MetadataEntry represents metadata about a project configuration.
type MetadataEntry struct {
	ProjectID    string `json:"project_id"`
	DatabaseName string `json:"database_name"`
	DatabasePath string `json:"database_path"`
	// Add other metadata fields as needed
}

// MetadataStore defines methods for interacting with MetadataEntry instances.
type MetadataStore interface {
	// InsertMetadata inserts a new MetadataEntry into the database.
	InsertMetadata(entry *MetadataEntry) error

	// GetMetadata retrieves the MetadataEntry with the given project ID from the database.
	GetMetadata(projectID string) (*MetadataEntry, error)

	// GetAllMetadata retrieves all MetadataEntry instances from the database.
	GetAllMetadata() ([]*MetadataEntry, error)

	// UpdateMetadata updates an existing MetadataEntry in the database.
	UpdateMetadata(entry *MetadataEntry) error

	// DeleteMetadata deletes the MetadataEntry with the given project ID from the database.
	DeleteMetadata(projectID string) error

	// SaveVersion saves a new tree version to disk.
	SaveVersion() ([]byte, int64, error)

	// Load loads the latest versioned tree from disk.
	Load() (int64, error)

	// LoadVersion loads a specific version of the tree from disk.
	LoadVersion(targetVersion int64) (int64, error)

	// String returns a string representation of the tree.
	String() (string, error)

	// WorkingVersion returns the current working version of the tree.
	WorkingVersion() int64

	// WorkingHash returns the root hash of the current working tree.
	WorkingHash() []byte

	// AvailableVersions returns a list of available versions.
	AvailableVersions() []int

	// IsEmpty checks if the database is empty.
	IsEmpty() bool

	// Close closes the database.
	Close() error
}

// MetadataStoreImpl represents a store for managing MetadataEntry instances.
type MetadataStoreImpl struct {
	db dbApi.Database
}

// NewMetadataStore creates a new MetadataStoreImpl instance with the provided Database.
func NewMetadataStore(db dbApi.Database) MetadataStore {
	return &MetadataStoreImpl{db: db}
}

// InsertMetadata inserts a new MetadataEntry into the database.
func (ms *MetadataStoreImpl) InsertMetadata(entry *MetadataEntry) error {
	// Serialize the entry
	serializedEntry, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	// Insert into the database
	return ms.db.Set([]byte(entry.ProjectID), serializedEntry)
}

// GetMetadata retrieves the MetadataEntry with the given project ID from the database.
func (ms *MetadataStoreImpl) GetMetadata(projectID string) (*MetadataEntry, error) {
	// Retrieve from the database
	data, err := ms.db.Get([]byte(projectID))
	if err != nil {
		return nil, err
	}
	// Deserialize the entry
	var entry MetadataEntry
	if err := json.Unmarshal(data, &entry); err != nil {
		return nil, err
	}
	return &entry, nil
}

// GetAllMetadata retrieves all MetadataEntry instances from the database.
func (ms *MetadataStoreImpl) GetAllMetadata() ([]*MetadataEntry, error) {
	// Initialize a slice to store all entries
	var entries []*MetadataEntry

	// Iterate over all keys in the database
	err := ms.db.Iterate(func(key, value []byte) bool {
		// Deserialize the value into a MetadataEntry
		var entry MetadataEntry
		if err := json.Unmarshal(value, &entry); err != nil {
			// Return true to stop iteration if unmarshaling fails
			return true
		}
		// Append the entry to the slice
		entries = append(entries, &entry)
		// Continue iteration
		return false
	})

	if err != nil {
		return entries, err
	}

	return entries, nil
}

// UpdateMetadata updates an existing MetadataEntry in the database.
func (ms *MetadataStoreImpl) UpdateMetadata(entry *MetadataEntry) error {
	// Serialize the entry
	serializedEntry, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	// Update in the database
	return ms.db.Set([]byte(entry.ProjectID), serializedEntry)
}

// DeleteMetadata deletes the MetadataEntry with the given project ID from the database.
func (ms *MetadataStoreImpl) DeleteMetadata(projectID string) error {
	// Delete from the database
	return ms.db.Delete([]byte(projectID))
}

// SaveVersion saves a new tree version to disk.
func (ms *MetadataStoreImpl) SaveVersion() ([]byte, int64, error) {
	return ms.db.SaveVersion()
}

// Load loads the latest versioned tree from disk.
func (ms *MetadataStoreImpl) Load() (int64, error) {
	return ms.db.Load()
}

// LoadVersion loads a specific version of the tree from disk.
func (ms *MetadataStoreImpl) LoadVersion(targetVersion int64) (int64, error) {
	return ms.db.LoadVersion(targetVersion)
}

// String returns a string representation of the tree.
func (ms *MetadataStoreImpl) String() (string, error) {
	return ms.db.String()
}

// WorkingVersion returns the current working version of the tree.
func (ms *MetadataStoreImpl) WorkingVersion() int64 {
	return ms.db.WorkingVersion()
}

// WorkingHash returns the root hash of the current working tree.
func (ms *MetadataStoreImpl) WorkingHash() []byte {
	return ms.db.WorkingHash()
}

// AvailableVersions returns a list of available versions.
func (ms *MetadataStoreImpl) AvailableVersions() []int {
	return ms.db.AvailableVersions()
}

// IsEmpty checks if the database is empty.
func (ms *MetadataStoreImpl) IsEmpty() bool {
	return ms.db.IsEmpty()
}

// Close closes the database.
func (ms *MetadataStoreImpl) Close() error {
	return ms.db.Close()
}
