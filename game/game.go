package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	g.camera.clear()

	g.background.ChangeScene(screen, &g.camera, g)

	g.player.Draw(screen, &g.camera)

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
