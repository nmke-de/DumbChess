package main

import (
	"fmt"
	"github.com/notnil/chess"
)

/*
Human player. Get valid user input and return it.
*/
type Human bool
func (this Human) move(game *chess.Game) *chess.Move {
	var an chess.AlgebraicNotation
	mh := game.Moves()
	if len(mh) > 0 {
		print("Played by adversary: ")
		println(an.Encode(game.Position(), mh[len(mh)-1]))
	}
	println(game.Position().Board().Draw())
	var input string
	for {
		print("Your move: ")
		n, err := fmt.Scanln(&input)
		if err != nil || n == 0 {
			println("Invalid input.")
			continue
		}
		err = game.MoveStr(input)
		if err == nil {
			return game.Moves()[len(game.Moves()) - 1]
		}
		println("Invalid input or move.")
	}
}
func (this Human) name() string {
	return "Human player"
}
