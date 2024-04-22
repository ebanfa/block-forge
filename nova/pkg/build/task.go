package build

import (
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// TaskInterface represents a build task.
type TaskInterface interface {
	system.SystemOperationInterface
}
