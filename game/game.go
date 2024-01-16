package game

import (
	"fmt"
	"image/color"
	"main/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	camera camera
	player Player
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

	if g.player.player.x < 3000 {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1, 0.8)
		op.GeoM.Translate(-1000, 0)
		g.camera.draw(assets.Ground, op)
	}

	if g.player.player.x < -7000 {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1, 0.8)
		op.GeoM.Translate(-2000, 0)
		g.camera.draw(assets.Ground, op)
	}

	if g.player.player.x < -14000 {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1, 0.8)
		op.GeoM.Translate(-3000, 0)
		g.camera.draw(assets.Ground, op)
	}

	if g.player.player.x > 3000 {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1, 0.8)
		op.GeoM.Translate(1000, 0)
		g.camera.draw(assets.Ground, op)
	}

	s := assets.IdleSprite
	if g.player.player.vx > 0 {
		s = assets.RightSprite
	} else if g.player.player.vx < 0 {
		s = assets.LeftSprite
	}

	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Scale(0.3, 0.3)
	op2.GeoM.Translate(float64(g.player.player.x)/unit, float64(g.player.player.y)/unit)
	g.camera.draw(s, op2)

	g.camera.render(screen)

	msg := fmt.Sprintf("Gopher X: %.2f, Y: %.2f", float64(g.player.player.x), float64(g.player.player.y))
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Update() error {
	g.player.Update()
	g.camera.setPos(g.player.player.x/unit-300, 0) //g.player.player.y/unit-400)
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
