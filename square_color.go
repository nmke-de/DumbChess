package main

import "github.com/notnil/chess"

func square_color(sq chess.Square) chess.Color {
	f := File2Int(sq.File())
	r := Rank2Int(sq.Rank())
	if r % 2 == f % 2 {
		return chess.Black
	}
	return chess.White
}
