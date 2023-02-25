package main

import (
	"github.com/notnil/chess"
	"time"
)

type Careful bool

func (this Careful) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	opponent := CCCP(true)
	var best []*chess.Move
	var next *chess.Game
	ceval := 110
	t := time.Now()
	for _, m := range valid {
		// Evaluate danger level
		eval := 0
		// Try a move
		next = game.Clone()
		next.Move(m)
		// Predict an aggressive opponent response, if the predicted game has not ended yet.
		if next.Outcome() == chess.NoOutcome {
			predicted := opponent.move(next.Clone())
			// Raise danger level for check
			if predicted.HasTag(chess.Check) {
				eval += 100
			}
			// Raise danger level for capture
			switch next.Position().Board().Piece(predicted.S2()).Type() {
			case chess.Pawn:
				eval += 1
			case chess.Bishop, chess.Knight:
				eval += 3
			case chess.Rook:
				eval += 5
			case chess.Queen:
				eval += 9
			default:
				eval += 0
			}
		}
		// Pick a move
		if eval < ceval {
			ceval = eval
			best = []*chess.Move{m}
		} else if eval == ceval {
			best = append(best, m)
		}
		// Check for time limit. This is deemed necessary to stop evaluations that take too long.
		if time.Now().Sub(t) > time.Second {
			print(".")
			break
		}
	}
	return TieBreak(best)
}

func (this Careful) name() string {
	return "Careful"
}
