package main

import "github.com/notnil/chess"

type Burglar bool
func (this Burglar) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	var best []*chess.Move
	ceval := 0
	for _, m := range valid {
		eval := -1
		if !m.HasTag(chess.Capture) && !m.HasTag(chess.Check) {
			continue
		}
		switch game.Position().Board().Piece(m.S2()).Type() {
			case chess.Pawn: eval = 1
			case chess.Bishop, chess.Knight: eval = 3
			case chess.Rook: eval = 5
			case chess.Queen: eval = 9
			default: eval = 0
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
func (this Burglar) name() string {
	return "Burglar"
}
