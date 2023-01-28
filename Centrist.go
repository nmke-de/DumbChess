package main

import "github.com/notnil/chess"

type Centrist bool

func (this Centrist) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	var best []*chess.Move
	ceval := 0
	for _, m := range valid {
		eval := Dist2Edge(m.S2())
		if eval > ceval {
			ceval = eval
			best = []*chess.Move{m}
		} else if eval == ceval {
			best = append(best, m)
		}
	}
	if len(best) == 0 {
		return TieBreak(valid)
	}
	return TieBreak(best)
}

func (this Centrist) name() string {
	return "Centrist"
}
