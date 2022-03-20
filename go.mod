module go-tic-tac-toe

go 1.17

require (
	go-tic-tac-toe/internal/field v0.0.0
	go-tic-tac-toe/internal/game v0.0.0
	go-tic-tac-toe/internal/httpGame v0.0.0
	go-tic-tac-toe/internal/player v0.0.0
)

replace (
	go-tic-tac-toe/internal/field v0.0.0 => ./internal/field
	go-tic-tac-toe/internal/game v0.0.0 => ./internal/game
	go-tic-tac-toe/internal/httpGame v0.0.0 => ./internal/httpGame
	go-tic-tac-toe/internal/player v0.0.0 => ./internal/player
)
