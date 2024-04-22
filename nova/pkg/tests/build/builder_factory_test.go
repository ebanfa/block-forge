package build

import (
	"testing"

	"github.com/edward1christian/block-forge/nova/pkg/build"
	"github.com/stretchr/testify/assert"
)

// TestPipelineBuilderFactory_CreatePipelineBuilder tests the functionality of creating pipeline builders
// for different pipeline types using the PipelineBuilderFactory.
func TestPipelineBuilderFactory_CreatePipelineBuilder(t *testing.T) {
	// Create a new instance of PipelineBuilderFactory
	factory := build.NewPipelineBuilderFactory()

	// Register builder creation functions for different pipeline types
	factory.RegisterPipelineBuilderFactory("type1", build.NewPipelineBuilder)
	factory.RegisterPipelineBuilderFactory("type2", build.NewPipelineBuilder)

	// Test creating pipeline builders for different types
	builder1, err := factory.CreatePipelineBuilder("Pipeline1", "type1")
	assert.NoError(t, err)
	assert.NotNil(t, builder1)

	builder2, err := factory.CreatePipelineBuilder("Pipeline2", "type2")
	assert.NoError(t, err)
	assert.NotNil(t, builder2)

}

// TestPipelineBuilderFactory_CreatePipelineBuilder_UnsupportedType tests the behavior when attempting
// to create a pipeline builder with an unsupported type.
func TestPipelineBuilderFactory_CreatePipelineBuilder_UnsupportedType(t *testing.T) {
	// Create a new instance of PipelineBuilderFactory
	factory := build.NewPipelineBuilderFactory()

	// Attempt to create pipeline builder with unsupported type
	builder, err := factory.CreatePipelineBuilder("Pipeline", "unsupported")
	assert.Nil(t, builder)
	assert.Error(t, err)
}
