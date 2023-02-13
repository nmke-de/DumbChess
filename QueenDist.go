package main

import "github.com/notnil/chess"

// Amount of moves a queen needs to move between two squares (over an empty board).
func QueenDist(s1 chess.Square, s2 chess.Square) int {
	return min(RookDist(s1, s2), BishopDist(s1, s2))
}
