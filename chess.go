package main

import (
	"math/rand"
	"time"
	//"strconv"
	"github.com/notnil/chess"
	//"github.com/notnil/chess/uci"
)

type MinOpponentMoves bool
func (this MinOpponentMoves) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	var chosen *chess.Move
	var next *chess.Game
	ceval := 8192
	for _, m := range valid {
		next = game.Clone()
		next.Move(m)
		eval := len(next.ValidMoves())
		if eval < ceval {
			ceval = eval
			chosen = m
		}
	}
	return chosen
}
func (this MinOpponentMoves) name() string {
	return "MinOpponentMoves"
}

type Pacifist bool
func (this Pacifist) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	var chosen *chess.Move
	chosen = valid[0]
	var ismate *chess.Game
	ceval := 0
	for _, m := range valid {
		eval := 0
		// Check for checkmate
		ismate = game.Clone()
		ismate.Move(m)
		if ismate.Method() == 1 {
			eval -= 110
		// Check for check
		} else if m.HasTag(chess.Check) {
			eval -= 100
		}
		// Check for capture
		if game.Position().Board().Piece(m.S2()).Type() != chess.NoPieceType {
			switch game.Position().Board().Piece(m.S2()).Type() {
				case chess.Pawn: eval -= 1
				case chess.Bishop, chess.Knight: eval -= 3
				case chess.Rook: eval -= 5
				case chess.Queen: eval -= 9
				default: eval -= 0
			}
		}
		if eval > ceval {
			ceval = eval
			chosen = m
		}
	}
	return chosen
}
func (this Pacifist) name() string {
	return "Pacifist"
}

/*type Stockfish bool
func (this Stockfish) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	eng, err := uci.New("stockfish")
	if err != nil {
		panic(err)
	}
	defer eng.Close()
	// initialize uci with new game
	if err := eng.Run(uci.CmdUCI, uci.CmdIsReady, uci.CmdUCINewGame); err != nil {
		panic(err)
	}
	cmdPos := uci.CmdPosition{Position: game.Position()}
	cmdGo := uci.CmdGo{MoveTime: time.Second / 100}
	if err := eng.Run(cmdPos, cmdGo); err != nil {
		return valid[rand.Intn(len(valid))]
	}
	move := eng.SearchResults().BestMove
	if err := game.Move(move); err != nil {
		return valid[rand.Intn(len(valid))]
	}
	return move
}
func (this Stockfish) name() string {
	return "Stockfish"
}*/

type MaxOpponentMoves bool
func (this MaxOpponentMoves) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	var chosen *chess.Move
	var next *chess.Game
	ceval := -1
	for _, m := range valid {
		next = game.Clone()
		next.Move(m)
		eval := len(next.ValidMoves())
		if eval > ceval {
			ceval = eval
			chosen = m
		}
	}
	return chosen
}
func (this MaxOpponentMoves) name() string {
	return "MaxOpponentMoves"
}

type SuicideKing bool
func (this SuicideKing) move(game *chess.Game) *chess.Move {
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
		// Only interested in the king
		if game.Position().Board().Piece(m.S1()).Type() != chess.King {
			continue
		}
		eval := KingDist(m.S2(), sq) * 10 + ManhattanDist(m.S2(), sq)
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
func (this SuicideKing) name() string {
	return "SuicideKing"
}

// Same as SuicideKing, but all pieces are to be moved.
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
