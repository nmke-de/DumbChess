package main

import "strings"

func findPlayer(name string) Player {
	for _, p := range playerList {
		if strings.Contains(strings.ToLower(p.name()), strings.ToLower(name)) {
			return p
		}
	}
	println("No such player. Choose player randomly.")
	return RandomPlayer()
}
