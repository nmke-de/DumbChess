package main

import (
	"math/rand"
	"time"
	"github.com/notnil/chess"
)

// Play match. Return outcome.
func match(white, black Player) chess.Outcome {
	rand.Seed(time.Now().UTC().UnixNano())
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
	return game.Outcome()
}