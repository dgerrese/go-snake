package input

import (
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"go-snake/go_snake/config"
	"go-snake/go_snake/events"
)

func GameControlInput() {
	PauseGame()
	RestartGame()
}

func PauseGame() {
	for _, k := range config.Keymap[config.ActionPauseGame] {
		if inpututil.IsKeyJustPressed(k) {
			events.ActionCh <- config.ActionPauseGame
			return
		}
	}
}

func RestartGame() {
	for _, k := range config.Keymap[config.ActionRestart] {
		if inpututil.IsKeyJustPressed(k) {
			events.ActionCh <- config.ActionRestart
			return
		}
	}
}
