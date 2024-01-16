package assets

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed *
var assets embed.FS

var IdleSprite = GetSingleImage("mainchar.png")
var RightSprite = GetSingleImage("right.png")
var LeftSprite = GetSingleImage("left.png")
var Platform = GetSingleImage("tile_0146.png")
var Ground = GetSingleImage("pbkg.png")

func GetSingleImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
