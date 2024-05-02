package database

import (
	"encoding/json"
	"time"
)

// MetadataEntry represents metadata about a project configuration
type MetadataEntry struct {
	ProjectID    string    `json:"project_id"`
	DatabaseName string    `json:"database_name"`
	DatabasePath string    `json:"database_path"`
	CreationDate time.Time `json:"creation_date"`
	LastUpdated  time.Time `json:"last_updated"`
	// Add other metadata fields as needed
}

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

	// WorkingVersion returns the current working version of the database.
	WorkingVersion() int64

	// SaveVersion saves a new version of the database to disk.
	SaveVersion() ([]byte, int64, error)

	// Load loads the latest versioned database from disk.
	Load() (int64, error)

	// LoadVersion loads a specific version of the database from disk.
	LoadVersion(targetVersion int64) (int64, error)

	// String returns a string representation of the database.
	String() (string, error)

	// WorkingHash returns the root hash of the current working database.
	WorkingHash() []byte

	// AvailableVersions returns a list of available versions of the database.
	AvailableVersions() []int

	// IsEmpty checks if the database is empty.
	IsEmpty() bool
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

// WorkingVersion returns the current working version of the database.
func (md *MetadataDatabase) WorkingVersion() int64 {
	return md.db.WorkingVersion()
}

// SaveVersion saves a new version of the database to disk.
func (md *MetadataDatabase) SaveVersion() ([]byte, int64, error) {
	return md.db.SaveVersion()
}

// Load loads the latest versioned database from disk.
func (md *MetadataDatabase) Load() (int64, error) {
	return md.db.Load()
}

// LoadVersion loads a specific version of the database from disk.
func (md *MetadataDatabase) LoadVersion(targetVersion int64) (int64, error) {
	return md.db.LoadVersion(targetVersion)
}

// String returns a string representation of the database.
func (md *MetadataDatabase) String() (string, error) {
	return md.db.String()
}

// WorkingHash returns the root hash of the current working database.
func (md *MetadataDatabase) WorkingHash() []byte {
	return md.db.WorkingHash()
}

// AvailableVersions returns a list of available versions of the database.
func (md *MetadataDatabase) AvailableVersions() []int {
	return md.db.AvailableVersions()
}

// IsEmpty checks if the database is empty.
func (md *MetadataDatabase) IsEmpty() bool {
	return md.db.IsEmpty()
}
