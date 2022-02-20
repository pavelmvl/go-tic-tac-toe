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
				fmt.Print("(", iter, ")(", string(p.GetMark()), ")Enter <column> <row>: ")
				fmt.Scan(&col, &row)
				// mark cell
				err := instField.AssignCell(col, row, p.GetMark())
				if err != nil {
					fmt.Println(err)
					fmt.Println("Try enter againg")
					continue
				}
				break
			}
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

/// NOTE This func is copied from tour of golang
// startBrowser tries to open the URL in a browser, and returns
// whether it succeed.
func startBrowser(url string) bool {
	// try to start the browser
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}
