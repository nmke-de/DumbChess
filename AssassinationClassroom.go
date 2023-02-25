package main

import "github.com/notnil/chess"

// Same as SuicideKing, but all pieces can be moved.
type AssassinationClassroom bool

func (this AssassinationClassroom) move(game *chess.Game) *chess.Move {
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
	var chosen *chess.Move
	ceval := 85
	for _, m := range valid {
		eval := KingDist(m.S2(), sq)*10 + ManhattanDist(m.S2(), sq)
		if eval < ceval {
			ceval = eval
			chosen = m
		}
	}
	return chosen
}
func (this AssassinationClassroom) name() string {
	return "AssassinationClassroom"
}
