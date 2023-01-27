run: DumbChess
	@./DumbChess

DumbChess: chess.go
	tinygo build -scheduler none -panic trap .
	strip DumbChess
