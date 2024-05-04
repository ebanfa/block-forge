package store_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/edward1christian/block-forge/nova/pkg/store"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestMetadataStoreImpl_InsertMetadata_Success tests the successful insertion of metadata into the store.
func TestMetadataStoreImpl_InsertMetadata_Success(t *testing.T) {
	// Arrange
	mockDB := &mocks.MockDatabase{}
	mockStore := store.NewMetadataStore(mockDB)
	entry := &store.MetadataEntry{
		ProjectID:    "testProject",
		DatabaseName: "testDB",
		DatabasePath: "/path/to/db",
	}
	serializedEntry, _ := json.Marshal(entry)
	mockDB.On("Set", []byte(entry.ProjectID), serializedEntry).Return(nil)

	// Act
	err := mockStore.InsertMetadata(entry)

	// Assert
	assert.NoError(t, err, "InsertMetadata should not return an error")
	mockDB.AssertExpectations(t)
}

// TestMetadataStoreImpl_InsertMetadata_Error tests the error scenario of inserting metadata into the store.
func TestMetadataStoreImpl_InsertMetadata_Error(t *testing.T) {
	// Arrange
	mockDB := &mocks.MockDatabase{}
	mockStore := store.NewMetadataStore(mockDB)
	entry := &store.MetadataEntry{
		ProjectID:    "testProject",
		DatabaseName: "testDB",
		DatabasePath: "/path/to/db",
	}
	expectedErr := errors.New("database error")
	mockDB.On("Set", []byte(entry.ProjectID), mock.Anything).Return(expectedErr)

	// Act
	err := mockStore.InsertMetadata(entry)

	// Assert
	assert.EqualError(t, err, expectedErr.Error(), "InsertMetadata should return expected error")
	mockDB.AssertExpectations(t)
}

// TestMetadataStoreImpl_GetMetadata_Success tests the successful retrieval of metadata from the store.
func TestMetadataStoreImpl_GetMetadata_Success(t *testing.T) {
	// Arrange
	mockDB := &mocks.MockDatabase{}
	mockStore := store.NewMetadataStore(mockDB)
	projectID := "testProject"
	entry := &store.MetadataEntry{
		ProjectID:    projectID,
		DatabaseName: "testDB",
		DatabasePath: "/path/to/db",
	}
	serializedEntry, _ := json.Marshal(entry)
	mockDB.On("Get", []byte(projectID)).Return(serializedEntry, nil)

	// Act
	result, err := mockStore.GetMetadata(projectID)

	// Assert
	assert.NoError(t, err, "GetMetadata should not return an error")
	assert.Equal(t, entry, result, "Retrieved entry should match expected entry")
	mockDB.AssertExpectations(t)
}

// TestMetadataStoreImpl_GetMetadata_Error tests the error scenario of retrieving metadata from the store.
func TestMetadataStoreImpl_GetMetadata_Error(t *testing.T) {
	// Arrange
	mockDB := &mocks.MockDatabase{}
	mockStore := store.NewMetadataStore(mockDB)
	projectID := "testProject"
	expectedErr := errors.New("database error")
	mockDB.On("Get", []byte(projectID)).Return([]byte{}, expectedErr)

	// Act
	_, err := mockStore.GetMetadata(projectID)

	// Assert
	assert.EqualError(t, err, expectedErr.Error(), "GetMetadata should return expected error")
	mockDB.AssertExpectations(t)
}

