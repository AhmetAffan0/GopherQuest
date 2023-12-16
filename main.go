package main

import (
	"log"
	"main/assets"
	"main/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// rootContainer := widget.NewContainer(
	// 	widget.ContainerOpts.Layout(widget.NewGridLayout(
	// 		widget.GridLayoutOpts.Columns(1),
	// 		widget.GridLayoutOpts.Stretch([]bool{true}, []bool{false, true, false}),
	// 		widget.GridLayoutOpts.Padding(widget.Insets{
	// 			Top:    200,
	// 			Bottom: 20,
	// 			Left:   100,
	// 		}),
	// 	)),
	// )

	// eui := &ebitenui.UI{
	// 	Container: rootContainer,
	// }

	// ttfFont, err := truetype.Parse(goregular.TTF)
	// if err != nil {
	// 	log.Fatal("Error Parsing Font", err)
	// }
	// fontFace := truetype.NewFace(ttfFont, &truetype.Options{
	// 	Size: 32,
	// })

	// label := widget.NewText(
	// 	widget.TextOpts.Text("Gopher Quest", fontFace, color.White),
	// )

	// rootContainer.AddChild(label)

	log.Println("Game Is Starting...")
	assets.WriteLine()
	g := game.NewGame()

	log.Println(g.Update())

	ebiten.SetTPS(144)
	ebiten.SetWindowTitle("Quest")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
