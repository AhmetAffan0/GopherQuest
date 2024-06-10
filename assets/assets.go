package assets

import (
	"bytes"
	"embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
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

//go:embed lofi.ogg
var audioBGM []byte

const SampleRate = 44100

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

type Sound struct {
	PlayerSound *audio.Player
}

func SoundFunc() {
	_ = audio.NewContext(SampleRate)
	stream, err := vorbis.DecodeWithSampleRate(SampleRate, bytes.NewReader(audioBGM))
	if err != nil {
		panic(err)
	}
	audioPlayer := audio.CurrentContext().NewPlayer(stream)

}
