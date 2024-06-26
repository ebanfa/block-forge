package system_test

import (
	"testing"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/component"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/stretchr/testify/assert"
)

// TestNewBaseSystemComponent tests the NewBaseSystemComponent function.
func TestNewBaseSystemComponent(t *testing.T) {
	// Call the NewBaseSystemComponent function to create a new BaseSystemComponent instance
	component := system.NewBaseSystemComponent("1", "Component1", "Description1")

	// Check if the instance is not nil
	assert.NotNil(t, component)

	// Check if the ID of the created instance matches the expected value
	assert.Equal(t, "1", component.ID())

	// Check if the Name of the created instance matches the expected value
	assert.Equal(t, "Component1", component.Name())

	// Check if the Description of the created instance matches the expected value
	assert.Equal(t, "Description1", component.Description())
}

// TestBaseSystemComponent_Type tests the Type method of BaseSystemComponent.
func TestBaseSystemComponent_Type(t *testing.T) {
	// Create a new BaseSystemComponent instance
	comp := system.NewBaseSystemComponent("1", "Component1", "Description1")

	// Call the Type method to get the component type
	componentType := comp.Type()

	// Check if the returned component type matches the expected value
	assert.Equal(t, component.SystemComponentType, componentType)
}

// TestBaseSystemComponent_ImplementingInterface tests if BaseSystemComponent implements the ComponentInterface.
func TestBaseSystemComponent_ImplementingInterface(t *testing.T) {
	// Ensure that BaseSystemComponent implements the ComponentInterface
	var _ component.ComponentInterface = (*system.BaseSystemComponent)(nil)
}

// TestBaseSystemComponent_Initialize tests the Initialize method of BaseSystemComponent.
func TestBaseSystemComponent_Initialize(t *testing.T) {
	// Create a new BaseSystemComponent instance
	component := system.NewBaseSystemComponent("1", "Component1", "Description1")

	// Create a mock context
	mockContext := &context.Context{}

	// Call the Initialize method
	err := component.Initialize(mockContext, nil)

	// Check if an error is returned
	assert.NoError(t, err)

}
