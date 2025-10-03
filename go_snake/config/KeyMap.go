package config

import "github.com/hajimehoshi/ebiten/v2"

type Action string

const (
	ActionMoveUp    Action = "up"
	ActionMoveDown  Action = "down"
	ActionMoveLeft  Action = "left"
	ActionMoveRight Action = "right"
)

var KeyMap = map[Action]ebiten.Key{
	ActionMoveUp:    ebiten.KeyArrowUp,
	ActionMoveDown:  ebiten.KeyArrowDown,
	ActionMoveLeft:  ebiten.KeyArrowLeft,
	ActionMoveRight: ebiten.KeyArrowRight,
}
