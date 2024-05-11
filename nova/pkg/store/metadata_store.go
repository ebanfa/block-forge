package store

import (
	"encoding/json"

	"github.com/edward1christian/block-forge/pkg/application/store"
)

// MetadataEntry represents metadata about a project configuration.
type MetadataEntry struct {
	ProjectID    string `json:"project_id"`
	ProjectName  string `json:"project_name"`
	DatabaseName string `json:"database_name"`
	DatabasePath string `json:"database_path"`
	// Add other metadata fields as needed
}

// MetadataStoreImpl represents a specialized Store for managing MetadataEntry instances.
type MetadataStoreImpl struct {
	store.Store // Embedding store.Store interface
}

// NewMetadataStore creates a new MetadataStoreImpl instance with the provided Store.
func NewMetadataStore(store store.Store) *MetadataStoreImpl {
	return &MetadataStoreImpl{Store: store}
}

// InsertMetadata inserts a new MetadataEntry into the database.
func (ms *MetadataStoreImpl) InsertMetadata(entry *MetadataEntry) error {
	// Serialize the entry
	serializedEntry, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	// Create a key based on the project ID
	key := []byte(entry.ProjectID)

	// Use the Store to insert the data
	err = ms.Set(key, serializedEntry)
	return err
}

// GetMetadata retrieves the MetadataEntry with the given project ID from the database.
func (ms *MetadataStoreImpl) GetMetadata(projectID string) (*MetadataEntry, error) {
	// Retrieve from the database using the Store
	data, err := ms.Get([]byte(projectID))
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
	err := ms.Iterate(func(key, value []byte) bool {
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
	// Update in the database using the Store
	err = ms.Set([]byte(entry.ProjectID), serializedEntry)
	return err
}

// DeleteMetadata deletes the MetadataEntry with the given project ID from the database.
func (ms *MetadataStoreImpl) DeleteMetadata(projectID string) error {
	// Delete from the database using the Store
	err := ms.Delete([]byte(projectID))
	return err
}
