package main

import (
	"math/rand"
	"time"
	//"strconv"
	"github.com/notnil/chess"
	//"github.com/notnil/chess/uci"
)

func main() {
	
	rand.Seed(time.Now().UTC().UnixNano())
	game := chess.NewGame()
	var white Player = RandomPlayer()
	var black Player = RandomPlayer()
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
