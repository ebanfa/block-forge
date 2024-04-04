package event

import (
	"fmt"

	"github.com/asaskevich/EventBus"
)

// SystemEventBus is a concrete implementation of the EventBusInterface.
type SystemEventBus struct {
	bus      EventBus.Bus           // Underlying third-party EventBus instance
	handlers map[string]interface{} // Map to store function references used for subscription
}

// NewSystemEventBus creates a new instance of the SystemEventBus.
func NewSystemEventBus() EventBusInterface {
	return &SystemEventBus{
		bus:      EventBus.New(),
		handlers: make(map[string]interface{}),
	}
}

// Subscribe subscribes to an event topic with the given parameters.
func (eb *SystemEventBus) Subscribe(params BusSubscriptionParams) error {
	// Create a wrapper function that constructs the Event and calls the provided EventHandler
	handler := func(args ...interface{}) {
		event := Event{
			Type: params.Topic,
			Data: args[0],
		}
		params.EventHandler(event)
	}

	// Store the function reference for later use in Unsubscribe
	eb.handlers[params.Topic] = handler

	// Subscribe to the underlying EventBus using the wrapper function
	return eb.bus.Subscribe(params.Topic, handler)
}

// SubscribeAsync subscribes to an event topic asynchronously with the given parameters.
func (eb *SystemEventBus) SubscribeAsync(params BusSubscriptionParams, transactional bool) error {
	// Create a wrapper function that constructs the Event and calls the provided EventHandler
	handler := func(args ...interface{}) {
		event := Event{
			Type: params.Topic,
			Data: args[0],
		}
		params.EventHandler(event)
	}

	// Store the function reference for later use in Unsubscribe
	eb.handlers[params.Topic] = handler

	// Subscribe asynchronously to the underlying EventBus using the wrapper function
	return eb.bus.SubscribeAsync(params.Topic, handler, transactional)
}

// SubscribeOnce subscribes to an event topic for a single event occurrence with the given parameters.
func (eb *SystemEventBus) SubscribeOnce(params BusSubscriptionParams) error {
	// Create a wrapper function that constructs the Event and calls the provided EventHandler
	handler := func(args ...interface{}) {
		event := Event{
			Type: params.Topic,
			Data: args[0],
		}
		params.EventHandler(event)
	}

	// Store the function reference for later use in Unsubscribe
	eb.handlers[params.Topic] = handler

	// Subscribe once to the underlying EventBus using the wrapper function
	return eb.bus.SubscribeOnce(params.Topic, handler)
}

// SubscribeOnceAsync subscribes to an event topic asynchronously for a single event occurrence with the given parameters.
func (eb *SystemEventBus) SubscribeOnceAsync(params BusSubscriptionParams) error {
	// Create a wrapper function that constructs the Event and calls the provided EventHandler
	handler := func(args ...interface{}) {
		event := Event{
			Type: params.Topic,
			Data: args[0],
		}
		params.EventHandler(event)
	}

	// Store the function reference for later use in Unsubscribe
	eb.handlers[params.Topic] = handler

	// Subscribe once asynchronously to the underlying EventBus using the wrapper function
	return eb.bus.SubscribeOnceAsync(params.Topic, handler)
}

// Unsubscribe unsubscribes from an event topic with the given parameters.
func (eb *SystemEventBus) Unsubscribe(params BusSubscriptionParams) error {
	// Retrieve the function reference used for subscription
	handler, ok := eb.handlers[params.Topic]
	if !ok {
		return fmt.Errorf("no handler found for topic %s", params.Topic)
	}

	// Unsubscribe from the underlying EventBus using the stored function reference
	err := eb.bus.Unsubscribe(params.Topic, handler)
	if err == nil {
		// Remove the function reference from the handlers map if unsubscription succeeded
		delete(eb.handlers, params.Topic)
	}
	return err
}

// Publish publishes an event to the event bus.
func (eb *SystemEventBus) Publish(event Event) {
	eb.bus.Publish(event.Type, event.Data)
}

// HasCallback checks if a handler is registered for the given topic.
func (eb *SystemEventBus) HasCallback(topic string) bool {
	return eb.bus.HasCallback(topic)
}

// WaitAsync blocks until all asynchronous operations are completed.
func (eb *SystemEventBus) WaitAsync() {
	eb.bus.WaitAsync()
}
