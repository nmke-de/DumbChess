package main

import "github.com/notnil/chess"

// Manhattan Distance
func ManhattanDist(s1 chess.Square, s2 chess.Square) int {
	dist_r := abs(Rank2Int(s1.Rank()) - Rank2Int(s2.Rank()))
	dist_f := abs(File2Int(s1.File()) - File2Int(s2.File()))
	return dist_r + dist_f
}
