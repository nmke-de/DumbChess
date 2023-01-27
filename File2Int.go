package main

import "github.com/notnil/chess"

// convert file to int
func File2Int(r chess.File) int {
	switch r {
		case chess.FileA: return 1
		case chess.FileB: return 2
		case chess.FileC: return 3
		case chess.FileD: return 4
		case chess.FileE: return 5
		case chess.FileF: return 6
		case chess.FileG: return 7
		case chess.FileH: return 8
		default: return 0
	}
}
