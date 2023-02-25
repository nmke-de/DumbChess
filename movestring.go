package main

import "github.com/notnil/chess"

func movestring(p *chess.Position, m *chess.Move) string {
	var an chess.AlgebraicNotation
	var piece string
	switch p.Board().Piece(m.S2()).Type() {
	case chess.Pawn:
		piece = ""
	case chess.Knight:
		piece = "N"
	case chess.Bishop:
		piece = "B"
	case chess.Rook:
		piece = "R"
	case chess.Queen:
		piece = "Q"
	case chess.King:
		piece = "K"
	}
	return piece + an.Encode(p, m)
}
