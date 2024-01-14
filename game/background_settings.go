package game

import (
	"image/color"
	"main/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Background struct{}

func (b *Background) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 0.8)
	screen.DrawImage(assets.Ground, op)
}
