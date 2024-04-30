package database

import (
	"sync"

	cosLogApi "cosmossdk.io/log"
	"github.com/cosmos/iavl"
	dbm "github.com/cosmos/iavl/db"
)

var (
	metaDB     *MetadataDatabase
	metaDBOnce sync.Once
)

var BackendTypeGoLevelDB = "goleveldb"

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
