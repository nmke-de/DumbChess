package main

import "github.com/notnil/chess"

// Amount of moves a bishop needs to move between two squares (over an empty board). Will return 65 if there is no path for a bishop to take.
func BishopDist(s1 chess.Square, s2 chess.Square) int {
	if s1 == s2 {
		return 0
	}
	if square_color(s1) != square_color(s2) {
		return 65
	}
	if abs(File2Int(s1.File())-File2Int(s2.File())) == abs(Rank2Int(s1.Rank())-Rank2Int(s2.Rank())) {
		return 1
	}
	return 2
}
