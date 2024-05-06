package mocks

import (
	"github.com/edward1christian/block-forge/pkg/application/db"
	"github.com/stretchr/testify/mock"
)

// MockDatabaseFactory is a mock implementation of the DatabaseFactory interface for testing purposes.
type MockDatabaseFactory struct {
	mock.Mock
}

// CreateDatabase is a mock method that simulates creating a database instance.
func (m *MockDatabaseFactory) CreateDatabase(name, path string) (db.Database, error) {
	args := m.Called(name, path)
	return args.Get(0).(db.Database), args.Error(1)
}
