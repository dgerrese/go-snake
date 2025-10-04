package config

import "github.com/hajimehoshi/ebiten/v2"

type Action byte

const (
	// Snake actions

	ActionTurnNorth Action = iota
	ActionTurnSouth
	ActionTurnWest
	ActionTurnEast

	// Game actions

	ActionPauseGame
	ActionRestart

	// Debug actions

	ActionToggleDebugMode
	ActionIncreaseGameSpeed
	ActionDecreaseGameSpeed
)

var Keymap = map[Action][]ebiten.Key{
	ActionTurnNorth:         {ebiten.KeyArrowUp, ebiten.KeyW},
	ActionTurnSouth:         {ebiten.KeyArrowDown, ebiten.KeyS},
	ActionTurnWest:          {ebiten.KeyArrowLeft, ebiten.KeyA},
	ActionTurnEast:          {ebiten.KeyArrowRight, ebiten.KeyD},
	ActionToggleDebugMode:   {ebiten.KeyF3},
	ActionPauseGame:         {ebiten.KeyP},
	ActionRestart:           {ebiten.KeyR},
	ActionIncreaseGameSpeed: {ebiten.KeyEqual, ebiten.KeyKPAdd},
	ActionDecreaseGameSpeed: {ebiten.KeyMinus, ebiten.KeyKPSubtract},
}
