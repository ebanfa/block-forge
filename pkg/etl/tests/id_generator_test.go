package tests

import (
	"testing"

	"github.com/edward1christian/block-forge/pkg/etl"
	"github.com/stretchr/testify/assert"
)

func TestGenerateID_Success(t *testing.T) {
	prefix := "test"
	gen := etl.NewProcessIDGenerator(prefix)

	_, err := gen.GenerateID()

	assert.NoError(t, err)
}

func TestGenerateID_EmptyPrefix(t *testing.T) {
	prefix := ""
	gen := etl.NewProcessIDGenerator(prefix)

	processID, err := gen.GenerateID()

	assert.EqualError(t, err, "prefix cannot be empty")
	assert.Empty(t, processID)
}
