package game

import (
	"main/assets"

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
	g := &Game{}
	g.camera.init()
	return g
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.camera.clear()
	g.camera.draw(assets.Ground, &ebiten.DrawImageOptions{})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.8, 0.8)
	op.GeoM.Translate(float64(g.player.player.x)/unit, float64(g.player.player.y)/unit)
	g.camera.draw(assets.IdleSprite, op)
	screen.DrawImage(assets.IdleSprite, op)
	//g.player.Draw(screen)
	g.camera.render(screen)

}

func (g *Game) Update() error {
	g.player.Update()
	g.camera.setPos(g.player.player.x/unit-300, g.player.player.y/unit-400)
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
