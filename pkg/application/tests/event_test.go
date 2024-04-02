package tests

import (
	"fmt"
	"sync"
	"testing"

	"github.com/edward1christian/block-forge/pkg/application/event"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSystemEventBus_Subscribe tests the Subscribe method of the SystemEventBus.
func TestSystemEventBus_Subscribe(t *testing.T) {
	// Create a new instance of SystemEventBus
	eb := event.NewSystemEventBus()

	// Define a test event handler
	handler := func(event event.Event) {
		assert.Equal(t, "test_topic", event.Type, "Unexpected event type")
		assert.Equal(t, "test_data", event.Data, "Unexpected event data")
	}

	// Subscribe to the test topic
	params := event.BusSubscriptionParams{
		Topic:        "test_topic",
		EventHandler: handler,
	}
	err := eb.Subscribe(params)
	require.NoError(t, err, "Failed to subscribe")

	// Publish an event to the test topic
	eb.Publish(event.Event{Type: "test_topic", Data: "test_data"})
}

// TestSystemEventBus_SubscribeAsync tests the SubscribeAsync method of the SystemEventBus.
func TestSystemEventBus_SubscribeAsync(t *testing.T) {
	// Create a new instance of SystemEventBus
	eb := event.NewSystemEventBus()

	// Define a test event handler
	handler := func(event event.Event) {
		assert.Equal(t, "test_topic", event.Type, "Unexpected event type")
		assert.Equal(t, "test_data", event.Data, "Unexpected event data")
	}

	// Subscribe to the test topic asynchronously
	params := event.BusSubscriptionParams{
		Topic:        "test_topic",
		EventHandler: handler,
	}
	err := eb.SubscribeAsync(params, false)
	require.NoError(t, err, "Failed to subscribe asynchronously")

	// Publish an event to the test topic
	eb.Publish(event.Event{Type: "test_topic", Data: "test_data"})

	// Wait for the asynchronous handler to complete
	eb.WaitAsync()
}

// TestSystemEventBus_SubscribeOnce tests the SubscribeOnce method of the SystemEventBus.
func TestSystemEventBus_SubscribeOnce(t *testing.T) {
	// Create a new instance of SystemEventBus
	eb := event.NewSystemEventBus()

	// Define a test event handler
	handlerCalled := false
	handler := func(event event.Event) {
		assert.Equal(t, "test_topic", event.Type, "Unexpected event type")
		assert.Equal(t, "test_data", event.Data, "Unexpected event data")
		handlerCalled = true
	}

	// Subscribe to the test topic for a single event occurrence
	params := event.BusSubscriptionParams{
		Topic:        "test_topic",
		EventHandler: handler,
	}
	err := eb.SubscribeOnce(params)
	require.NoError(t, err, "Failed to subscribe once")

	// Publish an event to the test topic
	eb.Publish(event.Event{Type: "test_topic", Data: "test_data"})

	// Verify that the handler was called
	assert.True(t, handlerCalled, "Handler was not called")

	// Publish another event to the test topic (handler should not be called)
	handlerCalled = false
	eb.Publish(event.Event{Type: "test_topic", Data: "test_data"})
	assert.False(t, handlerCalled, "Handler was called again after unsubscribing")
}

// TestSystemEventBus_SubscribeOnceAsync tests the SubscribeOnceAsync method of the SystemEventBus.
func TestSystemEventBus_SubscribeOnceAsync(t *testing.T) {
	// Create a new instance of SystemEventBus
	eb := event.NewSystemEventBus()

	// Define a test event handler
	handlerCalled := false
	handler := func(event event.Event) {
		assert.Equal(t, "test_topic", event.Type, "Unexpected event type")
		assert.Equal(t, "test_data", event.Data, "Unexpected event data")
		handlerCalled = true
	}

	// Subscribe to the test topic for a single event occurrence asynchronously
	params := event.BusSubscriptionParams{
		Topic:        "test_topic",
		EventHandler: handler,
	}
	err := eb.SubscribeOnceAsync(params)
	require.NoError(t, err, "Failed to subscribe once asynchronously")

	// Publish an event to the test topic
	eb.Publish(event.Event{Type: "test_topic", Data: "test_data"})

	// Wait for the asynchronous handler to complete
	eb.WaitAsync()

	// Verify that the handler was called
	assert.True(t, handlerCalled, "Handler was not called")

	// Publish another event to the test topic (handler should not be called)
	handlerCalled = false
	eb.Publish(event.Event{Type: "test_topic", Data: "test_data"})
	eb.WaitAsync()
	assert.False(t, handlerCalled, "Handler was called again after unsubscribing")
}

// TestSystemEventBus_Unsubscribe tests the Unsubscribe method of the SystemEventBus.
func TestSystemEventBus_Unsubscribe(t *testing.T) {
	// Create a new instance of SystemEventBus
	eb := event.NewSystemEventBus()

	// Define a test event handler
	handlerCalled := false
	handler := func(event event.Event) {
		handlerCalled = true
	}

	// Subscribe to the test topic
	params := event.BusSubscriptionParams{
		Topic:        "test_topic",
		EventHandler: handler,
	}
	err := eb.Subscribe(params)
	require.NoError(t, err, "Failed to subscribe")

	// Unsubscribe from the test topic
	err = eb.Unsubscribe(params)
	require.NoError(t, err, "Failed to unsubscribe")

	// Wait for any pending events to be processed
	eb.WaitAsync()

	// Publish an event to the test topic (handler should not be called)
	handlerCalled = false
	eb.Publish(event.Event{Type: "test_topic", Data: "test_data"})

	// Wait for any pending events to be processed
	eb.WaitAsync()

	assert.False(t, handlerCalled, "Handler was called after unsubscribing")
}

// TestSystemEventBus_HasCallback tests the HasCallback method of the SystemEventBus.
func TestSystemEventBus_HasCallback(t *testing.T) {
	// Create a new instance of SystemEventBus
	eb := event.NewSystemEventBus()

	// Define a test event handler
	handler := func(event event.Event) {}

	// Subscribe to the test topic
	params := event.BusSubscriptionParams{
		Topic:        "test_topic",
		EventHandler: handler,
	}
	err := eb.Subscribe(params)
	require.NoError(t, err, "Failed to subscribe")

	// Check if a callback is registered for the test topic
	hasCallback := eb.HasCallback("test_topic")
	assert.True(t, hasCallback, "Expected to have a callback registered for the test topic")

	// Unsubscribe from the test topic
	err = eb.Unsubscribe(params)
	require.NoError(t, err, "Failed to unsubscribe")

	// Wait for any pending events to be processed
	eb.WaitAsync()

	// Check if a callback is still registered for the test topic
	hasCallback = eb.HasCallback("test_topic")
	assert.False(t, hasCallback, "Expected not to have a callback registered for the test topic")
}

// TestSystemEventBus_ConcurrentPublish tests concurrent publishing to the SystemEventBus.
func TestSystemEventBus_ConcurrentPublish(t *testing.T) {
	// Create a new instance of SystemEventBus
	eb := event.NewSystemEventBus()

	// Define a test event handler
	var counter int
	var mutex sync.Mutex
	handler := func(event event.Event) {
		mutex.Lock()
		defer mutex.Unlock()
		counter++
	}

	// Subscribe to the test topic
	params := event.BusSubscriptionParams{
		Topic:        "test_topic",
		EventHandler: handler,
	}
	err := eb.Subscribe(params)
	require.NoError(t, err, "Failed to subscribe")

	// Publish events concurrently
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			eb.Publish(event.Event{Type: "test_topic", Data: fmt.Sprintf("event_%d", idx)})
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Check if the counter is equal to the number of events published
	mutex.Lock()
	defer mutex.Unlock()
	assert.Equal(t, 1000, counter, "Expected %d events, but received %d", 1000, counter)
}
