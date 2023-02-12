package main

import "github.com/notnil/chess"

// Amount of moves a rook needs to move between two squares (over an empty board)
func RookDist(s1 chess.Square, s2 chess.Square) int {
	dist := 2
	if s1.Rank() == s2.Rank() {
		dist -= 1
	}
	if s1.File() == s2.File() {
		dist -= 1
	}
	return dist
}
