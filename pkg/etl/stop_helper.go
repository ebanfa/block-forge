package etl

import (
	"github.com/edward1christian/block-forge/pkg/application/context"
)

// ETLProcessStopHelper provides helper functions for managing ETL processes.
type ETLProcessStopHelper struct {
}

// NewETLHelper creates a new instance of ETLProcessStopHelper.
func NewETLProcessStopHelper() *ETLProcessStopHelper {
	return &ETLProcessStopHelper{}
}

// StopETLProcess stops an ETL process with the given ID.
func (h *ETLProcessStopHelper) StopETLProcess(ctx *context.Context, process *ETLProcess) error {

	// Stop components in reverse order
	for i := len(process.Components) - 1; i >= 0; i-- {
		componentName := process.Config.Components[i].Name
		if component, ok := process.Components[componentName]; ok {
			if err := component.Stop(ctx); err != nil {
				return err
			}
		}
	}

	process.Status = ETLStatusStopped
	return nil
}
