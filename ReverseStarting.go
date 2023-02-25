package main

import "github.com/notnil/chess"

type ReverseStarting struct {}

func (this ReverseStarting) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	// Determine color
	color := game.Position().Board().Piece(valid[0].S1()).Color()
	var target chess.Piece
	if color == chess.White {
		target = chess.WhiteKing
	} else {
		target = chess.BlackKing
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
	ceval := 19
	for _, m := range valid {
		eval := 0
		switch game.Position().Board().Piece(m.S1()) {
			case chess.BlackPawn:
				// Ignore if target square has been reached
				if Rank2Int(m.S1().Rank()) == 2 {
					continue
				}
				eval = abs(Rank2Int(m.S2().Rank()) - 2)
				if eval < 0 {
					eval = 20
				}
			case chess.WhitePawn:
				// Ignore if target square has been reached
				if Rank2Int(m.S1().Rank()) == 7 {
					continue
				}
				eval = abs(7 - Rank2Int(m.S2().Rank()))
				if eval < 0 {
					eval = 20
				}
			case chess.BlackKnight:
				// Ignore if target square has been reached
				if m.S1() == chess.B1 || m.S2() == chess.G1 {
					continue
				}
				eval = min(KnightDist(m.S2(), chess.B1), KnightDist(m.S2(), chess.G1))
			case chess.WhiteKnight:
				// Ignore if target square has been reached
				if m.S1() == chess.B8 || m.S2() == chess.G8 {
					continue
				}
				eval = min(KnightDist(m.S2(), chess.B8), KnightDist(m.S2(), chess.G8))
			case chess.BlackBishop:
				// Ignore if target square has been reached
				if m.S1() == chess.C1 || m.S2() == chess.F1 {
					continue
				}
				if square_color(m.S2()) == chess.Black {
					eval = BishopDist(m.S2(), chess.C1)
				} else {
					eval = BishopDist(m.S2(), chess.F1)
				}
			case chess.WhiteBishop:
				// Ignore if target square has been reached
				if m.S1() == chess.C8 || m.S2() == chess.F8 {
					continue
				}
				if square_color(m.S2()) == chess.Black {
					eval = BishopDist(m.S2(), chess.F8)
				} else {
					eval = BishopDist(m.S2(), chess.C8)
				}
			case chess.BlackRook:
				// Ignore if target square has been reached
				if m.S1() == chess.A1 || m.S2() == chess.H1 {
					continue
				}
				eval = min(ManhattanDist(m.S2(), chess.A1), ManhattanDist(m.S2(), chess.H1))
			case chess.WhiteRook:
				// Ignore if target square has been reached
				if m.S1() == chess.A8 || m.S2() == chess.H8 {
					continue
				}
				eval = min(ManhattanDist(m.S2(), chess.A8), ManhattanDist(m.S2(), chess.H8))
			case chess.BlackQueen:
				// Ignore if target square has been reached
				if m.S1() == chess.D1 {
					continue
				}
				eval = KingDist(m.S2(), chess.D1)
			case chess.WhiteQueen:
				// Ignore if target square has been reached
				if m.S1() == chess.D8 {
					continue
				}
				eval = KingDist(m.S2(), chess.D8)
			case chess.BlackKing:
				// Ignore if target square has been reached
				if m.S1() == chess.E1 {
					continue
				}
				eval = KingDist(m.S2(), chess.E1)
			case chess.WhiteKing:
				// Ignore if target square has been reached
				if m.S1() == chess.E8 {
					continue
				}
				eval = KingDist(m.S2(), chess.E8)
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

func (this ReverseStarting) name() string {
	return "Reverse starting positions"
}
