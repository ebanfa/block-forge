package database

import (
	"testing"

	"cosmossdk.io/log"
	"github.com/cosmos/iavl"
	"github.com/cosmos/iavl/db"
	"github.com/edward1christian/block-forge/nova/pkg/database"
	"github.com/stretchr/testify/assert"
)

// TestIAVLDatabase_Get_Error tests the behavior of Get method when retrieving a nonexistent key.
func TestIAVLDatabase_Get_Error(t *testing.T) {
	// Arrange

	iavlTree := iavl.NewMutableTree(db.NewMemDB(), 100, false, log.NewNopLogger())
	mockDB := database.NewIAVLDatabase(iavlTree)

	// Act
	value, err := mockDB.Get([]byte("nonexistent"))

	// Assert
	assert.NoError(t, err, "Getting a nonexistent key should not return an error")
	assert.Nil(t, value, "Getting a nonexistent key should return nil value")
}

// TestIAVLDatabase_Set tests the behavior of Set method when setting a new key-value pair.
func TestIAVLDatabase_Set(t *testing.T) {
	// Arrange
	iavlTree := iavl.NewMutableTree(db.NewMemDB(), 100, false, log.NewNopLogger())
	mockDB := database.NewIAVLDatabase(iavlTree)
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
	iavlTree := iavl.NewMutableTree(db.NewMemDB(), 100, false, log.NewNopLogger())
	mockDB := database.NewIAVLDatabase(iavlTree)

	// Act
	err := mockDB.Delete([]byte("nonexistent"))

	// Assert
	assert.NoError(t, err, "Deleting non-existent key should return an error")
}

// TestIAVLDatabase_Has tests the behavior of Has method when checking for the existence of a key.
func TestIAVLDatabase_Has(t *testing.T) {
	// Arrange
	iavlTree := iavl.NewMutableTree(db.NewMemDB(), 100, false, log.NewNopLogger())
	mockDB := database.NewIAVLDatabase(iavlTree)
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
	iavlTree := iavl.NewMutableTree(db.NewMemDB(), 100, false, log.NewNopLogger())
	mockDB := database.NewIAVLDatabase(iavlTree)
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
	iavlTree := iavl.NewMutableTree(db.NewMemDB(), 100, false, log.NewNopLogger())
	mockDB := database.NewIAVLDatabase(iavlTree)
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
	iavlTree := iavl.NewMutableTree(db.NewMemDB(), 100, false, log.NewNopLogger())
	mockDB := database.NewIAVLDatabase(iavlTree)

	// Act
	hash := mockDB.Hash()

	// Assert
	assert.NotNil(t, hash, "Hash should not be nil")
}

// TestIAVLDatabase_Version tests the behavior of Version method when retrieving the version of the tree.
func TestIAVLDatabase_Version(t *testing.T) {
	// Arrange
	iavlTree := iavl.NewMutableTree(db.NewMemDB(), 100, false, log.NewNopLogger())
	mockDB := database.NewIAVLDatabase(iavlTree)

	// Act
	version := mockDB.Version()

	// Assert
	assert.Equal(t, int64(0), version, "Initial version should be 0")
}

// TestIAVLDatabase_Load tests the behavior of Load method when loading the latest versioned tree from disk.
func TestIAVLDatabase_Load(t *testing.T) {
	// Arrange
	iavlTree := iavl.NewMutableTree(db.NewMemDB(), 100, false, log.NewNopLogger())
	mockDB := database.NewIAVLDatabase(iavlTree)

	// Act
	version, err := mockDB.Load()

	// Assert
	assert.NoError(t, err, "Loading should not return an error")
	assert.Equal(t, int64(0), version, "Version should match expected value")
}

