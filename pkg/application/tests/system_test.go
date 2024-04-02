package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

func TestSystem_EventBus(t *testing.T) {
	// Create a mock event bus
	config := config.Configuration{}
	eventBus := &mocks.MockEventBus{}

	// Create a new instance of the System
	sys := system.NewSystem(eventBus, nil, nil, config)

	// Test EventBus method
	assert.Equal(t, eventBus, sys.EventBus(), "EventBus should return the expected EventBusInterface")
}

func TestSystem_Operations(t *testing.T) {
	// Create a mock operations instance
	config := config.Configuration{}
	operations := &mocks.MockOperations{}

	// Create a new instance of the System
	sys := system.NewSystem(nil, operations, nil, config)

	// Test Operations method
	assert.Equal(t, operations, sys.Operations(), "Operations should return the expected Operations")
}

func TestSystem_Logger(t *testing.T) {
	// Create a mock logger instance
	config := config.Configuration{}
	logger := &mocks.MockLogger{}

	// Create a new instance of the System
	sys := system.NewSystem(nil, nil, logger, config)

	// Test Logger method
	assert.Equal(t, logger, sys.Logger(), "Logger should return the expected LoggerInterface")
}

func TestSystem_Configuration(t *testing.T) {
	// Create a mock configuration instance
	config := config.Configuration{}

	// Create a new instance of the System
	sys := system.NewSystem(nil, nil, nil, config)
	fmt.Println("H>>>>>>>>>>>")
	// Test Configuration method
	assert.Equal(t, config, sys.Configuration(), "Configuration should return the expected Configuration")
}
