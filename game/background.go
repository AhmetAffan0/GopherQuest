package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	assets "github.com/lidldev/GameResources"
)

type Background struct {
	image *ebiten.Image

	isDrawed bool
}

func (b *Background) ChangeScene(screen *ebiten.Image, c *camera, g *Game) {
	b.image = assets.GopherWalkBackground

	if g.myBool {
		screen.Fill(Blackish)
	} else {
		screen.Fill(LightBlue)
	}

	g.Door = assets.Door

	op3 := &ebiten.DrawImageOptions{}
	op3.GeoM.Scale(0.45, 0.35)
	op3.GeoM.Translate(2000, 316)
	g.camera.draw(g.Door, op3)

	if g.myBool {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1, 0.79)
		op.GeoM.Translate(-3000, 0)
		c.draw(assets.GopherWalkBackground2, op)
		b.image.Clear()
		g.Door.Clear()
	} else {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1, 0.8)
		op.GeoM.Translate(-3000, 0)
		c.draw(b.image, op)
	}
}
