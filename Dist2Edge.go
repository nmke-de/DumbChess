package main

import "github.com/notnil/chess"

func Dist2Edge(sq chess.Square) int {
	r := Rank2Int(sq.Rank())
	f := File2Int(sq.File())
	return min(min(r-1, 8-r), min(f-1, 8-f))
}
