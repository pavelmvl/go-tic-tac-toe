package main

import (
	"flag"
	"fmt"
	"go-tic-tac-toe/internal/field"
	"go-tic-tac-toe/internal/player"
)

func main() {
	// init
	s := flag.Int("s", 3, "Setup square field size, Default is 3")
	p := flag.String("p", "", "Setup player mark. Default is empty")
	flag.Parse()
	field, err := field.New(*s)
	if err != nil {
		panic(err)
	}
	if *p == "" {
		fmt.Print("Enter your mark: ")
		fmt.Scan(p)
	}
	player := player.New(*p)
	_ = player
	// work
	var col int
	var row int
	var iter int = 0
	for {
		fmt.Print("(", iter, ")Enter <column> <row>: ")
		fmt.Scan(&col, &row)
		err := field.IsCellValid(col, row)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Try enter againg")
			continue
		}
		iter++
	}
}
