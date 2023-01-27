run: DumbChess
	@./DumbChess

DumbChess: *.go
	tinygo build -scheduler none -panic trap .
	strip DumbChess
