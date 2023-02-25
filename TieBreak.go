package main

import (
	"github.com/notnil/chess"
	"math/rand"
)

func TieBreak(moves []*chess.Move) *chess.Move {
	return moves[rand.Intn(len(moves))]
}
