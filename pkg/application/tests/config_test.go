package tests

import (
	"testing"

	"github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfigurationFromFile_Success(t *testing.T) {
	// Create a temporary test file with valid JSON data
	// Replace "/path/to/valid/config.json" with the path to your test JSON file
	testFilePath := "test_config.json"

	// Define a struct to unmarshal the JSON data into
	var targetConfig config.ApplicationConfig

	// Load configuration from the test file
	err := config.LoadConfigurationFromFile(testFilePath, &targetConfig)

	// Assert that there's no error
	assert.NoError(t, err, "loading configuration should not return an error")

	// Assert that the targetConfig struct is populated correctly
	assert.NotNil(t, targetConfig, "targetConfig should not be nil")
	// Add more assertions here based on your specific struct fields
}

func TestLoadConfigurationFromFile_Failure(t *testing.T) {
	// Create a temporary test file with invalid JSON data or non-existent file
	// Replace "/path/to/invalid/config.json" with the path to your test JSON file
	testFilePath := "/path/to/invalid/config.json"

	// Define a struct to unmarshal the JSON data into
	var targetConfig config.ApplicationConfig

	// Load configuration from the test file
	err := config.LoadConfigurationFromFile(testFilePath, &targetConfig)

	// Assert that an error occurred
	assert.Error(t, err, "loading configuration should return an error")
}
