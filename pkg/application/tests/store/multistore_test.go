package store

import (
	"testing"

	"github.com/edward1christian/block-forge/pkg/application/mocks"
	"github.com/edward1christian/block-forge/pkg/application/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestNewMultiStore tests the creation of a new MultiStore.
func TestNewMultiStore(t *testing.T) {
	// Arrange
	mockDbFactory := &mocks.MockDatabaseFactory{}
	mockDbFactory.On("CreateDatabase", mock.Anything, mock.Anything).Return(
		&mocks.MockDatabase{}, nil,
	)

	// Act
	ms, err := store.NewMultiStore(mockDbFactory)
	assert.NoError(t, err)

	// Assert
	assert.NotNil(t, ms)
	assert.Equal(t, 0, ms.GetStoreCount())
}

// TestGetStore_ExistingStore tests retrieving an existing store from MultiStore.
func TestGetStore_ExistingStore(t *testing.T) {
	// Arrange
	mockDbFactory := &mocks.MockDatabaseFactory{}
	mockDbFactory.On("CreateDatabase", mock.Anything, mock.Anything).Return(
		&mocks.MockDatabase{}, nil,
	)

	ms, err := store.NewMultiStore(mockDbFactory)
	assert.NoError(t, err)

	namespace := []byte("test")
	options := store.StoreOptions{}

	_, err = ms.CreateStore(namespace, options)
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
	mockDbFactory := &mocks.MockDatabaseFactory{}
	mockDbFactory.On("CreateDatabase", mock.Anything, mock.Anything).Return(
		&mocks.MockDatabase{}, nil,
	)
	ms, err := store.NewMultiStore(mockDbFactory)
	assert.NoError(t, err)

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
	mockDbFactory := &mocks.MockDatabaseFactory{}
	mockDbFactory.On("CreateDatabase", mock.Anything, mock.Anything).Return(
		&mocks.MockDatabase{}, nil,
	)
	ms, err := store.NewMultiStore(mockDbFactory)
	assert.NoError(t, err)

	namespace := []byte("test")
	options := store.StoreOptions{}

	// Act
	store, err := ms.CreateStore(namespace, options)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, store)
	assert.Equal(t, 1, ms.GetStoreCount())
}

// TestCreateStore_ExistingStore tests attempting to create an existing store in MultiStore.
func TestCreateStore_ExistingStore(t *testing.T) {
	// Arrange
	mockDbFactory := &mocks.MockDatabaseFactory{}
	mockDbFactory.On("CreateDatabase", mock.Anything, mock.Anything).Return(
		&mocks.MockDatabase{}, nil,
	)
	ms, err := store.NewMultiStore(mockDbFactory)
	assert.NoError(t, err)

	namespace := []byte("test")
	options := store.StoreOptions{}

	_, err = ms.CreateStore(namespace, options)
	assert.NoError(t, err)

	// Act
	store, err := ms.CreateStore(namespace, options)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, store)
	assert.Equal(t, 1, ms.GetStoreCount())
}

// TestGetStoreCount tests getting the count of stores in MultiStore.
func TestGetStoreCount(t *testing.T) {
	// Arrange
	mockDbFactory := &mocks.MockDatabaseFactory{}
	mockDbFactory.On("CreateDatabase", mock.Anything, mock.Anything).Return(
		&mocks.MockDatabase{}, nil,
	)
	ms, err := store.NewMultiStore(mockDbFactory)
	assert.NoError(t, err)

	_, err = ms.CreateStore([]byte("store1"), store.StoreOptions{})
	assert.NoError(t, err)

	_, err = ms.CreateStore([]byte("store2"), store.StoreOptions{})
	assert.NoError(t, err)

	// Act
	count := ms.GetStoreCount()

	// Assert
	assert.Equal(t, 2, count)
}
