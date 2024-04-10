package main

import "github.com/notnil/chess"

type TopDown struct{}

func (this TopDown) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	var best []*chess.Move
	ceval := 0
	for _, m := range valid {
		eval := 0
		if m.HasTag(chess.Capture) || m.HasTag(chess.Check) {
			eval = PieceValue(game.Position().Board().Piece(m.S1()).Type(), true) - PieceValue(game.Position().Board().Piece(m.S2()).Type(), true) + 1
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

func (this TopDown) name() string {
	return "TopDown"
}
