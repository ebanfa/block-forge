package provider_test

/* import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"

	provider "github.com/edward1christian/block-forge/nova/pkg"
	"github.com/edward1christian/block-forge/nova/pkg/config"
)

// TestInit_Success tests the initialization function with success.
func TestInit_Success(t *testing.T) {
	// Create a temporary directory for the dummy config file
	tempDir := t.TempDir()

	// Create a dummy blockchain configuration
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

	// Write the dummy blockchain configuration to a file
	blockchainConfigFile := filepath.Join(tempDir, "valid_config.json")
	writeJSONFile(t, blockchainConfigFile, blockchainConfig)

	// Prepare options with the path to the dummy config file
	_ = &provider.InitOptions{
		Debug:          true,
		Verbose:        true,
		ConfigFilePath: blockchainConfigFile,
	}

	// Initialize
	//provider.Init(options)
}

// TestInit_NoConfigFile_Error tests initialization with no config file, expecting an error.
func TestInit_NoConfigFile_Error(t *testing.T) {
	// Prepare options with no config file
	options := &provider.InitOptions{
		Debug:          true,
		Verbose:        true,
		ConfigFilePath: "", // Empty config file path
	}

	// Expect a panic with an error
	assert.PanicsWithError(t, "failed to load custom configuration: open : no such file or directory", func() {
		provider.Init(options)
	})
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

// TestProvideConfiguration_Success tests providing configuration with success.
func TestProvideConfiguration_Success(t *testing.T) {
	// Prepare options
	options := &provider.InitOptions{
		Debug:          true,
		Verbose:        true,
		ConfigFilePath: "valid_config.yaml",
	}

	// Provide configuration
	provideConfigFunc := provider.ProvideConfiguration(options)
	config, err := provideConfigFunc()

	// Assert no error and configuration is not nil
	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.True(t, config.Debug)
	assert.True(t, config.Verbose)
}

// TestProvideConfiguration_NoConfigFile_Error tests providing configuration with no config file, expecting an error.
func TestProvideConfiguration_NoConfigFile_Error(t *testing.T) {
	// Prepare options with no config file
	options := &provider.InitOptions{
		Debug:          true,
		Verbose:        true,
		ConfigFilePath: "", // Empty config file path
	}

	// Provide configuration
	provideConfigFunc := provider.ProvideConfiguration(options)
	config, err := provideConfigFunc()

	// Assert error and configuration is nil
	assert.Error(t, err)
	assert.Nil(t, config)
	assert.EqualError(t, err, "failed to load custom configuration: open : no such file or directory")
}

// TestProvideEventBus_Success tests providing an event bus with success.
func TestProvideEventBus_Success(t *testing.T) {
	// Provide event bus
	eventBus := provider.ProvideEventBus()

	// Assert event bus is not nil
	assert.NotNil(t, eventBus)
}

// TestProvideLogger_Success tests providing a logger with success.
func TestProvideLogger_Success(t *testing.T) {
	// Provide logger
	logger := provider.ProvideLogger()

	// Assert logger is not nil
	assert.NotNil(t, logger)
}

// TestProvidPluginManager_Success tests providing a plugin manager with success.
func TestProvidPluginManager_Success(t *testing.T) {
	// Provide plugin manager
	pluginManager := provider.ProvidPluginManager()

	// Assert plugin manager is not nil
	assert.NotNil(t, pluginManager)
}

// TestProvidComponentRegistrar_Success tests providing a component registrar with success.
func TestProvidComponentRegistrar_Success(t *testing.T) {
	// Provide component registrar
	componentRegistrar := provider.ProvidComponentRegistrar()

	// Assert component registrar is not nil
	assert.NotNil(t, componentRegistrar)
}

// TestProvideSystem_Success tests providing a system with success.
func TestProvideSystem_Success(t *testing.T) {
	/* // Create mock lifecycle
	lc := &fxtest.Lifecycle{}
	logger := logger.NewLogrusLogger()
	eventBus := event.NewSystemEventBus()
	configuration := &configApi.Configuration{}
	pluginManager := system.NewPluginManager()
	registrar := component.NewComponentRegistrar()

	// Provide system
	sys := provider.ProvideSystem(lc, logger, eventBus, configuration, pluginManager, registrar)

	// Assert system is not nil
	assert.NotNil(t, sys)
}

// MockLifecycle is a mock implementation of the fx.Lifecycle interface for testing purposes.
type MockLifecycle struct{}

func (m *MockLifecycle) Append(fx.Hook) {} */
