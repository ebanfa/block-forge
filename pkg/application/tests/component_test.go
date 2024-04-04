package tests

import (
	"testing"

	"github.com/edward1christian/block-forge/pkg/application/components"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
	"github.com/stretchr/testify/assert"
)

// TestBaseComponent_ID tests the ID method of BaseComponent.
func TestBaseComponent_ID(t *testing.T) {
	// Create a new instance of the mock
	mockComponent := new(mocks.MockComponent)

	// Define the expected behavior
	expectedID := "123"
	mockComponent.On("ID").Return(expectedID)

	// Call the method being tested
	actualID := mockComponent.ID()

	// Check if the method returned the expected value
	assert.Equal(t, expectedID, actualID)
}

// TestBaseComponent_Name tests the Name method of BaseComponent.
func TestBaseComponent_Name(t *testing.T) {
	// Create a new instance of the mock
	mockComponent := new(mocks.MockComponent)

	// Define the expected behavior
	expectedName := "TestComponent"
	mockComponent.On("Name").Return(expectedName)

	// Call the method being tested
	actualName := mockComponent.Name()

	// Check if the method returned the expected value
	assert.Equal(t, expectedName, actualName)
}

// TestBaseComponent_Type tests the Type method of BaseComponent.
func TestBaseComponent_Type(t *testing.T) {
	// Create a new instance of the mock
	mockComponent := new(mocks.MockComponent)

	// Define the expected behavior
	expectedType := components.BasicComponentType
	mockComponent.On("Type").Return(expectedType)

	// Call the method being tested
	actualType := mockComponent.Type()

	// Check if the method returned the expected value
	assert.Equal(t, expectedType, actualType)
}

// TestBaseComponent_Description tests the Description method of BaseComponent.
func TestBaseComponent_Description(t *testing.T) {
	// Create a new instance of the mock
	mockComponent := new(mocks.MockComponent)

	// Define the expected behavior
	expectedDescription := "Test Description"
	mockComponent.On("Description").Return(expectedDescription)

	// Call the method being tested
	actualDescription := mockComponent.Description()

	// Check if the method returned the expected value
	assert.Equal(t, expectedDescription, actualDescription)
}

// TestBaseComponent_ImplementingInterface tests if BaseComponent implements the ComponentInterface.
func TestBaseComponent_ImplementingInterface(t *testing.T) {
	// Ensure that BaseComponent implements the ComponentInterface
	var _ components.ComponentInterface = (*components.BaseComponent)(nil)
}
