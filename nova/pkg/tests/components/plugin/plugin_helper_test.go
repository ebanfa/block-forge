package plugin

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/edward1christian/block-forge/nova/pkg/common"
	"github.com/edward1christian/block-forge/nova/pkg/components/plugin"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
)

func TestCreateAndStartBuilderService_Success(t *testing.T) {
	// Setup
	ctx := context.Background()
	mockSystem := new(mocks.MockSystem)
	mockRegistrar := new(mocks.MockComponentRegistrar)
	mockBuildService := new(mocks.MockSystemService)

	mockBuildService.On("Start", ctx).Return(nil)
	mockSystem.On("Logger").Return(&mocks.MockLogger{})
	mockSystem.On("ComponentRegistry").Return(mockRegistrar)
	mockBuildService.On("ID").Return(common.IgniteBuildService)
	mockBuildService.On("Initialize", ctx, mockSystem).Return(nil)
	mockRegistrar.On("CreateComponent", ctx, mock.Anything).Return(mockBuildService, nil)

	// Execute
	err := plugin.StartBuildService(ctx, mockSystem)

	// Verify
	assert.NoError(t, err)
	mockRegistrar.AssertExpectations(t)
	mockBuildService.AssertExpectations(t)
}

func TestCreateAndStartBuilderService_Failure_CreateComponent(t *testing.T) {
	// Setup
	ctx := context.Background()
	expectedErr := errors.New("failed to create component")

	mockSystem := new(mocks.MockSystem)
	mockRegistrar := new(mocks.MockComponentRegistrar)
	mockBuildService := new(mocks.MockSystemService)

	mockBuildService.On("Start", ctx).Return(nil)
	mockSystem.On("Logger").Return(&mocks.MockLogger{})
	mockSystem.On("ComponentRegistry").Return(mockRegistrar)
	mockBuildService.On("ID").Return(common.IgniteBuildService)
	mockBuildService.On("Initialize", ctx, mockSystem).Return(nil)
	mockRegistrar.On("CreateComponent", ctx, mock.Anything).Return(mockBuildService, expectedErr)

	// Execute
	err := plugin.StartBuildService(ctx, mockSystem)

	// Verify
	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedErr.Error())
	mockRegistrar.AssertExpectations(t)
}

func TestCreateAndStartBuilderService_Failure_Start(t *testing.T) {
	// Setup
	ctx := context.Background()
	expectedErr := errors.New("failed to start service")

	mockSystem := new(mocks.MockSystem)
	mockRegistrar := new(mocks.MockComponentRegistrar)
	mockBuildService := new(mocks.MockSystemService)

	mockBuildService.On("Start", ctx).Return(expectedErr)
	mockSystem.On("Logger").Return(&mocks.MockLogger{})
	mockSystem.On("ComponentRegistry").Return(mockRegistrar)
	mockBuildService.On("ID").Return(common.IgniteBuildService)
	mockBuildService.On("Initialize", ctx, mockSystem).Return(nil)
	mockRegistrar.On("CreateComponent", ctx, mock.Anything).Return(mockBuildService, nil)

	// Execute
	err := plugin.StartBuildService(ctx, mockSystem)

	// Verify
	assert.Error(t, err)
	assert.Contains(t, err.Error(), expectedErr.Error())
	mockRegistrar.AssertExpectations(t)
	mockBuildService.AssertExpectations(t)
}
