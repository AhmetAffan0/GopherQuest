package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	assets "github.com/lidldev/GameResources"
)

type NPC struct {
	x         float64
	y         float64
	amogus    *ebiten.Image
	isPressed bool
}

const (
	normalFontSize = 24
	bigFontSize    = 48
)

const x = 220

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

var (
	textArr []string
)

func (n *NPC) mainText(screen *ebiten.Image, g *Game) {
	//impText := fmt.Sprintln("Press E To Interact")
	textArr = append(textArr, "Press E To Interact", "Have You Ever Heard Of Among Us?")

	op3 := &text.DrawOptions{}
	op3.GeoM.Translate(x, 102)
	op3.ColorScale.ScaleWithColor(color.RGBA{0x80, 0xa0, 0xc0, 0xff})

	op4 := &text.DrawOptions{}
	op4.GeoM.Translate(120, 102)
	op4.ColorScale.ScaleWithColor(color.RGBA{0x80, 0xa0, 0xc0, 0xff})

	if !n.isPressed {
		text.Draw(screen, textArr[0], &text.GoTextFace{
			Source: fontFaceSource2,
			Size:   normalFontSize,
		}, op3)
	} else {
		text.Draw(screen, textArr[1], &text.GoTextFace{
			Source: fontFaceSource2,
			Size:   normalFontSize,
		}, op4)
	}
}

func (n *NPC) textMode() bool {
	return n.isPressed
}

func (n *NPC) isTextMode(enabled bool) {
	if enabled {
		n.isPressed = true
	}
}

func (n *NPC) conversation(g *Game, screen *ebiten.Image) {
	rect := ebiten.NewImage(500, 100)
	rect.Fill(color.RGBA{100, 100, 100, 100})

	isEnabled := n.textMode()

	if g.myBool {
		if g.player.player.x < -18200 && g.player.player.x > -20900 {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(70, 70)
			screen.DrawImage(rect, op)

			n.mainText(screen, g)

			if inpututil.IsKeyJustPressed(ebiten.KeyE) {
				n.isTextMode(!isEnabled)
				n.mainText(screen, g)
			}
		}
	}
}
