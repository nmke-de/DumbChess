package main

import "github.com/notnil/chess"

type Pacifist bool
func (this Pacifist) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	var chosen *chess.Move
	var ismate *chess.Game
	ceval := 120
	for _, m := range valid {
		eval := 0
		// Check for checkmate
		ismate = game.Clone()
		ismate.Move(m)
		if ismate.Method() == 1 {
			eval += 110
		// Check for check
		} else if m.HasTag(chess.Check) {
			eval += 100
		}
		// Check for capture
		if game.Position().Board().Piece(m.S2()).Type() != chess.NoPieceType {
			switch game.Position().Board().Piece(m.S2()).Type() {
				case chess.Pawn: eval += 1
				case chess.Bishop, chess.Knight: eval += 3
				case chess.Rook: eval += 5
				case chess.Queen: eval += 9
				default: eval += 0
			}
		}
		if eval < ceval {
			ceval = eval
			chosen = m
		}
	}
	return chosen
}
func (this Pacifist) name() string {
	return "Pacifist"
}
