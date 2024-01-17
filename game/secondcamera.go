package game

import "github.com/hajimehoshi/ebiten/v2"

type camera2 struct {
	x int
	y int

	drawable *ebiten.Image // the image that the camera will draw
}

func (c *camera) setPos2(x, y int) {
	c.x = x
	c.y = y
}

func (c *camera) init2() {
	c.drawable = ebiten.NewImage(800, 600)
}

func (camera *camera) render2(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	screen.DrawImage(camera.drawable, op)
}

func (c *camera) draw2(image *ebiten.Image, op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(float64(-c.x), float64(-c.y))

	c.drawable.DrawImage(image, op)
}

func (c *camera) clear2() {
	c.drawable.Clear()
}
