package tests

import (
	goContext "context"
	"testing"
	"time"

	"github.com/edward1christian/block-forge/pkg/application/context"
	"github.com/stretchr/testify/assert"
)

func TestContext_WithValue(t *testing.T) {
	ctx := context.Background()

	// Add a key-value pair
	key := "key"
	value := "value"
	newCtx := ctx.WithValue(key, value)

	// Check if the value is correctly set
	assert.Equal(t, value, newCtx.Value(key))
}

func TestContext_WithPluginPaths(t *testing.T) {
	ctx := context.Background()

	// Add plugin paths
	paths := []string{"path1", "path2"}
	newCtx := ctx.WithPluginPaths(paths...)

	// Check if plugin paths are correctly set
	assert.Equal(t, paths, newCtx.PluginPaths)
}

func TestContext_WithRemotePluginPaths(t *testing.T) {
	ctx := context.Background()

	// Add plugin paths
	paths := []string{"path1", "path2"}
	newCtx := ctx.WithRemotePluginLocations(paths...)

	// Check if plugin paths are correctly set
	assert.Equal(t, paths, newCtx.RemotePluginLocations)
}

func TestContext_WithTraceID(t *testing.T) {
	ctx := context.Background()

	// Add a trace ID
	traceID := "12345"
	newCtx := ctx.WithTraceID(traceID)

	// Check if trace ID is correctly set
	assert.Equal(t, traceID, newCtx.Value("traceID"))
}

func TestBackground(t *testing.T) {
	// Test if interfaces.Background function returns a non-nil context
	assert.NotNil(t, context.Background())
}

func TestWithContext(t *testing.T) {
	// Test if WithContext function returns a non-nil context
	assert.NotNil(t, context.WithContext(context.Background()))
}

func TestWithTimeout(t *testing.T) {
	parentCtx := context.Background()
	timeoutDuration := 100 * time.Millisecond

	// Test if WithTimeout function returns a new context with timeout
	ctx, cancel := context.WithTimeout(parentCtx, timeoutDuration)
	defer cancel()

	// Check if the returned context has a timeout
	select {
	case <-ctx.Done():
		assert.Equal(t, goContext.DeadlineExceeded, ctx.Err())
	case <-time.After(timeoutDuration + 10*time.Millisecond):
		t.Error("Timeout not triggered")
	}
}
