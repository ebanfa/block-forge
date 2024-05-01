package database

import (
	"encoding/json"
	"time"
)

// MetadataDatabaseInterface defines methods for interacting with a metadata database.
type MetadataDatabaseInterface interface {
	// Insert inserts a new metadata entry into the metadata database
	Insert(entry *MetadataEntry) error

	// Get retrieves the metadata entry for the given project ID from the metadata database
	Get(projectID string) (*MetadataEntry, error)

	// GetAll retrieves all metadata entries from the metadata database
	GetAll() ([]*MetadataEntry, error)

	// Update updates an existing metadata entry in the metadata database
	Update(entry *MetadataEntry) error

	// Delete deletes the metadata entry for the given project ID from the metadata database
	Delete(projectID string) error
}

// MetadataEntry represents metadata about a project configuration
type MetadataEntry struct {
	ProjectID    string    `json:"project_id"`
	DatabaseName string    `json:"database_name"`
	DatabasePath string    `json:"database_path"`
	CreationDate time.Time `json:"creation_date"`
	LastUpdated  time.Time `json:"last_updated"`
	// Add other metadata fields as needed
}

// MetadataDatabase represents the metadata database stored using the Database interface
type MetadataDatabase struct {
	db Database
}

// NewMetadataDatabase creates a new MetadataDatabase instance
func NewMetadataDatabase(db Database) *MetadataDatabase {
	return &MetadataDatabase{db: db}
}

// Insert inserts a new metadata entry into the metadata database
func (md *MetadataDatabase) Insert(entry *MetadataEntry) error {
	// Serialize the entry
	serializedEntry, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	// Insert into the database
	return md.db.Set([]byte(entry.ProjectID), serializedEntry)
}

// Get retrieves the metadata entry for the given project ID from the metadata database
func (md *MetadataDatabase) Get(projectID string) (*MetadataEntry, error) {
	// Retrieve from the database
	data, err := md.db.Get([]byte(projectID))
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

// GetAll retrieves all metadata entries from the metadata database
func (md *MetadataDatabase) GetAll() ([]*MetadataEntry, error) {
	// Initialize a slice to store all entries
	var entries []*MetadataEntry

	// Iterate over all keys in the database
	err := md.db.Iterate(func(key, value []byte) bool {
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

// Update updates an existing metadata entry in the metadata database
func (md *MetadataDatabase) Update(entry *MetadataEntry) error {
	// Serialize the entry
	serializedEntry, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	// Update in the database
	return md.db.Set([]byte(entry.ProjectID), serializedEntry)
}

// Delete deletes the metadata entry for the given project ID from the metadata database
func (md *MetadataDatabase) Delete(projectID string) error {
	// Delete from the database
	return md.db.Delete([]byte(projectID))
}
