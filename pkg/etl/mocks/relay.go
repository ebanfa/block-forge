package mocks

import (
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
	blockchain "github.com/edward1christian/block-forge/pkg/blockchain/interfaces"
	"github.com/edward1christian/block-forge/pkg/etl"
	"github.com/stretchr/testify/mock"
)

type MockRelay struct {
	mock.Mock
	etl.BlockchainRelay
	processID  string
	blockchain blockchain.Blockchain
}

func (m *MockRelay) ID() string {
	return "mock_relay_id"
}

func (m *MockRelay) Name() string {
	return "Mock Relay"
}

func (m *MockRelay) Description() string {
	return "Mock Relay Description"
}

func (m *MockRelay) setProcessID(processID string) {
	m.processID = processID
}

func (m *MockRelay) getProcessID() string {
	return m.processID
}

func (m *MockRelay) Start(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockRelay) Stop(ctx *context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockRelay) Initialize(ctx *context.Context, system system.System) error {
	args := m.Called(ctx, system)
	return args.Error(0)
}

func (m *MockRelay) Blockchain() *blockchain.Blockchain {
	args := m.Called()
	// Dereference the pointer to get the actual value
	return args.Get(0).(*blockchain.Blockchain)
}

type MockBlockchainRelay struct {
	*MockRelay // Embedding the MockRelay struct
}

// Implementing the Blockchain() method from the BlockchainRelay interface
func (m *MockBlockchainRelay) Blockchain() *blockchain.Blockchain {
	args := m.Called()
	// Dereference the pointer to get the actual value
	return args.Get(0).(*blockchain.Blockchain)
}
