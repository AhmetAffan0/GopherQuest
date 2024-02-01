package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	assets "github.com/lidldev/GameResources"
)

type Background struct {
	image    *ebiten.Image
	isDrawed bool
}

func (b *Background) ChangeScene(c *camera) {
	b.image = assets.GopherWalkBackground

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 0.8)
	op.GeoM.Translate(-3000, 0)
	c.draw(b.image, op)

}
