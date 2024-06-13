package game

import (
	"bytes"
	_ "embed"
	"fmt"
	"image/color"
	"log"
	"main/assets"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed lofi.ogg
var audioBGM []byte

type Sound struct {
	player       *audio.Player
	audioContext *audio.Context
}

const (
	screenWidth    = 640
	screenHeight   = 480
	sampleRate     = 48000
	bytesPerSample = 4

	introLengthInSecond = 5
	loopLengthInSecond  = 180
)

func (s *Sound) SoundFunc() {
	if s.audioContext == nil {
		s.audioContext = audio.NewContext(sampleRate)
	}

	oggS, err := vorbis.DecodeWithoutResampling(bytes.NewReader(audioBGM))
	if err != nil {
		log.Fatal(err)
	}

	sound := audio.NewInfiniteLoopWithIntro(oggS, introLengthInSecond*bytesPerSample*sampleRate, loopLengthInSecond*bytesPerSample*sampleRate)

	s.player, err = s.audioContext.NewPlayer(sound)
	if err != nil {
		log.Fatal(err)
	}

	s.player.Play()
}

type Game struct {
	camera camera
	player Player
	npc    NPC
	s      Sound

	background        Background
	isDebugModeOn     bool
	myBool            bool
	menuOff           bool
	DoorForFirstScene *ebiten.Image
}

const (
	sW = 635
	sH = 475
)

var (
	//Background Colors
	LightBlue = color.RGBA{0x80, 0xa0, 0xc0, 0xff} //Basically morning
	Blackish  = color.RGBA{0, 0, 50, 250}          //Basically night
)

var (
	fontFaceSource  *text.GoTextFaceSource
	fontFaceSource2 *text.GoTextFaceSource
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(assets.Font_ttf))
	if err != nil {
		log.Fatal(err)
	}
	fontFaceSource = s

	s2, err := text.NewGoTextFaceSource(bytes.NewReader(assets.Sans_ttf))
	if err != nil {
		log.Fatal(err)
	}
	fontFaceSource2 = s2
}

func NewGame() *Game {
	g := &Game{}
	g.camera.init()
	g.s.SoundFunc()
	return g
}

const (
	mytext  = "Gopher Walk"
	mytext2 = "Press \"Enter\" to Start the Game"
)

func (g *Game) Draw(screen *ebiten.Image) {
	const (
		normalFontSize = 24
		bigFontSize    = 48
	)

	const x = 180

	op := &text.DrawOptions{}
	op.GeoM.Translate(x, 60)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, mytext, &text.GoTextFace{
		Source: fontFaceSource,
		Size:   bigFontSize,
	}, op)

	op = &text.DrawOptions{}
	op.GeoM.Translate(145, 250)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, mytext2, &text.GoTextFace{
		Source: fontFaceSource,
		Size:   normalFontSize,
	}, op)

	if g.menuOff {
		g.camera.clear()

		g.background.ChangeScene(screen, &g.camera, g)

		g.player.Draw(screen, &g.camera, *g)

		g.camera.render(screen)

		if g.isDebugModeOn {
			msg2 := fmt.Sprintf("TPS: %.2f, FPS: %.2f, VSync: %v",
				ebiten.ActualTPS(),
				ebiten.ActualFPS(),
				ebiten.IsVsyncEnabled())
			ebitenutil.DebugPrint(screen, msg2)
		}

		g.npc.conversation(g, screen)
	}
}

func (g *Game) debugMode() bool {
	return g.isDebugModeOn
}

func (g *Game) isDebugMode(enabled bool) {
	if enabled {
		g.isDebugModeOn = true
	} else {
		g.isDebugModeOn = false
	}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.menuOff = true
	}

	if g.menuOff {
		g.player.Update()
		g.camera.setPos(g.player.player.x/unit-300, 0)

		debugMode := g.debugMode()

		if inpututil.IsKeyJustPressed(ebiten.KeyL) {
			g.isDebugMode(!debugMode)
		}

		vsync := ebiten.IsVsyncEnabled()

		if inpututil.IsKeyJustPressed(ebiten.KeyV) {
			ebiten.SetVsyncEnabled(!vsync)
		}

		if g.myBool {
		} else {
			if g.player.player.x <= -27000 {
				g.camera.setPos(-3000, 0)
			}

			if g.player.player.x >= 26500 {
				g.camera.setPos(2350, 0)
			}
		}

		if g.player.player.x >= 19900 && g.player.player.x <= 20500 {
			if inpututil.IsKeyJustPressed(ebiten.KeyW) {
				g.myBool = true
				g.player.isBorder = true
				//g.background.menuOn = false
				g.player.player.x = 0
				time.Sleep(time.Second * 1)
			}
		}

		if g.player.player.x >= -36300 && g.player.player.x <= -35400 {
			if inpututil.IsKeyJustPressed(ebiten.KeyW) {
				g.myBool = false
				g.player.isBorder = false
				//g.background.menuOn = false
				g.player.player.x = 0
				time.Sleep(time.Second * 1)
			}
		}
	}

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return sW, sH
}
