package database

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/edward1christian/block-forge/nova/pkg/database"
	"github.com/edward1christian/block-forge/nova/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestMetadataDatabase_Insert_Success tests the Insert method of MetadataDatabase for success.
func TestMetadataDatabase_Insert_Success(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	metaDB := database.NewMetadataDatabase(db)
	entry := &database.MetadataEntry{
		ProjectID:    "project1",
		DatabaseName: "db1",
		DatabasePath: "/path/to/db1",
		CreationDate: time.Now(),
		LastUpdated:  time.Now(),
	}

	// Mock behavior
	db.On("Set", []byte(entry.ProjectID), mockJSON(entry)).Return(nil)

	// Act
	err := metaDB.Insert(entry)

	// Assert
	assert.NoError(t, err, "Inserting metadata entry should not return an error")
	db.AssertExpectations(t)
}

// TestMetadataDatabase_Insert_Error tests the Insert method of MetadataDatabase for error.
func TestMetadataDatabase_Insert_Error(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	metaDB := database.NewMetadataDatabase(db)
	entry := &database.MetadataEntry{
		ProjectID:    "project1",
		DatabaseName: "db1",
		DatabasePath: "/path/to/db1",
		CreationDate: time.Now(),
		LastUpdated:  time.Now(),
	}
	expectedErr := errors.New("database error")

	// Mock behavior
	db.On("Set", []byte(entry.ProjectID), mockJSON(entry)).Return(expectedErr)

	// Act
	err := metaDB.Insert(entry)

	// Assert
	assert.Error(t, err, "Inserting metadata entry should return an error")
	assert.EqualError(t, err, expectedErr.Error(), "Error should match expected error")
	db.AssertExpectations(t)
}

// TestMetadataDatabase_Get_Success tests the Get method of MetadataDatabase for success.
func TestMetadataDatabase_Get_Success(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	metaDB := database.NewMetadataDatabase(db)
	entry := &database.MetadataEntry{
		ProjectID:    "project1",
		DatabaseName: "db1",
		DatabasePath: "/path/to/db1",
		CreationDate: time.Now(),
		LastUpdated:  time.Now(),
	}
	expectedData, _ := json.Marshal(entry)

	// Mock behavior
	db.On("Get", []byte(entry.ProjectID)).Return(expectedData, nil)

	// Act
	result, err := metaDB.Get(entry.ProjectID)

	// Assert
	assert.NoError(t, err, "Getting metadata entry should not return an error")
	// Assert
	assert.NoError(t, err, "Getting metadata entry should not return an error")
	assert.Equal(t, entry.ProjectID, result.ProjectID, "ProjectID should match")
	assert.Equal(t, entry.DatabaseName, result.DatabaseName, "DatabaseName should match")
	assert.Equal(t, entry.DatabasePath, result.DatabasePath, "DatabasePath should match")
	assert.Equal(t, entry.CreationDate.Unix(), result.CreationDate.Unix(), "CreationDate should match")
	assert.Equal(t, entry.LastUpdated.Unix(), result.LastUpdated.Unix(), "LastUpdated should match")

	db.AssertExpectations(t)
}

// TestMetadataDatabase_Get_Error tests the Get method of MetadataDatabase for error.
func TestMetadataDatabase_Get_Error(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	metaDB := database.NewMetadataDatabase(db)
	projectID := "project1"
	expectedErr := errors.New("database error")

	// Mock behavior
	db.On("Get", []byte(projectID)).Return([]byte{}, expectedErr)

	// Act
	result, err := metaDB.Get(projectID)

	// Assert
	assert.Error(t, err, "Getting metadata entry should return an error")
	assert.Nil(t, result, "Result should be nil")
	assert.EqualError(t, err, expectedErr.Error(), "Error should match expected error")
	db.AssertExpectations(t)
}

// TestMetadataDatabase_GetAll_Success tests the GetAll method of MetadataDatabase for success.
func TestMetadataDatabase_GetAll_Success(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	metaDB := database.NewMetadataDatabase(db)
	entry1 := &database.MetadataEntry{
		ProjectID:    "project1",
		DatabaseName: "db1",
		DatabasePath: "/path/to/db1",
		CreationDate: time.Now(),
		LastUpdated:  time.Now(),
	}
	entry2 := &database.MetadataEntry{
		ProjectID:    "project2",
		DatabaseName: "db2",
		DatabasePath: "/path/to/db2",
		CreationDate: time.Now(),
		LastUpdated:  time.Now(),
	}
	expectedData1, _ := json.Marshal(entry1)
	expectedData2, _ := json.Marshal(entry2)

	// Mock behavior
	db.On("Iterate", mock.AnythingOfType("func([]uint8, []uint8) bool")).Return(nil).Run(func(args mock.Arguments) {
		fn := args.Get(0).(func([]byte, []byte) bool)
		fn([]byte("project1"), expectedData1)
		fn([]byte("project2"), expectedData2)
	})

	// Act
	results, err := metaDB.GetAll()

	// Assert
	assert.NoError(t, err, "Getting all metadata entries should not return an error")
	assert.NotNil(t, results, "Results should not be nil")
	assert.Len(t, results, 2, "Number of entries should match")
	assert.Equal(t, entry1.ProjectID, results[0].ProjectID, "ProjectID should match for entry 1")
	assert.Equal(t, entry2.ProjectID, results[1].ProjectID, "ProjectID should match for entry 2")

	db.AssertExpectations(t)
}

