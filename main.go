package main

import (
	"go-snake/go_snake/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
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
