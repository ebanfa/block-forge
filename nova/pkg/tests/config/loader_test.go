package config_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/edward1christian/block-forge/nova/pkg/config"
	"github.com/stretchr/testify/assert"
)

// TestReadBlockchainConfig tests the ReadBlockchainConfig function.
func TestReadBlockchainConfig(t *testing.T) {
	// Create a temporary test directory
	tempDir := t.TempDir()

	// Create a test module directory and configuration file
	moduleDir := filepath.Join(tempDir, "test-module")
	err := os.Mkdir(moduleDir, 0755)
	assert.NoError(t, err)

	moduleConfig := config.ModuleConfig{
		Name:             "test-module",
		Version:          "1.0",
		Dependencies:     []string{"dependency1", "dependency2"},
		EntityConfigDir:  "entities",
		MessageConfigDir: "messages",
		QueryConfigDir:   "queries",
	}
	moduleConfigFile := filepath.Join(moduleDir, "module.json")
	writeJSONFile(t, moduleConfigFile, moduleConfig)

	// Create a test blockchain configuration file
	blockchainConfig := config.BlockchainConfig{
		Name: "test-chain",
		Modules: []config.ModuleConfig{
			{
				Name:             "test-module",
				Version:          "",  // Should be empty string in the expected struct
				Dependencies:     nil, // Should be nil in the expected struct
				EntityConfigDir:  "entities",
				MessageConfigDir: "messages",
				QueryConfigDir:   "queries",
			},
		},
	}
	blockchainConfigFile := filepath.Join(tempDir, "blockchain.json")
	writeJSONFile(t, blockchainConfigFile, blockchainConfig)

	// Create directories for entities, messages, and queries
	entityDir, messageDir, queryDir := createTestDirs(moduleDir, t)

	// Create test entity, message, and query configuration files
	createTestConfigs(entityDir, t, messageDir, queryDir)

	// Read the blockchain configuration
	readConfig, err := config.ReadBlockchainConfig(blockchainConfigFile)

	// Verify no error occurred and the read configuration matches the expected configuration
	assert.NoError(t, err)
	assert.Equal(t, blockchainConfig.Name, readConfig.Name)
}

// TestReadModuleConfig tests the ReadModuleConfig function.
func TestReadModuleConfig(t *testing.T) {
	// Create a temporary test directory
	tempDir := t.TempDir()

	// Create a test module configuration file
	moduleConfig := config.ModuleConfig{
		Name:             "test-module",
		Version:          "1.0",
		Dependencies:     []string{"dependency1", "dependency2"},
		EntityConfigDir:  "entities", // Define directories for entities, messages, and queries
		MessageConfigDir: "messages",
		QueryConfigDir:   "queries",
	}
	moduleConfigFile := filepath.Join(tempDir, "module.json")
	writeJSONFile(t, moduleConfigFile, moduleConfig)

	// Create directories for entities, messages, and queries
	entityDir, messageDir, queryDir := createTestDirs(tempDir, t)

	// Create test entity, message, and query configuration files
	createTestConfigs(entityDir, t, messageDir, queryDir)

	// Read the module configuration
	readConfig, err := config.ReadModuleConfig(moduleConfigFile)

	// Verify no error occurred and the read configuration matches the expected configuration
	assert.NoError(t, err)
	assert.Equal(t, moduleConfig.Name, readConfig.Name)
}

func createTestConfigs(entityDir string, t *testing.T, messageDir string, queryDir string) {
	entityConfig := config.EntityConfig{Name: "test-entity"}
	entityConfigFile := filepath.Join(entityDir, "entity.json")
	writeJSONFile(t, entityConfigFile, entityConfig)

	messageConfig := config.MessageConfig{Name: "test-message"}
	messageConfigFile := filepath.Join(messageDir, "message.json")
	writeJSONFile(t, messageConfigFile, messageConfig)

	queryConfig := config.QueryConfig{Name: "test-query"}
	queryConfigFile := filepath.Join(queryDir, "query.json")
	writeJSONFile(t, queryConfigFile, queryConfig)
}

func createTestDirs(tempDir string, t *testing.T) (string, string, string) {
	entityDir := filepath.Join(tempDir, "entities")
	messageDir := filepath.Join(tempDir, "messages")
	queryDir := filepath.Join(tempDir, "queries")
	err := os.MkdirAll(entityDir, 0755)
	assert.NoError(t, err)
	err = os.MkdirAll(messageDir, 0755)
	assert.NoError(t, err)
	err = os.MkdirAll(queryDir, 0755)
	assert.NoError(t, err)
	return entityDir, messageDir, queryDir
}

// writeJSONFile writes a JSON object to a file.
func writeJSONFile(t *testing.T, filename string, obj interface{}) {
	file, err := os.Create(filename)
	assert.NoError(t, err)
	defer file.Close()
	fmt.Printf("Writing to file: %s\n", filename)
	encoder := json.NewEncoder(file)
	assert.NoError(t, encoder.Encode(obj))
}

// TestReadEntityConfig tests the ReadEntityConfig function.
func TestReadEntityConfig(t *testing.T) {
	// Create a temporary test directory
	tempDir := t.TempDir()

	// Create a test entity configuration file
	entityConfig := config.EntityConfig{
		Name: "test-entity",
		// Add other fields as needed
	}
	entityConfigFile := filepath.Join(tempDir, "entity.json")
	writeJSONFile(t, entityConfigFile, entityConfig)

	// Read the entity configuration
	readConfig, err := config.ReadEntityConfig(entityConfigFile)

	// Verify no error occurred and the read configuration matches the expected configuration
	assert.NoError(t, err)
	assert.Equal(t, entityConfig, readConfig)
}

// TestReadMessageConfig tests the ReadMessageConfig function.
func TestReadMessageConfig(t *testing.T) {
	// Create a temporary test directory
	tempDir := t.TempDir()

	// Create a test message configuration file
	messageConfig := config.MessageConfig{
		Name: "test-message",
		// Add other fields as needed
	}
	messageConfigFile := filepath.Join(tempDir, "message.json")
	writeJSONFile(t, messageConfigFile, messageConfig)

	// Read the message configuration
	readConfig, err := config.ReadMessageConfig(messageConfigFile)

	// Verify no error occurred and the read configuration matches the expected configuration
	assert.NoError(t, err)
	assert.Equal(t, messageConfig, readConfig)
}

// TestReadQueryConfig tests the ReadQueryConfig function.
func TestReadQueryConfig(t *testing.T) {
	// Create a temporary test directory
	tempDir := t.TempDir()

	// Create a test query configuration file
	queryConfig := config.QueryConfig{
		Name: "test-query",
		// Add other fields as needed
	}
	queryConfigFile := filepath.Join(tempDir, "query.json")
	writeJSONFile(t, queryConfigFile, queryConfig)

	// Read the query configuration
	readConfig, err := config.ReadQueryConfig(queryConfigFile)

	// Verify no error occurred and the read configuration matches the expected configuration
	assert.NoError(t, err)
	assert.Equal(t, queryConfig, readConfig)
}
