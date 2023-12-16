package game

import (
	"image/color"
	"main/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Background struct{}

func (b *Background) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	screen.DrawImage(assets.Ground, op)
}
