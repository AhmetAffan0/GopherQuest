package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	assets "github.com/lidldev/GameResources"
)

type Background struct {
	image *ebiten.Image

	isDrawed bool
}

func (b *Background) ChangeScene(c *camera, g *Game) {
	b.image = assets.GopherWalkBackground

	if g.myBool {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1, 0.4)
		op.GeoM.Translate(-3000, 0)
		c.draw(assets.GopherJumpBackground, op)
		b.image.Clear()
		g.Door.Clear()
	} else {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1, 0.8)
		op.GeoM.Translate(-3000, 0)
		c.draw(b.image, op)
	}
}
