package snake

import (
	"log"

	"go-snake/go_snake/config"
	"go-snake/go_snake/events"
	"go-snake/go_snake/types/direction"
)

func (s *Snake) listenForTurnEvents(stopCh chan any) {
	for {
		select {
		case <-stopCh:
			log.Default().Printf("Stopped listening for turn events")
			return
		case e := <-events.ActionCh:
			switch e {
			case config.ActionTurnNorth:
				s.turn(direction.North)
			case config.ActionTurnEast:
				s.turn(direction.East)
			case config.ActionTurnSouth:
				s.turn(direction.South)
			case config.ActionTurnWest:
				s.turn(direction.West)
			default:
				// Put unrecognized events back to the channel
				events.ActionCh <- e
			}
		}
	}
}
