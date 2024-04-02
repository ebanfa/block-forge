package etl

import (
	"github.com/edward1christian/block-forge/pkg/application/context"
)

// ETLProcessInitHelper provides helper functions for managing ETL processes.
type ETLProcessInitHelper struct {
	idGenerator ProcessIDGenerator
}

// NewETLHelper creates a new instance of ETLProcessInitHelper.
func NewETLProcessInitHelper(idGenerator ProcessIDGenerator) *ETLProcessInitHelper {
	return &ETLProcessInitHelper{
		idGenerator: idGenerator,
	}
}

// CreateETLProcess creates an ETL process with the provided configuration.
func (h *ETLProcessInitHelper) CreateETLProcess(ctx *context.Context, config *ETLConfig) (*ETLProcess, error) {
	// Generate the process ID using the injected generator
	processID, err := h.idGenerator.GenerateID()
	if err != nil {
		return nil, err
	}

	if h.isEmptyConfig(config) {
		return nil, ErrInvalidETLProcessConfig
	}
	// Create the process
	process := &ETLProcess{
		ID:         processID, // assuming generateProcessID() generates unique process ID
		Config:     config,
		Status:     ETLStatusInitialized,
		Components: make(map[string]ETLComponent),
	}

	return process, nil
}

// RollbackComponents stops initialized components.
func (h *ETLProcessInitHelper) RollbackComponents(ctx *context.Context, process *ETLProcess) {
	for _, component := range process.Components {
		_ = component.Stop(ctx)
	}
}

// Check if the ETLConfig is empty
func (h *ETLProcessInitHelper) isEmptyConfig(config *ETLConfig) bool {
	// Check if all fields are empty or zero-valued
	return len(config.Components) == 0
}
