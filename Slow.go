package main

import "github.com/notnil/chess"

// Move pieces as far as possible
type Slow bool

func (this Slow) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	var chosen *chess.Move
	ceval := 16
	for _, m := range valid {
		eval := KingDist(m.S1(), m.S2())
		if eval <= ceval {
			ceval = eval
			chosen = m
		}
	}
	return chosen
}

func (this Slow) name() string {
	return "Slow"
}
