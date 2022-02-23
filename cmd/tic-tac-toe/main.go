package main

import (
	"flag"
	"fmt"
	"go-tic-tac-toe/internal/cliGame"
	"go-tic-tac-toe/internal/common"
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
	a := flag.Bool("a", false, "Ask config by cli or webUI")
	s := flag.Int("s", 3, "Setup square field size, Default is 3")
	p := flag.String("p", "X", "Setup player mark. Default is 'X'")
	v := flag.String("v", "cli", "Setup using http variant of game, default is cli, valid values is 'h'(http) or 'c'(cli)")
	flag.Parse()
	if *a {
		fmt.Print("Enter your field size: ")
		fmt.Scan(s)
		fmt.Print("Enter your mark: ")
		fmt.Scan(p)
		fmt.Print("Enter game variant: ")
		fmt.Scan(v)
	}
	instField, errField := field.New(*s)
	if errField != nil {
		panic(errField)
	}
	players := make([]common.IPlayer, 0, 2)
	players = append(players, player.New(*p))
	if players[0].GetMark() == 'X' {
		players = append(players, player.New("O"))
	} else {
		players = append(players, player.New("X"))
	}
	instGame := game.NewGame(instField, players...)
	switch *v {
	case "h", "http":
		startBrowser("http://127.0.0.1:8080")
		httpGame.NewHttpGame(&instGame)
	default:
		cliGame.PlayCliGame(&instGame, os.Stdin, os.Stdout, os.Stderr)
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
