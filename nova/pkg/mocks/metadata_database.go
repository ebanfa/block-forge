package mocks

import (
	"github.com/edward1christian/block-forge/nova/pkg/database"
	"github.com/stretchr/testify/mock"
)

// MockMetadataDatabase is a mock struct for MetadataDatabase
type MockMetadataDatabase struct {
	mock.Mock
}

// Insert mocks the Insert method
func (m *MockMetadataDatabase) Insert(entry *database.MetadataEntry) error {
	args := m.Called(entry)
	return args.Error(0)
}

// Get mocks the Get method
func (m *MockMetadataDatabase) Get(projectID string) (*database.MetadataEntry, error) {
	args := m.Called(projectID)
	return args.Get(0).(*database.MetadataEntry), args.Error(1)
}

// GetAll mocks the GetAll method
func (m *MockMetadataDatabase) GetAll() ([]*database.MetadataEntry, error) {
	args := m.Called()
	return args.Get(0).([]*database.MetadataEntry), args.Error(1)
}

// Update mocks the Update method
func (m *MockMetadataDatabase) Update(entry *database.MetadataEntry) error {
	args := m.Called(entry)
	return args.Error(0)
}

// Delete mocks the Delete method
func (m *MockMetadataDatabase) Delete(projectID string) error {
	args := m.Called(projectID)
	return args.Error(0)
}

// WorkingVersion mocks the WorkingVersion method
func (m *MockMetadataDatabase) WorkingVersion() int64 {
	args := m.Called()
	return args.Get(0).(int64)
}

// SaveVersion mocks the SaveVersion method
func (m *MockMetadataDatabase) SaveVersion() ([]byte, int64, error) {
	args := m.Called()
	return args.Get(0).([]byte), args.Get(1).(int64), args.Error(2)
}

// Load mocks the Load method
func (m *MockMetadataDatabase) Load() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

// LoadVersion mocks the LoadVersion method
func (m *MockMetadataDatabase) LoadVersion(targetVersion int64) (int64, error) {
	args := m.Called(targetVersion)
	return args.Get(0).(int64), args.Error(1)
}

// String mocks the String method
func (m *MockMetadataDatabase) String() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

// WorkingHash mocks the WorkingHash method
func (m *MockMetadataDatabase) WorkingHash() []byte {
	args := m.Called()
	return args.Get(0).([]byte)
}

// AvailableVersions mocks the AvailableVersions method
func (m *MockMetadataDatabase) AvailableVersions() []int {
	args := m.Called()
	return args.Get(0).([]int)
}

// IsEmpty mocks the IsEmpty method
func (m *MockMetadataDatabase) IsEmpty() bool {
	args := m.Called()
	return args.Bool(0)
}