// TestMetadataStoreImpl_GetAllMetadata_Success tests the successful retrieval of all metadata entries from the store.
func TestMetadataStoreImpl_GetAllMetadata_Success(t *testing.T) {
	// Arrange
	mockDB := &mocks.MockDatabase{}
	mockStore := store.NewMetadataStore(mockDB)
	entries := []*store.MetadataEntry{
		{
			ProjectID:    "project1",
			DatabaseName: "db1",
			DatabasePath: "/path/to/db1",
		},
		{
			ProjectID:    "project2",
			DatabaseName: "db2",
			DatabasePath: "/path/to/db2",
		},
	}
	serializedEntries := make([][]byte, len(entries))
	for i, entry := range entries {
		serializedEntry, _ := json.Marshal(entry)
		serializedEntries[i] = serializedEntry
	}
	mockDB.On("Iterate", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		fn := args.Get(0).(func(key, value []byte) bool)
		for _, serializedEntry := range serializedEntries {
			fn([]byte("some_key"), serializedEntry)
		}
	})

	// Act
	result, err := mockStore.GetAllMetadata()

	// Assert
	assert.NoError(t, err, "GetAllMetadata should not return an error")
	assert.ElementsMatch(t, entries, result, "Retrieved entries should match expected entries")
	mockDB.AssertExpectations(t)
}

// TestMetadataStoreImpl_GetAllMetadata_Error tests the error scenario of retrieving all metadata entries from the store.
func TestMetadataStoreImpl_GetAllMetadata_Error(t *testing.T) {
	// Arrange
	mockDB := &mocks.MockDatabase{}
	mockStore := store.NewMetadataStore(mockDB)
	expectedErr := errors.New("database error")
	mockDB.On("Iterate", mock.Anything).Return(expectedErr)

	// Act
	_, err := mockStore.GetAllMetadata()

	// Assert
	assert.EqualError(t, err, expectedErr.Error(), "GetAllMetadata should return expected error")
	mockDB.AssertExpectations(t)
}

// TestMetadataStoreImpl_UpdateMetadata_Success tests the successful update of metadata in the store.
func TestMetadataStoreImpl_UpdateMetadata_Success(t *testing.T) {
	// Arrange
	mockDB := &mocks.MockDatabase{}
	mockStore := store.NewMetadataStore(mockDB)
	entry := &store.MetadataEntry{
		ProjectID:    "testProject",
		DatabaseName: "testDB",
		DatabasePath: "/path/to/db",
	}
	serializedEntry, _ := json.Marshal(entry)
	mockDB.On("Set", []byte(entry.ProjectID), serializedEntry).Return(nil)

	// Act
	err := mockStore.UpdateMetadata(entry)

	// Assert
	assert.NoError(t, err, "UpdateMetadata should not return an error")
	mockDB.AssertExpectations(t)
}

// TestMetadataStoreImpl_UpdateMetadata_Error tests the error scenario of updating metadata in the store.
func TestMetadataStoreImpl_UpdateMetadata_Error(t *testing.T) {
	// Arrange
	mockDB := &mocks.MockDatabase{}
	mockStore := store.NewMetadataStore(mockDB)
	entry := &store.MetadataEntry{
		ProjectID:    "testProject",
		DatabaseName: "testDB",
		DatabasePath: "/path/to/db",
	}
	expectedErr := errors.New("database error")
	mockDB.On("Set", []byte(entry.ProjectID), mock.Anything).Return(expectedErr)

	// Act
	err := mockStore.UpdateMetadata(entry)

	// Assert
	assert.EqualError(t, err, expectedErr.Error(), "UpdateMetadata should return expected error")
	mockDB.AssertExpectations(t)
}

// TestMetadataStoreImpl_DeleteMetadata_Success tests the successful deletion of metadata from the store.
func TestMetadataStoreImpl_DeleteMetadata_Success(t *testing.T) {
	// Arrange
	mockDB := &mocks.MockDatabase{}
	mockStore := store.NewMetadataStore(mockDB)
	projectID := "testProject"
	mockDB.On("Delete", []byte(projectID)).Return(nil)

	// Act
	err := mockStore.DeleteMetadata(projectID)

	// Assert
	assert.NoError(t, err, "DeleteMetadata should not return an error")
	mockDB.AssertExpectations(t)
}

