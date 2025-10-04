package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-snake/go_snake/apple"
	"image"
)

type World interface {
	Size() (int, int)
	Render() *ebiten.Image
	SetDebugMode(mode bool)
	GetAppleAt(p image.Point) *apple.Apple
	RemoveApple(a *apple.Apple)
}
