DumbChess: *.go
	go build .
	@#tinygo build -scheduler none -panic trap .
	strip DumbChess

run: DumbChess
	@./DumbChess
