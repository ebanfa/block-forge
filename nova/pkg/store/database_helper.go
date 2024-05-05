package store

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/edward1christian/block-forge/nova/pkg/utils"
	"github.com/edward1christian/block-forge/pkg/application/db"
)

const (
	NovaHomeDirName = ".nova"
	MetadataDbName  = "MetadataStore"
)

// GetDatabasePath returns the database path based on the user's home directory and project ID.
func GetDefaultDatabasePath(artifactID string) (string, error) {
	if artifactID == "" {
		return "", errors.New("artifactID cannot be empty")
	}
	// Get the home directory of the current user
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// Define the database path within the .nova directory
	dbPath := filepath.Join(homeDir, NovaHomeDirName, "databases", artifactID+".db")

	return dbPath, nil
}

// GetMetadataStoreInstance creates a new instance of MetadataStore.
// The database will only be created if it does not already exist.
func GetMetadataStoreInstance(name, path string) (MetadataStore, error) {
	// Initialize the database
	db, err := db.CreateIAVLDatabase(name, path)
	if err != nil {
		return nil, err
	}

	// Create the metadata store instance
	return NewMetadataStore(db), nil
}

func GetDefaultMetadataDB(name string) (MetadataStore, error) {
	databaseId := utils.HashSHA256(name)
	// Get the default database path for metadata
	metaDBPath, err := GetDefaultDatabasePath(databaseId)
	if err != nil {
		return nil, err
	}

	// Get the MetadataDatabase instance
	metaDB, err := GetMetadataStoreInstance(databaseId, metaDBPath)
	if err != nil {
		return nil, err
	}

	return metaDB, nil
}
