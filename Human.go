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
	if len(game.Moves()) > 0 {
		mh := game.Moves()
		print("Played by adversary: ")
		println(movestring(game.Position(), mh[len(mh)-1]))
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
