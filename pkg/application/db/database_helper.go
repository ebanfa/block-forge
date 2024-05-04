package db

import (
	cosLogApi "cosmossdk.io/log"
	"github.com/cosmos/iavl"
	dbm "github.com/cosmos/iavl/db"
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
