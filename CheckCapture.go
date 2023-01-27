package main

import "github.com/notnil/chess"

type CheckCapture bool
func (this CheckCapture) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	chosen := valid[0]
	chosen_eval := 0
	for _, m := range valid {
		if chosen_eval < 2 && m.HasTag(chess.Check) {
			return m
		} else if chosen_eval < 1 && m.HasTag(chess.Capture) {
			chosen_eval = 1
			chosen = m
		}
	}
	return chosen
}
func (this CheckCapture) name() string {
	return "CheckCapture"
}
