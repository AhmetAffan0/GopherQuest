package main

import (
	"log"
	"main/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	log.Println("Game Is Starting...")
	g := game.NewGame()

	ebiten.SetTPS(144)
	ebiten.SetWindowTitle("Quest")
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
