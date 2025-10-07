package input

import (
	"go-snake/go_snake/config"
	"go-snake/go_snake/events"

	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func SnakeInput() {
	SnakeTurnNorth()
	SnakeTurnSouth()
	SnakeTurnWest()
	SnakeTurnEast()
}

func SnakeTurnNorth() {
	for _, k := range config.Keymap[config.ActionTurnNorth] {
		if inpututil.IsKeyJustPressed(k) {
			events.ActionCh <- config.ActionTurnNorth
			return
		}
	}
}

func SnakeTurnSouth() {
	for _, k := range config.Keymap[config.ActionTurnSouth] {
		if inpututil.IsKeyJustPressed(k) {
			events.ActionCh <- config.ActionTurnSouth
			return
		}
	}
}

func SnakeTurnWest() {
	for _, k := range config.Keymap[config.ActionTurnWest] {
		if inpututil.IsKeyJustPressed(k) {
			events.ActionCh <- config.ActionTurnWest
			return
		}
	}
}

func SnakeTurnEast() {
	for _, k := range config.Keymap[config.ActionTurnEast] {
		if inpututil.IsKeyJustPressed(k) {
			events.ActionCh <- config.ActionTurnEast
			return
		}
	}
}
