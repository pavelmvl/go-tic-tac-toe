package main

import (
	"fmt"
	"go-tic-tac-toe/internal/initPlayer"
)

func main() {
	player := initPlayer.InitPlayer()
	fmt.Println(player)
}
