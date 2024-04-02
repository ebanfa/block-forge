package system

import (
	"github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/edward1christian/block-forge/pkg/application/event"
	"github.com/edward1christian/block-forge/pkg/application/logger"
)

// SystemImpl represents the core system in the application.
type SystemImpl struct {
	eventBus      event.EventBusInterface
	operations    Operations
	logger        logger.LoggerInterface
	configuration config.Configuration
}

// NewSystem creates a new instance of the SystemImpl.
func NewSystem(
	eventBus event.EventBusInterface,
	operations Operations,
	logger logger.LoggerInterface,
	configuration config.Configuration) *SystemImpl {
	return &SystemImpl{
		eventBus:      eventBus,
		operations:    operations,
		logger:        logger,
		configuration: configuration,
	}
}

// EventBus returns the event bus associated with the system.
func (s *SystemImpl) EventBus() event.EventBusInterface {
	return s.eventBus
}

// Operations returns the operations associated with the system.
func (s *SystemImpl) Operations() Operations {
	return s.operations
}

// Logger returns the logger associated with the system.
func (s *SystemImpl) Logger() logger.LoggerInterface {
	return s.logger
}

// Configuration returns the configuration associated with the system.
func (s *SystemImpl) Configuration() config.Configuration {
	return s.configuration
}
