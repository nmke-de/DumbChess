package main

import "github.com/notnil/chess"

type KnightlyOrder bool

func (this KnightlyOrder) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	var best []*chess.Move
	ceval := 0
	for _, m := range valid {
		eval := 0
		if m.Promo() == chess.Knight {
			eval += 1
		}
		if game.Position().Board().Piece(m.S1()).Type() == chess.Knight {
			eval += 3
		}
		if game.Position().Board().Piece(m.S2()).Type() == chess.Knight {
			eval -= 1
		} else if m.HasTag(chess.Capture) {
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
func (this KnightlyOrder) name() string {
	return "KnightlyOrder"
}
