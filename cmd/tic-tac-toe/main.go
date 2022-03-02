package main

import (
	"flag"
	"fmt"
	"go-tic-tac-toe/internal/cliGame"
	"go-tic-tac-toe/internal/common"
	"go-tic-tac-toe/internal/config"
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
	cfg := config.NewConfigFromEnv()
	flag.StringVar(&cfg.GameVariant, "v", "http", "Setup using http variant of game, valid values is 'h'(http) or 'c'(cli)")
	flag.StringVar(&cfg.PlayersSides, "p", "XO", "Setup players marks and order")
	flag.IntVar(&cfg.FieldSize, "s", 3, "Setup square field size")
	flag.IntVar(&cfg.FieldWinSeq, "w", 3, "Setup win sequence length, should be greater than 2 and less or equal field size")
	configJsonFileName := flag.String("c", "", "Path to file with config in json")
	interactive := flag.Bool("i", false, "Ask reenter config from console")
	flag.Parse()
	if file, err := os.Open(*configJsonFileName); err == nil {
		cfg.ReadJsonAndMerge(file)
		file.Close()
	} else {
		if *configJsonFileName != "" {
			fmt.Println("Skip file reading:", err)
		}
	}
	if *interactive {
		fmt.Print("Enter game variant: ")
		fmt.Scan(&cfg.GameVariant)
		fmt.Print("Enter your field size: ")
		fmt.Scan(&cfg.FieldSize)
		fmt.Print("Enter players marks and order: ")
		fmt.Scan(&cfg.PlayersSides)
	}
	instField, errField := field.New(cfg.FieldSize, cfg.FieldWinSeq)
	if errField != nil {
		panic(errField)
	}
	// TODO trim cfg.PlayersSides
	sides := []rune(cfg.PlayersSides)
	players := make([]common.IPlayer, len(sides))
	for i := range players {
		players[i] = player.New(string(sides[i]))
	}
	instGame := game.NewGame(instField, players...)
	switch cfg.GameVariant {
	case "h", "http", "html":
		startBrowser("http://127.0.0.1:8080")
		httpGame.NewHttpGame(&instGame, field.NewIField, player.NewIPlayer)
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
