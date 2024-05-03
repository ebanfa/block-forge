package database

import (
	"errors"
	"os"
	"path/filepath"
	"sync"

	cosLogApi "cosmossdk.io/log"
	"github.com/cosmos/iavl"
	dbm "github.com/cosmos/iavl/db"
	"github.com/edward1christian/block-forge/pkg/application/db"
)

var (
	metaDB     *MetadataDatabase
	metaDBOnce sync.Once
)

var NovaHomeDirNm = ".nova"
var MetadataDatabaseID = "MetadataDatabase"
var BackendTypeGoLevelDB = "goleveldb"

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
	dbPath := filepath.Join(homeDir, NovaHomeDirNm, "databases", artifactID+".db")

	return dbPath, nil
}

// InitializeLevelDB initializes and returns a LevelDB instance
func CreateBackendLevelDB(name, path string) (dbm.DB, error) {
	db, err := dbm.NewDB(name, BackendTypeGoLevelDB, path)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// CreateIAVLDatabase initializes the IAVLDB instance and returns it
func CreateIAVLDatabase(name, path string) (*db.IAVLDatabase, error) {
	// Initialize the LevelDB instance
	ldb, err := CreateBackendLevelDB(name, path)
	if err != nil {
		return nil, err
	}

	// Initialize the IAVLDB instance
	iavlTree := iavl.NewMutableTree(ldb, 100, false, cosLogApi.NewNopLogger())
	iavlDB := db.NewIAVLDatabase(iavlTree)

	return iavlDB, nil
}

// GetMetadataDBInstance returns the singleton instance of MetadataDatabase
func GetMetadataDBInstance(name, path string) (*MetadataDatabase, error) {
	var err error
	metaDBOnce.Do(func() {
		// Initialize the LevelDB and IAVLDB
		iavlDB, dbErr := CreateIAVLDatabase(name, path)
		if dbErr != nil {
			// Assign the error
			err = dbErr
			return
		}

		// Create the metadata database instance
		metaDB = NewMetadataDatabase(iavlDB)
	})

	// Check if there was an error during initialization
	if err != nil {
		return nil, err
	}

	return metaDB, nil
}

// GetProjectDatabaseInstance creates an new instance of ProjectDatabase.
// The database will only be created if it does not already exist
func GetProjectDatabaseInstance(name, path string) (*ProjectDatabase, error) {
	// Initialize the database
	db, err := CreateIAVLDatabase(name, path)
	if err != nil {
		return nil, err
	}

	// Create the project database instance
	return NewProjectDatabase(db), nil
}
