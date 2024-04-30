package database

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/edward1christian/block-forge/nova/pkg/database"
	"github.com/edward1christian/block-forge/nova/pkg/mocks"
	"github.com/stretchr/testify/assert"
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
