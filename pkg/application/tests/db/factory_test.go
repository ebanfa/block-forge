package db

import (
	"errors"
	"testing"

	dbm "github.com/cosmos/iavl/db"
	"github.com/edward1christian/block-forge/pkg/application/db"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock creator function for the database

// TestIAVLDatabaseFactory_CreateDatabase_Success tests the CreateDatabase method of IAVLDatabaseFactory for success.
func TestIAVLDatabaseFactory_CreateDatabase_Success(t *testing.T) {
	// Arrange
	mockDbm := &mocks.MockDBM{}
	mockDbm.On("Get", mock.Anything).Return([]byte{}, nil)
	mockDbm.On("NewBatchWithSize", mock.Anything).Return(&mocks.MockBatch{}, nil)

	factory := db.NewIAVLDatabaseFactory(func(name, backendType, path string) (dbm.DB, error) {
		return mockDbm, nil // Implement mock behavior here if necessary
	})

	// Act
	db, err := factory.CreateDatabase("test", "/path/to/db")

	// Assert
	assert.NoError(t, err, "Creating database should not return an error")
	assert.NotNil(t, db, "Database should not be nil")
	// Additional assertions specific to IAVL database can be added here if necessary
}

// TestIAVLDatabaseFactory_CreateDatabase_Error tests the CreateDatabase method of IAVLDatabaseFactory for error.
func TestIAVLDatabaseFactory_CreateDatabase_Error(t *testing.T) {
	// Arrange
	expectedErr := errors.New("error creating database")
	factory := db.NewIAVLDatabaseFactory(func(name, backendType, path string) (dbm.DB, error) {
		return nil, expectedErr
	})

	// Act
	db, err := factory.CreateDatabase("test", "/path/to/db")

	// Assert
	assert.Error(t, err, "Creating database should return an error")
	assert.Nil(t, db, "Database should be nil")
	assert.EqualError(t, err, expectedErr.Error(), "Error message should match")
}
