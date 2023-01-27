package main

import "github.com/notnil/chess"

type RandomMove bool
func (this RandomMove) move(game *chess.Game) *chess.Move {
	return TieBreak(game.ValidMoves())
}
func (this RandomMove) name() string {
	return "RandomMove"
}
