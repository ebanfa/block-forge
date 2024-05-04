package store

import (
	"testing"

	"github.com/edward1christian/block-forge/pkg/application/mocks"
	"github.com/edward1christian/block-forge/pkg/application/store"
	"github.com/stretchr/testify/assert"
)

// TestNewMultiStore tests the creation of a new MultiStore.
func TestNewMultiStore(t *testing.T) {
	// Arrange
	mockDb := &mocks.MockDatabase{}

	// Act
	ms := store.NewMultiStore(mockDb)

	// Assert
	assert.NotNil(t, ms)
	assert.Equal(t, 0, ms.GetStoreCount())
}

// TestGetStore_ExistingStore tests retrieving an existing store from MultiStore.
func TestGetStore_ExistingStore(t *testing.T) {
	// Arrange
	mockDb := &mocks.MockDatabase{}
	ms := store.NewMultiStore(mockDb)
	namespace := []byte("test")
	options := store.StoreOptions{}
	_, err := ms.CreateStore(namespace, options, mockDb)
	assert.NoError(t, err)

	// Act
	store, err := ms.GetStore(namespace, options)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, store)
}

// TestGetStore_NonExistingStore tests attempting to retrieve a non-existing store from MultiStore.
func TestGetStore_NonExistingStore(t *testing.T) {
	// Arrange
	mockDb := &mocks.MockDatabase{}
	ms := store.NewMultiStore(mockDb)
	namespace := []byte("test")
	options := store.StoreOptions{}

	// Act
	store, err := ms.GetStore(namespace, options)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, store)
}

// TestCreateStore_NewStore tests creating a new store in MultiStore.
func TestCreateStore_NewStore(t *testing.T) {
	// Arrange
	mockDb := &mocks.MockDatabase{}
	ms := store.NewMultiStore(mockDb)
	namespace := []byte("test")
	options := store.StoreOptions{}

	// Act
	store, err := ms.CreateStore(namespace, options, mockDb)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, store)
	assert.Equal(t, 1, ms.GetStoreCount())
}

// TestCreateStore_ExistingStore tests attempting to create an existing store in MultiStore.
func TestCreateStore_ExistingStore(t *testing.T) {
	// Arrange
	mockDb := &mocks.MockDatabase{}
	ms := store.NewMultiStore(mockDb)
	namespace := []byte("test")
	options := store.StoreOptions{}
	_, err := ms.CreateStore(namespace, options, mockDb)
	assert.NoError(t, err)

	// Act
	store, err := ms.CreateStore(namespace, options, mockDb)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, store)
	assert.Equal(t, 1, ms.GetStoreCount())
}

// TestGetStoreCount tests getting the count of stores in MultiStore.
func TestGetStoreCount(t *testing.T) {
	// Arrange
	mockDb := &mocks.MockDatabase{}
	ms := store.NewMultiStore(mockDb)
	_, err := ms.CreateStore([]byte("store1"), store.StoreOptions{}, mockDb)
	assert.NoError(t, err)
	_, err = ms.CreateStore([]byte("store2"), store.StoreOptions{}, mockDb)
	assert.NoError(t, err)

	// Act
	count := ms.GetStoreCount()

	// Assert
	assert.Equal(t, 2, count)
}
