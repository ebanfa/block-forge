package database

import (
	"encoding/json"
	"log"
	"sync"
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

var (
	metaDB     *MetadataDatabase
	metaDBOnce sync.Once
)

// GetMetadataDBInstance returns the singleton instance of MetadataDatabase
func GetMetadataDBInstance(name, path string) *MetadataDatabase {
	metaDBOnce.Do(func() {
		// Initialize the metadata database with the appropriate database instance
		db, err := InitializeLevelDB(name, path)
		if err != nil {
			// Handle error if needed
			log.Fatal(err)
		}

		iavlDB := GetIAVLDatabase(db)
		metaDB = NewMetadataDatabase(iavlDB)
	})
	return metaDB
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
