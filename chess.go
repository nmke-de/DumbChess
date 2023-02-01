package main

import (
	"math/rand"
	"time"
	"os"
	"github.com/notnil/chess"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var white Player
	var black Player
	args := os.Args[1:]
	if len(args) > 0 {
		white = findPlayer(args[0])
	} else {
		white = RandomPlayer()
	}
	if len(args) > 1 {
		black = findPlayer(args[1])
	} else {
		black = RandomPlayer()
	}
	game := chess.NewGame()
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
