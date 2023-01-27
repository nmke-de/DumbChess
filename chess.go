package main

import (
	"math/rand"
	"time"
	//"strconv"
	"github.com/notnil/chess"
	//"github.com/notnil/chess/uci"
)

// Same as SuicideKing, but all pieces can be moved.
type AssassinationClassroom bool
func (this AssassinationClassroom) move(game *chess.Game) *chess.Move {
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
	// Evaluate all valid moves
	var chosen *chess.Move
	ceval := 85
	for _, m := range valid {
		eval := KingDist(m.S2(), sq) * 10 + ManhattanDist(m.S2(), sq)
		if eval < ceval {
			ceval = eval
			chosen = m
		}
	}
	return chosen
}
func (this AssassinationClassroom) name() string {
	return "AssassinationClassroom"
}

type Coward bool
func (this Coward) move(game *chess.Game) *chess.Move {
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
	// Evaluate all valid moves
	var chosen *chess.Move
	ceval := 0
	for _, m := range valid {
		eval := KingDist(m.S2(), sq) * 10 + ManhattanDist(m.S2(), sq)
		if eval > ceval {
			ceval = eval
			chosen = m
		}
	}
	return chosen
}
func (this Coward) name() string {
	return "Coward!"
}

type Huddle bool
func (this Huddle) move(game *chess.Game) *chess.Move {
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
	// Evaluate all valid moves
	var chosen *chess.Move
	ceval := 85
	for _, m := range valid {
		// Only interested in not the king
		if game.Position().Board().Piece(m.S1()).Type() == chess.King {
			continue
		}
		eval := KingDist(m.S2(), sq)
		if eval < ceval {
			ceval = eval
			chosen = m
		}
	}
	// Choose randomly if no move was interesting
	if chosen == nil {
		return valid[rand.Intn(len(valid))]
	}
	return chosen
}
func (this Huddle) name() string {
	return "Huddle"
}

type Swarm bool
func (this Swarm) move(game *chess.Game) *chess.Move {
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
	// Evaluate all valid moves
	var chosen *chess.Move
	ceval := 0
	for _, m := range valid {
		// Only interested in not the king
		if game.Position().Board().Piece(m.S1()).Type() == chess.King {
			continue
		}
		eval := KingDist(m.S2(), sq)
		if eval > ceval {
			ceval = eval
			chosen = m
		}
	}
	// Choose randomly if no move was interesting
	if chosen == nil {
		return valid[rand.Intn(len(valid))]
	}
	return chosen
}
func (this Swarm) name() string {
	return "Swarm"
}


func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	game := chess.NewGame()
	var white Player = Human(true)
	var black Player = MinOpponentMoves(true)
	// Introduce players
	print("White: ")
	println(white.name())
	print("Black: ")
	println(black.name())
	// Determine active player
	var current Player = white
	for game.Outcome() == chess.NoOutcome {
		game.Move(current.move(game.Clone()))
		if current == white {
			current = black
		} else {
			current = white
		}
	}
	println(game.Position().Board().Draw())
	print("Outcome: ")
	println(game.Outcome())
	print("Cause: ")
	println(EvalMethod(game.Method()))
	print("Moves: ")
	println(game.String())
}
