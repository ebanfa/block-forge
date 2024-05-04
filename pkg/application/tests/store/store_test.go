package store

import (
	"testing"

	"github.com/edward1christian/block-forge/pkg/application/db"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
	"github.com/edward1christian/block-forge/pkg/application/store"
	"github.com/stretchr/testify/assert"
)

// TestNewStoreImpl_Success tests creating a new StoreImpl instance successfully.
func TestNewStoreImpl_Success(t *testing.T) {
	// Arrange
	mockDatabase := &mocks.MockDatabase{} // Create a mock implementation of db.Database
	expectedStore := &store.StoreImpl{
		Database: mockDatabase,
	}

	// Act
	store, err := store.NewStoreImpl(mockDatabase)

	// Assert
	assert.NoError(t, err, "NewStoreImpl should not return an error")
	assert.NotNil(t, store, "NewStoreImpl should not return nil")
	assert.Equal(t, expectedStore.Database, store.Database, "NewStoreImpl should initialize the correct database")
}

// TestNewStoreImpl_Error tests creating a new StoreImpl instance with an error.
func TestNewStoreImpl_Error(t *testing.T) {
	// Arrange
	var nilDatabase db.Database

	// Act
	store, err := store.NewStoreImpl(nilDatabase)

	// Assert
	assert.Error(t, err, "NewStoreImpl should return an error when database is nil")
	assert.Nil(t, store, "NewStoreImpl should return nil when database is nil")
}
