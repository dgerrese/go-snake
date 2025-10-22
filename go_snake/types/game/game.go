package game

import "github.com/hajimehoshi/ebiten/v2"

type Game interface {
	ebiten.Game
	Start()
}
