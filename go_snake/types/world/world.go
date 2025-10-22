package world

import (
	"image"

	"go-snake/go_snake/apple"

	"github.com/hajimehoshi/ebiten/v2"
)

type World interface {
	Size() (int, int)
	Render() *ebiten.Image
	SetDebugMode(mode bool)
	GetAppleAt(p image.Point) *apple.Apple
	RemoveApple(a *apple.Apple)
}
