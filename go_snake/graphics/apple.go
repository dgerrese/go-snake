package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Apple int

const (
	AppleRed    Apple = 6
	AppleGold   Apple = 16 + 6
	AppleCherry Apple = (3 * 16) + 6
	AppleCookie Apple = (3 * 16) + 7
	AppleGreen  Apple = (4 * 16) + 6
)

func (a Apple) Image() *ebiten.Image {
	return getTile(int(a))
}
