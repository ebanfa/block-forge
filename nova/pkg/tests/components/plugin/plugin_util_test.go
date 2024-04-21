package plugin_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/edward1christian/block-forge/nova/pkg/components/plugin"
	configApi "github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
)

func TestRegisterAndCreateComponent_Success(t *testing.T) {
	// Mock component registrar
	mockRegistrar := &mocks.MockComponentRegistrar{}

	mockRegistrar.On("RegisterFactory", "test_factory", mock.Anything).Return(nil)
	mockRegistrar.On("CreateComponent", &configApi.ComponentConfig{
		ID:        "test_component",
		FactoryID: "test_factory",
	}).Return(&mocks.MockComponent{}, nil)

	// Prepare configuration
	config := &configApi.ComponentConfig{
		ID:        "test_component",
		FactoryID: "test_factory",
	}

	// Call the function
	err := plugin.RegisterAndCreateComponent(mockRegistrar, config)

	// Assert no error
	assert.NoError(t, err)
}

func TestRegisterAndCreateComponent_Failure_RegistrationError(t *testing.T) {
	// Mock component registrar with registration error
	mockRegistrar := &mocks.MockComponentRegistrar{}

	mockRegistrar.On("RegisterFactory", "test_factory", mock.Anything).Return(errors.New("failed to register factory"))

	// Prepare configuration
	config := &configApi.ComponentConfig{
		ID:        "test_component",
		FactoryID: "test_factory",
	}

	// Call the function
	err := plugin.RegisterAndCreateComponent(mockRegistrar, config)

	// Assert error with correct message
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to register factory")
}

func TestRegisterAndCreateComponent_Failure_CreationError(t *testing.T) {
	// Mock component registrar with creation error
	mockRegistrar := &mocks.MockComponentRegistrar{}

	mockRegistrar.On("RegisterFactory", "test_factory", mock.Anything).Return(nil)
	mockRegistrar.On("CreateComponent", &configApi.ComponentConfig{
		ID:        "test_component",
		FactoryID: "test_factory",
	}).Return(&mocks.MockComponent{}, errors.New("failed to create component"))

	// Prepare configuration
	config := &configApi.ComponentConfig{
		ID:        "test_component",
		FactoryID: "test_factory",
	}

	// Call the function
	err := plugin.RegisterAndCreateComponent(mockRegistrar, config)

	// Assert error with correct message
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create component")
}
