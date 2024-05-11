package commands

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/edward1christian/block-forge/nova/pkg/components/operations/commands"
	novaConfigApi "github.com/edward1christian/block-forge/nova/pkg/config"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/component"
	"github.com/edward1christian/block-forge/pkg/application/config"
	"github.com/edward1christian/block-forge/pkg/application/mocks"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestNewCreateConfigurationOp tests the NewCreateConfigurationOp function.
func TestNewCreateConfigurationOp(t *testing.T) {
	// Arrange
	id := "test-id"
	name := "Test Config Op"
	description := "Test description"

	// Act
	op := commands.NewCreateConfigurationOp(id, name, description)

	// Assert
	assert.NotNil(t, op)
	assert.Equal(t, id, op.Id)
	assert.Equal(t, name, op.Nm)
	assert.Equal(t, description, op.Desc)
	assert.Equal(t, component.BasicComponentType, op.Type())
}

// TestCreateConfigurationOp_Execute_Success tests the successful execution of the CreateConfigurationOp.
func TestCreateConfigurationOp_Execute_Success(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockSystem := &mocks.MockSystem{}
	mockStore := &mocks.MockStore{}
	mockMultiStore := &mocks.MockMultiStore{}

	mockConfig := &config.Configuration{
		CustomConfig: novaConfigApi.NovaConfig{
			UserHomeDir:      os.TempDir(),
			MultiStoreDbName: novaConfigApi.MultiStoreDbName,
		},
	}

	input := &system.SystemOperationInput{Data: "Test Project"}
	op := commands.NewCreateConfigurationOp("test-id", "Test Config Op", "Test description")

	mockStore.On("Load").Return(int64(1), nil)
	mockMultiStore.On("Load").Return(int64(1), nil)

	mockStore.On("SaveVersion").Return([]byte{}, int64(1), nil)
	mockMultiStore.On("SaveVersion").Return([]byte{}, int64(1), nil)

	mockStore.On("Set", mock.Anything, mock.Anything).Return(nil)
	mockMultiStore.On("CreateStore", mock.Anything).Return(mockStore, true, nil)

	mockSystem.On("Configuration").Return(mockConfig)
	mockSystem.On("MultiStore").Return(mockMultiStore)

	op.Initialize(ctx, mockSystem)

	// Act
	output, err := op.Execute(ctx, input)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, output)
}

// TestCreateConfigurationOp_Execute_InvalidInputData tests the case when the input data format is invalid.
func TestCreateConfigurationOp_Execute_InvalidInputData(t *testing.T) {
	// Arrange
	ctx := context.Background()
	mockSystem := &mocks.MockSystem{}
	mockConfig := &config.Configuration{
		CustomConfig: novaConfigApi.NovaConfig{
			UserHomeDir:      os.TempDir(),
			MultiStoreDbName: novaConfigApi.MultiStoreDbName,
		},
	}
	input := &system.SystemOperationInput{Data: 123}
	op := commands.NewCreateConfigurationOp("test-id", "Test Config Op", "Test description")

	mockSystem.On("Configuration").Return(mockConfig)

	op.Initialize(ctx, mockSystem)

	// Act
	output, err := op.Execute(ctx, input)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, output)
	assert.EqualError(t, err, "failed to create project. Invalid input data")
}

// ClearTempDir removes all files and subdirectories within the temporary directory.
func ClearTempDir() error {
	tempDir := os.TempDir()
	entries, err := ioutil.ReadDir(tempDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		err = os.RemoveAll(filepath.Join(tempDir, entry.Name()))
		if err != nil {
			return err
		}
	}
	return nil
}
