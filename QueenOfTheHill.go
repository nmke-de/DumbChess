package main

import "github.com/notnil/chess"

type QueenOfTheHill bool

func (this QueenOfTheHill) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	var best []*chess.Move
	ceval := 0
	for _, m := range valid {
		eval := Dist2Edge(m.S2())
		// Favour the queen(s(?))
		if game.Position().Board().Piece(m.S1()).Type() == chess.Queen {
			eval += 2
		}
		// Favour promoting to queen
		if m.Promo() == chess.Queen {
			eval += 1
		}
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

func (this QueenOfTheHill) name() string {
	return "Queen of The Hill"
}
