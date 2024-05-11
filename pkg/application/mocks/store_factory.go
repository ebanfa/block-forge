package mocks

import (
	"github.com/edward1christian/block-forge/pkg/application/db"
	"github.com/edward1christian/block-forge/pkg/application/store"
	"github.com/stretchr/testify/mock"
)

// MockStoreFactory is a mock implementation of the StoreFactory struct
type MockStoreFactory struct {
	mock.Mock
}

// NewStoreFactory is a mocked method for creating a new StoreFactory instance
func (m *MockStoreFactory) NewStoreFactory(databasesDir string, dbFactory db.DatabaseFactory) *store.StoreFactory {
	args := m.Called(databasesDir, dbFactory)
	return args.Get(0).(*store.StoreFactory)
}

// CreateStore is a mocked method for creating a store
func (m *MockStoreFactory) CreateStore(name string) (store.Store, error) {
	args := m.Called(name)
	return args.Get(0).(store.Store), args.Error(1)
}
