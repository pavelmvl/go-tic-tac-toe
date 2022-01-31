package main

import (
	"fmt"
	"go-tic-tac-toe/internal/pkg/initPlayer"
)

func main() {
	player := initPlayer.InitPlayer()
	fmt.Println(player)
}
