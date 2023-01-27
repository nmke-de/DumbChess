package main

import "github.com/notnil/chess"

type LastValid bool
func (this LastValid) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	return valid[len(valid) - 1]
}
func (this LastValid) name() string {
	return "LastValid"
}
