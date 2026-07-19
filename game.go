package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameState int

const (
	StateMenu GameState = iota
	StatePlayerTurn
	StateDealerTurn
	StateGameOver
)

type Game struct {
	playerHealth int
	dealerHealth int

	dealerTimer int

	shells    []Shell
	playerWon bool
	state     GameState

	startButton Button
}

func (g *Game) Update() error {
	switch g.state {
	case StateMenu:
		if g.startButton.isClicked() {
			g.state = StatePlayerTurn
		}

	case StatePlayerTurn:
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			shell := g.takeShell()
			if shell.Live {
				g.dealerHealth--
				g.state = StateDealerTurn
			} else {
				g.state = StateDealerTurn
			}

		} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
			shell := g.takeShell()
			if shell.Live {
				g.playerHealth--
				g.state = StateDealerTurn
			} else {
				g.state = StatePlayerTurn
			}
		}

	case StateDealerTurn:
		g.dealerTimer++

		if g.dealerTimer < 60 {
			return nil
		}

		g.dealerTimer = 0

		choice := rand.N(2)

		if choice == 0 {
			shell := g.takeShell()
			if shell.Live {
				g.dealerHealth--
				g.state = StatePlayerTurn
			} else {
				g.state = StateDealerTurn

			}
		} else if choice == 1 {
			shell := g.takeShell()
			if shell.Live {
				g.playerHealth--
				g.state = StatePlayerTurn
			} else {
				g.state = StatePlayerTurn
			}
		}

	case StateGameOver:
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			g.restart()
		}
	}

	g.statusGameOver()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	if g.state == StatePlayerTurn || g.state == StateDealerTurn {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Player: %d", g.playerHealth), 0, 0)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Dealer: %d", g.dealerHealth), 0, 10)

	}
	if g.state == StateMenu {
		g.startButton.draw(screen)
		ebitenutil.DebugPrintAt(screen, "EXIT", 430, 350)
	}

	if g.state == StatePlayerTurn {
		ebitenutil.DebugPrintAt(screen, "DEALER", 430, 330)
		ebitenutil.DebugPrintAt(screen, "YOU", 430, 350)
	}

	if g.state == StateDealerTurn {
		ebitenutil.DebugPrintAt(screen, "DEALER TURN", 430, 330)
	}

	if g.state == StateGameOver {
		if g.playerWon == true {
			ebitenutil.DebugPrintAt(screen, "YOU WIN", 430, 330)
		} else {
			ebitenutil.DebugPrintAt(screen, "YOU LOSE", 430, 330)
		}
		ebitenutil.DebugPrintAt(screen, "RETURN PRESS R", 430, 350)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) statusGameOver() bool {
	if g.dealerHealth <= 0 {
		g.playerWon = true
		g.state = StateGameOver
		return true
	}

	if g.playerHealth <= 0 {
		g.playerWon = false
		g.state = StateGameOver
		return true
	}

	return false
}

func (g *Game) restart() {
	if g.state == StateGameOver {
		g.playerHealth = 3
		g.dealerHealth = 3
		g.dealerTimer = 0
		g.shells = []Shell{
			{true},
			{false},
			{false},
			{true},
			{true},
			{true},
			{true},
		}
		g.playerWon = false
		g.state = StateMenu
	}
}

func newGame() *Game {
	return &Game{
		playerHealth: 3,
		dealerHealth: 3,
		dealerTimer:  0,
		shells: []Shell{
			{true},
			{false},
			{false},
			{true},
			{true},
			{true},
			{true},
		},
		playerWon: false,
		state:     StateMenu,
		startButton: Button{
			X:      400,
			Y:      300,
			Width:  160,
			Height: 50,
			Text:   "Start",
		},
	}

}
