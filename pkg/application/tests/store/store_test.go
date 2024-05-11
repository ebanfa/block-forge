package store

import (
	"testing"

	"github.com/edward1christian/block-forge/pkg/application/db"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
	"github.com/edward1christian/block-forge/pkg/application/store"
	"github.com/stretchr/testify/assert"
)

const (
	MockDbName = "MockDb"
	MockDbPath = "MockDbPath"
)

// TestNewStoreImpl_Success tests creating a new StoreImpl instance successfully.
func TestNewStoreImpl_Success(t *testing.T) {
	// Arrange
	mockDatabase := &mocks.MockDatabase{}

	expectedStore := &store.StoreImpl{
		Database: mockDatabase,
	}
	// Act
	store, err := store.NewStoreImpl(MockDbName, MockDbPath, mockDatabase)

	// Assert
	assert.NoError(t, err, "NewStoreImpl should not return an error")
	assert.NotNil(t, store, "NewStoreImpl should not return nil")
	assert.Equal(t, expectedStore.Database, store.Database, "NewStoreImpl should initialize the correct database")
}

// TestNewStoreImpl_Error tests creating a new StoreImpl instance with an error.
func TestNewStoreImpl_Error(t *testing.T) {
	// Arrange
	var nilDb db.Database

	// Act
	store, err := store.NewStoreImpl(MockDbName, MockDbPath, nilDb)

	// Assert
	assert.Error(t, err, "NewStoreImpl should return an error when database is nil")
	assert.Nil(t, store, "NewStoreImpl should return nil when database is nil")
}
