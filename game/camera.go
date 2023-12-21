package game

import "github.com/hajimehoshi/ebiten/v2"

type camera struct {
	x        int
	y        int
	drawable *ebiten.Image
}

func (c *camera) setInitialImg() {
	c.drawable = ebiten.NewImage(screenWidth, screenHeight)
}

func (c *camera) draw(screen *ebiten.Image) {
	c.drawable.DrawImage(screen, &ebiten.DrawImageOptions{})
}
