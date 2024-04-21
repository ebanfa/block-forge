package builders

import "github.com/edward1christian/block-forge/nova/pkg/build"

func CosmosSDKBlockchainPipelineBuilder(name string) build.PipelineBuilderInterface {
	builder := build.NewPipelineBuilder(name)
	builder.AddStage("Test")
	return builder
}
