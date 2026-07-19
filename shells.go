package main

import (
	"math/rand/v2"
)

type Shell struct {
	Live bool
}

func (g *Game) loadShells() {
	g.shells = []Shell{
		{true},
		{true},
		{false},
		{false},
	}

	rand.Shuffle(len(g.shells), func(i, j int) {
		g.shells[i], g.shells[j] = g.shells[j], g.shells[i]
	})
}

func (g *Game) takeShell() Shell {
	if len(g.shells) == 0 {
		g.loadShells()
	}

	shell := g.shells[0]
	g.shells = g.shells[1:]

	return shell
}
