package main

import "math/rand/v2"

type Shell struct {
	Live bool
}

func (g *Game) loadShells() {
	if g.round == 0 {
		g.shells = []Shell{
			{true},
			{false},
			{true},
			{false},
		}
	}
	if g.round == 1 {
		g.shells = []Shell{
			{true},
			{false},
			{true},
			{false},
			{true},
			{false},
			{true},
			{false},
		}

		rand.Shuffle(len(g.shells), func(i, j int) {
			g.shells[i], g.shells[j] = g.shells[j], g.shells[i]
		})
	}

	if g.round == 2 {
		g.shells = []Shell{
			{true},
			{false},
			{true},
			{false},
			{true},
			{false},
			{true},
			{false},
			{true},
			{false},
			{true},
			{false},
		}

		rand.Shuffle(len(g.shells), func(i, j int) {
			g.shells[i], g.shells[j] = g.shells[j], g.shells[i]
		})
	}

}

func (g *Game) takeShell() Shell {
	if len(g.shells) == 0 {
		g.loadShells()
	}

	shell := g.shells[0]
	g.shells = g.shells[1:]

	return shell
}
