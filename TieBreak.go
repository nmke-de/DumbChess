package main

import (
	"math/rand"
	"github.com/notnil/chess"
)

func TieBreak(moves []*chess.Move) *chess.Move {
	return moves[rand.Intn(len(moves))]
}
