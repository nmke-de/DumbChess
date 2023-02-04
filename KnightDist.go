package main

import "github.com/notnil/chess"

func KnightDist (s1 chess.Square, s2 chess.Square) int {
	if s1 == s2 {
		return 0
	}
	dist := 1
	visited := make(map[chess.Square]bool)
	visited[s1] = true
	frontier := []chess.Square{s1}
	next_frontier := []chess.Square{}
	for ; dist < 7; dist++ {
		for _, s := range frontier {
			next := knight_moves(s)
			for _, next_square := range next {
				if next_square == s2 {
					return dist
				}
				if visited[next_square] {
					next_frontier = append(next_frontier, next_square)
					visited[next_square] = true
				}
			}
		}
		frontier = next_frontier
		next_frontier = []chess.Square{}
	}
	// Dummy
	return 0
}

func knight_moves (sq chess.Square) []chess.Square {
	f := File2Int(sq.File())
	r := Rank2Int(sq.Rank())
	var result []chess.Square
	if f - 2 > 0 {
		if r - 1 > 0 {
			// Subtract one to count from zero, not one
			result = append(result, chess.NewSquare(chess.File(f - 3), chess.Rank(r - 2)))
		}
		if r + 1 < 9 {
			result = append(result, chess.NewSquare(chess.File(f - 3), chess.Rank(r)))
		}
	}
	if r + 2 < 9 {
		if f - 1 > 0 {
			result = append(result, chess.NewSquare(chess.File(f - 2), chess.Rank(r + 1)))
		}
		if f + 1 < 9 {
			result = append(result, chess.NewSquare(chess.File(f), chess.Rank(r + 1)))
		}
	}
	if f + 2 < 9 {
		if r - 1 > 0 {
			result = append(result, chess.NewSquare(chess.File(f + 1), chess.Rank(r - 2)))
		}
		if r + 1 < 9 {
			result = append(result, chess.NewSquare(chess.File(f + 1), chess.Rank(r)))
		}
	}
	if r - 2 > 0 {
		if f - 1 > 0 {
			result = append(result, chess.NewSquare(chess.File(f - 2), chess.Rank(r - 3)))
		}
		if f + 1 < 9 {
			result = append(result, chess.NewSquare(chess.File(f), chess.Rank(r - 3)))
		}
	}
	return result
}
