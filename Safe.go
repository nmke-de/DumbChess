package main

import "github.com/notnil/chess"

type Safe struct {}

func (this Safe) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	var best []*chess.Move
	var next *chess.Game
	ceval := 20
	for _, m := range valid {
		// Danger level
		eval := 0
		// Try a move
		next = game.Clone()
		next.Move(m)
		// See if piece can be captured
		if next.Outcome() == chess.NoOutcome {
			for _, predicted := range next.ValidMoves() {
				// Consider check
				if predicted.HasTag(chess.Check) {
					eval += 10
				}
				if predicted.S2() == m.S2() {
					switch game.Position().Board().Piece(m.S1()).Type() {
						case chess.Pawn: eval += 1
						case chess.Bishop, chess.Knight: eval += 3
						case chess.Rook: eval += 5
						case chess.Queen: eval += 9
						default: eval += 0
					}
					break
				}
			}
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
		return TieBreak(valid)
	}
	return TieBreak(best)
}

func (this Safe) name() string {
	return "Safe"
}
