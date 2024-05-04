package db

import (
	"errors"
	"os"
	"path/filepath"

	cosLogApi "cosmossdk.io/log"
	"github.com/cosmos/iavl"
	dbm "github.com/cosmos/iavl/db"
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
func CreateIAVLDatabase(name, path string) (*IAVLDatabase, error) {
	// Initialize the LevelDB instance
	ldb, err := CreateBackendLevelDB(name, path)
	if err != nil {
		return nil, err
	}

	// Initialize the IAVLDB instance
	iavlTree := iavl.NewMutableTree(ldb, 100, false, cosLogApi.NewNopLogger())
	iavlDB := NewIAVLDatabase(iavlTree)

	return iavlDB, nil
}
