package mocks

import (
	"github.com/edward1christian/block-forge/nova/pkg/store"
	"github.com/stretchr/testify/mock"
)

// MockMetadataStore is a mock implementation of MetadataStore.
type MockMetadataStore struct {
	mock.Mock
}

// InsertMetadata mocks the InsertMetadata method of MetadataStore.
func (m *MockMetadataStore) InsertMetadata(entry *store.MetadataEntry) error {
	args := m.Called(entry)
	return args.Error(0)
}

// GetMetadata mocks the GetMetadata method of MetadataStore.
func (m *MockMetadataStore) GetMetadata(projectID string) (*store.MetadataEntry, error) {
	args := m.Called(projectID)
	return args.Get(0).(*store.MetadataEntry), args.Error(1)
}

// GetAllMetadata mocks the GetAllMetadata method of MetadataStore.
func (m *MockMetadataStore) GetAllMetadata() ([]*store.MetadataEntry, error) {
	args := m.Called()
	return args.Get(0).([]*store.MetadataEntry), args.Error(1)
}

// UpdateMetadata mocks the UpdateMetadata method of MetadataStore.
func (m *MockMetadataStore) UpdateMetadata(entry *store.MetadataEntry) error {
	args := m.Called(entry)
	return args.Error(0)
}

// DeleteMetadata mocks the DeleteMetadata method of MetadataStore.
func (m *MockMetadataStore) DeleteMetadata(projectID string) error {
	args := m.Called(projectID)
	return args.Error(0)
}

// SaveVersion mocks the SaveVersion method of MetadataStore.
func (m *MockMetadataStore) SaveVersion() ([]byte, int64, error) {
	args := m.Called()
	return args.Get(0).([]byte), args.Get(1).(int64), args.Error(2)
}

// Load mocks the Load method of MetadataStore.
func (m *MockMetadataStore) Load() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

// LoadVersion mocks the LoadVersion method of MetadataStore.
func (m *MockMetadataStore) LoadVersion(targetVersion int64) (int64, error) {
	args := m.Called(targetVersion)
	return args.Get(0).(int64), args.Error(1)
}

// String mocks the String method of MetadataStore.
func (m *MockMetadataStore) String() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

// WorkingVersion mocks the WorkingVersion method of MetadataStore.
func (m *MockMetadataStore) WorkingVersion() int64 {
	args := m.Called()
	return args.Get(0).(int64)
}

// WorkingHash mocks the WorkingHash method of MetadataStore.
func (m *MockMetadataStore) WorkingHash() []byte {
	args := m.Called()
	return args.Get(0).([]byte)
}

// AvailableVersions mocks the AvailableVersions method of MetadataStore.
func (m *MockMetadataStore) AvailableVersions() []int {
	args := m.Called()
	return args.Get(0).([]int)
}

// IsEmpty mocks the IsEmpty method of MetadataStore.
func (m *MockMetadataStore) IsEmpty() bool {
	args := m.Called()
	return args.Bool(0)
}

// Close mocks the Close method of MetadataStore.
func (m *MockMetadataStore) Close() error {
	args := m.Called()
	return args.Error(0)
}
