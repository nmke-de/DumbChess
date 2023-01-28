package main

import "github.com/notnil/chess"

type EdgierEdgelord bool

func (this EdgierEdgelord) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	var best []*chess.Move
	ceval := 5
	for _, m := range valid {
		eval := Dist2Edge(m.S2())
		if m.HasTag(chess.Capture) {
			eval -= 1
		}
		if eval < ceval {
			ceval = eval
			best = []*chess.Move{m}
		} else if eval == ceval {
			best = append(best, m)
		}
	}
	if len(best) == 0 {
		return TieBreak(valid)
	}
	return TieBreak(best)
}

func (this EdgierEdgelord) name() string {
	return "Even edgier (read: more edgy) edgelord"
}
