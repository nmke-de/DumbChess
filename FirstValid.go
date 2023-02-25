package main

import "github.com/notnil/chess"

type FirstValid bool

func (this FirstValid) move(game *chess.Game) *chess.Move {
	return game.ValidMoves()[0]
}
func (this FirstValid) name() string {
	return "FirstValid"
}
