package main

import "github.com/notnil/chess"

type Player interface {
	move(game *chess.Game) *chess.Move
	name() string
}
