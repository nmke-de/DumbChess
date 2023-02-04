package main

import (
	"github.com/notnil/chess"
)

type NotSinglePiece struct {
	sq *chess.Square
}

func newNotSinglePiece() NotSinglePiece {
	default_square := chess.E4
	return NotSinglePiece{&default_square}
}

func (this NotSinglePiece) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	var chosen *chess.Move
	var best []*chess.Move
	ceval := 0
	for _, m := range valid {
		eval := 0
		if ManhattanDist(m.S1(), *this.sq) == 0 {
			eval += 10
		}
		if m.HasTag(chess.Check) {
			eval += 5
		}
		if m.HasTag(chess.Capture) {
			eval += 2
		}
		if KingDist(*this.sq, m.S2()) > 3 {
			eval += 1
		}
		if KingDist(*this.sq, m.S1()) < 2 {
			eval += 1
		}
		// Pick a move
		if eval < ceval {
			ceval = eval
			best = []*chess.Move{m}
		} else if eval == ceval {
			best = append(best, m)
		}
	}
	if len(best) == 0 {
		chosen = TieBreak(valid)
	} else {
		chosen = TieBreak(best)
	}
	*this.sq = chosen.S2()
	return chosen
}

func (this NotSinglePiece) name() string {
	return "NotSinglePiece"
}
