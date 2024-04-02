package mocks

import (
	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
	blockchain "github.com/edward1christian/block-forge/pkg/blockchain/interfaces"
	"github.com/edward1christian/block-forge/pkg/etl"
	"github.com/stretchr/testify/mock"
)

type MockETLComponent struct {
	mock.Mock
	etl.ETLComponent
	processID string
}

func (m *MockETLComponent) ID() string {
	return "mock_adapter_id"
}

func (m *MockETLComponent) Name() string {
	return "Mock Adapter"
}

func (m *MockETLComponent) Description() string {
	return "Mock Adapter Description"
}

func (m *MockETLComponent) setProcessID(processID string) {
	m.processID = processID
}

func (m *MockETLComponent) getProcessID() string {
	return m.processID
}

func (m *MockETLComponent) Start(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockETLComponent) Stop(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockETLComponent) Initialize(ctx *context.Context, system system.System) error {
	args := m.Called(ctx, system)
	return args.Error(0)
}

// Implementing the Blockchain() method from the BlockchainRelay interface
func (m *MockETLComponent) Blockchain() *blockchain.Blockchain {
	args := m.Called()
	// Dereference the pointer to get the actual value
	return args.Get(0).(*blockchain.Blockchain)
}
