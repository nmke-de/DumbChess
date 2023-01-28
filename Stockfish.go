package main


import (
	"math/rand"
	"time"
	"github.com/notnil/chess"
	"github.com/notnil/chess/uci"
)


// (stateless) Stockfish. Requires stockfish to be installed.
type Stockfish bool

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
}
