package database

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/edward1christian/block-forge/nova/pkg/config"
	"github.com/edward1christian/block-forge/nova/pkg/database"
	"github.com/edward1christian/block-forge/nova/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestProjectDatabase_Insert_Success tests the Insert method of ProjectDatabase for success.
func TestProjectDatabase_Insert_Success(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	projectDB := database.NewProjectDatabase(db)
	project := &config.Project{
		ID:   "project1",
		Name: "Project 1",
	}

	// Mock behavior
	db.On("Set", []byte(project.ID), mockJSON(project)).Return(nil)

	// Act
	err := projectDB.Insert(project)

	// Assert
	assert.NoError(t, err, "Inserting project entry should not return an error")
	db.AssertExpectations(t)
}

// TestProjectDatabase_Insert_Error tests the Insert method of ProjectDatabase for error.
func TestProjectDatabase_Insert_Error(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	projectDB := database.NewProjectDatabase(db)
	project := &config.Project{
		ID:   "project1",
		Name: "Project 1",
	}
	expectedErr := errors.New("database error")

	// Mock behavior
	db.On("Set", []byte(project.ID), mockJSON(project)).Return(expectedErr)

	// Act
	err := projectDB.Insert(project)

	// Assert
	assert.Error(t, err, "Inserting project entry should return an error")
	assert.EqualError(t, err, expectedErr.Error(), "Error should match expected error")
	db.AssertExpectations(t)
}

// TestProjectDatabase_Get_Success tests the Get method of ProjectDatabase for success.
func TestProjectDatabase_Get_Success(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	projectDB := database.NewProjectDatabase(db)
	project := &config.Project{
		ID:   "project1",
		Name: "Project 1",
	}
	expectedData, _ := json.Marshal(project)

	// Mock behavior
	db.On("Get", []byte(project.ID)).Return(expectedData, nil)

	// Act
	result, err := projectDB.Get(project.ID)

	// Assert
	assert.NoError(t, err, "Getting project entry should not return an error")
	assert.NotNil(t, result, "Result should not be nil")
	assert.Equal(t, project.ID, result.ID, "Project ID should match")
	assert.Equal(t, project.Name, result.Name, "Project Name should match")
	db.AssertExpectations(t)
}

// TestProjectDatabase_Get_Error tests the Get method of ProjectDatabase for error.
func TestProjectDatabase_Get_Error(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	projectDB := database.NewProjectDatabase(db)
	projectID := "project1"
	expectedErr := errors.New("database error")

	// Mock behavior
	db.On("Get", []byte(projectID)).Return([]byte{}, expectedErr)

	// Act
	result, err := projectDB.Get(projectID)

	// Assert
	assert.Error(t, err, "Getting project entry should return an error")
	assert.Nil(t, result, "Result should be nil")
	assert.EqualError(t, err, expectedErr.Error(), "Error should match expected error")
	db.AssertExpectations(t)
}

// TestProjectDatabase_GetAll_Success tests the GetAll method of ProjectDatabase for success.
func TestProjectDatabase_GetAll_Success(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	projectDB := database.NewProjectDatabase(db)
	project1 := &config.Project{
		ID:   "project1",
		Name: "Project 1",
	}
	project2 := &config.Project{
		ID:   "project2",
		Name: "Project 2",
	}
	expectedData1, _ := json.Marshal(project1)
	expectedData2, _ := json.Marshal(project2)

	// Mock behavior
	db.On("Iterate", mock.AnythingOfType("func([]uint8, []uint8) bool")).Return(nil).Run(func(args mock.Arguments) {
		fn := args.Get(0).(func([]byte, []byte) bool)
		fn([]byte("project1"), expectedData1)
		fn([]byte("project2"), expectedData2)
	})

	// Act
	results, err := projectDB.GetAll()

	// Assert
	assert.NoError(t, err, "Getting all project entries should not return an error")
	assert.NotNil(t, results, "Results should not be nil")
	assert.Len(t, results, 2, "Number of projects should match")
	assert.Equal(t, project1.ID, results[0].ID, "Project ID should match for project 1")
	assert.Equal(t, project2.ID, results[1].ID, "Project ID should match for project 2")
	db.AssertExpectations(t)
}

// TestProjectDatabase_GetAll_Error tests the GetAll method of ProjectDatabase for error.
func TestProjectDatabase_GetAll_Error(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	projectDB := database.NewProjectDatabase(db)
	expectedErr := errors.New("database error")

	// Mock behavior
	db.On("Iterate", mock.AnythingOfType("func([]uint8, []uint8) bool")).Return(expectedErr)

	// Act
	results, err := projectDB.GetAll()

	// Assert
	assert.Error(t, err, "Getting all project entries should return an error")
	assert.Nil(t, results, "Results should be nil")
	assert.EqualError(t, err, expectedErr.Error(), "Error should match expected error")
	db.AssertExpectations(t)
}

