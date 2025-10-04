package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-snake/go_snake/game"
	"log"
)

func main() {
	ebiten.SetWindowSize(1080, 720)
	ebiten.SetWindowTitle("Snake Game in Go")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	err := ebiten.RunGame(game.NewGame(120))

	if err != nil {
		log.Fatal(err)
	}
}
