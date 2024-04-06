package common

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// IDGeneratorInterface interface defines the behavior of a process ID generator.
type IDGeneratorInterface interface {
	GenerateID() (string, error)
}

// RandomProcessIDGenerator provides functionality to generate unique process IDs.
type ProcessIDGenerator struct {
	IDGeneratorInterface
	prefix string
}

// NewRandomProcessIDGenerator creates a new instance of RandomProcessIDGenerator with the given prefix.
func NewProcessIDGenerator(prefix string) *ProcessIDGenerator {
	return &ProcessIDGenerator{
		prefix: prefix,
	}
}

// GenerateID generates a unique process ID.
func (gen *ProcessIDGenerator) GenerateID() (string, error) {
	// Check if the prefix is empty
	if gen.prefix == "" {
		return "", fmt.Errorf("prefix cannot be empty")
	}

	// Generate a random number to ensure uniqueness
	randomNum, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return "", fmt.Errorf("failed to generate random number: %w", err)
	}

	// Combine prefix and random number to create the process ID
	processID := fmt.Sprintf("%s-%d", gen.prefix, randomNum)

	return processID, nil
}
