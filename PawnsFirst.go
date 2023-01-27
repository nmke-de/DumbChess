package main

import "github.com/notnil/chess"

type PawnsFirst bool
func (this PawnsFirst) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	var chosen *chess.Move
	ceval := 0
	for _, m := range valid {
		eval := 0
		if game.Position().Board().Piece(m.S1()).Type() == chess.Pawn {
			eval += 3
		}
		if game.Position().Board().Piece(m.S2()).Type() == chess.Pawn {
			eval -= 1
		} else if m.HasTag(chess.Capture) {
			eval += 1
		}
		if eval > ceval {
			ceval = eval
			chosen = m
		}
	}
	if chosen == nil {
		return TieBreak(valid)
	}
	return chosen
}
func (this PawnsFirst) name() string {
	return "PawnsFirst"
}
