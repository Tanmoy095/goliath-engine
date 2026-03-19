package runtime

import "sync"

type EventType uint8

const (
	EventTaskCompleted EventType = iota
	EventTaskFailed
	EventTaskReady
)

// Event is the tiny message sent through the bus.
// It just contains the type of event and the ID of the task it relates to
type Event struct {
	Type   EventType
	TaskID uint64
}
type EventBus struct {
	mu          sync.RWMutex
	subscribers map[EventType][]chan Event
}

func (e *EventBus) Subscribe(eventType EventType) chan Event {
	e.mu.Lock() //RWMutex.Lock because we are modifying the Subscribers map. so no one can read or write to it while we are modifying it.
	defer e.mu.Unlock()
	//make new channel for every subscribers, so they can receive events independently.
	ch := make(chan Event, 100)                                     //buffered channel to avoid blocking the publisher if the subscriber is not ready to receive the event.
	e.subscribers[eventType] = append(e.subscribers[eventType], ch) //append the new channel to the list of subscribers for the given event type.
	return ch
}

func (e *EventBus) Publish(event Event) {
	e.mu.RLock() //RWMutex.RLock for only reading the Subscribers map. so multiple goroutines can read from it concurrently.
	defer e.mu.RUnlock()
	for _, ch := range e.subscribers[event.Type] {
		//non blocking send to the channel, if the subscriber is not ready to receive the event, we just skip it.
		select {
		case ch <- event:
			//event sent successfully to the channel . message sent to the subscriber.
		default:
			//subscriber is not ready to receive the event, we just skip it.
		}

	}

}