// TestMetadataDatabase_GetAll_Error tests the GetAll method of MetadataDatabase for error.
func TestMetadataDatabase_GetAll_Error(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	metaDB := database.NewMetadataDatabase(db)
	expectedErr := errors.New("database error")

	// Mock behavior
	db.On("Iterate", mock.AnythingOfType("func([]uint8, []uint8) bool")).Return(expectedErr)

	// Act
	results, err := metaDB.GetAll()

	// Assert
	assert.Error(t, err, "Getting all metadata entries should return an error")
	assert.Nil(t, results, "Results should be nil")
	assert.EqualError(t, err, expectedErr.Error(), "Error should match expected error")

	db.AssertExpectations(t)
}

// TestMetadataDatabase_Update_Success tests the Update method of MetadataDatabase for success.
func TestMetadataDatabase_Update_Success(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	metaDB := database.NewMetadataDatabase(db)
	entry := &database.MetadataEntry{
		ProjectID:    "project1",
		DatabaseName: "db1",
		DatabasePath: "/path/to/db1",
		CreationDate: time.Now(),
		LastUpdated:  time.Now(),
	}

	// Mock behavior
	db.On("Set", []byte(entry.ProjectID), mockJSON(entry)).Return(nil)

	// Act
	err := metaDB.Update(entry)

	// Assert
	assert.NoError(t, err, "Updating metadata entry should not return an error")
	db.AssertExpectations(t)
}

// TestMetadataDatabase_Update_Error tests the Update method of MetadataDatabase for error.
func TestMetadataDatabase_Update_Error(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	metaDB := database.NewMetadataDatabase(db)
	entry := &database.MetadataEntry{
		ProjectID:    "project1",
		DatabaseName: "db1",
		DatabasePath: "/path/to/db1",
		CreationDate: time.Now(),
		LastUpdated:  time.Now(),
	}
	expectedErr := errors.New("database error")

	// Mock behavior
	db.On("Set", []byte(entry.ProjectID), mockJSON(entry)).Return(expectedErr)

	// Act
	err := metaDB.Update(entry)

	// Assert
	assert.Error(t, err, "Updating metadata entry should return an error")
	assert.EqualError(t, err, expectedErr.Error(), "Error should match expected error")
	db.AssertExpectations(t)
}

// TestMetadataDatabase_Delete_Success tests the Delete method of MetadataDatabase for success.
func TestMetadataDatabase_Delete_Success(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	metaDB := database.NewMetadataDatabase(db)
	projectID := "project1"

	// Mock behavior
	db.On("Delete", []byte(projectID)).Return(nil)

	// Act
	err := metaDB.Delete(projectID)

	// Assert
	assert.NoError(t, err, "Deleting metadata entry should not return an error")
	db.AssertExpectations(t)
}

// TestMetadataDatabase_Delete_Error tests the Delete method of MetadataDatabase for error.
func TestMetadataDatabase_Delete_Error(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	metaDB := database.NewMetadataDatabase(db)
	projectID := "project1"
	expectedErr := errors.New("database error")

	// Mock behavior
	db.On("Delete", []byte(projectID)).Return(expectedErr)

	// Act
	err := metaDB.Delete(projectID)

	// Assert
	assert.Error(t, err, "Deleting metadata entry should return an error")
	assert.EqualError(t, err, expectedErr.Error(), "Error should match expected error")
	db.AssertExpectations(t)
}

// Additional test cases for other methods can be added similarly.

// mockJSON serializes the provided object to JSON.
func mockJSON(obj interface{}) []byte {
	data, _ := json.Marshal(obj)
	return data
}

