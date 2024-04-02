package etl

import (
	"github.com/edward1christian/block-forge/pkg/application/context"
)

// ETLProcessStartHelper provides helper functions for managing ETL processes.
type ETLProcessStartHelper struct {
}

// NewETLHelper creates a new instance of ETLProcessStartHelper.
func NewETLProcessStartHelper(system *ETLSystemImpl) *ETLProcessStartHelper {
	return &ETLProcessStartHelper{}
}

// StartETLProcess starts an ETL process with the given ID.
func (h *ETLProcessStartHelper) StartETLProcess(ctx *context.Context, process *ETLProcess) error {

	// Start component
	for _, component := range process.Components {
		if err := component.Start(ctx); err != nil {
			// Rollback: stop already started component
			for _, startedAdapter := range process.Components {
				_ = startedAdapter.Stop(ctx)
			}
			return err
		}
	}

	process.Status = ETLStatusRunning
	return nil
}