// TestProjectDatabase_Update_Success tests the Update method of ProjectDatabase for success.
func TestProjectDatabase_Update_Success(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	projectDB := database.NewProjectDatabase(db)
	project := &config.Project{
		ID:   "project1",
		Name: "Project 1",
	}

	// Mock behavior
	db.On("Set", []byte(project.ID), mockJSON(project)).Return(nil)

	// Act
	err := projectDB.Update(project)

	// Assert
	assert.NoError(t, err, "Updating project entry should not return an error")
	db.AssertExpectations(t)
}

// TestProjectDatabase_Update_Error tests the Update method of ProjectDatabase for error.
func TestProjectDatabase_Update_Error(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	projectDB := database.NewProjectDatabase(db)
	project := &config.Project{
		ID:   "project1",
		Name: "Project 1",
	}
	expectedErr := errors.New("database error")

	// Mock behavior
	db.On("Set", []byte(project.ID), mockJSON(project)).Return(expectedErr)

	// Act
	err := projectDB.Update(project)

	// Assert
	assert.Error(t, err, "Updating project entry should return an error")
	assert.EqualError(t, err, expectedErr.Error(), "Error should match expected error")
	db.AssertExpectations(t)
}

// TestProjectDatabase_Delete_Success tests the Delete method of ProjectDatabase for success.
func TestProjectDatabase_Delete_Success(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	projectDB := database.NewProjectDatabase(db)
	projectID := "project1"

	// Mock behavior
	db.On("Delete", []byte(projectID)).Return(nil)

	// Act
	err := projectDB.Delete(projectID)

	// Assert
	assert.NoError(t, err, "Deleting project entry should not return an error")
	db.AssertExpectations(t)
}

// TestProjectDatabase_Delete_Error tests the Delete method of ProjectDatabase for error.
func TestProjectDatabase_Delete_Error(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	projectDB := database.NewProjectDatabase(db)
	projectID := "project1"
	expectedErr := errors.New("database error")

	// Mock behavior
	db.On("Delete", []byte(projectID)).Return(expectedErr)

	// Act
	err := projectDB.Delete(projectID)

	// Assert
	assert.Error(t, err, "Deleting project entry should return an error")
	assert.EqualError(t, err, expectedErr.Error(), "Error should match expected error")
	db.AssertExpectations(t)
}

// TestProjectDatabase_SaveVersion_Success tests the SaveVersion method of ProjectDatabase for success.
func TestProjectDatabase_SaveVersion_Success(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	projectDB := database.NewProjectDatabase(db)
	expectedVersion := int64(123)
	expectedData := []byte("some_data")

	// Mock behavior
	db.On("SaveVersion").Return(expectedData, expectedVersion, nil)

	// Act
	data, version, err := projectDB.SaveVersion()

	// Assert
	assert.NoError(t, err, "Saving version should not return an error")
	assert.Equal(t, expectedVersion, version, "Version should match")
	assert.Equal(t, expectedData, data, "Data should match")
	db.AssertExpectations(t)
}

// TestProjectDatabase_SaveVersion_Error tests the SaveVersion method of ProjectDatabase for error.
func TestProjectDatabase_SaveVersion_Error(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	projectDB := database.NewProjectDatabase(db)
	expectedErr := errors.New("database error")

	// Mock behavior
	db.On("SaveVersion").Return([]byte{}, int64(0), expectedErr)

	// Act
	data, version, err := projectDB.SaveVersion()

	// Assert
	assert.Error(t, err, "Saving version should return an error")
	assert.Zero(t, version, "Version should be zero")
	assert.Equal(t, data, []byte{})
	assert.EqualError(t, err, expectedErr.Error(), "Error should match expected error")
	db.AssertExpectations(t)
}

// TestProjectDatabase_Load_Success tests the Load method of ProjectDatabase for success.
func TestProjectDatabase_Load_Success(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	projectDB := database.NewProjectDatabase(db)
	expectedVersion := int64(123)

	// Mock behavior
	db.On("Load").Return(expectedVersion, nil)

	// Act
	version, err := projectDB.Load()

	// Assert
	assert.NoError(t, err, "Loading version should not return an error")
	assert.Equal(t, expectedVersion, version, "Version should match")
	db.AssertExpectations(t)
}

// TestProjectDatabase_Load_Error tests the Load method of ProjectDatabase for error.
func TestProjectDatabase_Load_Error(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	projectDB := database.NewProjectDatabase(db)
	expectedErr := errors.New("database error")

	// Mock behavior
	db.On("Load").Return(int64(0), expectedErr)

	// Act
	version, err := projectDB.Load()

	// Assert
	assert.Error(t, err, "Loading version should return an error")
	assert.Zero(t, version, "Version should be zero")
	assert.EqualError(t, err, expectedErr.Error(), "Error should match expected error")
	db.AssertExpectations(t)
}

