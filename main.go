package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 960
	screenHeight = 540
)

func main() {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Buckshot")

	game := newGame()

	err := ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}
}
