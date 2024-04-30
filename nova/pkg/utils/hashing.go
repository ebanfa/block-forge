package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// hashSHA256 calculates the SHA-256 hash of the input string and returns the hexadecimal representation.
func HashSHA256(input string) string {
	// Create a new SHA-256 hasher
	hasher := sha256.New()

	// Write the input string to the hasher
	hasher.Write([]byte(input))

	// Calculate the hash sum
	hashSum := hasher.Sum(nil)

	// Convert the hash sum to a hexadecimal string
	hashHex := hex.EncodeToString(hashSum)

	return hashHex
}
