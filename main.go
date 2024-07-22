package main

import (
	"log"
	"main/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	log.Println("Game Is Starting...")
	g := game.NewGame()

	ebiten.SetWindowTitle("Quest")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

//niğgffg