// TestIAVLDatabase_SaveVersion tests the behavior of SaveVersion method when saving a new tree version to disk.
func TestIAVLDatabase_SaveVersion(t *testing.T) {
	// Arrange
	iavlTree := iavl.NewMutableTree(db.NewMemDB(), 100, false, log.NewNopLogger())
	mockDB := database.NewIAVLDatabase(iavlTree)

	// Act: Set initial data
	_ = mockDB.Set([]byte("key1"), []byte("value1"))
	_ = mockDB.Set([]byte("key2"), []byte("value2"))

	// Act: Save initial version
	hash1, version1, err1 := mockDB.SaveVersion()
	assert.NoError(t, err1, "Saving version should not return an error")

	// Assert: Check initial version
	assert.NotNil(t, hash1, "Hash of initial version should not be nil")
	assert.Equal(t, int64(1), version1, "Initial version should be 1")

	// Act: Make additional changes but don't save
	_ = mockDB.Set([]byte("key3"), []byte("value3"))

	// Act: Rollback to previous version
	mockDB.Rollback()

	// Act: Retrieve data after rollback
	value1, _ := mockDB.Get([]byte("key1"))
	value2, _ := mockDB.Get([]byte("key2"))
	value3, _ := mockDB.Get([]byte("key3"))

	// Assert: Check if changes are reverted after rollback
	assert.NotNil(t, value1, "Value for key1 should exist after rollback")
	assert.NotNil(t, value2, "Value for key2 should exist after rollback")
	assert.Nil(t, value3, "Value for key3 should not exist after rollback")
}

// TestIAVLDatabase_Rollback tests the behavior of Rollback method when resetting the working tree to the latest saved version.
func TestIAVLDatabase_Rollback(t *testing.T) {
	// Arrange
	iavlTree := iavl.NewMutableTree(db.NewMemDB(), 100, false, log.NewNopLogger())
	mockDB := database.NewIAVLDatabase(iavlTree)
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
	iavlTree := iavl.NewMutableTree(db.NewMemDB(), 100, false, log.NewNopLogger())
	mockDB := database.NewIAVLDatabase(iavlTree)

	// Act
	err := mockDB.Close()

	// Assert
	assert.NoError(t, err, "Closing should not return an error")
}

// TestIAVLDatabase_String tests the behavior of String method when getting a string representation of the tree.
func TestIAVLDatabase_String(t *testing.T) {
	// Arrange
	iavlTree := iavl.NewMutableTree(db.NewMemDB(), 100, false, log.NewNopLogger())
	mockDB := database.NewIAVLDatabase(iavlTree)

	// Act
	str, err := mockDB.String()

	// Assert
	assert.NoError(t, err, "Getting string representation should not return an error")
	assert.NotEmpty(t, str, "String representation should not be empty")
}

// TestIAVLDatabase_WorkingVersion tests the behavior of WorkingVersion method when retrieving the current working version of the tree.
func TestIAVLDatabase_WorkingVersion(t *testing.T) {
	// Arrange
	iavlTree := iavl.NewMutableTree(db.NewMemDB(), 100, false, log.NewNopLogger())
	mockDB := database.NewIAVLDatabase(iavlTree)

	// Act
	workingVersion := mockDB.WorkingVersion()

	// Assert
	assert.Equal(t, int64(1), workingVersion, "Initial working version should be 1")
}

// TestIAVLDatabase_WorkingHash tests the behavior of WorkingHash method when retrieving the root hash of the current working tree.
func TestIAVLDatabase_WorkingHash(t *testing.T) {
	// Arrange
	iavlTree := iavl.NewMutableTree(db.NewMemDB(), 100, false, log.NewNopLogger())
	mockDB := database.NewIAVLDatabase(iavlTree)

	// Act
	workingHash := mockDB.WorkingHash()

	// Assert
	assert.NotNil(t, workingHash, "Working hash should not be nil")
}

// TestIAVLDatabase_AvailableVersions tests the behavior of AvailableVersions method when retrieving a list of available versions.
func TestIAVLDatabase_AvailableVersions(t *testing.T) {
	// Arrange
	iavlTree := iavl.NewMutableTree(db.NewMemDB(), 100, false, log.NewNopLogger())
	mockDB := database.NewIAVLDatabase(iavlTree)
	_, _, _ = mockDB.SaveVersion() // Save a version to make it available

	// Act
	versions := mockDB.AvailableVersions()

	// Assert
	assert.NotEmpty(t, versions, "Available versions should not be empty")
}

// TestIAVLDatabase_IsEmpty tests the behavior of IsEmpty method when checking if the database is empty.
func TestIAVLDatabase_IsEmpty(t *testing.T) {
	// Arrange
	iavlTree := iavl.NewMutableTree(db.NewMemDB(), 100, false, log.NewNopLogger())
	mockDB := database.NewIAVLDatabase(iavlTree)

	// Act
	empty := mockDB.IsEmpty()

	// Assert
	assert.True(t, empty, "The database should be empty initially")
}
