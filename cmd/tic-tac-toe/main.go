package main

import (
	"flag"
	"fmt"
	"go-tic-tac-toe/internal/field"
	"go-tic-tac-toe/internal/player"
)

func main() {
	p := flag.String("p", "", "Setup player mark. Default is 'x'")
	Rows := flag.Int("r", 3, "Setup rows size. Default is '3'")
	Cols := flag.Int("c", 3, "Setup columns size. Default is '3'")
	flag.Parse()
	if *p == "" {
		fmt.Print("Enter your mark: ")
		fmt.Scan(p)
	}
	field, err := field.New(*Rows, *Cols)
	if err != nil {
		panic(err)
	}
	field.AddPlayer(player.New(*p))
	fmt.Println(field)
	var col int
	var row int
	var work bool = true
	for iter := 0; work; iter++ {
		fmt.Print("(", iter, ")Enter <column> <row>: ")
		fmt.Scan(&col, &row)
		err := field.Check(col, row)
		if err != nil {
			fmt.Print(err)
			work = false
		}
	}
}
