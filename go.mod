module go-tic-tac-toe

go 1.17

replace go-tic-tac-toe/internal/player v0.0.0 => ./internal/player

replace go-tic-tac-toe/internal/field v0.0.0 => ./internal/field

require go-tic-tac-toe/internal/player v0.0.0

require go-tic-tac-toe/internal/field v0.0.0
