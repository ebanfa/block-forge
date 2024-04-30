package database

import (
	"cosmossdk.io/log"
	"github.com/cosmos/iavl"
	dbm "github.com/cosmos/iavl/db"
)

var BackendTypeGoLevelDB = "goleveldb"

// InitializeLevelDB initializes and returns a LevelDB instance
func InitializeLevelDB(name, path string) (dbm.DB, error) {
	db, err := dbm.NewDB(name, BackendTypeGoLevelDB, path)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// GetIAVLDatabase initializes and returns an IAVL database instance
func GetIAVLDatabase(db dbm.DB) *IAVLDatabase {
	iavlTree := iavl.NewMutableTree(db, 100, false, log.NewNopLogger())
	return NewIAVLDatabase(iavlTree)
}
