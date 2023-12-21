package game

import (
	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1000
	screenHeight = 600
)

type Game struct {
	camera     *camera
	background Background
	player     Player
	ui         *ebitenui.UI
	print      bool
}

func NewGame() *Game {
	g := &Game{
		//ui: ,
	}

	return g
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.background.Draw(screen)
	g.player.Draw(screen)
}

func (g *Game) Update() error {
	g.player.Update()
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
