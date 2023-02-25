package main

import "github.com/notnil/chess"

type Rookies struct{}

func (this Rookies) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	// Determine color and target king
	color := game.Position().Board().Piece(valid[0].S1()).Color()
	var target chess.Piece
	if color == chess.White {
		target = chess.BlackKing
	} else {
		target = chess.WhiteKing
	}
	// Determine position of target king
	var sq chess.Square
	for i := 0; i < 64; i++ {
		sq = chess.Square(i)
		if game.Position().Board().Piece(sq) == target {
			break
		}
	}
	var best []*chess.Move
	ceval := 0
	for _, m := range valid {
		eval := 0
		if m.Promo() == chess.Rook {
			eval += 2 * (3 - RookDist(m.S2(), sq))
		}
		switch game.Position().Board().Piece(m.S1()).Type() {
		case chess.Rook:
			eval += 2 * (3 - RookDist(m.S2(), sq))
		case chess.Pawn:
			eval += 4 - Dist2Edge(m.S1())
		}
		switch game.Position().Board().Piece(m.S2()).Type() {
		case chess.NoPieceType:
			eval += 0
		case chess.Rook:
			eval += 1
		default:
			eval += 2
		}
		if eval > ceval {
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

func (this Rookies) name() string {
	return "Rookies"
}
