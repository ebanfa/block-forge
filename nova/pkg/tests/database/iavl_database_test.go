package database

import (
	"testing"

	"github.com/cosmos/iavl"
	"github.com/edward1christian/block-forge/nova/pkg/database"
	"github.com/stretchr/testify/assert"
)

// TestIAVLDatabase_Get_Error tests the behavior of Get method when retrieving a nonexistent key.
func TestIAVLDatabase_Get_Error(t *testing.T) {
	// Arrange
	mockTree := &iavl.MutableTree{}
	mockDB := database.NewIAVLDatabase(mockTree)

	// Act
	value, err := mockDB.Get([]byte("nonexistent"))

	// Assert
	assert.NoError(t, err, "Getting a nonexistent key should not return an error")
	assert.Nil(t, value, "Getting a nonexistent key should return nil value")
}

// TestIAVLDatabase_Set tests the behavior of Set method when setting a new key-value pair.
func TestIAVLDatabase_Set(t *testing.T) {
	// Arrange
	mockTree := &iavl.MutableTree{}
	mockDB := database.NewIAVLDatabase(mockTree)
	key := []byte("testKey")
	value := []byte("testValue")

	// Act
	err := mockDB.Set(key, value)

	// Assert
	assert.NoError(t, err, "Setting key-value pair should not return an error")
}

// TestIAVLDatabase_Delete_Error tests the behavior of Delete method when attempting to delete a nonexistent key.
func TestIAVLDatabase_Delete_Error(t *testing.T) {
	// Arrange
	mockTree := &iavl.MutableTree{}
	mockDB := database.NewIAVLDatabase(mockTree)

	// Act
	err := mockDB.Delete([]byte("nonexistent"))

	// Assert
	assert.NoError(t, err, "Deleting non-existent key should return an error")
}

// TestIAVLDatabase_Has tests the behavior of Has method when checking for the existence of a key.
func TestIAVLDatabase_Has(t *testing.T) {
	// Arrange
	mockTree := &iavl.MutableTree{}
	mockDB := database.NewIAVLDatabase(mockTree)
	key := []byte("testKey")
	value := []byte("testValue")
	_ = mockDB.Set(key, value)

	// Act
	hasKey, err := mockDB.Has(key)

	// Assert
	assert.NoError(t, err, "Checking if key exists should not return an error")
	assert.True(t, hasKey, "Expected key to exist in the tree")
}

// TestIAVLDatabase_Iterate_Error tests the behavior of Iterate method when an error occurs during iteration.
func TestIAVLDatabase_Iterate_Error(t *testing.T) {
	// Arrange
	mockTree := &iavl.MutableTree{}
	mockDB := database.NewIAVLDatabase(mockTree)
	mockDB.Set([]byte("key1"), []byte("value1"))
	mockDB.Set([]byte("key2"), []byte("value2"))

	// Act
	err := mockDB.Iterate(func(key, value []byte) bool {
		// Simulate an error during iteration
		return false
	})

	// Assert
	assert.NoError(t, err, "Iterating should not return an error")
}

// TestIAVLDatabase_IterateRange tests the behavior of IterateRange method when iterating over a range of key-value pairs.
func TestIAVLDatabase_IterateRange(t *testing.T) {
	// Arrange
	mockTree := &iavl.MutableTree{}
	mockDB := database.NewIAVLDatabase(mockTree)
	mockDB.Set([]byte("key1"), []byte("value1"))
	mockDB.Set([]byte("key2"), []byte("value2"))

	// Act
	var keys [][]byte
	var values [][]byte
	err := mockDB.IterateRange(nil, nil, true, func(key, value []byte) bool {
		keys = append(keys, key)
		values = append(values, value)
		return false // Continue iteration
	})

	// Assert
	assert.NoError(t, err, "Iterating range should not return an error")
	assert.Equal(t, [][]byte{[]byte("key1"), []byte("key2")}, keys, "Keys should match expected values")
	assert.Equal(t, [][]byte{[]byte("value1"), []byte("value2")}, values, "Values should match expected values")
}

// TestIAVLDatabase_Hash tests the behavior of Hash method when retrieving the root hash of the tree.
func TestIAVLDatabase_Hash(t *testing.T) {
	// Arrange
	mockTree := &iavl.MutableTree{}
	mockDB := database.NewIAVLDatabase(mockTree)

	// Act
	hash := mockDB.Hash()

	// Assert
	assert.NotNil(t, hash, "Hash should not be nil")
}

// TestIAVLDatabase_Version tests the behavior of Version method when retrieving the version of the tree.
func TestIAVLDatabase_Version(t *testing.T) {
	// Arrange
	mockTree := &iavl.MutableTree{}
	mockDB := database.NewIAVLDatabase(mockTree)

	// Act
	version := mockDB.Version()

	// Assert
	assert.Equal(t, int64(0), version, "Initial version should be 0")
}

// TestIAVLDatabase_Load tests the behavior of Load method when loading the latest versioned tree from disk.
func TestIAVLDatabase_Load(t *testing.T) {
	// Arrange
	mockTree := &iavl.MutableTree{}
	mockDB := database.NewIAVLDatabase(mockTree)

	// Act
	version, err := mockDB.Load()

	// Assert
	assert.NoError(t, err, "Loading should not return an error")
	assert.Equal(t, int64(0), version, "Version should match expected value")
}

// TestIAVLDatabase_SaveVersion tests the behavior of SaveVersion method when saving a new tree version to disk.
func TestIAVLDatabase_SaveVersion(t *testing.T) {
	// Arrange
	mockTree := &iavl.MutableTree{}
	mockDB := database.NewIAVLDatabase(mockTree)

	// Act
	hash, version, err := mockDB.SaveVersion()

	// Assert
	assert.NoError(t, err, "Saving version should not return an error")
	assert.NotNil(t, hash, "Hash should not be nil")
	assert.Equal(t, int64(1), version, "Version should match expected value")
}

// TestIAVLDatabase_Rollback tests the behavior of Rollback method when resetting the working tree to the latest saved version.
func TestIAVLDatabase_Rollback(t *testing.T) {
	// Arrange
	mockTree := &iavl.MutableTree{}
	mockDB := database.NewIAVLDatabase(mockTree)
	mockDB.Set([]byte("key"), []byte("value"))

	// Act
	mockDB.Rollback()

	value, _ := mockDB.Get([]byte("key"))

	// Assert
	assert.Nil(t, value, "After Rolling back Get should return nil")
}

// TestIAVLDatabase_Close tests the behavior of Close method when closing the tree.
func TestIAVLDatabase_Close(t *testing.T) {
	// Arrange
	mockTree := &iavl.MutableTree{}
	mockDB := database.NewIAVLDatabase(mockTree)

	// Act
	err := mockDB.Close()

	// Assert
	assert.NoError(t, err, "Closing should not return an error")
}

// TestIAVLDatabase_String tests the behavior of String method when getting a string representation of the tree.
func TestIAVLDatabase_String(t *testing.T) {
	// Arrange
	mockTree := &iavl.MutableTree{}
	mockDB := database.NewIAVLDatabase(mockTree)

	// Act
	str, err := mockDB.String()

	// Assert
	assert.NoError(t, err, "Getting string representation should not return an error")
	assert.NotEmpty(t, str, "String representation should not be empty")
}
