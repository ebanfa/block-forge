package store

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/edward1christian/block-forge/nova/pkg/store"
	"github.com/edward1christian/block-forge/pkg/application/db"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockOS is a mock implementation of the OS interface
type MockOS struct {
	mock.Mock
}

// UserHomeDir is a mocked method for getting the user's home directory
func (m *MockOS) UserHomeDir() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

// TestGetDefaultDatabasePath_Success tests the successful construction of the default database path.
func TestGetDefaultDatabasePath_Success(t *testing.T) {
	// Create a temporary directory for testing
	homeDir, err := os.UserHomeDir()
	assert.NoError(t, err)

	// Expected database path
	expectedPath := filepath.Join(homeDir,
		store.NovaHomeDirNm, "databases", "testArtifact.db")

	// Call the function
	path, err := store.GetDefaultDatabasePath("testArtifact")

	// Assertions
	assert.NoError(t, err, "No error should occur during path construction")
	assert.Equal(t, expectedPath, path, "Constructed path should match expected path")
}

// TestCreateBackendLevelDB_Success tests the successful initialization of a LevelDB instance.
func TestCreateBackendLevelDB_Success(t *testing.T) {
	// Create a temporary directory for the database
	tmpDir := t.TempDir()
	// Arrange
	name := "testDB"

	// Act
	db, err := db.CreateBackendLevelDB(name, tmpDir)

	// Assert
	assert.NoError(t, err, "No error should occur during initialization")
	assert.NotNil(t, db, "Database instance should not be nil")
}

// TestCreateBackendLevelDB_Error tests the error handling during LevelDB initialization.
func TestCreateBackendLevelDB_Error(t *testing.T) {
	// Arrange
	name := "testDB"
	// Providing an invalid path to simulate an error
	path := "/invalid/path"

	// Act
	_, err := db.CreateBackendLevelDB(name, path)

	// Assert
	assert.Error(t, err, "Error should occur during initialization")
}

// TestCreateIAVLDatabase tests the initialization of an IAVL store.
func TestCreateIAVLDatabase(t *testing.T) {
	// Mock a database instance for testing
	mockDB := &mocks.MockDatabase{}
	mockDB.On("Close").Return(nil)
	mockDB.On("Get", mock.Anything).Return([]byte{}, nil)

	// Act
	path := filepath.Join(t.TempDir(), "subdirectory")
	iavlDB, err := db.CreateIAVLDatabase("testDB", path)

	// Assert
	assert.NoError(t, err, "No error should occur during initialization")
	assert.NotNil(t, iavlDB, "IAVL database instance should not be nil")
}

// TestGetMetadataDBInstance_Success tests the successful initialization of the MetadataDatabase instance.
func TestGetMetadataDBInstance_Success(t *testing.T) {
	// Arrange
	name := "testDB"
	tmpDir := t.TempDir()

	// Act
	metaDB, err := store.GetMetadataStoreInstance(name, tmpDir)

	// Assert
	assert.NoError(t, err, "No error should occur during initialization")
	assert.NotNil(t, metaDB, "MetadataDatabase instance should not be nil")
}
