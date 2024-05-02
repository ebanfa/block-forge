package database

import (
	"encoding/json"

	"github.com/edward1christian/block-forge/nova/pkg/config"
)

// ProjectDatabaseInterface defines methods for interacting with a project database.
type ProjectDatabaseInterface interface {
	// Insert inserts a new project entry into the project database
	Insert(project *config.Project) error

	// Get retrieves the project entry for the given ID from the project database
	Get(projectID string) (*config.Project, error)

	// GetAll retrieves all project entries from the project database
	GetAll() ([]*config.Project, error)

	// Update updates an existing project entry in the project database
	Update(project *config.Project) error

	// Delete deletes the project entry for the given ID from the project database
	Delete(projectID string) error

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
}

// ProjectDatabase represents the project database stored using the Database interface
type ProjectDatabase struct {
	db Database
}

// NewProjectDatabase creates a new ProjectDatabase instance
func NewProjectDatabase(db Database) *ProjectDatabase {
	return &ProjectDatabase{db: db}
}

// Insert inserts a new project entry into the project database
func (pd *ProjectDatabase) Insert(project *config.Project) error {
	// Serialize the project
	serializedProject, err := json.Marshal(project)
	if err != nil {
		return err
	}
	// Insert into the database
	return pd.db.Set([]byte(project.ID), serializedProject)
}

// Get retrieves the project entry for the given ID from the project database
func (pd *ProjectDatabase) Get(projectID string) (*config.Project, error) {
	// Retrieve from the database
	data, err := pd.db.Get([]byte(projectID))
	if err != nil {
		return nil, err
	}
	// Deserialize the entry
	var project config.Project
	if err := json.Unmarshal(data, &project); err != nil {
		return nil, err
	}
	return &project, nil
}

// GetAll retrieves all project entries from the project database
func (pd *ProjectDatabase) GetAll() ([]*config.Project, error) {
	// Initialize a slice to store all entries
	var projects []*config.Project

	// Iterate over all keys in the database
	err := pd.db.Iterate(func(key, value []byte) bool {
		// Deserialize the value into a Project
		var project config.Project
		if err := json.Unmarshal(value, &project); err != nil {
			// Return true to stop iteration if unmarshaling fails
			return true
		}
		// Append the project to the slice
		projects = append(projects, &project)
		// Continue iteration
		return false
	})

	if err != nil {
		return projects, err
	}

	return projects, nil
}

// Update updates an existing project entry in the project database
func (pd *ProjectDatabase) Update(project *config.Project) error {
	// Serialize the project
	serializedProject, err := json.Marshal(project)
	if err != nil {
		return err
	}
	// Update in the database
	return pd.db.Set([]byte(project.ID), serializedProject)
}

// Delete deletes the project entry for the given ID from the project database
func (pd *ProjectDatabase) Delete(projectID string) error {
	// Delete from the database
	return pd.db.Delete([]byte(projectID))
}

// SaveVersion saves a new tree version to disk.
func (pd *ProjectDatabase) SaveVersion() ([]byte, int64, error) {
	return pd.db.SaveVersion()
}

// Load loads the latest versioned tree from disk.
func (pd *ProjectDatabase) Load() (int64, error) {
	return pd.db.Load()
}

// LoadVersion loads a specific version of the tree from disk.
func (pd *ProjectDatabase) LoadVersion(targetVersion int64) (int64, error) {
	return pd.db.LoadVersion(targetVersion)
}

// String returns a string representation of the tree.
func (pd *ProjectDatabase) String() (string, error) {
	return pd.db.String()
}

// WorkingVersion returns the current working version of the tree.
func (pd *ProjectDatabase) WorkingVersion() int64 {
	return pd.db.WorkingVersion()
}

// WorkingHash returns the root hash of the current working tree.
func (pd *ProjectDatabase) WorkingHash() []byte {
	return pd.db.WorkingHash()
}

// AvailableVersions returns a list of available versions.
func (pd *ProjectDatabase) AvailableVersions() []int {
	return pd.db.AvailableVersions()
}

// IsEmpty checks if the database is empty.
func (pd *ProjectDatabase) IsEmpty() bool {
	return pd.db.IsEmpty()
}
