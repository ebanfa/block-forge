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
