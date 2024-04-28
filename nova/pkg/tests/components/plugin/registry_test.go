package plugin

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/edward1christian/block-forge/nova/pkg/components/plugin"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
)

func TestRegisterComponents_Success(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	mockContext := &context.Context{} // Mock context
	mockSystem := &mocks.MockSystem{} // Mock system
	mockComponent := &mocks.MockSystemService{}

	mockRegistrar := &mocks.MockComponentRegistrar{}

	mockSystem.On("ComponentRegistry").Return(mockRegistrar)
	mockRegistrar.On("RegisterFactory", ctx, mock.Anything, mock.Anything).Return(nil)
	mockRegistrar.On("CreateComponent", ctx, mock.Anything).Return(mockComponent, nil)

	// Act
	err := plugin.RegisterComponents(mockContext, mockSystem)

	// Assert
	assert.NoError(t, err, "RegisterComponents should not return an error")
}

func TestRegisterComponents_Error(t *testing.T) {
	// Arrange
	ctx := &context.Context{}
	mockContext := &context.Context{} // Mock context
	mockSystem := &mocks.MockSystem{} // Mock system
	mockComponent := &mocks.MockSystemService{}

	mockRegistrar := &mocks.MockComponentRegistrar{}

	mockSystem.On("ComponentRegistry").Return(mockRegistrar)
	mockRegistrar.On("RegisterFactory", ctx, mock.Anything, mock.Anything).Return(errors.New("Error"))
	mockRegistrar.On("CreateComponent", ctx, mock.Anything).Return(mockComponent, nil)

	// Act
	err := plugin.RegisterComponents(mockContext, mockSystem)

	// Assert
	assert.Error(t, err, "RegisterComponents should return an error")
}
