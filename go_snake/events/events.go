package events

import (
	"go-snake/go_snake/apple"
	"go-snake/go_snake/config"
	"image"
)

var ActionCh = make(chan config.Action)
var MovementCh = make(chan image.Point)
var EatingCh = make(chan *apple.Apple)

type gameEvent byte

const (
	DeathEvent gameEvent = iota
)

var GameCh = make(chan gameEvent)
