package database_test

import (
	"path/filepath"
	"testing"

	"github.com/edward1christian/block-forge/nova/pkg/database"
	"github.com/edward1christian/block-forge/nova/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestCreateBackendLevelDB_Success tests the successful initialization of a LevelDB instance.
func TestCreateBackendLevelDB_Success(t *testing.T) {
	// Create a temporary directory for the database
	tmpDir := t.TempDir()
	// Arrange
	name := "testDB"

	// Act
	db, err := database.CreateBackendLevelDB(name, tmpDir)

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
	_, err := database.CreateBackendLevelDB(name, path)

	// Assert
	assert.Error(t, err, "Error should occur during initialization")
}

// TestCreateIAVLDatabase tests the initialization of an IAVL database.
func TestCreateIAVLDatabase(t *testing.T) {
	// Mock a database instance for testing
	mockDB := &mocks.MockDB{}
	mockDB.On("Close").Return(nil)
	mockDB.On("NewBatch").Return(&mocks.MockBatch{})
	mockDB.On("NewBatchWithSize", mock.Anything).Return(&mocks.MockBatch{})
	mockDB.On("Get", mock.Anything).Return([]byte{}, nil)

	// Act
	path := filepath.Join(t.TempDir(), "subdirectory")
	iavlDB, err := database.CreateIAVLDatabase("testDB", path)

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
	metaDB, err := database.GetMetadataDBInstance(name, tmpDir)

	// Assert
	assert.NoError(t, err, "No error should occur during initialization")
	assert.NotNil(t, metaDB, "MetadataDatabase instance should not be nil")
}
