package utils

import "github.com/edward1christian/block-forge/pkg/etl"

// Check if the ETLConfig is empty
func IsEmptyConfig(config *etl.ETLConfig) bool {
	// Check if all fields are empty or zero-valued
	return len(config.Components) == 0
}
