package blockchain

// Blockchain represents a blockchain network.
type Blockchain interface {
	// Name returns the name of the blockchain.
	Name() string

	// Version returns the version of the blockchain.
	Version() string
}
