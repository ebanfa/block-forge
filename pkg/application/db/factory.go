package db

import (
	cosLogApi "cosmossdk.io/log"
	"github.com/cosmos/iavl"
	dbm "github.com/cosmos/iavl/db"
)

// DatabaseFactory is an interface for creating databases.
type DatabaseFactory interface {
	// CreateDatabase creates and initializes a database instance with the given name and path.
	CreateDatabase(name, path string) (Database, error)
}

// IAVLDatabaseFactory is a concrete implementation of the DatabaseFactory interface
// that creates IAVL database instances.
type IAVLDatabaseFactory struct {
	dbCreator func(name, backendType, path string) (dbm.DB, error)
}

// NewIAVLDatabaseFactory creates a new instance of IAVLDatabaseFactory with the given DB creator function.
func NewIAVLDatabaseFactory(dbCreator func(name, backendType, path string) (dbm.DB, error)) *IAVLDatabaseFactory {
	return &IAVLDatabaseFactory{
		dbCreator: dbCreator,
	}
}

// CreateDatabase creates and initializes an IAVL database instance with the given name and path.
func (f *IAVLDatabaseFactory) CreateDatabase(name, path string) (Database, error) {
	// Initialize the LevelDB instance using the injected creator function
	ldb, err := f.dbCreator(name, BackendTypeGoLevelDB, path)
	if err != nil {
		return nil, err
	}

	// Initialize the IAVLDB instance
	iavlTree := iavl.NewMutableTree(ldb, 100, false, cosLogApi.NewNopLogger())
	iavlDB := NewIAVLDatabase(iavlTree)

	return iavlDB, nil
}
