package main

import "github.com/notnil/chess"

type OppositeColor struct{}

func (this OppositeColor) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	color := game.Position().Board().Piece(valid[0].S1()).Color()
	var best []*chess.Move
	for _, m := range valid {
		if square_color(m.S2()) != color {
			best = append(best, m)
		}
	}
	if len(best) == 0 {
		return TieBreak(valid)
	}
	return TieBreak(best)
}

func (this OppositeColor) name() string {
	return "OppositeColor"
}
