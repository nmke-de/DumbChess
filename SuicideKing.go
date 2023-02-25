package main

import "github.com/notnil/chess"

type SuicideKing bool

func (this SuicideKing) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	// Determine color and target king
	color := game.Position().Board().Piece(valid[0].S1()).Color()
	var target chess.Piece
	if color == chess.White {
		target = chess.BlackKing
	} else {
		target = chess.WhiteKing
	}
	// Determine position of target king
	var sq chess.Square
	for i := 0; i < 64; i++ {
		sq = chess.Square(i)
		if game.Position().Board().Piece(sq) == target {
			break
		}
	}
	// Evaluate all valid moves
	var best []*chess.Move
	ceval := 85
	for _, m := range valid {
		// Only interested in the king
		if game.Position().Board().Piece(m.S1()).Type() != chess.King {
			continue
		}
		eval := KingDist(m.S2(), sq)*10 + ManhattanDist(m.S2(), sq)
		if eval < ceval {
			ceval = eval
			best = []*chess.Move{m}
		} else if eval == ceval {
			best = append(best, m)
		}
	}
	// Choose randomly if no move was interesting
	if len(best) == 0 {
		return TieBreak(valid)
	}
	return TieBreak(best)
}

func (this SuicideKing) name() string {
	return "SuicideKing"
}
