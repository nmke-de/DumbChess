package main

import "math/rand"

// Choose player randomly
func RandomPlayer() Player {
	return playerList[rand.Intn(len(playerList))]
}
