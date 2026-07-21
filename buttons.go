package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	debugCharWidth  = 6
	debugCharHeight = 12
)

type Button struct {
	X    int
	Y    int
	Text string
}

func (b *Button) width() int {
	return len(b.Text) * debugCharWidth
}

func (b *Button) draw(screen *ebiten.Image) {
	mouseX, mouseY := ebiten.CursorPosition()

	if b.contains(mouseX, mouseY) {
		ebitenutil.DebugPrintAt(screen, ">", b.X-12, b.Y)
	}

	ebitenutil.DebugPrintAt(screen, b.Text, b.X, b.Y)

}

func (b *Button) contains(x int, y int) bool {
	return x >= b.X &&
		x < b.X+b.width() &&
		y >= b.Y &&
		y < b.Y+debugCharHeight
}

func (b *Button) isClicked() bool {
	mouseX, mouseY := ebiten.CursorPosition()

	return b.contains(mouseX, mouseY) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)

}