// TestProjectDatabase_LoadVersion_Success tests the LoadVersion method of ProjectDatabase for success.
func TestProjectDatabase_LoadVersion_Success(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	projectDB := database.NewProjectDatabase(db)
	targetVersion := int64(123)
	expectedVersion := int64(456)

	// Mock behavior
	db.On("LoadVersion", targetVersion).Return(expectedVersion, nil)

	// Act
	version, err := projectDB.LoadVersion(targetVersion)

	// Assert
	assert.NoError(t, err, "Loading version should not return an error")
	assert.Equal(t, expectedVersion, version, "Version should match")
	db.AssertExpectations(t)
}

// TestProjectDatabase_LoadVersion_Error tests the LoadVersion method of ProjectDatabase for error.
func TestProjectDatabase_LoadVersion_Error(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	projectDB := database.NewProjectDatabase(db)
	targetVersion := int64(123)
	expectedErr := errors.New("database error")

	// Mock behavior
	db.On("LoadVersion", targetVersion).Return(int64(0), expectedErr)

	// Act
	version, err := projectDB.LoadVersion(targetVersion)

	// Assert
	assert.Error(t, err, "Loading version should return an error")
	assert.Zero(t, version, "Version should be zero")
	assert.EqualError(t, err, expectedErr.Error(), "Error should match expected error")
	db.AssertExpectations(t)
}

// TestProjectDatabase_String_Success tests the String method of ProjectDatabase for success.
func TestProjectDatabase_String_Success(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	projectDB := database.NewProjectDatabase(db)
	expectedString := "database_string_representation"

	// Mock behavior
	db.On("String").Return(expectedString, nil)

	// Act
	str, err := projectDB.String()

	// Assert
	assert.NoError(t, err, "String method should not return an error")
	assert.Equal(t, expectedString, str, "String representation should match")
	db.AssertExpectations(t)
}

// TestProjectDatabase_String_Error tests the String method of ProjectDatabase for error.
func TestProjectDatabase_String_Error(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	projectDB := database.NewProjectDatabase(db)
	expectedErr := errors.New("database error")

	// Mock behavior
	db.On("String").Return("", expectedErr)

	// Act
	str, err := projectDB.String()

	// Assert
	assert.Error(t, err, "String method should return an error")
	assert.Empty(t, str, "String representation should be empty")
	assert.EqualError(t, err, expectedErr.Error(), "Error should match expected error")
	db.AssertExpectations(t)
}

// TestProjectDatabase_WorkingVersion tests the WorkingVersion method of ProjectDatabase.
func TestProjectDatabase_WorkingVersion(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	projectDB := database.NewProjectDatabase(db)
	expectedVersion := int64(123)

	// Mock behavior
	db.On("WorkingVersion").Return(expectedVersion)

	// Act
	version := projectDB.WorkingVersion()

	// Assert
	assert.Equal(t, expectedVersion, version, "Working version should match")
	db.AssertExpectations(t)
}

// TestProjectDatabase_WorkingHash tests the WorkingHash method of ProjectDatabase.
func TestProjectDatabase_WorkingHash(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	projectDB := database.NewProjectDatabase(db)
	expectedHash := []byte{1, 2, 3}

	// Mock behavior
	db.On("WorkingHash").Return(expectedHash)

	// Act
	hash := projectDB.WorkingHash()

	// Assert
	assert.Equal(t, expectedHash, hash, "Working hash should match")
	db.AssertExpectations(t)
}

// TestProjectDatabase_AvailableVersions tests the AvailableVersions method of ProjectDatabase.
func TestProjectDatabase_AvailableVersions(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	projectDB := database.NewProjectDatabase(db)
	expectedVersions := []int{1, 2, 3}

	// Mock behavior
	db.On("AvailableVersions").Return(expectedVersions)

	// Act
	versions := projectDB.AvailableVersions()

	// Assert
	assert.Equal(t, expectedVersions, versions, "Available versions should match")
	db.AssertExpectations(t)
}

// TestProjectDatabase_IsEmpty tests the IsEmpty method of ProjectDatabase.
func TestProjectDatabase_IsEmpty(t *testing.T) {
	// Arrange
	db := &mocks.MockDatabase{} // Replace mockDatabase with your actual database mock
	projectDB := database.NewProjectDatabase(db)
	expectedEmpty := true

	// Mock behavior
	db.On("IsEmpty").Return(expectedEmpty)

	// Act
	isEmpty := projectDB.IsEmpty()

	// Assert
	assert.Equal(t, expectedEmpty, isEmpty, "IsEmpty result should match")
	db.AssertExpectations(t)
}
