package main

import "github.com/notnil/chess"

// Move pieces as far as possible
type Traveller bool

func (this Traveller) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	var best []*chess.Move
	ceval := 0
	for _, m := range valid {
		eval := ManhattanDist(m.S1(), m.S2())
		if eval > ceval {
			ceval = eval
			best = []*chess.Move{m}
		} else if eval == ceval {
			best = append(best, m)
		}
	}
	return TieBreak(best)
}

func (this Traveller) name() string {
	return "Traveller"
}
