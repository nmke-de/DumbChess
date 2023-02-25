package main

import (
	"github.com/notnil/chess"
	"github.com/notnil/chess/uci"
	"math/rand"
	"time"
)

// (stateless) UCI player. Requires the related engine (i.e. Stockfish) to be installed.
type UCI struct {
	exec string
}

func newUCI(exec string) UCI {
	return UCI{exec}
}

func (this UCI) move(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	eng, err := uci.New(this.exec)
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

func (this UCI) name() string {
	return this.exec
}
