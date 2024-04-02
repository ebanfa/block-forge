package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edward1christian/block-forge/pkg/application"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

func TestInit(t *testing.T) {
	// Test initialization of the Fx application
	// Note: It's difficult to test the Init function directly as it creates an Fx application.
	// Therefore, manual testing or integration tests may be necessary.
	// Add integration tests as needed...
}

func TestProvideEventBus(t *testing.T) {
	// Test ProvideEventBus function
	bus := application.ProvideEventBus()
	assert.NotNil(t, bus)
}

func TestProvideModuleManager(t *testing.T) {
	// Test ProvideModuleManager function
	manager := application.ProvideModuleManager("MMID", "ModuleManager", "Module Manager")
	assert.NotNil(t, manager)
}

func TestProvideSystem(t *testing.T) {
	// Test ProvideSystem function
	// Mock dependencies
	config := system.Configuration{}
	mockLogger := new(mocks.MockLogger)
	mockEventBus := new(mocks.MockEventBus)

	// Test ProvideSystem function
	sys := application.ProvideSystem(mockEventBus, mockLogger, config)
	assert.NotNil(t, sys)
	// Add more assertions as needed...
}

func TestProvideConfiguration(t *testing.T) {
	// Test ProvideConfiguration function
	t.Run("WithValidFiles", func(t *testing.T) {
		// Test ProvideConfiguration function with valid files
		provideFn := application.ProvideConfiguration("test_custom_config.json", "test_config.json")
		config, err := provideFn()
		assert.NoError(t, err)
		assert.NotNil(t, config)
		// Add more assertions as needed...
	})

	t.Run("WithInvalidFiles", func(t *testing.T) {
		// Test ProvideConfiguration function with invalid files
		provideFn := application.ProvideConfiguration("invalid_file.json", "invalid_file.json")
		config, err := provideFn()
		assert.Error(t, err)
		assert.Nil(t, config)
		// Add more assertions as needed...
	})
}

func TestProvideLogger(t *testing.T) {
	// Test ProvideLogger function
	logger := application.ProvideLogger()
	assert.NotNil(t, logger)
	// Add more assertions as needed...
}

func TestProvideApplication(t *testing.T) {
	// Test ProvideApplication function
	// Mock dependencies
	mockModuleManager := new(mocks.MockModuleManager)
	mockSystem := new(mocks.MockSystem)

	// Test ProvideApplication function
	app := application.ProvideApplication("appl", "appl", "appl", mockModuleManager, mockSystem)
	assert.NotNil(t, app)
	// Add more assertions as needed...
}
