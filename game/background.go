package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	assets "github.com/lidldev/GameResources"
)

type Background struct {
	n NPC
}

func (b *Background) ChangeScene(screen *ebiten.Image, c *camera, g *Game) {
	screen.Fill(color.Black)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 0.79)
	op.GeoM.Translate(-3000, 0)

	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Scale(1, 0.8)
	op2.GeoM.Translate(-3000, 0)

	op3 := &ebiten.DrawImageOptions{}
	op3.GeoM.Scale(0.45, 0.35)
	op3.GeoM.Translate(-3600, 316)

	op5 := &ebiten.DrawImageOptions{}
	op5.GeoM.Scale(0.45, 0.35)
	op5.GeoM.Translate(2000, 316)

	if g.menuOff {
		if g.myBool {
			screen.Fill(Blackish)
		} else {
			screen.Fill(LightBlue)
		}

		g.DoorForFirstScene = assets.Door

		if g.myBool {
			c.draw(assets.GopherWalkBackground2, op)
			c.draw(assets.Door, op3)
			b.n.drawAmogus(*g, *c)
		} else if !g.myBool {
			c.draw(assets.GopherWalkBackground, op2)
			g.camera.draw(assets.Door, op5)
		}
	}

}