// TestMetadataStoreImpl_DeleteMetadata_Error tests the error scenario of deleting metadata from the store.
func TestMetadataStoreImpl_DeleteMetadata_Error(t *testing.T) {
	// Arrange
	mockDB := &mocks.MockDatabase{}
	mockStore := store.NewMetadataStore(mockDB)
	projectID := "testProject"
	expectedErr := errors.New("database error")
	mockDB.On("Delete", []byte(projectID)).Return(expectedErr)

	// Act
	err := mockStore.DeleteMetadata(projectID)

	// Assert
	assert.EqualError(t, err, expectedErr.Error(), "DeleteMetadata should return expected error")
	mockDB.AssertExpectations(t)
}

// TestMetadataStoreImpl_SaveVersion_Success tests the successful saving of a version in the store.
func TestMetadataStoreImpl_SaveVersion_Success(t *testing.T) {
	// Arrange
	mockDB := &mocks.MockDatabase{}
	mockStore := store.NewMetadataStore(mockDB)
	expectedHash := []byte("hash")
	expectedVersion := int64(123)
	mockDB.On("SaveVersion").Return(expectedHash, expectedVersion, nil)

	// Act
	hash, version, err := mockStore.SaveVersion()

	// Assert
	assert.NoError(t, err, "SaveVersion should not return an error")
	assert.Equal(t, expectedHash, hash, "Returned hash should match expected hash")
	assert.Equal(t, expectedVersion, version, "Returned version should match expected version")
	mockDB.AssertExpectations(t)
}

// TestMetadataStoreImpl_SaveVersion_Error tests the error scenario of saving a version in the store.
func TestMetadataStoreImpl_SaveVersion_Error(t *testing.T) {
	// Arrange
	mockDB := &mocks.MockDatabase{}
	mockStore := store.NewMetadataStore(mockDB)
	expectedErr := errors.New("database error")
	mockDB.On("SaveVersion").Return([]byte{}, int64(0), expectedErr)

	// Act
	_, _, err := mockStore.SaveVersion()

	// Assert
	assert.EqualError(t, err, expectedErr.Error(), "SaveVersion should return expected error")
	mockDB.AssertExpectations(t)
}

// TestMetadataStoreImpl_Load_Success tests the successful loading of metadata from the store.
func TestMetadataStoreImpl_Load_Success(t *testing.T) {
	// Arrange
	mockDB := &mocks.MockDatabase{}
	mockStore := store.NewMetadataStore(mockDB)
	expectedVersion := int64(123)
	mockDB.On("Load").Return(expectedVersion, nil)

	// Act
	version, err := mockStore.Load()

	// Assert
	assert.NoError(t, err, "Load should not return an error")
	assert.Equal(t, expectedVersion, version, "Returned version should match expected version")
	mockDB.AssertExpectations(t)
}

// TestMetadataStoreImpl_Load_Error tests the error scenario of loading metadata from the store.
func TestMetadataStoreImpl_Load_Error(t *testing.T) {
	// Arrange
	mockDB := &mocks.MockDatabase{}
	mockStore := store.NewMetadataStore(mockDB)
	expectedErr := errors.New("database error")
	mockDB.On("Load").Return(int64(0), expectedErr)

	// Act
	_, err := mockStore.Load()

	// Assert
	assert.EqualError(t, err, expectedErr.Error(), "Load should return expected error")
	mockDB.AssertExpectations(t)
}

// TestMetadataStoreImpl_LoadVersion_Success tests the successful loading of a specific version from the store.
func TestMetadataStoreImpl_LoadVersion_Success(t *testing.T) {
	// Arrange
	mockDB := &mocks.MockDatabase{}
	mockStore := store.NewMetadataStore(mockDB)
	targetVersion := int64(123)
	expectedVersion := int64(456)
	mockDB.On("LoadVersion", targetVersion).Return(expectedVersion, nil)

	// Act
	version, err := mockStore.LoadVersion(targetVersion)

	// Assert
	assert.NoError(t, err, "LoadVersion should not return an error")
	assert.Equal(t, expectedVersion, version, "Returned version should match expected version")
	mockDB.AssertExpectations(t)
}

