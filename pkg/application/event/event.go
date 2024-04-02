package event

const (
	// EventTypeDataExtracted represents an event emitted when data is extracted from a source blockchain.
	EventTypeDataExtracted string = "data_extracted"

	// EventTypeDataTransformed represents an event emitted when data is transformed by a pipeline stage.
	EventTypeDataTransformed string = "data_transformed"

	// EventTypeDataLoaded represents an event emitted when data is loaded into a target blockchain.
	EventTypeDataLoaded string = "data_loaded"
)

// Event represents an event within the system.
type Event struct {
	// Type is the type or identifier of the event.
	Type string

	// Data is the payload or data associated with the event.
	Data interface{}
}

// EventHandler defines the signature for an event handler function.
type EventHandler func(event Event)

// BusSubscriptionParams represents the parameters for subscribing to an event topic.
type BusSubscriptionParams struct {
	// Topic is the event topic to subscribe to.
	Topic string

	// EventHandler is the function that will handle the events for the subscribed topic.
	EventHandler EventHandler
}

// BusSubscriber defines subscription-related bus behavior.
type BusSubscriber interface {
	// Subscribe subscribes to an event topic with the given parameters.
	Subscribe(params BusSubscriptionParams) error

	// SubscribeAsync subscribes to an event topic asynchronously with the given parameters.
	SubscribeAsync(params BusSubscriptionParams, transactional bool) error

	// SubscribeOnce subscribes to an event topic for a single event occurrence with the given parameters.
	SubscribeOnce(params BusSubscriptionParams) error

	// SubscribeOnceAsync subscribes to an event topic asynchronously for a single event occurrence with the given parameters.
	SubscribeOnceAsync(params BusSubscriptionParams) error

	// Unsubscribe unsubscribes from an event topic with the given parameters.
	Unsubscribe(params BusSubscriptionParams) error
}

// BusPublisher defines publishing-related bus behavior.
type BusPublisher interface {
	// Publish publishes an event to the event bus.
	Publish(event Event)
}

// BusController defines bus control behavior (checking handler's presence, synchronization).
type BusController interface {
	// HasCallback checks if a handler is registered for the given topic.
	HasCallback(topic string) bool

	// WaitAsync blocks until all asynchronous operations are completed.
	WaitAsync()
}

// EventBusInterface englobes global (subscribe, publish, control) bus behavior.
type EventBusInterface interface {
	BusController
	BusSubscriber
	BusPublisher
}
