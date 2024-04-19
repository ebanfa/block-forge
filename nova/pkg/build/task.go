package build

import (
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// BuildTaskInterface represents a build task.
type BuildTaskInterface interface {
	system.SystemOperationInterface

	// GetName returns the name of the build task.
	GetName() string
}
