package main

import "github.com/notnil/chess"

type Coward bool

func (this Coward) move(game *chess.Game) *chess.Move {
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
	ceval := 0
	for _, m := range valid {
		eval := KingDist(m.S2(), sq)*10 + ManhattanDist(m.S2(), sq)
		if eval > ceval {
			ceval = eval
			chosen = m
		}
	}
	return chosen
}
func (this Coward) name() string {
	return "Coward!"
}
