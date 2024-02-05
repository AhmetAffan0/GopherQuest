package game

import (
	"bytes"
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	assets "github.com/lidldev/GameResources"
)

type Game struct {
	camera camera
	player Player

	background        Background
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
	fontFaceSource *text.GoTextFaceSource
)

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(assets.Font_ttf))
	if err != nil {
		log.Fatal(err)
	}
	fontFaceSource = s
}

func NewGame() *Game {
	g := &Game{}
	g.camera.init()
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

		g.player.Draw(screen, &g.camera)

		g.camera.render(screen)

		msg := fmt.Sprintf("Gopher X: %.2f, Y: %.2f",
			float64(g.player.player.x),
			float64(g.player.player.y))
		ebitenutil.DebugPrint(screen, msg)

		msg2 := fmt.Sprintf("\nTPS: %.2f, FPS: %.2f, VSync: %v",
			ebiten.ActualTPS(),
			ebiten.ActualFPS(),
			ebiten.IsVsyncEnabled())
		ebitenutil.DebugPrint(screen, msg2)
	}
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.menuOff = true
	}
	if g.menuOff {
		g.player.Update()
		g.camera.setPos(g.player.player.x/unit-300, 0)

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
				g.player.isBorder = true
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
