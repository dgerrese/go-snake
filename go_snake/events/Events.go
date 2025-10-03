package events

type EventType byte

const (
	EventTypeAppleEaten EventType = iota
	EventSnakeTurn
)

var eventNames = map[EventType]string{
	EventTypeAppleEaten: "AppleEaten",
}

func (e EventType) String() string {
	return eventNames[e]
}
