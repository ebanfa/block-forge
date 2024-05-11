package store

import (
	"testing"

	"github.com/edward1christian/block-forge/pkg/application/mocks"
	"github.com/edward1christian/block-forge/pkg/application/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	MultiDbName = "MockDb"
	MultiDbPath = "MockDbPath"
)

// TestNewMultiStore tests the creation of a new MultiStore.
func TestNewMultiStore(t *testing.T) {
	// Arrange
	mockStore := &mocks.MockStore{}
	mockDbFactory := &mocks.MockDatabaseFactory{}
	mockStoreFactory := &mocks.MockStoreFactory{}
	mockDbFactory.On("CreateDatabase", mock.Anything, mock.Anything).Return(
		&mocks.MockDatabase{}, nil,
	)

	// Act
	ms, err := store.NewMultiStore(mockStore, mockStoreFactory)
	assert.NoError(t, err)

	// Assert
	assert.NotNil(t, ms)
	assert.Equal(t, 0, ms.GetStoreCount())
}

// TestCreateStore_Success tests the successful creation of a store.
func TestCreateStore_Success(t *testing.T) {
	// Arrange
	mockStore := &mocks.MockStore{}
	mockStoreFactory := &mocks.MockStoreFactory{}

	mockStoreFactory.On("CreateStore", mock.Anything).Return(mockStore, nil)

	ms, err := store.NewMultiStore(mockStore, mockStoreFactory)
	assert.NoError(t, err)

	// Act
	createdStore, created, err := ms.CreateStore("mockNamespace")

	// Assert
	assert.NoError(t, err, "Creating store should not return an error")
	assert.True(t, created, "Creating a new store should return true")
	assert.Equal(t, mockStore, createdStore, "Created store should match the mock store")

	// Verify mock behavior
	mockStoreFactory.AssertCalled(t, "CreateStore", mock.Anything)
}

// TestCreateStore_Error_InvalidNamespace tests error when creating a store with an invalid namespace.
func TestCreateStore_Error_InvalidNamespace(t *testing.T) {
	// Arrange
	mockStore := &mocks.MockStore{}
	mockStoreFactory := &mocks.MockStoreFactory{}

	mockStoreFactory.On("CreateStore", mock.Anything).Return(mockStore, nil)

	ms, err := store.NewMultiStore(mockStore, mockStoreFactory)
	assert.NoError(t, err)

	// Act
	createdStore, created, err := ms.CreateStore("")

	// Assert
	assert.Error(t, err, "Creating store with empty namespace should return an error")
	assert.False(t, created, "Creating a store with empty namespace should return false")
	assert.Nil(t, createdStore, "Created store should be nil")

	// Verify mock behavior (no call to StoreFactory)
	mockStoreFactory.AssertNotCalled(t, "CreateStore", mock.Anything)
}

// TestCreateStore_NewStore tests creating a new store in MultiStore.
func TestCreateStore_NewStore(t *testing.T) {
	// Arrange
	mockStore := &mocks.MockStore{}
	mockStoreFactory := &mocks.MockStoreFactory{}

	mockStoreFactory.On("CreateStore", mock.Anything).Return(mockStore, nil)

	ms, err := store.NewMultiStore(mockStore, mockStoreFactory)
	assert.NoError(t, err)

	// Act
	store, _, err := ms.CreateStore("test")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, store)
	assert.Equal(t, 1, ms.GetStoreCount())
}

// TestCreateStore_ExistingStore tests attempting to create an existing store in MultiStore.
func TestCreateStore_ExistingStore(t *testing.T) {
	// Arrange
	mockStore := &mocks.MockStore{}
	mockStoreFactory := &mocks.MockStoreFactory{}

	mockStoreFactory.On("CreateStore", mock.Anything).Return(mockStore, nil)

	ms, err := store.NewMultiStore(mockStore, mockStoreFactory)
	assert.NoError(t, err)

	store, _, err := ms.CreateStore("test")
	assert.NoError(t, err)

	// Act
	_, created, err := ms.CreateStore("test")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, store)
	assert.False(t, created)
	assert.Equal(t, 1, ms.GetStoreCount())
}

