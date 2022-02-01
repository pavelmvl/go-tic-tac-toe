package main

import (
	"flag"
	"fmt"
	"go-tic-tac-toe/internal/player"
)

func main() {
	p := flag.String("p", "", "Setup player mark. Default is 'x'")
	flag.Parse()
	if *p == "" {
		fmt.Print("Enter your mark: ")
		fmt.Scan(p)
	}
	player := player.New(*p)
	fmt.Println(player)
}
