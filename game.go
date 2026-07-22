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
	StateRound1
	StateRound2
	StateRound3
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

	startButton  Button
	exitButton   Button
	playerButton Button
	dealerButton Button

	round int
}

func (g *Game) Update() error {
	switch g.state {
	case StateMenu:
		if g.startButton.isClicked() {
			g.state = StateRound1
		} else if g.exitButton.isClicked() {
			return ebiten.Termination
		}

	case StatePlayerTurn:
		if g.dealerButton.isClicked() {
			shell := g.takeShell()
			if shell.Live {
				g.dealerHealth--
				g.state = StateDealerTurn
			} else {
				g.state = StateDealerTurn
			}

		} else if g.playerButton.isClicked() {
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

	case StateRound1:
		g.playerHealth = 2
		g.dealerHealth = 2
		g.state = StatePlayerTurn

	case StateRound2:
		g.playerHealth = 4
		g.dealerHealth = 4
		g.state = StatePlayerTurn
		g.loadShells()

	case StateRound3:
		g.playerHealth = 6
		g.dealerHealth = 6
		g.state = StatePlayerTurn
		g.loadShells()
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
		g.exitButton.draw(screen)
	}

	if g.state == StatePlayerTurn {
		g.dealerButton.draw(screen)
		g.playerButton.draw(screen)
	}

	if g.state == StateDealerTurn {
		ebitenutil.DebugPrintAt(screen, "DEALER TURN", (screenWidth-len("DEALER TURN")*debugCharWidth)/2, 350)
	}

	if g.round == 0 && g.state != StateMenu {
		ebitenutil.DebugPrintAt(screen, "ROUND1", (screenWidth-len("ROUND1")*debugCharWidth)/2, 0)
	}

	if g.round == 1 {
		ebitenutil.DebugPrintAt(screen, "ROUND2", (screenWidth-len("ROUND2")*debugCharWidth)/2, 0)
	}

	if g.round == 2 {
		ebitenutil.DebugPrintAt(screen, "ROUND3", (screenWidth-len("ROUND3")*debugCharWidth)/2, 0)
	}

	if g.state == StateGameOver {
		if g.playerWon == true {
			ebitenutil.DebugPrintAt(screen, "YOU WIN", (screenWidth-len("YOU WIN")*debugCharWidth)/2, 350)
		} else {
			ebitenutil.DebugPrintAt(screen, "YOU LOSE", (screenWidth-len("YOU LOSE")*debugCharWidth)/2, 350)
		}
		ebitenutil.DebugPrintAt(screen, "RETURN PRESS R", (screenWidth-len("RETURN PRESS R")*debugCharWidth)/2, 375)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) statusGameOver() bool {

	if g.dealerHealth <= 0 && g.round == 0 {
		g.round++
		g.playerWon = true
		g.state = StateRound2
		return true
	} else if g.dealerHealth <= 0 && g.round == 1 {
		g.round++
		g.playerWon = true
		g.state = StateRound3
		return true
	} else if g.dealerHealth <= 0 && g.round == 2 {
		g.round++
		g.playerWon = true
		g.state = StateGameOver
		return true
	}

	if g.playerHealth <= 0 && g.round == 0 {
		g.round++
		g.playerWon = false
		g.state = StateGameOver
		return false
	} else if g.playerHealth <= 0 && g.round == 1 {
		g.round++
		g.playerWon = false
		g.state = StateGameOver
		return false
	} else if g.playerHealth <= 0 && g.round == 2 {
		g.round++
		g.playerWon = false
		g.state = StateGameOver
		return false
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
		g.startButton = Button{
			X:    (screenWidth - len(startText)*debugCharWidth) / 2,
			Y:    350,
			Text: startText,
		}
		g.exitButton = Button{
			X:    (screenWidth - len(exitText)*debugCharWidth) / 2,
			Y:    375,
			Text: exitText,
		}
		g.playerButton = Button{
			X:    (screenWidth - len(playerText)*debugCharWidth) / 2,
			Y:    375,
			Text: playerText,
		}
		g.dealerButton = Button{
			X:    (screenWidth - len(dealerText)*debugCharWidth) / 2,
			Y:    350,
			Text: dealerText,
		}
		g.round = 0
	}
}

const (
	startText  = "START"
	exitText   = "EXIT"
	playerText = "YOU"
	dealerText = "DEALER"
)

func newGame() *Game {
	return &Game{
		playerHealth: 3,
		dealerHealth: 3,
		dealerTimer:  0,
		shells: []Shell{
			{true},
			{false},
			{true},
			{false},
		},
		playerWon: false,
		state:     StateMenu,
		startButton: Button{
			X:    (screenWidth - len(startText)*debugCharWidth) / 2,
			Y:    350,
			Text: startText,
		},
		exitButton: Button{
			X:    (screenWidth - len(exitText)*debugCharWidth) / 2,
			Y:    375,
			Text: exitText,
		},
		playerButton: Button{
			X:    (screenWidth - len(playerText)*debugCharWidth) / 2,
			Y:    375,
			Text: playerText,
		},
		dealerButton: Button{
			X:    (screenWidth - len(dealerText)*debugCharWidth) / 2,
			Y:    350,
			Text: dealerText,
		},
		round: 0,
	}

}