// TestMetadataStoreImpl_LoadVersion_Error tests the error scenario of loading a specific version from the store.
func TestMetadataStoreImpl_LoadVersion_Error(t *testing.T) {
	// Arrange
	mockDB := &mocks.MockDatabase{}
	mockStore := store.NewMetadataStore(mockDB)
	targetVersion := int64(123)
	expectedErr := errors.New("database error")
	mockDB.On("LoadVersion", targetVersion).Return(int64(0), expectedErr)

	// Act
	_, err := mockStore.LoadVersion(targetVersion)

	// Assert
	assert.EqualError(t, err, expectedErr.Error(), "LoadVersion should return expected error")
	mockDB.AssertExpectations(t)
}

// TestMetadataStoreImpl_WorkingVersion_Success tests the successful retrieval of the working version from the store.
func TestMetadataStoreImpl_WorkingVersion_Success(t *testing.T) {
	// Arrange
	mockDB := &mocks.MockDatabase{}
	mockStore := store.NewMetadataStore(mockDB)
	expectedVersion := int64(123)
	mockDB.On("WorkingVersion").Return(expectedVersion)

	// Act
	version := mockStore.WorkingVersion()

	// Assert
	assert.Equal(t, expectedVersion, version, "WorkingVersion should return expected version")
	mockDB.AssertExpectations(t)
}

// TestMetadataStoreImpl_WorkingHash_Success tests the successful retrieval of the working hash from the store.
func TestMetadataStoreImpl_WorkingHash_Success(t *testing.T) {
	// Arrange
	mockDB := &mocks.MockDatabase{}
	mockStore := store.NewMetadataStore(mockDB)
	expectedHash := []byte("hash")
	mockDB.On("WorkingHash").Return(expectedHash)

	// Act
	hash := mockStore.WorkingHash()

	// Assert
	assert.Equal(t, expectedHash, hash, "WorkingHash should return expected hash")
	mockDB.AssertExpectations(t)
}

// TestMetadataStoreImpl_AvailableVersions_Success tests the successful retrieval of available versions from the store.
func TestMetadataStoreImpl_AvailableVersions_Success(t *testing.T) {
	// Arrange
	mockDB := &mocks.MockDatabase{}
	mockStore := store.NewMetadataStore(mockDB)
	expectedVersions := []int{1, 2, 3}
	mockDB.On("AvailableVersions").Return(expectedVersions)

	// Act
	versions := mockStore.AvailableVersions()

	// Assert
	assert.Equal(t, expectedVersions, versions, "AvailableVersions should return expected versions")
	mockDB.AssertExpectations(t)
}

// TestMetadataStoreImpl_IsEmpty_Success tests the successful check if the store is empty.
func TestMetadataStoreImpl_IsEmpty_Success(t *testing.T) {
	// Arrange
	mockDB := &mocks.MockDatabase{}
	mockStore := store.NewMetadataStore(mockDB)
	mockDB.On("IsEmpty").Return(true)

	// Act
	isEmpty := mockStore.IsEmpty()

	// Assert
	assert.True(t, isEmpty, "IsEmpty should return true")
	mockDB.AssertExpectations(t)
}

// TestMetadataStoreImpl_Close_Success tests the successful closing of the store.
func TestMetadataStoreImpl_Close_Success(t *testing.T) {
	// Arrange
	mockDB := &mocks.MockDatabase{}
	mockStore := store.NewMetadataStore(mockDB)
	mockDB.On("Close").Return(nil)

	// Act
	err := mockStore.Close()

	// Assert
	assert.NoError(t, err, "Close should not return an error")
	mockDB.AssertExpectations(t)
}
