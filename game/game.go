package game

import (
	"image/color"
	"main/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1000
	screenHeight = 600
)

type Game struct {
	camera     camera
	background Background
	player     Player
}

func NewGame() *Game {
	g := &Game{}
	g.camera.init()
	return g
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	g.camera.clear()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 0.8)
	g.camera.draw(assets.Ground, op)
	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Scale(0.3, 0.3)
	op2.GeoM.Translate(float64(g.player.player.x)/unit, float64(g.player.player.y)/unit)
	g.camera.draw(assets.IdleSprite, op2)
	//g.player.Draw(screen)
	g.camera.render(screen)

}

func (g *Game) Update() error {
	g.player.Update()
	g.camera.setPos(g.player.player.x/unit-300, 0) //g.player.player.y/unit-400)
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
