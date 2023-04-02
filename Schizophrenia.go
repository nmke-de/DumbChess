package main

import "github.com/notnil/chess"

type Schizophrenia bool

func (this Schizophrenia) move(game *chess.Game) *chess.Move {
	var p Player
	h := Human(true)
	for p == nil || p.name() == h.name() || p.name() == this.name() {
		p = RandomPlayer()
	}
	return p.move(game)
}

func (this Schizophrenia) name() string {
	return "Schizophrenia"
}
