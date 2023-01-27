package main

import "github.com/notnil/chess"

type CCCP bool
func (this CCCP) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	var ismate *chess.Game
	chosen := valid[0]
	color := game.Position().Board().Piece(chosen.S1()).Color()
	ceval := 0
	pushpiece := true
	for _, m := range valid {
		eval := 0
		// Check for checkmate
		ismate = game.Clone()
		ismate.Move(m)
		if ismate.Method() == 1 {
			return m
		}
		// Check for check
		if m.HasTag(chess.Check) {
			if pushpiece {
				pushpiece = false
				ceval = 0
			}
			eval += 10
		}
		// Check for capture
		if pushpiece && game.Position().Board().Piece(m.S2()).Type() != chess.NoPieceType {
			pushpiece = false
			ceval = 0
		}
		switch game.Position().Board().Piece(m.S2()).Type() {
			case chess.Pawn: eval += 1
			case chess.Bishop, chess.Knight: eval += 3
			case chess.Rook: eval += 5
			case chess.Queen: eval += 9
			default: eval += 0
		}
		// Otherwise push as deep as possible
		if pushpiece && eval == 0 {
			rank := Rank2Int(m.S2().Rank())
			switch color {
				case chess.White: eval += rank
				case chess.Black: eval += 9 - rank
			}
		}
		if eval > ceval {
			ceval = eval
			chosen = m
		}
	}
	return chosen
}
func (this CCCP) name() string {
	return "CCCP"
}
