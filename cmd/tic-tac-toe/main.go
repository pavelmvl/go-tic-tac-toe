package main

import (
	"flag"
	"fmt"
	"go-tic-tac-toe/internal/field"
	"go-tic-tac-toe/internal/game"
	"go-tic-tac-toe/internal/httpGame"
	"go-tic-tac-toe/internal/player"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	// init
	s := flag.Int("s", 3, "Setup square field size, Default is 3")
	p := flag.String("p", "", "Setup player mark. Default is empty")
	http := flag.Bool("h", false, "Setup using http variant of game, default is false")
	flag.Parse()
	instField, errField := field.New(*s)
	if errField != nil {
		panic(errField)
	}
	if *p == "" {
		fmt.Print("Enter your mark: ")
		fmt.Scan(p)
	}
	players := make([]game.IPlayer, 0, 2)
	players = append(players, player.New(*p))
	if players[0].GetMark() == 'X' {
		players = append(players, player.New("O"))
	} else {
		players = append(players, player.New("X"))
	}
	instGame := game.NewGame(instField, players...)
	if *http {
		startBrowser("http://127.0.0.1:8080")
		err := httpGame.NewHttpGame(&instGame)
		fmt.Println(err)
		os.Exit(0)
	}
	// work
	var col, row int
	for {
		// Enter and validate cell coord
		for {
			p := instGame.GetCurrentPlayer()
			fmt.Print("(", instGame.GetIter(), ")(", string(p.GetMark()), ")Enter <column> <row>: ")
			fmt.Scan(&col, &row)
			// mark cell
			err := instGame.NextMove(col, row)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Try enter againg")
				continue
			}
			break
		}
		fmt.Print(instGame.ToString())
		winnerString, winnerErr := instGame.GetWinnerString()
		if winnerErr == nil {
			fmt.Println(winnerString)
			return
		}
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
