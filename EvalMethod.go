package main

import "github.com/notnil/chess"

// eval Method
func EvalMethod(this chess.Method) string {
	switch this {
		case 0: return "Not finished or no outcome"
		case 1: return "Checkmate"
		case 2: return "Resignation"
		case 3: return "Offered draw"
		case 4: return "Stalemate"
		case 5: return "Threefold repetition of game state"
		case 6: return "Fivefold repetition of game state"
		case 7: return "FiftyMoveRule"
		case 8: return "SeventyFiveMoveRule"
		case 9: return "Insufficient material"
		default: return "Invalid method"
	}
}