// TestGetStore_ExistingStore tests retrieving an existing store from MultiStore.
func TestGetStore_ExistingStore(t *testing.T) {
	// Arrange
	mockStore := &mocks.MockStore{}
	mockStoreFactory := &mocks.MockStoreFactory{}

	mockStoreFactory.On("CreateStore", mock.Anything).Return(mockStore, nil)

	// Act
	ms, err := store.NewMultiStore(mockStore, mockStoreFactory)
	assert.NoError(t, err)

	namespace := []byte("test")

	_, _, err = ms.CreateStore(string(namespace))
	assert.NoError(t, err)

	// Act
	testStore := ms.GetStore([]byte(store.GenerateStoreId("test")))

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, testStore)
}

// TestGetStore_NonExistingStore tests attempting to retrieve a non-existing store from MultiStore.
func TestGetStore_NonExistingStore(t *testing.T) {
	// Arrange
	mockStore := &mocks.MockStore{}
	mockStoreFactory := &mocks.MockStoreFactory{}

	mockStoreFactory.On("CreateStore", mock.Anything).Return(mockStore, nil)

	// Act
	ms, err := store.NewMultiStore(mockStore, mockStoreFactory)
	assert.NoError(t, err)

	// Act
	store := ms.GetStore([]byte(store.GenerateStoreId("test")))

	// Assert
	assert.NoError(t, err)
	assert.Nil(t, store)
}

// TestGetStoreCount tests getting the count of stores in MultiStore.
func TestGetStoreCount(t *testing.T) {
	// Arrange
	mockStore := &mocks.MockStore{}
	mockStoreFactory := &mocks.MockStoreFactory{}

	mockStoreFactory.On("CreateStore", mock.Anything).Return(mockStore, nil)

	ms, err := store.NewMultiStore(mockStore, mockStoreFactory)
	assert.NoError(t, err)

	_, _, err = ms.CreateStore("test1")
	assert.NoError(t, err)

	_, _, err = ms.CreateStore("test2")
	assert.NoError(t, err)

	// Act
	count := ms.GetStoreCount()

	// Assert
	assert.Equal(t, 2, count)
}

// TestLoad tests loading the latest versioned database from disk.
func TestLoad(t *testing.T) {
	// Arrange
	mockStore := &mocks.MockStore{}
	mockStoreFactory := &mocks.MockStoreFactory{}

	mockStore.On("Load").Return(int64(1), nil)
	mockStoreFactory.On("CreateStore", mock.Anything).Return(mockStore, nil)

	ms, err := store.NewMultiStore(mockStore, mockStoreFactory)
	assert.NoError(t, err)

	// Mock database iteration
	mockStore.On("Iterate", mock.Anything).Return(nil)

	// Act
	version, err := ms.Load()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, int64(1), version)
}

// TestSaveVersion tests saving a new version of the database to disk.
func TestSaveVersion(t *testing.T) {
	// Arrange
	mockDb := &mocks.MockDatabase{}
	mockStore := &mocks.MockStore{}
	mockDbFactory := &mocks.MockDatabaseFactory{}
	mockStoreFactory := &mocks.MockStoreFactory{}

	mockDb.On("Set", mock.Anything, mock.Anything).Return(nil)
	mockStore.On("SaveVersion").Return([]byte("data"), int64(1), nil)
	mockDbFactory.On("CreateDatabase", mock.Anything, mock.Anything).Return(mockDb, nil)

	// Create StoreOptions from StoreMetaData
	ms, err := store.NewMultiStore(mockStore, mockStoreFactory)
	assert.NoError(t, err)

	// Act
	data, version, err := ms.SaveVersion()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, []byte("data"), data)
	assert.Equal(t, int64(1), version)
}
