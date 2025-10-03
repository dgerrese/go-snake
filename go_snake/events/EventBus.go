package events

import "log"

type EventBus struct {
	subscribers map[EventType][]chan any
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[EventType][]chan any),
	}
}

func (eb *EventBus) Subscribe(eventType EventType, ch chan any) {
	eb.subscribers[eventType] = append(eb.subscribers[eventType], ch)
}

func (eb *EventBus) Unsubscribe(eventType EventType, ch chan any) {
	subscribers := eb.subscribers[eventType]
	for i, subscriber := range subscribers {
		if subscriber == ch {
			eb.subscribers[eventType] = append(subscribers[:i], subscribers[i+1:]...)
			break
		}
	}
}

func (eb *EventBus) Publish(eventType EventType, data any) {
	log.Default().Printf("Publishing event %s with data: %v", eventType, data)

	if subscribers, found := eb.subscribers[eventType]; found {
		for _, ch := range subscribers {
			ch <- data
		}
	}
}
