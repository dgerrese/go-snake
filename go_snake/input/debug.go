package input

import (
	"go-snake/go_snake/config"
	"go-snake/go_snake/events"

	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func DebugInput() {
	ToggleDebugMode()
	IncreaseGameSpeed()
	DecreaseGameSpeed()
}

func ToggleDebugMode() {
	for _, k := range config.Keymap[config.ActionToggleDebugMode] {
		if inpututil.IsKeyJustPressed(k) {
			events.ActionCh <- config.ActionToggleDebugMode
			return
		}
	}
}

func IncreaseGameSpeed() {
	for _, k := range config.Keymap[config.ActionIncreaseGameSpeed] {
		if inpututil.IsKeyJustPressed(k) {
			events.ActionCh <- config.ActionIncreaseGameSpeed
			return
		}
	}
}

func DecreaseGameSpeed() {
	for _, k := range config.Keymap[config.ActionDecreaseGameSpeed] {
		if inpututil.IsKeyJustPressed(k) {
			events.ActionCh <- config.ActionDecreaseGameSpeed
			return
		}
	}
}
