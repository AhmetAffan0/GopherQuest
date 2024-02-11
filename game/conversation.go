package game

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	assets "github.com/lidldev/GameResources"
)

type NPC struct {
	x      float64
	y      float64
	amogus *ebiten.Image
}

func (n *NPC) AmogusPos(x, y float64) {
	n.x = x
	n.y = y
}

func (n *NPC) drawAmogus(c camera) {
	n.amogus = assets.AmongUsChar
	n.AmogusPos(2000, -350)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.35, 0.35)
	op.GeoM.Translate(-n.x, -n.y)
	c.draw(n.amogus, op)
}

func (n *NPC) conversation(g Game, screen *ebiten.Image) {
	d := math.Sqrt(math.Pow(n.y-float64(g.player.player.y), 2)/unit + math.Pow(n.x-float64(g.player.player.x), 2)/unit)

	msg := fmt.Sprintf("\n\n\n\n%.2f", d)
	ebitenutil.DebugPrint(screen, msg)

}
