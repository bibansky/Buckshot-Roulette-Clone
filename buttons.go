package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Button struct {
	X      int
	Y      int
	Width  int
	Height int
	Text   string
}

func (b *Button) draw(screen *ebiten.Image) {

	vector.FillRect(screen, float32(b.X), float32(b.Y), float32(b.Width), float32(b.Height), color.RGBA{255, 255, 255, 1}, false)

	ebitenutil.DebugPrintAt(
		screen,
		b.Text,
		b.X+15,
		b.Y+20,
	)
}

func (b *Button) contains(x int, y int) bool {
	return x >= b.X &&
		x < b.X+b.Width &&
		y >= b.Y &&
		y < b.Y+b.Height
}

func (b *Button) isClicked() bool {
	mouseX, mouseY := ebiten.CursorPosition()

	return b.contains(mouseX, mouseY) && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)

}
