package database

import (
	"testing"

	"github.com/edward1christian/block-forge/nova/pkg/database"
	"github.com/edward1christian/block-forge/nova/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInitializeLevelDB_Success(t *testing.T) {
	// Create a temporary directory for the database
	tmpDir := t.TempDir()
	// Arrange
	name := "testDB"

	// Act
	db, err := database.InitializeLevelDB(name, tmpDir)

	// Assert
	assert.NoError(t, err, "No error should occur during initialization")
	assert.NotNil(t, db, "Database instance should not be nil")
}

func TestInitializeLevelDB_Error(t *testing.T) {
	// Arrange
	name := "testDB"
	// Providing an invalid path to simulate an error
	path := "/invalid/path"

	// Act
	_, err := database.InitializeLevelDB(name, path)

	// Assert
	assert.Error(t, err, "Error should occur during initialization")
}

func TestGetIAVLDatabase(t *testing.T) {
	// Arrange
	// Create a temporary directory for the database
	//tmpDir := t.TempDir()

	// Mock a database instance for testing
	mockDB := &mocks.MockDB{}
	mockDB.On("Close").Return(nil)
	mockDB.On("NewBatch").Return(&mocks.MockBatch{})
	mockDB.On("NewBatchWithSize", mock.Anything).Return(&mocks.MockBatch{})
	mockDB.On("Get", mock.Anything).Return([]byte{}, nil)

	// Act
	iavlDB := database.GetIAVLDatabase(mockDB)

	// Assert
	assert.NotNil(t, iavlDB, "IAVL database instance should not be nil")
}