// TestMetadataDatabase_WorkingVersion_Success tests the WorkingVersion method of MetadataDatabase for success.
func TestMetadataDatabase_WorkingVersion_Success(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	metaDB := database.NewMetadataDatabase(db)
	expectedVersion := int64(123)

	// Mock behavior
	db.On("WorkingVersion").Return(expectedVersion)

	// Act
	version := metaDB.WorkingVersion()

	// Assert
	assert.Equal(t, expectedVersion, version, "WorkingVersion should return the expected version")
	db.AssertExpectations(t)
}

// TestMetadataDatabase_SaveVersion_Success tests the SaveVersion method of MetadataDatabase for success.
func TestMetadataDatabase_SaveVersion_Success(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	metaDB := database.NewMetadataDatabase(db)
	expectedHash := []byte{0x12, 0x34, 0x56}
	expectedVersion := int64(123)
	expectedError := error(nil)

	// Mock behavior
	db.On("SaveVersion").Return(expectedHash, expectedVersion, expectedError)

	// Act
	hash, version, err := metaDB.SaveVersion()

	// Assert
	assert.NoError(t, err, "SaveVersion should not return an error")
	assert.Equal(t, expectedHash, hash, "Returned hash should match expected hash")
	assert.Equal(t, expectedVersion, version, "Returned version should match expected version")
	db.AssertExpectations(t)
}

// TestMetadataDatabase_Load_Success tests the Load method of MetadataDatabase for success.
func TestMetadataDatabase_Load_Success(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	metaDB := database.NewMetadataDatabase(db)
	expectedVersion := int64(123)
	expectedError := error(nil)

	// Mock behavior
	db.On("Load").Return(expectedVersion, expectedError)

	// Act
	version, err := metaDB.Load()

	// Assert
	assert.NoError(t, err, "Load should not return an error")
	assert.Equal(t, expectedVersion, version, "Returned version should match expected version")
	db.AssertExpectations(t)
}

// TestMetadataDatabase_LoadVersion_Success tests the LoadVersion method of MetadataDatabase for success.
func TestMetadataDatabase_LoadVersion_Success(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	metaDB := database.NewMetadataDatabase(db)
	targetVersion := int64(123)
	expectedVersion := int64(456)
	expectedError := error(nil)

	// Mock behavior
	db.On("LoadVersion", targetVersion).Return(expectedVersion, expectedError)

	// Act
	version, err := metaDB.LoadVersion(targetVersion)

	// Assert
	assert.NoError(t, err, "LoadVersion should not return an error")
	assert.Equal(t, expectedVersion, version, "Returned version should match expected version")
	db.AssertExpectations(t)
}

// TestMetadataDatabase_String_Success tests the String method of MetadataDatabase for success.
func TestMetadataDatabase_String_Success(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	metaDB := database.NewMetadataDatabase(db)
	expectedString := "metadata_database_string"
	expectedError := error(nil)

	// Mock behavior
	db.On("String").Return(expectedString, expectedError)

	// Act
	str, err := metaDB.String()

	// Assert
	assert.NoError(t, err, "String should not return an error")
	assert.Equal(t, expectedString, str, "Returned string should match expected string")
	db.AssertExpectations(t)
}

// TestMetadataDatabase_WorkingHash_Success tests the WorkingHash method of MetadataDatabase for success.
func TestMetadataDatabase_WorkingHash_Success(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	metaDB := database.NewMetadataDatabase(db)
	expectedHash := []byte{0x12, 0x34, 0x56}

	// Mock behavior
	db.On("WorkingHash").Return(expectedHash)

	// Act
	hash := metaDB.WorkingHash()

	// Assert
	assert.Equal(t, expectedHash, hash, "WorkingHash should return the expected hash")
	db.AssertExpectations(t)
}

// TestMetadataDatabase_AvailableVersions_Success tests the AvailableVersions method of MetadataDatabase for success.
func TestMetadataDatabase_AvailableVersions_Success(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	metaDB := database.NewMetadataDatabase(db)
	expectedVersions := []int{1, 2, 3}

	// Mock behavior
	db.On("AvailableVersions").Return(expectedVersions)

	// Act
	versions := metaDB.AvailableVersions()

	// Assert
	assert.Equal(t, expectedVersions, versions, "AvailableVersions should return the expected versions")
	db.AssertExpectations(t)
}

// TestMetadataDatabase_IsEmpty_Success tests the IsEmpty method of MetadataDatabase for success.
func TestMetadataDatabase_IsEmpty_Success(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	metaDB := database.NewMetadataDatabase(db)
	expectedResult := true

	// Mock behavior
	db.On("IsEmpty").Return(expectedResult)

	// Act
	isEmpty := metaDB.IsEmpty()

	// Assert
	assert.Equal(t, expectedResult, isEmpty, "IsEmpty should return the expected result")
	db.AssertExpectations(t)
}
