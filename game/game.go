package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	assets "github.com/lidldev/GameResources"
)

type Game struct {
	camera camera
	player Player

	background Background
	myBool     bool

	Door *ebiten.Image
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

func NewGame() *Game {
	g := &Game{}
	g.camera.init()
	return g
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.myBool {
		screen.Fill(Blackish)
	} else {
		screen.Fill(LightBlue)
	}

	g.camera.clear()

	g.background.ChangeScene(&g.camera, g)

	g.Door = assets.Door

	op3 := &ebiten.DrawImageOptions{}
	op3.GeoM.Scale(0.45, 0.35)
	op3.GeoM.Translate(2000, 316)
	g.camera.draw(g.Door, op3)

	s := assets.GopherIdle
	if g.player.player.vx > 0 {
		s = assets.GopherRight
	} else if g.player.player.vx < 0 {
		s = assets.GopherLeft
	}

	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Scale(0.3, 0.3)
	op2.GeoM.Translate(float64(g.player.player.x)/unit, float64(g.player.player.y)/unit)
	g.camera.draw(s, op2)

	g.camera.render(screen)

	msg := fmt.Sprintf("Gopher X: %.2f, Y: %.2f", float64(g.player.player.x), float64(g.player.player.y))
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Update() error {
	g.player.Update()
	g.camera.setPos(g.player.player.x/unit-300, 0)

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
			g.background.isDrawed = false
			g.player.player.x = 0
		}
	}

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return sW, sH
}
