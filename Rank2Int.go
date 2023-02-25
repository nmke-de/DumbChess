package main

import "github.com/notnil/chess"

// convert rank to int
func Rank2Int(r chess.Rank) int {
	switch r {
	case chess.Rank1:
		return 1
	case chess.Rank2:
		return 2
	case chess.Rank3:
		return 3
	case chess.Rank4:
		return 4
	case chess.Rank5:
		return 5
	case chess.Rank6:
		return 6
	case chess.Rank7:
		return 7
	case chess.Rank8:
		return 8
	default:
		return 0
	}
}
