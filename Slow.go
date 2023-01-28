package main

import "github.com/notnil/chess"

// Move pieces not as far as possible
type Slow bool

func (this Slow) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	var best []*chess.Move
	ceval := 16
	for _, m := range valid {
		eval := KingDist(m.S1(), m.S2())
		if eval < ceval {
			ceval = eval
			best = []*chess.Move{m}
		} else if eval == ceval {
			best = append(best, m)
		}
	}
	return TieBreak(best)
}

func (this Slow) name() string {
	return "Slow"
}
