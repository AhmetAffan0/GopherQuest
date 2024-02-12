package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type debugMode struct {
	isDebugModeOn bool
}

func (d *debugMode) debugMode(g Game) {
	if inpututil.IsKeyJustPressed(ebiten.KeySlash) {
		d.isDebugModeOn = true
	}

}
func (d *debugMode) debugUtils(screen *ebiten.Image, g Game, p Player, c camera) {
	msg := fmt.Sprintf("Gopher X: %.2f, Y: %.2f",
		float64(p.player.x),
		float64(p.player.y))
	ebitenutil.DebugPrintAt(screen, msg, p.player.x/unit-c.x-70, p.player.y/unit-c.y-10)

	msg2 := fmt.Sprintf("TPS: %.2f, FPS: %.2f, VSync: %v",
		ebiten.ActualTPS(),
		ebiten.ActualFPS(),
		ebiten.IsVsyncEnabled())
	ebitenutil.DebugPrint(screen, msg2)

	// msg3 := fmt.Sprintf("\n\nNPC X: %.2f, NPC Y: %.2f",
	// 	g.npc.x,
	// 	g.npc.y)
	// ebitenutil.DebugPrint(screen, msg3)
	// g.npc.conversation(*g, screen)

}
