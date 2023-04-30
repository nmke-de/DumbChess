package main

import "github.com/notnil/chess"

type KingOfTheHill bool

func (this KingOfTheHill) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	// Determine color and target king
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
	ceval := 0
	for _, m := range valid {
		eval := 0
		// Only interested in the king (unless he can't move)
		if game.Position().Board().Piece(m.S1()).Type() != chess.King {
			if abs(File2Int(m.S1().File()) - 5) < 2 {
				eval = 1
			}
		} else {
			// Only move the king if he is not king of the hill already
			if Dist2Edge(m.S1()) == 3 {
				continue
			}
			eval = Dist2Edge(m.S2()) + 2
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

func (this KingOfTheHill) name() string {
	return "King of The Hill"
}
