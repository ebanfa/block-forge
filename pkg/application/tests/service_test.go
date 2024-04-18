package tests

import (
	"testing"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/components"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/stretchr/testify/assert"
)

// TestNewBaseSystemService tests the NewBaseSystemService function.
func TestNewBaseSystemService(t *testing.T) {
	// Call the NewBaseSystemService function to create a new BaseSystemService instance
	service := systemApi.NewBaseSystemService("1", "Service1", "Description1")

	// Check if the instance is not nil
	assert.NotNil(t, service)

	// Check if the ID of the created instance matches the expected value
	assert.Equal(t, "1", service.ID())

	// Check if the Name of the created instance matches the expected value
	assert.Equal(t, "Service1", service.Name())

	// Check if the Description of the created instance matches the expected value
	assert.Equal(t, "Description1", service.Description())
}

// TestBaseSystemService_Type tests the Type method of BaseSystemService.
func TestBaseSystemService_Type(t *testing.T) {
	// Create a new BaseSystemService instance
	service := systemApi.NewBaseSystemService("1", "Service1", "Description1")

	// Call the Type method to get the component type
	componentType := service.Type()

	// Check if the returned component type matches the expected value
	assert.Equal(t, components.ServiceType, componentType)
}

// TestBaseSystemService_ImplementingInterface tests if BaseSystemService implements the StartableInterface.
func TestBaseSystemService_ImplementingInterface(t *testing.T) {
	// Ensure that BaseSystemService implements the StartableInterface
	var _ components.StartableInterface = (*systemApi.BaseSystemService)(nil)
}

// TestBaseSystemService_Initialize tests the Initialize method of BaseSystemService.
func TestBaseSystemService_Initialize(t *testing.T) {
	// Create a new BaseSystemService instance
	service := systemApi.NewBaseSystemService("1", "Service1", "Description1")

	// Create a mock context
	mockContext := &context.Context{}

	// Call the Initialize method
	err := service.Initialize(mockContext, nil)

	// Check if an error is returned
	assert.NoError(t, err)
}

// TestBaseSystemService_Start tests the Start method of BaseSystemService.
func TestBaseSystemService_Start(t *testing.T) {
	// Create a new BaseSystemService instance
	service := systemApi.NewBaseSystemService("1", "Service1", "Description1")

	// Create a mock context
	mockContext := &context.Context{}

	// Call the Start method
	err := service.Start(mockContext)

	// Check if an error is returned
	assert.Error(t, err)

	// Check if the error message matches the expected value
	assert.EqualError(t, err, "service not implemented")
}

// TestBaseSystemService_Stop tests the Stop method of BaseSystemService.
func TestBaseSystemService_Stop(t *testing.T) {
	// Create a new BaseSystemService instance
	service := systemApi.NewBaseSystemService("1", "Service1", "Description1")

	// Create a mock context
	mockContext := &context.Context{}

	// Call the Stop method
	err := service.Stop(mockContext)

	// Check if an error is returned
	assert.Error(t, err)

	// Check if the error message matches the expected value
	assert.EqualError(t, err, "service not implemented")
}
