package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"go-snake/go_snake"
	"log"
)

func main() {
	ebiten.SetWindowSize(1080, 720)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	g := go_snake.NewGame(32, 24, 1000)
	g.EnableDebugMode()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
