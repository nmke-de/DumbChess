package main

import (
	"os"
)

func main() {
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
	outcome := match(white, black)
	println(outcome.String())
}
