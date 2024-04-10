package main

import "github.com/notnil/chess"

func PieceValue(t chess.PieceType, considerKing bool) int {
	switch t {
	case chess.Pawn:
		return 1
	case chess.Bishop, chess.Knight:
		return 3
	case chess.Rook:
		return 5
	case chess.Queen:
		return 9
	case chess.King:
		if considerKing {
			return 10
		} else {
			return 0
		}
	default:
		return 0
	}
}
