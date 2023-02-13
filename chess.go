package main

import (
	"os"
	"strconv"
	"github.com/notnil/chess"
)

func main() {
	var white Player
	var black Player
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "-t" {
		// tournament mode
		whitePlayers := make(map[Player]int)
		blackPlayers := make(map[Player]int)
		for _, w := range playerList {
			if w.name() == "Human player" {
				continue
			}
			for _, b := range playerList {
				if b.name() == "Human player" {
					continue
				}
				switch match(w, b) {
					case chess.WhiteWon:
						whitePlayers[w] += 2
					case chess.BlackWon:
						blackPlayers[b] += 2
					default:
						whitePlayers[w] += 1
						blackPlayers[b] += 1
				}
			}
		}
		println("As white:\n")
		for player, score := range whitePlayers {
			println("" + strconv.Itoa(score) + "\t" + player.name())
		}
		println("\n\nAs black:\n")
		for player, score := range blackPlayers {
			println("" + strconv.Itoa(score) + "\t" + player.name())
		}
		return
	} else if len(args) > 0 && args[0] == "-l" {
		for _, p := range playerList {
			println(p.name())
		}
		return
	} else if len(args) > 0 {
		white = findPlayer(args[0])
	} else {
		white = RandomPlayer()
	}
	if len(args) > 1 {
		black = findPlayer(args[1])
	} else {
		black = RandomPlayer()
	}
	match(white, black)
}
