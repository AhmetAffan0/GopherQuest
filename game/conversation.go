package game

import (
	"fmt"

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
	//d := math.Sqrt(math.Pow(n.y-float64(g.player.player.y), 2)/unit + math.Pow(n.x-float64(g.player.player.x), 2)/unit)

	if g.myBool {
		if g.player.player.x < -18200 && g.player.player.x > -20900 {
			msg := fmt.Sprintln("\nI love ebitengine community")
			ebitenutil.DebugPrint(screen, msg)
		}
	}
}
