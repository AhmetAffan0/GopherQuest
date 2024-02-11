package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	assets "github.com/lidldev/GameResources"
)

type NPC struct {
	x      float64
	y      float64
	amogus *ebiten.Image
}

func (n *NPC) AmogusPos(x, y int) {
	n.x = float64(x)
	n.y = float64(y)
}

func (n *NPC) drawAmogus(g Game, c camera) {
	n.amogus = assets.AmongUsChar

	n.x = 2000
	n.y = 350

	op4 := &ebiten.DrawImageOptions{}
	op4.GeoM.Scale(0.35, 0.35)
	op4.GeoM.Translate(-n.x, n.y)
	c.draw(n.amogus, op4)
}
