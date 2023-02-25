package main

import "github.com/notnil/chess"

type MaxOpponentMoves bool

func (this MaxOpponentMoves) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	var chosen *chess.Move
	var next *chess.Game
	ceval := -1
	for _, m := range valid {
		next = game.Clone()
		next.Move(m)
		eval := len(next.ValidMoves())
		if eval > ceval {
			ceval = eval
			chosen = m
		}
	}
	return chosen
}
func (this MaxOpponentMoves) name() string {
	return "MaxOpponentMoves"
}
