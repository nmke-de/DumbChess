package main

import "math/rand"

// Choose player randomly
func RandomPlayer() Player {
	player_list := []Player{
		AssassinationClassroom(true),
		Burglar(true),
		CCCP(true),
		CheckCapture(true),
		Coward(true),
		FirstValid(true),
		Huddle(true),
		Human(true),
		KnightlyOrder(true),
		LastValid(true),
		MaxOpponentMoves(true),
		MinOpponentMoves(true),
		Pacifist(true),
		PawnsFirst(true),
		RandomMove(true),
		Schizophrenia(true),
		//Stockfish(true),
		SuicideKing(true),
		Swarm(true),
	}
	return player_list[rand.Intn(len(player_list))]
}
