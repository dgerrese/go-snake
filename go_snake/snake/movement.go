package snake

import (
	"image"

	"go-snake/go_snake/events"
	"go-snake/go_snake/types/direction"
)

const directionQueueMaxLength = 3

var validTurns = map[direction.Direction][]direction.Direction{
	direction.North:     {direction.East, direction.West},
	direction.East:      {direction.North, direction.South},
	direction.South:     {direction.East, direction.West},
	direction.West:      {direction.North, direction.South},
	direction.Undefined: {direction.North, direction.East, direction.South, direction.West},
}

func (s *Snake) Move() {
	if len(s.directionQueue) > 0 {
		s.direction = s.directionQueue[0]
		s.directionQueue = s.directionQueue[1:]
	}

	nh := s.NextHead()

	if s.Head() == nh {
		// snake is not moving
		return
	}

	if s.eat(nh) {
		return
	}

	s.body = append([]image.Point{nh}, s.body[:len(s.body)-1]...)

	events.MovementCh <- nh
}

func (s *Snake) turn(d direction.Direction) {
	if len(s.directionQueue) >= directionQueueMaxLength {
		return
	}

	var ld direction.Direction

	if len(s.directionQueue) > 0 {
		ld = s.directionQueue[len(s.directionQueue)-1]
	} else {
		ld = s.direction
	}

	for _, v := range validTurns[ld] {
		if d == v {
			s.directionQueue = append(s.directionQueue, d)
			return
		}
	}
}
