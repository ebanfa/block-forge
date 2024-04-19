package build_test

import (
	"errors"
	"testing"

	"github.com/edward1christian/block-forge/nova/pkg/mocks"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestMockBuildTask_GetName tests the GetName method of MockBuildTask.
func TestMockBuildTask_GetName(t *testing.T) {
	// Arrange
	task := &mocks.MockBuildTask{}

	// Mock behavior
	task.On("GetName").Return("TestTask")

	// Act
	name := task.GetName()

	// Assert
	assert.Equal(t, name, "TestTask")
}

// TestMockBuildTask_Execute_Success tests the Execute method of MockBuildTask for success.
func TestMockBuildTask_Execute_Success(t *testing.T) {
	// Arrange
	task := &mocks.MockBuildTask{}

	// Mock behavior
	task.On("GetName").Return("TestTask")
	task.On("Execute", mock.Anything, mock.Anything).Return(&system.SystemOperationOutput{}, nil)

	// Act
	_, err := task.Execute(nil, nil)

	// Assert
	assert.NoError(t, err, "Execute should not return an error")
}

// TestMockBuildTask_Execute_Error tests the Execute method of MockBuildTask for an error case.
func TestMockBuildTask_Execute_Error(t *testing.T) {
	// Arrange
	task := &mocks.MockBuildTask{}
	expectedErr := errors.New("mock error")

	// Mock behavior
	task.On("GetName").Return("TestTask")
	task.On("Execute", mock.Anything, mock.Anything).Return(&system.SystemOperationOutput{}, expectedErr)

	// Act
	_, err := task.Execute(nil, nil)

	// Assert
	assert.Error(t, err, "Execute should return an error")
	assert.EqualError(t, err, expectedErr.Error(), "Error message should match")
}
