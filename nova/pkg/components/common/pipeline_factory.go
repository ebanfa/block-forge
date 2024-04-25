package common

import (
	"github.com/edward1christian/block-forge/nova/pkg/common"
	"github.com/edward1christian/block-forge/nova/pkg/types"
	"github.com/edward1christian/block-forge/pkg/application/component"
	configApi "github.com/edward1christian/block-forge/pkg/application/config"
)

// PipelineFactory is responsible for creating instances of Pipeline.
type PipelineFactory struct {
}

// CreateComponent creates a new instance of the Pipeline.
func (bpf *PipelineFactory) CreateComponent(config *configApi.ComponentConfig) (component.ComponentInterface, error) {
	// Construct the service
	return NewPipelineBuilder(config).
		AddStage("IgniteStage").
		AddTask(types.NewTask(common.CreateWorkspaceTask)).
		Build(), nil
}
