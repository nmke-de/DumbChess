package main

import "github.com/notnil/chess"

type MinOpponentMoves bool

func (this MinOpponentMoves) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	var chosen *chess.Move
	var next *chess.Game
	ceval := 8192
	for _, m := range valid {
		next = game.Clone()
		next.Move(m)
		eval := len(next.ValidMoves())
		if eval < ceval {
			ceval = eval
			chosen = m
		}
	}
	return chosen
}

func (this MinOpponentMoves) name() string {
	return "MinOpponentMoves"
}
