package main

import "math/rand"

// Choose player randomly
func RandomPlayer() Player {
	player_list := []Player{
		AssassinationClassroom(true),
		Burglar(true),
		CCCP(true),
		Centrist(true),
		CheckCapture(true),
		Coward(true),
		Edgelord(true),
		EdgierEdgelord(true),
		FirstValid(true),
		Huddle(true),
		Human(true),
		KingOfTheHill(true),
		KnightlyOrder(true),
		LastValid(true),
		MaxOpponentMoves(true),
		MinOpponentMoves(true),
		Pacifist(true),
		PawnsFirst(true),
		QueenOfTheHill(true),
		RandomMove(true),
		Schizophrenia(true),
		Slow(true),
		newUCI("stockfish"),
		SuicideKing(true),
		Swarm(true),
		Traveller(true),
	}
	return player_list[rand.Intn(len(player_list))]
}
