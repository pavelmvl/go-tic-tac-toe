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
	instField, err := field.New(*s)
	if err != nil {
		panic(err)
	}
	if *p == "" {
		fmt.Print("Enter your mark: ")
		fmt.Scan(p)
	}
	players := make([]player.Player, 0, 2)
	players = append(players, player.New(*p))
	if players[0].Mark == 'X' {
		players = append(players, player.New("O"))
	} else {
		players = append(players, player.New("X"))
	}
	// work
	var col int
	var row int
	var iter int = 0
	for {
		for _, p := range players {
			// Enter and validate cell coord
			for {
				fmt.Print("(", iter, ")(", string(p.Mark), ")Enter <column> <row>: ")
				fmt.Scan(&col, &row)
				err := instField.IsCellValid(col, row)
				if err != nil {
					fmt.Println(err)
					fmt.Println("Try enter againg")
					continue
				}
				err = instField.IsCellFree(col, row)
				if err != nil {
					fmt.Println(err)
					fmt.Println("Try enter againg")
					continue
				}
				break
			}
			// mark cell
			instField.AssignCell(col, row, p.Mark)
			// print current field
			instField.Print()
			// check winners
			winner := instField.IsCellWinner(col, row)
			if winner != field.NoWinner {
				fmt.Println(string(winner), "is won")
				return
			}
			// check draw
			if instField.IsFieldFull() == true {
				fmt.Println("friendship is won")
				return
			}
		}
		iter++
	}
}
