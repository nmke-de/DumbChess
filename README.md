# DumbChess

**DumbChess** is a collection of deliberately stupid and not-so-bright minded chess bots to pit against each other or play against. Inspired by [Tom7](http://tom7.org/chess/).

## But why?

1. Fun.
2. Learning [Go](https://go.dev).

## Dependencies

Some engines currently require [Stockfish](https://stockfishchess.org/). The program will start without it too, but sometimes crash.

### Build Dependencies

- [Go](https://go.dev)
- [This chess library](https://github.com/notnil/chess)

## Build (and Run)

Type `make`.

## Usage

Type `./DumbChess` or `make` in the directory where this project is located. The program will then randomly choose two players (that could also be the human player) to pit against each other. After a while (this may take a few minutes depending on the players), an ouput akin to this will show:

```
White: stockfish
Black: MinOpponentMoves

 A B C D E F G H
8- - - - - ♝ ♞ ♜
7♟ ♖ - - - ♟ ♟ ♟
6- - ♟ - ♚ - - -
5- - - - ♕ - - -
4- - ♟ ♙ - - - -
3- - ♙ - ♙ - - ♘
2♙ - - - - ♙ ♙ ♙
1- - ♗ - ♔ - - ♖

Outcome: 1-0
Cause: Checkmate
Moves:
1. d4 d5 2. c4 Bh3 3. Nxh3 dxc4 4. e3 Qd5 5. Nc3 Qa5 6. Bxc4 Qxc3+ 7. bxc3 b5 8. Qf3 bxc4 9. Qxa8 c6 10. Qxb8+ Kd7 11. Rb1 e5 12. Rb7+ Ke6 13. Qxe5# 1-0
```

If the human player is picked, you will be prompted to enter a move in the algebraic notation (i.e. `e4` or `Nxf3`). Only valid moves will be accepted.

To take a better look at a match move by move, I will refer you to copy the text behind `Moves:` to `https://www.chess.com/analysis`.

## Dumb chess engines (and Stockfish)

For a list of implemented chess engines, refer to [RandomPlayer.go](RandomPlayer.go).
