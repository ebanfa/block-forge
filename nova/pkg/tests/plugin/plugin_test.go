package plugin

import (
	"testing"

	"github.com/edward1christian/block-forge/nova/pkg/components/plugin"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
	"github.com/stretchr/testify/assert"
)

func TestPluginInterface_Initialize_Success(t *testing.T) {
	// Arrange
	mockSystem := new(mocks.MockSystem)
	plugin := &plugin.NovaPlugin{}

	// Mock behavior
	mockSystem.On("SomeMethod").Return(nil)

	// Act
	err := plugin.Initialize(context.Background(), mockSystem)

	// Assert
	assert.NoError(t, err, "Initialize should not return an error")
}

func TestPluginInterface_RegisterResources_Success(t *testing.T) {
	// Arrange
	plugin := &plugin.NovaPlugin{}

	// Act
	err := plugin.RegisterResources(context.Background())

	// Assert
	assert.NoError(t, err, "RegisterResources should not return an error")
}

func TestPluginInterface_Start_Success(t *testing.T) {
	// Arrange
	plugin := &plugin.NovaPlugin{}

	// Act
	err := plugin.Start(context.Background())

	// Assert
	assert.NoError(t, err, "Start should not return an error")
}

func TestPluginInterface_Stop_Success(t *testing.T) {
	// Arrange
	plugin := &plugin.NovaPlugin{}

	// Act
	err := plugin.Stop(context.Background())

	// Assert
	assert.NoError(t, err, "Stop should not return an error")
}
