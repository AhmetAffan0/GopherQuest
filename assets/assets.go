package assets

import (
	"embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed *
var assets embed.FS

// Gopher Part
var GopherIdle = GetSingleImage("mainchar.png")
var GopherRight = GetSingleImage("right.png")
var GopherLeft = GetSingleImage("left.png")

// Chars part
var AmongUsChar = GetSingleImage("impostor.png")
var Door = GetSingleImage("door.png")

// Background Part
var GopherWalkBackground = GetSingleImage("GopherWalk.png")
var GopherWalkBackground2 = GetSingleImage("GopherWalk2.png")

//go:embed Honk-Regular.ttf
var Font_ttf []byte

//go:embed OpenSans-Medium.ttf
var Sans_ttf []byte

func GetSingleImage(name string) *ebiten.Image {
	file, err := assets.Open(name)
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
